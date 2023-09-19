#!/usr/bin/env bash

go test -v -short -timeout 30s -covermode=atomic -coverprofile=coverage.txt ./... > test/out.txt 2> test/err.txt

mkdir -p test/logs
zip -9 -r test/logs/logs-${{ github.run_number }}.zip test/out.txt test/err.txt
go run tbx.go dev ci artifact up -local-path test/logs/logs-${{ github.run_number }}.zip -dropbox-path "/watermint-toolbox-build/test-logs/${{ github.run_number }}"

if [ $? -ne 0 ]; then
  echo "Failed to upload logs: $?"
  go run tbx.go job log last -kind capture -quiet
  exit 1
fi
