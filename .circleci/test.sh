#!/usr/bin/env bash

set -e
echo "" > coverage.txt
echo "" > testreport.txt

for d in $(go list ./... | grep -v vendor); do
  echo Testing: $d
  CGO_ENABLED=0 go test -short -v -coverprofile=profile.out -covermode=atomic $d 2>&1 | tee test.out
  if [ -f profile.out ]; then
    cat profile.out >> coverage.txt
    rm profile.out
  fi
  if [ -f test.out ]; then
    cat test.out >> testreport.txt
    rm test.out
  fi
done
