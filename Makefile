config ?= "dev.json"
action ?= "apply"

all: build

setup:
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/lint/golint

build:
	go build -o build/micro cisco/micro/micro

run: build
	./build/micro -config=$(config) -debug=true $(action)

test:
	mkdir -p build
	go test ./...

coverage:
	./goclean.sh

clean:
	-rm -r ./build

ci: clean coverage

.PHONY: build run test ci clean
