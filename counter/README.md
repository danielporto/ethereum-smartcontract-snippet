How to create commands and subcommands:


https://www.linode.com/docs/guides/using-cobra/
## Deploy a contract
It returns a contract address (use it for printing the counter at the end)
```
./counter deploy --url "http://146.193.41.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```


## Print the content of a counter
```
./counter  --url "http://146.193.41.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0" print --contract "0x3debb8a04d6311cfdd72b44ac2db2c4920e9db2d" 
```

## Run increment/decrement
It requires a contract address. 
```
./counter increment --contract "0x3debb8a04d6311cfdd72b44ac2db2c4920e9db2d"  --url "http://146.193.41.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```

## Run workload
It creates a contract and run operations on it
```
./counter workload -o increment -c 10000  --url "http://146.193.41.166:22000" --key "1be3b50b31734be48452c29d714941ba165ef0cbf3ccea8ca16c45e3d8d45fb0"
```
