APP_NAME=date-manager-api

build:
	GOOS=linux GOARCH=amd64 go build -o $(APP_NAME)_linux_amd64 cmd/main.go
	GOOS=windows GOARCH=amd64 go build -o $(APP_NAME)_windows_amd64.exe cmd/main.go
	GOOS=darwin GOARCH=amd64 go build -o $(APP_NAME)_darwin_amd64 cmd/main.go

run:
	go run cmd/main.go

clean:
	rm -f $(APP_NAME)_linux_amd64 $(APP_NAME)_windows_amd64.exe $(APP_NAME)_darwin_amd64
