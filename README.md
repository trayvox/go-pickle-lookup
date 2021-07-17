# go-pickle-lookup

Tool written in Golang to retrieve your holdings in Pickle.Finance.

Enter a `rpc_url` along with addresses to retrieve positions for in the `Config.json` file.

The provided `Config.json` use addresses from the latest snapshot vote.

To run:

```
// Build
go build

// Execute 
./go-pickle-lookup
```

