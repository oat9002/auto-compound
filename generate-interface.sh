#!/bin/bash
cd ./pancake-farm

yarn go:interface

cd ..

docker run --rm -v /home/oat/Desktop/jobcl/auto-compound/pancake-farm/abigenBindings:/sources -v /home/oat/Desktop/jobcl/auto-compound/contracts:/output ethereum/client-go:alltools-latest abigen --bin=/sources/bin/MasterChef.bin --abi=/sources/abi/MasterChef.abi --pkg=masterChef --out=/output/MasterChef.go

sudo chmod 755 ./contracts/MasterChef.go

