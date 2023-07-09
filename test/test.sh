#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

onExit() {
  cd "${ORIG_DIR}"
}
trap onExit EXIT

# Exit immediately if a command fails
set -e

echo "Test service root"
curl -v webapp:8080

echo "Test API root"
curl -v webapp:8080/api

echo "Test APP root"
curl -v webapp:8080/app

echo "Test error case (wrong port)"
set +e
curl -v webapp:8081
let status=$?
set -e
test $status -eq 7
