# toolbox

[![Build Status](https://travis-ci.org/watermint/toolbox.svg?branch=master)](https://travis-ci.org/watermint/toolbox)
[![Coverage Status](https://coveralls.io/repos/github/watermint/toolbox/badge.svg?branch=master)](https://coveralls.io/github/watermint/toolbox?branch=master)

Tools for Dropbox and Dropbox Business

# Usage

`tbx` have various features. Run without an option for a list of supported commands and options.
Released package contains binaries for each operating system. Each binary are named like `tbx-[version]-[os]-[system]`. E.g. if the binary is for macOS, then the name will like `tbx-32.1.0.0-darwin-10.6-amd64`.

```bash
% ./tbx-32.1.0.0-darwin-10.6-amd64

Usage: ./tbx-32.1.0.0-darwin-10.6-amd64 COMMAND

Available commands:
  file       File operation


Run './tbx-32.1.0.0-darwin-10.6-amd64 COMMAND' for more information on a command.

please specify sub command
```

# Build

## Build script

```bash
$ docker build -t toolbox . && rm -fr /tmp/dist && docker run -v /tmp/dist:/dist:rw --rm toolbox
```
