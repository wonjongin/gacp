run:
	@go run gacp.go

build:
	@go build -o build/gacp gacp.go
clean:
	@rm -r build 
	@rm -r dist
cc:
	@echo "Cross Compile"
	@mkdir build/gacp-darwin-amd64 build/gacp-linux-amd64 build/gacp-windows-amd64
	GOOS=darwin GOARCH=amd64 go build -o build/gacp-darwin-amd64/gacp gacp.go
	GOOS=linux GOARCH=amd64 go build -o build/gacp-linux-amd64/gacp gacp.go
	GOOS=windows GOARCH=amd64 go build -o build/gacp-windows-amd64/gacp.exe gacp.go
pkg:
	@echo "Packging"
	@mkdir dist build build/gacp-darwin-amd64 build/gacp-linux-amd64 build/gacp-windows-amd64
	GOOS=darwin GOARCH=amd64 go build -o build/gacp-darwin-amd64/gacp gacp.go
	GOOS=linux GOARCH=amd64 go build -o build/gacp-linux-amd64/gacp gacp.go
	GOOS=windows GOARCH=amd64 go build -o build/gacp-windows-amd64/gacp.exe gacp.go
	@pwd
	@zip -j dist/gacp-darwin-amd64.zip build/gacp-darwin-amd64/*
	@zip -j dist/gacp-linux-amd64.zip build/gacp-linux-amd64/*
	@zip -j dist/gacp-windows-amd64.zip build/gacp-windows-amd64/*


ccwindow: 
	@echo "Cross Compile"
	set GOOS=darwin& set GOARCH=amd64& go build -o build/gacp-darwin-amd64 gacp.go
	set GOOS=linux& set GOARCH=amd64& go build -o build/gacp-linux-amd64 gacp.go
	set GOOS=windows& set GOARCH=amd64& go build -o build/gacp-windows-amd64.exe gacp.go

install:
	@go build gacp.go -o build/gacp
	@cp build/gacp /usr/local/bin/gacp
