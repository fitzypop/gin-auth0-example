build:
    go build -o build/api

run: build
    ./build/api

test:
    go test -v ./...
