default:
  just --list

build: install
    go build -o build/goapi main.go && chmod +x build/goapi

run: build
    ./build/goapi

install:
    go mod download

update:
    go get -u .

test:
    go test ./...
