#!/usr/bin/env bash
# Prepare for release

mkdir -p $HOME/.toolbox/secrets
echo "$TOOLBOX_APPKEYS" > resources/toolbox.appkeys
echo "$TOOLBOX_ENDTOEND" > $HOME/.toolbox/secrets/end_to_end_test.tokens

#if [ x"master" != x"$CIRCLE_BRANCH" ]; then
#  echo Skip release testing: current branch "$CIRCLE_BRANCH"
#  exit 0
#fi

if [ x"" == x"$TOOLBOX_TESTRESOURCE_URL" ]; then
  echo Skip testing with supplemental test resource, test resource url missing
  exit 0
fi
if [ x"" == x"$TOOLBOX_TESTRESOURCE" ]; then
  echo Skip testing with supplemental test resource, test resource path missing
  exit 0
fi

curl -L "$TOOLBOX_TESTRESOURCE_URL" > "$TOOLBOX_TESTRESOURCE"
