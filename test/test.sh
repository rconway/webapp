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

echo -e "\n\n### Test service root"
curl -v webapp:8080

echo -e "\n\n### Test API root"
curl -v webapp:8080/api

echo -e "\n\n### Test APP root"
curl -v webapp:8080/app

echo -e "\n\n### Test error case (wrong port)"
set +e
curl -v webapp:8081
let status=$?
set -e
test $status -eq 7
