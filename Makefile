config ?= "dev.json"
args ?= "apply"

all: build

setup:
	go get -u github.com/bradfitz/goimports
	go get -u github.com/golang/lint
	go get -u golang.org/x/tools/cmd/cover
	go get -u github.com/golang/lint/golint
	go get -u github.com/axw/gocov/gocov
	go get -u github.com/mattn/goveralls

build:
	go build -o build/micro ./micro

run: build
	./build/micro -debug=true $(args)

test:
	mkdir -p build
	go test ./...

coverage: build
	./goclean.sh

clean:
	-rm -r ./build

ci: setup clean coverage

.PHONY: build run test ci clean
