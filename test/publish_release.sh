#!/usr/bin/env bash

TEST_ARGS=""
if [ x"$TOOLBOX_TEST_RESOURCE" != x"" ]; then
  TEST_ARGS="-test-resource $TOOLBOX_TEST_RESOURCE"
  echo Release proceeding with the test resource: $TOOLBOX_TEST_RESOURCE
else
  echo Abort: Test resource is not found
  exit 1
fi

mkdir -p publish/secrets

LC_ALL=C ./tbx dev ci auth export -quiet > /dev/null
LC_ALL=C ./tbx dev ci auth import -workspace $PWD/publish 
LC_ALL=C ./tbx dev release publish -workspace $PWD/publish -artifact-path $PWD $TEST_ARGS

