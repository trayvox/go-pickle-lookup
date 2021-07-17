package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	api "github.com/trayvox/go-eth-multicall"
	"github.com/trayvox/go-pickle-lookup/api/FeeDistributor"
	"github.com/trayvox/go-pickle-lookup/api/Gauge"
	"github.com/trayvox/go-pickle-lookup/api/IERC20"
	"github.com/trayvox/go-pickle-lookup/api/PickleJar"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

func formatAmount(amount *big.Int, decimals int, precision int) string {
	res, _, err := big.ParseFloat(amount.String(), 10, 0, big.ToZero)
	if err != nil {
		panic(err)
	}
	temp := new(big.Float).Quo(res, big.NewFloat(math.Pow10(decimals)))
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%."+strconv.Itoa(precision)+"f", temp), "0"), ".")

}

var erc20Abi, _ = abi.JSON(strings.NewReader(IERC20.IERC20ABI))
var gaugeAbi, _ = abi.JSON(strings.NewReader(Gauge.GaugeABI))
var jarAbi, _ = abi.JSON(strings.NewReader(PickleJar.PickleJarABI))
var feeDistributorABI, _ = abi.JSON(strings.NewReader(FeeDistributor.FeeDistributorABI))

var ETHER, _ = new(big.Int).SetString("1000000000000000000", 10)

func GetBalanceCall(name string, tokenAddress common.Address, userAddress common.Address) api.Call {
	parsedData, err := erc20Abi.Pack("balanceOf", userAddress)
	if err != nil {
		panic(err)
	}
	return api.Call{
		Target:   tokenAddress,
		CallData: parsedData,
		Name:     name,
	}
}

func GetGaugeEarnedCall(name string, contractAddress common.Address, userAddress common.Address) api.Call {
	parsedData, err := gaugeAbi.Pack("earned", userAddress)
	if err != nil {
		panic(err)
	}
	return api.Call{
		Target:   contractAddress,
		CallData: parsedData,
		Name:     name,
	}
}

func GetFarmRatioCall(name string, contractAddress common.Address) api.Call {
	parsedData, err := jarAbi.Pack("getRatio")
	if err != nil {
		panic(err)
	}
	return api.Call{
		Target:   contractAddress,
		CallData: parsedData,
		Name:     name,
	}
}

type Position struct {
	Farm     common.Address `json:"farm"`
	Jar      common.Address `json:"jar"`
	Token    string         `json:"token"`
	Decimals int            `json:"decimals"`
}

type Data struct {
	Pickle         common.Address      `json:"pickle"`
	Dill           common.Address      `json:"dill"`
	FeeDistributor common.Address      `json:"FeeDistributor"`
	Positions      map[string]Position `json:"positions"`
}

type Config struct {
	RPCUrl string           `json:"rpc_url"`
	Users  []common.Address `json:"users"`
}

func loadConfig() Config {
	jsonFile, err := os.Open("Config.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	var config Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func loadSetup() Data {
	jsonFile, err := os.Open("PickleData.json")
	if err != nil {
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)

	var data Data

	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		panic(err)
	}

	return data
}

func LookupUsers(users []common.Address, caller api.EthMultiCaller) {
	green := color.New(color.FgGreen).SprintFunc()
	boldGreen := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	red := color.New(color.FgHiRed).SprintFunc()
	data := loadSetup()

	// We need to load the ratios between the farm tokens and the underlying token
	var ratioCalls = make([]api.Call, 0)
	for name, position := range data.Positions {
		ratioCalls = append(ratioCalls, GetFarmRatioCall(name, position.Jar))
	}
	ratioResponse := caller.Execute(ratioCalls)

	for _, user := range users {
		var userCalls = make([]api.Call, 0)
		userCalls = append(userCalls, GetBalanceCall("DillBalance", data.Dill, user))
		userCalls = append(userCalls, GetBalanceCall("PickleBalance", data.Pickle, user))
		// Underlying token name?

		for name, position := range data.Positions {
			userCalls = append(userCalls, GetBalanceCall(name+"_farm_bal", position.Farm, user))
			userCalls = append(userCalls, GetGaugeEarnedCall(name+"_farm_earn", position.Farm, user))
			userCalls = append(userCalls, GetBalanceCall(name+"_jar_bal", position.Jar, user))
		}

		start := time.Now()
		multiCallResponse := caller.Execute(userCalls)
		elapsed := time.Since(start)

		dillBalance := new(big.Int).SetBytes(multiCallResponse["DillBalance"].ReturnData)
		pickleBalance := new(big.Int).SetBytes(multiCallResponse["PickleBalance"].ReturnData)

		fmt.Printf("%s (%s pickle, %s dill) have following positions (%s ms):\n",
			boldGreen(user.Hex()),
			green(formatAmount(pickleBalance, 18, 3)),
			boldGreen(formatAmount(dillBalance, 18, 3)),
			red(elapsed.Milliseconds()),
		)

		// Need to look at the dill & pickle balance as well

		totalClaimablePickle := big.NewInt(0)

		// Loop over the positions
		for name, position := range data.Positions {
			ratio := new(big.Int).SetBytes(ratioResponse[name].ReturnData)
			farmBalance := new(big.Int).SetBytes(multiCallResponse[name+"_farm_bal"].ReturnData)
			pickleEarned := new(big.Int).SetBytes(multiCallResponse[name+"_farm_earn"].ReturnData)
			jarBalance := new(big.Int).SetBytes(multiCallResponse[name+"_jar_bal"].ReturnData)
			totalClaimablePickle = new(big.Int).Add(totalClaimablePickle, pickleEarned)

			// If user has balances or claimable pickle print the position:
			if farmBalance.Cmp(common.Big0) > 0 || pickleEarned.Cmp(common.Big0) > 0 || jarBalance.Cmp(common.Big0) > 0 {
				farmBal := new(big.Int).Div(new(big.Int).Mul(farmBalance, ratio), ETHER)
				jarBal := new(big.Int).Div(new(big.Int).Mul(jarBalance, ratio), ETHER)
				fmt.Printf("\t%s farming in Pickle: %s (%s in jar). Claimable Pickle: %s\n",
					green(position.Token),
					formatAmount(farmBal, position.Decimals, 3),
					formatAmount(jarBal, position.Decimals, 3),
					formatAmount(pickleEarned, 18, 3),
				)
			}
		}

		distributorCallData, _ := feeDistributorABI.Pack("claim")
		feeDistributorMsg := ethereum.CallMsg{
			From: user,
			To:   &data.FeeDistributor,
			Gas:  0,
			Data: distributorCallData,
		}
		feeRes, _ := caller.Client.CallContract(context.Background(), feeDistributorMsg, nil)
		unpacked, _ := feeDistributorABI.Unpack("claim", feeRes)
		pickleFromFeeDistributor := unpacked[0].(*big.Int)
		totalClaimablePickle = new(big.Int).Add(totalClaimablePickle, pickleFromFeeDistributor)
		if pickleFromFeeDistributor.Cmp(common.Big0) > 0 {
			fmt.Printf("\tPickle from dill distributor: %s pickles\n",
				formatAmount(pickleFromFeeDistributor, 18, 3),
			)
		}

		fmt.Printf("\tTotal claimable pickle: %s pickles\n\n", formatAmount(totalClaimablePickle, 18, 3))
	}
}

func main() {
	boldGreen := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	config := loadConfig()
	caller := api.New(config.RPCUrl)
	fmt.Printf("Looking for positions in %s:\n", boldGreen("Pickle Finance"))
	LookupUsers(config.Users, caller)
}
