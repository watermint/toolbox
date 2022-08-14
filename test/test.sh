#!/usr/bin/env bash

TEST_RESULTS=test/results
TEST_DEBUG=test/debug
TEST_OUT=$TEST_DEBUG/all.out
TEST_ERR=$TEST_DEBUG/err.out
TEST_PROFILE=coverage.txt

mkdir -p resources/keys
touch resources/keys/toolbox.build
mkdir -p $TEST_RESULTS
mkdir -p $TEST_DEBUG

echo TEST: Run tests
go test -v -short -timeout 30s -covermode=atomic -coverprofile=$TEST_PROFILE ./... >"$TEST_OUT" 2>"$TEST_ERR"
TEST_EXIT_CODE=$?

if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo Test failed: $TEST_EXIT_CODE
  if [ "$CIRCLE_BUILD_NUM"x != ""x ]; then
    echo TEST: Uploading logs
    go run tbx.go dev ci artifact up -budget-memory low -local-path $TEST_DEBUG -dropbox-path /watermint-toolbox-build/test-logs/$CIRCLE_BUILD_NUM -peer-name deploy -quiet
    go run tbx.go dev ci artifact up -budget-memory low -local-path $HOME/.toolbox/jobs -dropbox-path /watermint-toolbox-build/test-logs/$CIRCLE_BUILD_NUM -peer-name deploy -quiet
  fi
  exit 1
fi

echo TEST: Finished
