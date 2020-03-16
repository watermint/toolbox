#!/usr/bin/env bash

TEST_ARGS=""
if [ x"$TOOLBOX_TEST_RESOURCE" != x"" ]; then
  TEST_ARGS="-test-resource $TOOLBOX_TEST_RESOURCE"
  echo Release proceeding with the test resource: $TOOLBOX_TEST_RESOURCE
else
  echo Release proceeding without test resources
fi

mkdir -p publish/secrets
cp $HOME/.toolbox/secrets/end_to_end_test.tokens publish/secrets

cp $HOME/.toolbox/release/test_resource.json     publish/test_resource.json

./tbx dev release publish -workspace $PWD/publish -artifact-path $PWD
