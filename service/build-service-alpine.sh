#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-webapp-alpine}"

# go cache directory
GOROOT="goroot-alpine"
mkdir -p ${GOROOT}/.cache

docker run --user $(id -u):$(id -g) --rm -it \
  -v $PWD/..:/src \
  -v $PWD/${GOROOT}/.cache:/.cache \
  -e GOPATH_BUILD=/src/service/${GOROOT}/go \
  golang:alpine \
  /src/service/build-service.sh "${EXENAME}"
