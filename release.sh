#!/usr/bin/env bash
# Prepare for release

if [ x"master" != x"$CIRCLE_BRANCH" ]; then
  echo Skip release testing: current branch "$CIRCLE_BRANCH"
  exit 0
fi

if [ x"" == x"$TOOLBOX_TESTRESOURCE" ]; then
  echo Skip testing with supplemental test resource
  exit 0
fi
