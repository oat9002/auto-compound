# auto-compound

-   ~~Query and Auto compound following cron for cake-cake pool~~ (CakePool is migrated to MasterChefV2, no manual compound)
-   Notify with Line Notify or Telegram

## Development in local

Environment is defined in `.env` or using command-line argument, `-dev` in case using the binary file. Default is `production`.

## How to generate smart contract interface

### Prerequisite

-   docker

### Step

1. find any smart contract you'd like to work with

2. copy abi and binary code for that smart contract into your machine, for instance CakePool contract, you can find it from bscscan.

3. Run

```
docker run --rm -v <path-which-have-abi-and-bin>:/sources -v <path-to-project>/auto-compound/contracts:/output ethereum/client-go:alltools-latest abigen --bin=/sources/<bin-file> --abi=/sources/<abi-file> --pkg=<contract-name> --out=/output/<contract-name>.go
```

4. Run `sudo chmod 755 ./contracts/<contract-name>.go`
