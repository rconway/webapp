#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-$(basename "$PWD")-alpine}"

# React app...
./ui/build-app.sh

# static go build...
./service/build-service.sh
