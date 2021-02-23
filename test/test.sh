#!/usr/bin/env bash

TEST_RESULTS=test/results
TEST_DEBUG=test/debug
TEST_SUBSET=$TEST_DEBUG/subset.txt
TEST_OUT=$TEST_DEBUG/all.out
TEST_ERR=$TEST_DEBUG/err.out
TEST_REPORT=$TEST_RESULTS/all.xml
TEST_PROFILE=coverage.txt
LAUNCHABLE_BUILD_NAME=watermint

mkdir -p resources/keys
mkdir -p $TEST_RESULTS
mkdir -p $TEST_DEBUG

if [ "$LAUNCHABLE_TOKEN"x != x"" ]; then
  echo TEST: Preparing launchable resources
  launchable verify
  launchable record build --name $LAUNCHABLE_BUILD_NAME
  go test -list ./... | launchable subset --build $LAUNCHABLE_BUILD_NAME --target 10% go-test > $TEST_SUBSET
fi

echo TEST: Run tests $(cat $TEST_SUBSET)
if [ "$(cat test/debug/subset.txt)"x = ""x ]; then
  go test -v -covermode=atomic -coverprofile=$TEST_PROFILE ./... > "$TEST_OUT" 2> "$TEST_ERR"
  TEST_EXIT_CODE=$?
else
  go test -v -covermode=atomic -coverprofile=$TEST_PROFILE -run $(cat $TEST_SUBSET) ./... > "$TEST_OUT" 2> "$TEST_ERR"
  TEST_EXIT_CODE=$?
fi

echo TEST: Generate JUnit style report
hash go-junit-report 2>/dev/null

if [ "$?" -eq "0" ]; then
  cat $TEST_OUT | go-junit-report >$TEST_REPORT
  cp $TEST_REPORT $PWD/result.xml
  cp $TEST_REPORT $TEST_DEBUG/result.xml

  if [ "$LAUNCHABLE_TOKEN"x != x"" ]; then
    launchable record tests --build $LAUNCHABLE_BUILD_NAME go-test $PWD
  fi

  if [ "$CIRCLE_BUILD_NUM"x != ""x ]; then
    echo TEST: Uploading logs
    go run tbx.go dev ci artifact up -budget-memory low -local-path $TEST_DEBUG -dropbox-path /watermint-toolbox-build/test-logs/$CIRCLE_BUILD_NUM -peer-name deploy -quiet
    go run tbx.go dev ci artifact up -budget-memory low -local-path $HOME/.toolbox/jobs -dropbox-path /watermint-toolbox-build/test-logs/$CIRCLE_BUILD_NUM -peer-name deploy -quiet
  fi
fi

if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo Test failed: $TEST_EXIT_CODE
  exit 1
fi

echo TEST: Finished
