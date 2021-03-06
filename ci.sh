#!/bin/bash
# The script does automatic checking on a Go package and its sub-packages, including:
# 1. gofmt         (http://golang.org/cmd/gofmt/)
# 2. goimports     (https://github.com/bradfitz/goimports)
# 3. golint        (https://github.com/golang/lint)
# 4. go vet        (http://golang.org/cmd/vet)
# 5. race detector (http://blog.golang.org/race-detector)
# 6. test coverage (http://blog.golang.org/cover)

set -e

go get -u github.com/bradfitz/goimports
go get -u github.com/golang/lint/golint
go get -u golang.org/x/tools/cmd/cover
go get -u github.com/mattn/goveralls

mkdir -p build

# Automatic checks
test -z "$(gofmt -l -w .     | tee /dev/stderr)"
test -z "$(goimports -l -w . | tee /dev/stderr)"
test -z "$(golint .          | tee /dev/stderr)"

go vet ./...
go test -race ./...

# Run test coverage on each subdirectories and merge the coverage profile.

echo "mode: count" > build/coverage.cov

# Standard go tooling behavior is to ignore dirs with leading underscors
for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path '*/_*' -type d);
do
if ls $dir/*.go &> /dev/null; then
    go test -covermode=count -coverprofile=$dir/profile.tmp $dir
    if [ -f $dir/profile.tmp ]
    then
        cat $dir/profile.tmp | tail -n +2 >> build/coverage.cov
        rm $dir/profile.tmp
    fi
fi
done

go tool cover -func build/coverage.cov

# Submit the test coverage result to coveralls.io
goveralls -coverprofile=build/coverage.cov -service=drone.io -repotoken $COVERALLS_TOKEN

