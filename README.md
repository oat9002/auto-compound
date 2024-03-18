# auto-compound

- Query and Auto compound following cron for VE Cake
- Notify with Line Notify or Telegram

## Development in local

Environment is defined in `.env` or using command-line argument, `-dev` in case using the binary file. Default is `production`.

## How to generate contract interface

### Prerequisite

- docker

### Step

1. Find any contract you like, both abi and binary find and save in to your local machine

2. Start docker and run this command `(<...> fill in this first)`

```
docker run --rm -v <path-to-your-contract-folder>:/sources -v <path-to-your-contract-folder>:/output ethereum/client-go:alltools-latest abigen --bin=/sources/<your-contract-binary-file> --abi=/sources/<your-contract-abit> --pkg=<package-name> --out=/output/<go-file-name>
```

3. Run `sudo chmod 755 <path-to-your-contract-folder>/<go-file-name>`

4. Copy your go file to `contracts` folder in the project
