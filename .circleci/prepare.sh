#!/usr/bin/env bash
# Prepare for release

mkdir -p $HOME/.toolbox/secrets
mkdir -p resources/keys
echo "$TOOLBOX_APPKEYS" > resources/keys/toolbox.appkeys

if [ x"" == x"$TOOLBOX_TESTRESOURCE_URL" ]; then
  echo Skip testing with supplemental test resource, test resource url missing
  exit 0
fi
if [ x"" == x"$TOOLBOX_TESTRESOURCE" ]; then
  echo Skip testing with supplemental test resource, test resource path missing
  exit 0
fi

curl -L "$TOOLBOX_TESTRESOURCE_URL" > "$TOOLBOX_TESTRESOURCE"
