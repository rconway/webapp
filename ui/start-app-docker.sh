#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

docker run --rm -it -u node -v ${PWD}:/app -w /app -p 3000:3000 node ./start-app.sh
