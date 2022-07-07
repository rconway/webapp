#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-$(basename "$PWD")}"

# React app...
./build-app-docker.sh

# Go service
./build-service-docker.sh
