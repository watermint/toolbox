#!/usr/bin/env bash

mkdir -p publish/secrets
cp $HOME/.toolbox/secrets/end_to_end_test.tokens publish/secrets
cp $HOME/.toolbox/release/test_resource.json     publish/test_resource.json

./tbx dev release publish -workspace $PWD/publish -artifact-path $PWD -test-resource $PWD/publish/test_resource.json
