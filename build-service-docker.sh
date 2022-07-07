#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-$(basename "$PWD")}"

# go cache directory
GOROOT="goroot"
mkdir -p ${GOROOT}/.cache

docker run --user $(id -u):$(id -g) --rm -it \
  -v $PWD/..:/src \
  -v $PWD/${GOROOT}/.cache:/.cache \
  -e GOPATH_BUILD=/src/webapp/${GOROOT}/go \
  golang \
  sh /src/webapp/build-exe.sh "${EXENAME}"
