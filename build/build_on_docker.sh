#!/usr/bin/env bash

# Prerequisite: set absolute path to project root to $PROJECT_ROOT

BUILD_PATH=/tmp/out
DIST_PATH=/dist

TARGET_OS="windows darwin linux"
TARGET_ARCH="amd64 386"
TARGET_TOOLS="dupload"

BUILD_VERSION=$(cat $PROJECT_ROOT/version)

function test {
    echo BUILD: Testing..

    cd $PROJECT_ROOT
    go test $(glide novendor)
}

function build {
    echo BUILD: Building tool: $1, target OS: $2, target Arch: $3

    TOOL=$1
    export GOOS=$2
    export GOARCH=$3

    X_APP_NAME="-X github.com/watermint/toolbox/infra/knowledge.AppName=$TOOL"
    X_APP_VERSION="-X github.com/watermint/toolbox/infra/knowledge.AppVersion=$BUILD_VERSION"
    X_APP_CREDENTIALS=""

    if [ -e $PROJECT_ROOT/tools/$TOOL/credentials.secret ]; then
      X_APP_CREDENTIALS=$(cat $PROJECT_ROOT/tools/$TOOL/credentials.secret | xargs)
    fi

    TOOL_BUILD_PATH=$BUILD_PATH/$TOOL/$GOOS/$GOARCH
    mkdir -p $TOOL_BUILD_PATH
    cd $TOOL_BUILD_PATH
    go build -ldflags "$X_APP_NAME $X_APP_VERSION $X_APP_CREDENTIALS" github.com/watermint/toolbox/tools/$TOOL
}

test

for t in $TARGET_TOOLS; do
  for os in $TARGET_OS; do
    for a in $TARGET_ARCH; do
      build $t $os $a
    done
  done
  cd $BUILD_PATH/$t
  zip -9 -r $DIST_PATH/$t-$BUILD_VERSION.zip .
done
