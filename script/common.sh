#!/bin/sh
set -euo pipefail
ROOT=`git rev-parse --show-toplevel`
MOUNTOPTION=""
if [ "$(uname -s)" = "Darwin" ]; then
    MOUNTOPTION=":delegated"
fi

#TODO: replace the PACKAGE env variable with the name of your project
#      Also make sure the GOIMAGE variable is pointed to our most recent one
#PACKAGE=git.rsglab.com/rsg/ephemera
#GOIMAGE=dockerfactory.rsglab.com/rsg/centos/7/golang:1.10.3-2b9b417

#dockerwrap() {
#    docker run -t --rm \
#        -u $(id -u):$(id -g) \
#        -v "${ROOT}:/go/src/${PACKAGE}${MOUNTOPTION}" \
#        -w "/go/src/${PACKAGE}" \
#        -e GOPATH=/go \
#        -e CGO_ENABLED=0 \
#        -e GOOS=linux \
#        -e GOARCH=amd64 \
#        ${GOIMAGE} \
#        "$@"
#}
