config ?= "dev.json"

all: build

setup:
	go get golang.org/x/tools/cmd/cover

build:
	go build -o ./build/micro ./src

run:
	go run ./src/main.go -config=$(config)

test:
	mkdir -p build
	go test -coverprofile=build/coverage.out ./src/

coverage: test
	go

clean:
	-rm -r ./build
	-rm terraform.tfstate
	-rm terraform.tfstate.backup

ci: clean test

.PHONY: build run test ci clean
