build_executeable_file: 
	env GOOS=linux GOARCH=amd64 go build
	env GOOS=linux GOARCH=arm go build -o auto-compound-arm32
	env GOOS=linux GOARCH=arm64 go build -o auto-compound-arm64
	env GOOS=windows GOARCH=amd64 go build

