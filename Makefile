icon:
        rsrc -ico go.ico

build:
        go get -u -v all
        go build -ldflags "-s -w" -o bin/ .

run:
        go run .

compile:
        go get -u -v all
        GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/main-linux-arm .
        GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o bin/main-linux-arm64 .
        GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/main-windows-amd64 .
        GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/app-amd64-macos .

all: icon build
