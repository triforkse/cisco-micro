config ?= "dev.json"
action ?= "apply"

all: build

setup:
	go get golang.org/x/tools/cmd/cover

build:
	go build -o build/micro cisco/micro/micro

run: build
	./build/micro -config=$(config) -debug=true $(action)

test:
	mkdir -p build
	go test -coverprofile=build/coverage.out

coverage: test
	go tool cover -html=build/coverage.out

clean:
	-rm -r ./build

ci: clean test

.PHONY: build run test ci clean
