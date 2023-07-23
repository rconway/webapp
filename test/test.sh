#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

onExit() {
  cd "${ORIG_DIR}"
}
trap onExit EXIT

heading() {
  echo -e "\n\n##########"
  echo "### ${1}"
  echo -e "##########\n"
}

# Exit immediately if a command fails
set -e

heading "Test service root"
curl -v webapp:8080

heading "Test API root"
curl -v webapp:8080/api

heading "Test APP root"
curl -v webapp:8080/app

heading "Test error case (wrong port)"
set +e
curl -v webapp:8081
let status=$?
set -e
test $status -eq 7
