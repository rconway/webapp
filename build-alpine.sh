#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-$(basename "$PWD")-alpine}"

docker run --rm -it -v $PWD/..:/src -e GOPATH_BUILD=/src/webapp/go-alpine golang:alpine sh /src/webapp/build.sh "${EXENAME}"
