#!/usr/bin/env bash

if [ -z "$RUN_NUMBER" ]; then
  echo "RUN_NUMBER is not set"
  RUN_NUMBER=0
fi

echo "Run test"
go test -v -short -timeout 30s -covermode=atomic -coverprofile=coverage.txt ./... > test/out.txt 2> test/err.txt

echo "Collect logs"
mkdir -p test/logs
zip -9 -r test/logs/logs-$RUN_NUMBER.zip test/out.txt test/err.txt

echo "Upload logs"
go run tbx.go dev ci artifact up -local-path test/logs/logs-$RUN_NUMBER.zip -dropbox-path "/watermint-toolbox-build/test-logs/$RUN_NUMBER"

if [ $? -ne 0 ]; then
  echo "Failed to upload logs: $?"
  go run tbx.go job log last -kind capture -quiet
  exit 1
fi
