#!/usr/bin/env bash

TEST_ARGS=""
if [ x"$TOOLBOX_TEST_RESOURCE" != x"" ]; then
  TEST_ARGS="-test-resource $TOOLBOX_TEST_RESOURCE"
  echo PUBLISH : Release proceeding with the test resource: $TOOLBOX_TEST_RESOURCE
else
  echo PUBLISH : Abort: Test resource is not found
  exit 1
fi

mkdir -p publish/secrets

echo PUBLISH : Import test tokens from default location
TOOLBOX_ENDTOEND=$(go run tbx.go dev ci auth export -quiet) go run tbx.go dev ci auth import -workspace $PWD/publish

echo PUBLISH : Prepare for publish
LC_ALL=C ./tbx dev release publish -workspace $PWD/publish -artifact-path $PWD $TEST_ARGS
