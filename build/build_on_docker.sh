#!/usr/bin/env bash

# Prerequisite: set absolute path to project root to $PROJECT_ROOT

BUILD_PATH=/tmp/out
DIST_PATH=/dist

TARGET_TOOLS="dupload dsharedlink dteammember dcmp"
TARGET_PLATFORM="windows/386,windows/amd64,darwin/amd64,linux/386,linux/amd64"

BUILD_VERSION=$(cat $PROJECT_ROOT/version)
BUILD_HASH=$(cd $PROJECT_ROOT && git rev-parse HEAD)

echo --------------------
echo BUILD: Testing..

cd $PROJECT_ROOT
go test $(glide novendor)


for t in $TARGET_TOOLS; do
  echo --------------------
  echo BUILD: Building tool $t

  X_APP_NAME="-X github.com/watermint/toolbox/infra/knowledge.AppName=$t"
  X_APP_VERSION="-X github.com/watermint/toolbox/infra/knowledge.AppVersion=$BUILD_VERSION"
  X_APP_HASH="-X github.com/watermint/toolbox/infra/knowledge.AppHash=$BUILD_HASH"
  X_APP_CREDENTIALS=""
  if [ -e $PROJECT_ROOT/credentials.secret ]; then
    X_APP_CREDENTIALS=$(cat $PROJECT_ROOT/credentials.secret | xargs)
  fi
  LD_FLAGS="$X_APP_NAME $X_APP_VERSION $X_APP_HASH $X_APP_CREDENTIALS"

  xgo --ldflags="$LD_FLAGS" -out $t-$BUILD_VERSION -targets $TARGET_PLATFORM github.com/watermint/toolbox/tools/$t 
done

echo --------------------
echo BUILD: Packaging
cd /build
zip -9 -r $DIST_PATH/toolbox-$BUILD_VERSION.zip .
