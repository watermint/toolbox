#!/usr/bin/env bash

if [ x"" = x"$PROJECT_ROOT" ]; then
  PROJECT_ROOT=$(dirname $0)
fi
cd "$PROJECT_ROOT"

OUT_RESULTS=test/results
OUT_TEST=$OUT_RESULTS/all.out
OUT_TEST_REPORT=$OUT_RESULTS/all.xml
OUT_PROFILE=$OUT_RESULTS/profile.out
OUT_COVERAGE=coverage.txt

mkdir -p resources/keys
mkdir -p $OUT_RESULTS

echo "" >$OUT_COVERAGE
echo "" >$OUT_TEST

go test -v -covermode=atomic -coverprofile=$OUT_PROFILE  ./... > "$OUT_TEST"

if [ "$?" -ne "0" ]; then
  echo Test failed: $?
  exit 1
fi

hash go-junit-report 2>/dev/null

if [ "$?" -eq "0" ]; then
  cat $OUT_TEST | go-junit-report >$OUT_TEST_REPORT
fi

