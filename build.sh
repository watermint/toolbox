#!/usr/bin/env bash

if [ x"" = x"$PROJECT_ROOT" ]; then
  # Configure for regular build
  PROJECT_ROOT=$PWD
  BUILD_PATH=$PWD/build
  DIST_PATH=$PWD/dist
else
  # Configure for Docker build
  BUILD_PATH=$PROJECT_ROOT/build
  DIST_PATH=/dist
fi

if [ x"" = x"$1" ]; then
  TARGET=linux
else
  TARGET=$1
fi

go run tbx.go dev build target          \
  -dist-path "$DIST_PATH"                 \
  -build-path "$BUILD_PATH"               \
  -deploy-path /watermint-toolbox-build \
  -target-name "$TARGET"

bash "$BUILD_PATH"/build-target.sh
