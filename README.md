# auto-compound

Currently support only staking in cake pool

## Development in local

Environment is defined in `.env` or using command-line argument, `-dev` in case using the binary file. Default is `production`.

## How to generate pancake contract interface

### Prerequisite

-   docker

### Step

1. Clone pancake contract into this repository [link](https://github.com/pancakeswap/pancake-farm.git)
2. Inside `pancake-farm` folder, add this config to `truffle-config.js`

```
plugins: ["@chainsafe/truffle-plugin-abigen"]
```

3. Also add this command in `package.json`, under `scripts` section

```
"go:interface": "truffle compile && truffle run abigen MasterChef"
```

4. Run `yarn install`

5. Run `yarn go:interface`

6. Go back to project root by `cd ..`

7. Run

```
docker run --rm -v <path-to-project>/auto-compound/pancake-farm/abigenBindings:/sources -v <path-to-project>/auto-compound/contracts:/output ethereum/client-go:alltools-latest abigen --bin=/sources/bin/MasterChef.bin --abi=/sources/abi/MasterChef.abi --pkg=masterChef --out=/output/MasterChef.go
```

8. Run `sudo chmod 755 ./contracts/MasterChef.go`
