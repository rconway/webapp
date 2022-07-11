#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

# Running go tests
echo "Running service tests..."
go test -v ./api/...
let status=$?
echo "...service tests done"

exit $status
