update-deps:
	go get -u ./...

.PHONY: clean
clean: 
	rm -rf ./dest

.PHONY: build
build: clean linux-build win-build mac-build

.PHONY: linux-build
linux-build: main.go
	GOOS=linux GOARCH=amd64 go build -o dest/linux/health-check-cli main.go

.PHONY: win-build
win-build: main.go
	GOOS=windows GOARCH=amd64 go build -o dest/windows/health-check-cli.exe main.go

.PHONY: mac-build
mac-build: main.go
	GOOS=darwin GOARCH=amd64 go build -o dest/darwin/health-check-cli main.go