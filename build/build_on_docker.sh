#!/usr/bin/env bash

# Prerequisite: set absolute path to project root to $PROJECT_ROOT

BUILD_PATH=/tmp/out
DIST_PATH=/dist

TARGET_TOOLS="dtm dfm"
#TARGET_PLATFORM="windows/386,windows/amd64,darwin/amd64,linux/386,linux/amd64"
TARGET_PLATFORM="windows/386,darwin/amd64,linux/386"

BUILD_MAJOR_VERSION=$(cat $PROJECT_ROOT/version)
BUILD_HASH=$(cd $PROJECT_ROOT && git rev-parse HEAD)

if [ "$TOOLBOX_BUILD_ID"x = ""x ]; then
  TOOLBOX_BUILD_ID=0.0
fi
BUILD_VERSION=$BUILD_MAJOR_VERSION.$TOOLBOX_BUILD_ID

echo --------------------
echo BUILD: Start building version: $BUILD_VERSION

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
  if [ "$TOOLBOX_APP_CREDENTIALS"x != ""x ]; then
    X_APP_CREDENTIALS="$TOOLBOX_APP_CREDENTIALS"
  fi
  LD_FLAGS="$X_APP_NAME $X_APP_VERSION $X_APP_HASH $X_APP_CREDENTIALS"

  xgo --ldflags="$LD_FLAGS" -out $t-$BUILD_VERSION -targets $TARGET_PLATFORM github.com/watermint/toolbox/tools/$t

  ( cd /build && zip -9 -r $DIST_PATH/$t-$BUILD_VERSION.zip $t-* )
done
