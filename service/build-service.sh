#!/usr/bin/env sh

ORIG_DIR="$(pwd)"
cd "$(dirname "$0")"
BIN_DIR="$(pwd)"

trap "cd '${ORIG_DIR}'" EXIT

EXENAME="${1:-webapp}"

export GOPATH="${GOPATH_BUILD:-$(pwd)/go}"
export PATH="${GOPATH}/bin:${PATH}"

# copy the UI app build files
rm -rf cmd/webapp/app
cp -r ../ui/app/build cmd/webapp/app

# static go build
echo "Creating static binary..."
CGO_ENABLED=0 go build -o "${EXENAME}" ./cmd/webapp
status=$?
echo "...created binary '${EXENAME}'"

exit $status
