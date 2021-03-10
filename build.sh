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

mkdir -p resources/data
mkdir -p resources/images
mkdir -p resources/keys
mkdir -p resources/messages
mkdir -p resources/templates
mkdir -p resources/web

BUILD_MAJOR_VERSION=$(cat "$PROJECT_ROOT"/version)
BUILD_HASH=$(cd "$PROJECT_ROOT" && git rev-parse HEAD)
BUILD_TIMESTAMP=$(date --iso-8601=seconds)

if [ ! -d $BUILD_PATH ]; then
  mkdir -p $BUILD_PATH
  for p in win win64 mac linux; do
    mkdir -p $BUILD_PATH/$p
  done
fi
if [ ! -d $DIST_PATH ]; then
  mkdir -p $DIST_PATH
fi
if [ "$TOOLBOX_BUILD_ID"x = ""x ]; then
  # Circle CI
  if [ "$CIRCLE_BUILD_NUM"x != ""x ]; then
    if [ "$CIRCLE_BRANCH"x = "main"x ]; then
      TOOLBOX_BUILD_ID=8.$CIRCLE_BUILD_NUM
    elif [ "$CIRCLE_BRANCH"x = "master"x ]; then
      TOOLBOX_BUILD_ID=4.$CIRCLE_BUILD_NUM
    else
      TOOLBOX_BUILD_ID=2.$CIRCLE_BUILD_NUM
    fi

  # GitHub Actions
  elif [ "$GITHUB_RUN_ID"x != ""x ]; then
    TOOLBOX_BUILD_ID=3.$GITHUB_RUN_ID

  # Gitlab
  elif [ "$CI_PIPELINE_IID" ]; then
    TOOLBOX_BUILD_ID=1.$CI_PIPELINE_IID

  # Docker
  else
    TOOLBOX_BUILD_ID=0.$(date +%Y%m%d%H%M%S)
  fi
fi
BUILD_VERSION=$BUILD_MAJOR_VERSION.$TOOLBOX_BUILD_ID

echo --------------------
echo BUILD: Start building version: $BUILD_VERSION

echo --------------------
echo BUILD: Preparing resources

mkdir -p resources/keys
if [ x"" != x"$TOOLBOX_APPKEYS" ]; then
  echo "$TOOLBOX_APPKEYS" >resources/keys/toolbox.appkeys
fi

if [ -e "resources/keys/toolbox.appkeys" ]; then
  echo App keys file found. Verify app key file...
  cat resources/keys/toolbox.appkeys | jq type >/dev/null
  if [[ $? == 0 ]]; then
    echo Valid
  else
    echo Invalid. return code: $?
  fi

  go run infra/security/sc_zap_tool/main.go
  if [[ $? == 0 ]]; then
    rm resources/keys/toolbox.appkeys
  else
    echo Zap exit with code $?
    exit $?
  fi
  TOOLBOX_ZAP=$(cat /tmp/toolbox.zap)
else
  echo ERR: No app key file found
  exit 1
fi

if [ x"" = x"$TOOLBOX_BUILDERKEY" ]; then
  if [ -e resources/keys/toolbox.buildkey ]; then
    TOOLBOX_BUILDERKEY=$(cat resources/keys/toolbox.buildkey)
  else
    TOOLBOX_BUILDERKEY="watermint-toolbox-default"
  fi
fi

echo --------------------
echo BUILD: License information

mkdir $BUILD_PATH/license
go-licenses csv github.com/watermint/toolbox 2>/dev/null >"$BUILD_PATH/license/licenses.csv"
go-licenses save github.com/watermint/toolbox --save_path "$BUILD_PATH/license/licenses" 2>/dev/null
go run tbx.go dev build license -quiet -source-path $BUILD_PATH/license -dest-path "$PROJECT_ROOT/resources/data/licenses.json"

echo BUILD: Building tool

rice embed-go

X_APP_VERSION="-X github.com/watermint/toolbox/infra/app.Version=$BUILD_VERSION"
X_APP_HASH="-X github.com/watermint/toolbox/infra/app.Hash=$BUILD_HASH"
X_APP_ZAP="-X github.com/watermint/toolbox/infra/app.Zap=$TOOLBOX_ZAP"
X_APP_BUILDERKEY="-X github.com/watermint/toolbox/infra/app.BuilderKey=$TOOLBOX_BUILDERKEY"
X_APP_BUILDTIMESTAMP="-X github.com/watermint/toolbox/infra/app.BuildTimestamp=$BUILD_TIMESTAMP"
X_APP_BRANCH="-X github.com/watermint/toolbox/infra/app.Branch=$CIRCLE_BRANCH"
LD_FLAGS="$X_APP_VERSION $X_APP_HASH $X_APP_ZAP $X_APP_BUILDERKEY $X_APP_BUILDTIMESTAMP $X_APP_BRANCH"

echo Building: Windows 386
CGO_ENABLED=0 GOOS=windows GOARCH=386 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/win/tbx.exe github.com/watermint/toolbox
echo Building: Windows amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/win64/tbx.exe github.com/watermint/toolbox
echo Building: Linux
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/linux/tbx github.com/watermint/toolbox
echo Building: Darwin
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build --ldflags "$LD_FLAGS" -o $BUILD_PATH/mac/tbx github.com/watermint/toolbox

echo --------------------
echo BUILD: Testing binary

$BUILD_PATH/linux/tbx dev test resources -quiet
if [[ $? == 0 ]]; then
  echo Success: resources test
else
  echo Unable to pass binary resources test: code=$?
  exit $?
fi

if [ ! -s $BUILD_PATH/win/tbx.exe ]; then
  echo Failed to build Windows binary
  exit 1
fi

if [ ! -s $BUILD_PATH/win64/tbx.exe ]; then
  echo Failed to build Windows x64 binary
  exit 1
fi

if [ ! -s $BUILD_PATH/linux/tbx ]; then
  echo Failed to build Linux binary
  exit 1
fi

if [ ! -s $BUILD_PATH/mac/tbx ]; then
  echo Failed to build macOS x64 binary
  exit 1
fi

echo --------------------
echo BUILD: Generating documents

$BUILD_PATH/linux/tbx license -output markdown >$BUILD_PATH/LICENSE.txt
$BUILD_PATH/linux/tbx dev build readme -path $BUILD_PATH/README.txt

if [ ! -s $BUILD_PATH/LICENSE.txt ]; then
  echo Failed to generate LICENSE
  exit 1
fi

if [ ! -s $BUILD_PATH/README.txt ]; then
  echo Failed to generate README
  exit 1
fi

echo --------------------
echo BUILD: Packaging
for p in win win64 mac linux; do
  echo BUILD: Packaging $p
  cp $BUILD_PATH/LICENSE.txt $BUILD_PATH/"$p"/LICENSE.txt
  cp $BUILD_PATH/README.txt $BUILD_PATH/"$p"/README.txt
  (cd $BUILD_PATH/"$p" && zip -9 -r $BUILD_PATH/tbx-"$BUILD_VERSION"-"$p".zip .)
done
echo BUILD: Packaging all
(cd $BUILD_PATH && zip -0 $DIST_PATH/tbx-"$BUILD_VERSION".zip *.zip)

if [ x"$TOOLBOX_DEPLOY_TOKEN" = x"" ]; then
  exit 0
fi

echo --------------------
echo BUILD: Deploying

cd $PROJECT_ROOT
$BUILD_PATH/linux/tbx dev ci artifact up -budget-memory low -dropbox-path /watermint-toolbox-build -local-path $DIST_PATH -peer-name deploy -quiet
