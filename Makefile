config ?= "dev.json"

all: build

setup:
	go get golang.org/x/tools/cmd/cover

build:
	go build -o ./build/micro

run:
	go run main.go -config=$(config)

test:
	mkdir -p build
	go test -coverprofile=build/coverage.out

coverage: test
	go tool cover -html=build/coverage.out

clean:
	-rm -r ./build
	-rm terraform.tfstate
	-rm terraform.tfstate.backup

ci: clean test

.PHONY: build run test ci clean
