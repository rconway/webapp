#!/usr/bin/env bash

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

pushd app
echo "Building the React app..."
if [ ! -d node_modules ]; then
  echo "  Installing dependencies..."
  npm install
  let status=$?
fi
npm run build
let status=$status+$?
echo "...built app."
popd

exit $status
