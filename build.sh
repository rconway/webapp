#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-$(basename "$PWD")}"

if ! hash pkger >/dev/null 2>&1; then
  echo "Install pkger tool with 'go get'..."
  go get -v github.com/markbates/pkger/cmd/pkger
fi

# pkger generate
if test ! -f pkged.go; then
  echo "Running pkger to bundle static assets..."
  pkger
fi

# static go build
echo "Creating static binary..."
GOPATH=$(pwd)/go \
CGO_ENABLED=0 \
  go build -o "${EXENAME}"
echo "...created binary '${EXENAME}'"
