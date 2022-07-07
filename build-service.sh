#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-$(basename "$PWD")}"

export GOPATH="${GOPATH_BUILD:-$(pwd)/go}"
export PATH="${GOPATH}/bin:${PATH}"

# static go build
echo "Creating static binary..."
CGO_ENABLED=0 go build -o "${EXENAME}"
echo "...created binary '${EXENAME}'"
