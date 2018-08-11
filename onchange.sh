#!/bin/bash -e

path=$1
dir=$(dirname "$path")
filename=$(basename "$path")
extension="${filename##*.}"
nameonly="${filename%.*}"

case $extension in
    go)
        gofmt -w $path
        ;;
esac

case $filename in
    token.go)
	go generate
	;;
esac

go vet
go test -coverprofile /tmp/c.out .
uncover /tmp/c.out
