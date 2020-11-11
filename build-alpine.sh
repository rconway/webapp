#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

docker run --rm -it -v $PWD/..:/src -e GOPATH_BUILD=/src/webapp/go-alpine golang:alpine sh /src/webapp/build.sh webapp-alpine
