config ?= "dev.json"

all: build

setup:
	go get golang.org/x/tools/cmd/cover

build:
	go build -o ./build/micro

run: build
	./build/micro -config=$(config) -debug=true

test:
	mkdir -p build
	go test -coverprofile=build/coverage.out

coverage: test
	go tool cover -html=build/coverage.out

clean:
	-rm -r ./build

ci: clean test

.PHONY: build run test ci clean
