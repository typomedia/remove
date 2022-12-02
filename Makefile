icon:
        rsrc -ico go.ico

build:
        go get -u -v all
        go build -ldflags "-s -w" -o bin/ .

run:
        go run .

compile:
        go get -u -v all
        GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/remove-linux-arm .
        GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/remove-linux-amd64 .
        GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/remove-windows-amd64 .
        GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o bin/remove-macos-amd64 .

all: icon build
