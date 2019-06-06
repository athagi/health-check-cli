deps:
	go get -u ./...

clean: 
	rm -rf ./dest
	
build:
	GOOS=linux GOARCH=amd64 go build -o dest/health-check-cli main.go

win-build:
	GOOS=windows GOARCH=amd64 go build -o dest/health-check-cli.exe main.go