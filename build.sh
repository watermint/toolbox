#!/usr/bin/env bash

if [ x"" = x"$PROJECT_ROOT" ]; then
  # Configure for regular build
  PROJECT_ROOT=$PWD
  BUILD_PATH=$PWD/build
  DIST_PATH=$PWD/dist
else
  # Configure for Docker build
  BUILD_PATH=/build
  DIST_PATH=/dist
fi

BUILD_MAJOR_VERSION=$(cat $PROJECT_ROOT/version)
BUILD_HASH=$(cd $PROJECT_ROOT && git rev-parse HEAD)

if [ ! -d $BUILD_PATH ]; then
  mkdir -p $BUILD_PATH
fi
if [ ! -d $DIST_PATH ]; then
  mkdir -p $DIST_PATH
fi
if [ "$TOOLBOX_BUILD_ID"x = ""x ]; then
  TOOLBOX_BUILD_ID=0.0
fi
BUILD_VERSION=$BUILD_MAJOR_VERSION.$TOOLBOX_BUILD_ID

echo --------------------
echo BUILD: Start building version: $BUILD_VERSION

echo --------------------
echo BUILD: Testing..

cd $PROJECT_ROOT
rice embed-go
go test $(glide novendor)


echo --------------------
echo BUILD: Building tool

if [ -e "resources/toolbox.appkeys" ]; then
  echo App keys file found. Verify app key file...
  cat resources/toolbox.appkeys | jq type
  if [[ $? = 0 ]]; then
    echo valid
  else
    echo invalid. return code: $?
  fi
fi

X_APP_NAME="-X github.com/watermint/toolbox/app.AppName=toolbox"
X_APP_VERSION="-X github.com/watermint/toolbox/app.AppVersion=$BUILD_VERSION"
X_APP_HASH="-X github.com/watermint/toolbox/app.AppHash=$BUILD_HASH"
LD_FLAGS="$X_APP_NAME $X_APP_VERSION $X_APP_HASH"

echo Building: Windows
GOOS=windows GOARCH=386   go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/tbx-$BUILD_VERSION-win.exe github.com/watermint/toolbox
echo Building: Linux
GOOS=linux   GOARCH=386   go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/tbx-$BUILD_VERSION-linux   github.com/watermint/toolbox
echo Building: Darwin
GOOS=darwin  GOARCH=amd64 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/tbx-$BUILD_VERSION-macos   github.com/watermint/toolbox


echo --------------------
echo BUILD: Packaging
( cd $BUILD_PATH && zip -9 -r $DIST_PATH/tbx-$BUILD_VERSION.zip tbx-* )
