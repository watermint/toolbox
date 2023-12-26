#!/usr/bin/env bash

if [ -z "$RUN_NUMBER" ]; then
  echo "RUN_NUMBER is not set"
  RUN_NUMBER=0
fi

echo "Prepare toolbox"
mkdir -p build
curl -L -o build/tbx.url https://raw.githubusercontent.com/watermint/homebrew-toolbox/master/latest/linux-intel.url
curl -L -o build/tbx.zip "$(cat build/tbx.url)"
unzip build/tbx.zip -d build
if [ ! -e build/tbx ]; then
  echo "Failed to download toolbox"
  exit 1
fi

echo "Determine GOPATH"
if [ -z "$GOPATH" ]; then
  if [ -z "$GOROOT" ]; then
    echo "Use home as GOPATH($HOME/go)"
    GOPATH=$HOME/go
  else
    echo "Use GOROOT($GOROOT) as GOPATH($GOPATH)"
    GOPATH=$GOROOT
  fi
fi

echo "Run test"
go test -v -short -covermode=atomic -coverprofile=coverage.txt ./... > test/out.txt 2> test/err.txt

echo "Collect logs"
mkdir -p test/logs
zip -9 -r test/logs/logs-$RUN_NUMBER.zip test/out.txt test/err.txt

echo "Upload logs"
if [ x"$TOOLBOX_DEPLOY_DB" != x"" ]; then
  echo "Prepare deploy db"
  mkdir -p build/db
  echo "$TOOLBOX_DEPLOY_DB" | base64 -d > build/db/deploy.db.gz
  gzip -d build/db/deploy.db.gz

  build/tbx file sync up -auth-database build/db/deploy.db -local-path test/logs/logs-$RUN_NUMBER.zip -dropbox-path "/watermint-toolbox-build/test-logs/$RUN_NUMBER" -experiment suppress_progress

  if [ $? -ne 0 ]; then
    echo "Failed to upload logs: $?"
    go run tbx.go job log last -kind capture -quiet
    exit 1
  fi
fi

