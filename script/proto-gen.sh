#!/bin/bash
set -eu
set -o pipefail
set -x

source "$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )/common.sh"

IMAGE="us.gcr.io/rsg-base-prod/grpc-generator:v0.1.0"

module=$(awk '/module/ {print $2}' go.mod)

rm -rf ./internal/proto
mkdir -p ./internal/proto

for dir in $(find proto -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
    dockerwrap $IMAGE protoc -I ./proto --go_out=./internal/proto --go-grpc_out=./internal/proto "$dir"/*.proto
done

mv "internal/proto/$module"/internal/proto/* ./internal/proto/
find ./internal/proto -empty -type d -delete
