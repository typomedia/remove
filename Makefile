icon:
	rsrc -ico go.ico

build:
	go build -ldflags "-s -w" -o bin/ .

run:
	go run .

compile:
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o bin/main-linux-arm .
	GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o bin/main-linux-arm64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o bin/main-windows-amd64 .

all: icon build
