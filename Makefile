config ?= "dev.json"
args ?= "apply"

all: build

setup:
	go get -u github.com/bradfitz/goimports
	go get -u github.com/golang/lint
	go get -u golang.org/x/tools/cmd/cover

build:
	go build -o build/micro ./micro

run: build
	./build/micro -debug=true $(args)

run-build:
	make run args="build"

test:
	mkdir -p build
	go test ./...

clean:
	-rm -r ./build

ci: setup clean
	go get -u github.com/mattn/goveralls
	./ci.sh

.PHONY: setup build run run-build test ci clean
