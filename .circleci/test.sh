#!/usr/bin/env bash

set -e
echo "" >coverage.txt
echo "" >testreport.txt

for d in $(go list ./... | grep -v vendor); do
  echo Testing: $d
  CGO_ENABLED=0 go test -short -v -coverprofile=profile.out -covermode=atomic $d 2>&1 | tee test.out
  if [ "$?" -ne "0" ]; then
    echo Test failed: $?
    cat test.out
    exit $?
  fi
  if [ -f profile.out ]; then
    cat profile.out >>coverage.txt
    rm profile.out
  fi
  if [ -f test.out ]; then
    cat test.out >>testreport.txt
    rm test.out
  fi
done

cat testreport.txt | go-junit-report >test/results/all.xml
cp testreport.txt test/results/out.txt
