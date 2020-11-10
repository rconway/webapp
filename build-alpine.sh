#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

docker run --rm -it -v $HOME/go-alpine:/go -v $PWD/..:/src -e GOPATH_BUILD=/go golang:alpine sh /src/webapp/build.sh
