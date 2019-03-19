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

cd $PROJECT_ROOT

echo BUILD: Preparing license information
for l in $(find vendor -name LICENSE\*); do
  pkg=$(dirname $l | sed 's/.*\/vendor\///')
  pf=$(echo $pkg | sed 's/\//-/g')
  jq -Rn "{\"$pkg\":[inputs]}" $l > $BUILD_PATH/$pf.lic
done
jq -Rn '{"github.com/watermint/toolbox":[inputs]}' LICENSE.md > $BUILD_PATH/github.com-watermint-toolbox.lic
jq -s add $BUILD_PATH/*.lic > resources/licenses.json


echo BUILD: Building tool

if [ -e "resources/toolbox.appkeys" ]; then
  echo App keys file found. Verify app key file...
  cat resources/toolbox.appkeys | jq type > /dev/null
  if [[ $? = 0 ]]; then
    echo Valid
  else
    echo Invalid. return code: $?
  fi

  go run app/app_zap/app_zap_tool/main.go
  if [[ $? = 0 ]]; then
    rm resources/toolbox.appkeys
  fi
  TOOLBOX_ZAP=$(cat /tmp/toolbox.zap)
fi
rice embed-go

X_APP_NAME="-X github.com/watermint/toolbox/app.AppName=toolbox"
X_APP_VERSION="-X github.com/watermint/toolbox/app.AppVersion=$BUILD_VERSION"
X_APP_HASH="-X github.com/watermint/toolbox/app.AppHash=$BUILD_HASH"
X_APP_ZAP="-X github.com/watermint/toolbox/app.AppZap=$TOOLBOX_ZAP"
LD_FLAGS="$X_APP_NAME $X_APP_VERSION $X_APP_HASH $X_APP_ZAP"

echo Building: Windows
GOOS=windows GOARCH=386   go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/tbx-$BUILD_VERSION-win.exe github.com/watermint/toolbox
echo Building: Linux
GOOS=linux   GOARCH=386   go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/tbx-$BUILD_VERSION-linux   github.com/watermint/toolbox
echo Building: Darwin
GOOS=darwin  GOARCH=amd64 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/tbx-$BUILD_VERSION-macos   github.com/watermint/toolbox


echo --------------------
echo BUILD: Packaging
( cd $BUILD_PATH && zip -9 -r $DIST_PATH/tbx-$BUILD_VERSION.zip tbx-* )
