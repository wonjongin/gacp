run:
	@go run gacp.go

build:
	@go build gacp.go -o build/gacp

cc:
	@echo "Cross Compile"
	GOOS=darwin GOARCH=amd64 go build -o build/gacp-darwin-amd64.exe gacp.go
	GOOS=linux GOARCH=amd64 go build -o build/gacp-linux-amd64.exe gacp.go
	GOOS=windows GOARCH=amd64 go build -o build/gacp-windows-amd64.exe gacp.go

ccwindow: 
	@echo "Cross Compile"
	set GOOS=darwin& set GOARCH=amd64& go build -o build/gacp-darwin-amd64.exe gacp.go
	set GOOS=linux& set GOARCH=amd64& go build -o build/gacp-linux-amd64.exe gacp.go
	set GOOS=windows& set GOARCH=amd64& go build -o build/gacp-windows-amd64.exe gacp.go

install:
	@go build gacp.go -o build/gacp
	@cp build/gacp /usr/local/bin/gacp