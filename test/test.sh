#!/usr/bin/env bash

TEST_RESULTS=test/results
TEST_DEBUG=test/debug
TEST_OUT=$TEST_DEBUG/all.out
TEST_ERR=$TEST_DEBUG/err.out
TEST_REPORT=$TEST_RESULTS/all.xml
TEST_PROFILE=coverage.txt

mkdir -p resources/keys
mkdir -p $TEST_RESULTS
mkdir -p $TEST_DEBUG

echo TEST: Run tests
go test -v -timeout 20s -covermode=atomic -coverprofile=$TEST_PROFILE ./... >"$TEST_OUT" 2>"$TEST_ERR"
TEST_EXIT_CODE=$?

echo TEST: Generate JUnit style report
hash go-junit-report 2>/dev/null

if [ "$?" -eq "0" ]; then
  cat $TEST_OUT | go-junit-report >$TEST_REPORT
  cp $TEST_REPORT $PWD/result.xml
  cp $TEST_REPORT $TEST_DEBUG/result.xml

  if [ "$CIRCLE_BUILD_NUM"x != ""x ]; then
    if [ $TEST_EXIT_CODE -ne 0 ]; then
      echo TEST: Uploading logs
      go run tbx.go dev ci artifact up -budget-memory low -local-path $TEST_DEBUG -dropbox-path /watermint-toolbox-build/test-logs/$CIRCLE_BUILD_NUM -peer-name deploy -quiet
      go run tbx.go dev ci artifact up -budget-memory low -local-path $HOME/.toolbox/jobs -dropbox-path /watermint-toolbox-build/test-logs/$CIRCLE_BUILD_NUM -peer-name deploy -quiet
    fi
  fi
fi

if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo Test failed: $TEST_EXIT_CODE
  exit 1
fi

echo TEST: Finished
