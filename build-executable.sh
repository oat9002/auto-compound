#!/bin/bash

echo 'Build linux amd 64'
env GOOS=linux GOARCH=amd64 go build
chmod +x ./auto-compound

echo 'Build linux arm 32'
env GOOS=linux GOARCH=arm go build -o auto-compound-arm32
chmod +x ./auto-compound-arm32

echo 'Build linux arm 64'
env GOOS=linux GOARCH=arm64 go build -o auto-compound-arm64
chmod +x ./auto-compound-arm64

echo 'Build windows amd 64'
env GOOS=windows GOARCH=amd64 go build
chmod +x ./auto-compound.exe


