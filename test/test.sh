#!/usr/bin/env bash

TEST_RESULTS=test/results
TEST_OUT=$TEST_RESULTS/all.out
TEST_ERR=$TEST_RESULTS/err.out
TEST_REPORT=$TEST_RESULTS/all.xml
TEST_PROFILE=coverage.txt

mkdir -p resources/keys
mkdir -p $TEST_RESULTS

go test -v -covermode=atomic -coverprofile=$TEST_PROFILE  ./... > "$TEST_OUT" 2> "$TEST_ERR"
TEST_EXIT_CODE=$?

if [ $TEST_EXIT_CODE -ne 0 ]; then
  echo Test failed: $TEST_EXIT_CODE
  exit 1
fi

hash go-junit-report 2>/dev/null

if [ "$?" -eq "0" ]; then
  cat $TEST_OUT | go-junit-report >$TEST_REPORT
fi
