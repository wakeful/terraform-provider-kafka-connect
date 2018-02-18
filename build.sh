#!/usr/bin/env bash

set -e

GO_BUILD_CMD="go build -a -installsuffix cgo"
GO_BUILD_LDFLAGS="-s -w"

BUILD_PLATFORMS="linux darwin"
BUILD_ARCHS="amd64"

mkdir -p release

for OS in ${BUILD_PLATFORMS[@]}; do
  for ARCH in ${BUILD_ARCHS[@]}; do
    NAME="terraform-provider-kafka-connect-$OS-$ARCH"
    echo "Building for $OS/$ARCH"
    GOARCH=$ARCH GOOS=$OS CGO_ENABLED=0 $GO_BUILD_CMD -ldflags "$GO_BUILD_LDFLAGS"\
     -o "release/$NAME" .
    shasum -a 256 "release/$NAME" > "release/$NAME".sha256
  done
done
