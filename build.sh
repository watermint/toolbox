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

if [ x"" = x"$TOOLBOX_BUILDERKEY" ]; then
  if [ -e resources/toolbox.buildkey ]; then
    TOOLBOX_BUILDERKEY=$(cat resources/toolbox.buildkey)
  else
    TOOLBOX_BUILDERKEY="watermint-toolbox-default"
  fi
fi

BUILD_MAJOR_VERSION=$(cat "$PROJECT_ROOT"/version)
BUILD_HASH=$(cd "$PROJECT_ROOT" && git rev-parse HEAD)

if [ ! -d $BUILD_PATH ]; then
  mkdir -p $BUILD_PATH
  for p in win mac linux; do
    mkdir -p $BUILD_PATH/$p
  done
fi
if [ ! -d $DIST_PATH ]; then
  mkdir -p $DIST_PATH
fi
if [ "$TOOLBOX_BUILD_ID"x = ""x ]; then
  # Circle CI
  if [ "$CIRCLE_BUILD_NUM"x != ""x ]; then
    if [ "$CIRCLE_BRANCH"x = "master"x ]; then
      TOOLBOX_BUILD_ID=4.$CIRCLE_BUILD_NUM
    else
      TOOLBOX_BUILD_ID=2.$CIRCLE_BUILD_NUM
    fi

  # Gitlab
  elif [ "$CI_PIPELINE_IID" ]; then
    TOOLBOX_BUILD_ID=1.$CI_PIPELINE_IID
  else
    TOOLBOX_BUILD_ID=0.0
  fi
fi
BUILD_VERSION=$BUILD_MAJOR_VERSION.$TOOLBOX_BUILD_ID

echo --------------------
echo BUILD: Start building version: $BUILD_VERSION

cd "$PROJECT_ROOT"
echo BUILD: Preparing license information
for m in $(go list -m all | awk '{print $1}'); do
  d=$(go list -json $m | jq -r .Module.Dir)
  if [ x"" != x"$d" ]; then
    l=$(find "$d" -maxdepth 1 -iname LICENSE\*)
    if [ x"" != x"$l" ]; then
      p=$(echo $m | sed 's/\//-/g')
      jq -Rn "{\"$m\":[inputs]}" "$l" > $BUILD_PATH/$p.lic
    fi
  fi
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

  go run infra/security/sc_zap_tool/main.go
  if [[ $? = 0 ]]; then
    rm resources/toolbox.appkeys
  else
    echo Zap exit with code $?
    exit $?
  fi
  TOOLBOX_ZAP=$(cat /tmp/toolbox.zap)
else
  echo ERR: No app key file found
  exit 1
fi
rice embed-go

X_APP_NAME="-X github.com/watermint/toolbox/infra/app.Name=toolbox"
X_APP_VERSION="-X github.com/watermint/toolbox/infra/app.Version=$BUILD_VERSION"
X_APP_HASH="-X github.com/watermint/toolbox/infra/app.Hash=$BUILD_HASH"
X_APP_ZAP="-X github.com/watermint/toolbox/infra/app.Zap=$TOOLBOX_ZAP"
X_APP_BUILDERKEY="-X github.com/watermint/toolbox/infra/app.BuilderKey=$TOOLBOX_BUILDERKEY"
LD_FLAGS="$X_APP_NAME $X_APP_VERSION $X_APP_HASH $X_APP_ZAP $X_APP_BUILDERKEY"

echo Building: Windows 386
CGO_ENABLED=0 GOOS=windows GOARCH=386   go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/win/tbx.exe github.com/watermint/toolbox
echo Building: Windows amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/win/tbx64.exe github.com/watermint/toolbox
echo Building: Linux
CGO_ENABLED=0 GOOS=linux   GOARCH=386   go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/linux/tbx   github.com/watermint/toolbox
echo Building: Darwin
CGO_ENABLED=0 GOOS=darwin  GOARCH=amd64 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/mac/tbx     github.com/watermint/toolbox

echo --------------------
echo BUILD: Testing binary

$BUILD_PATH/linux/tbx dev test resources -quiet
if [[ $? = 0 ]]; then
  echo Success: resources test
else
  echo Unable to pass binary resources test: code=$?
  exit $?
fi

echo --------------------
echo BUILD: Generating documents

$BUILD_PATH/linux/tbx license -quiet                            > $BUILD_PATH/LICENSE.txt
$BUILD_PATH/linux/tbx dev doc -filename README.txt -badge=false > $BUILD_PATH/README.txt

echo --------------------
echo BUILD: Packaging
for p in win mac linux; do
  echo BUILD: Packaging $p
  cp $BUILD_PATH/LICENSE.txt $BUILD_PATH/"$p"/
  cp README.txt  $BUILD_PATH/"$p"/README.txt
  ( cd $BUILD_PATH/"$p" && zip -9 -r $BUILD_PATH/tbx-"$BUILD_VERSION"-"$p".zip . )
done
( cd $BUILD_PATH && zip -0 $DIST_PATH/tbx-"$BUILD_VERSION".zip *.zip )

if [ ""x != "$CIRCLE_BUILD_URL"x ]; then
  echo Artifact: "$CIRCLE_BUILD_URL"/"$CIRCLE_NODE_INDEX"/$DIST_PATH/tbx-"$BUILD_VERSION".zip
fi
