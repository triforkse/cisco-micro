config ?= "dev.json"
args ?= "apply"

all: build

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
	-rm -rf ./build

ci: clean
	./ci.sh

.PHONY: build run run-build test ci clean
