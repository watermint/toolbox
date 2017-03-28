# toolbox

[![Build Status](https://travis-ci.org/watermint/toolbox.svg?branch=master)](https://travis-ci.org/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg?branch=master)](https://coveralls.io/github/watermint/toolbox?branch=master)

Tools for Dropbox and Dropbox Business

# Build

## Credentials

* Copy `credentials.sample` with name `credentials.secret`.
* Update `ApiKey` and `ApiSecret` for your Application ID.
* Build entire project using Dockerfile on top of the project.

## Build script

```sh
$ docker build -t toolbox . && rm -fr /tmp/dist && docker run -v /tmp/dist:/dist:rw --rm toolbox
```

# Available Tools

* [dcmp](tools/dcmp) ... file compare utility.
* [dupload](tools/dupload/) ... file/directory uploader.
* [dsharedlink](tools/dsharedlink) ... Shared link utility.
* [dteammember](tools/dteammember) ... Team member management.
