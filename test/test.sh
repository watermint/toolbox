#!/usr/bin/env bash

OUT_RESULTS=test/results
OUT_TEST=$OUT_RESULTS/all.out
OUT_TEST_REPORT=$OUT_RESULTS/all.xml
OUT_PROFILE=$OUT_RESULTS/profile.out
OUT_COVERAGE=coverage.txt

mkdir -p resources/keys
mkdir -p $OUT_RESULTS

echo "" >$OUT_COVERAGE
echo "" >$OUT_TEST

go test -v -covermode=atomic -coverprofile=$OUT_PROFILE  ./... 2>&1 > "$OUT_TEST"

if [ "$?" -ne "0" ]; then
  echo Test failed: $?
  exit 1
fi

hash go-junit-report 2>/dev/null

if [ "$?" -eq "0" ]; then
  cat $OUT_TEST | go-junit-report >$OUT_TEST_REPORT
fi

