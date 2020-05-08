#!/usr/bin/env bash

OUT_RESULTS=test/results
OUT_TEST=$OUT_RESULTS/last.out
OUT_TEST_ALL=$OUT_RESULTS/all.out
OUT_TEST_REPORT=$OUT_RESULTS/all.xml
OUT_PROFILE=$OUT_RESULTS/profile.out
OUT_COVERAGE=coverage.txt

echo "" >$OUT_COVERAGE
echo "" >$OUT_TEST_ALL

for d in $(go list ./... | grep -v vendor); do
  echo Testing: $d
  CGO_ENABLED=0 go test -short -v -coverprofile=$OUT_PROFILE -covermode=atomic $d 2>&1 >$OUT_TEST
  if [ "$?" -ne "0" ]; then
    echo Test failed: $?
    echo ---------------
    cat $OUT_TEST
    echo ---------------
    exit 1
  fi
  if [ -f $OUT_PROFILE ]; then
    cat $OUT_PROFILE >>$OUT_COVERAGE
    rm $OUT_PROFILE
  fi
  if [ -f $OUT_TEST ]; then
    cat $OUT_TEST >>$OUT_TEST_ALL
    rm $OUT_TEST
  fi
done

cat $OUT_TEST_ALL | go-junit-report >$OUT_TEST_REPORT
