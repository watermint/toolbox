---
layout: home
title: watermint toolbox
lang: en
---

# watermint toolbox

![watermint toolbox]({{ site.baseurl }}/images/logo.png){: width="160" }

The multi-purpose utility command-line tool for web services including Dropbox, Dropbox Business, Google, GitHub, etc.

# Do more with the watermint toolbox

The watermint toolbox has 191 commands to solve your daily tasks. For example, if you are an admin of Dropbox Business and need to manage a group. You can bulk create groups or bulk add members to groups via group commands.

![Demo]({{ site.baseurl }}/images/demo.gif)

The watermint toolbox runs on Windows, macOS (Darwin), and Linux without any additional libraries. You can run a command immediately after you download and extract the binary.

Please see command references for more detail.

| Reference                                                                    |
|------------------------------------------------------------------------------|
| [Commands]({{ site.baseurl }}/commands/toc.html)                             |
| [Dropbox Business commands]({{ site.baseurl }}/guides/dropbox-business.html) |

# Built executable

Pre-compiled binaries can be found in [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are building directly from the source, please refer [BUILD.md](BUILD.md).

## Installing using Homebrew on macOS

First, you need to install Homebrew. Please refer the instruction on [the official site](https://brew.sh/). Then run following commands to install watermint toolbox.
```
brew tap watermint/toolbox
brew install toolbox
```

# Licensing & Disclaimers

watermint toolbox is licensed under the MIT license.
Please see LICENSE.md or LICENSE.txt for more detail.

Please carefully note:
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

# Security and privacy

## Information Not Collected 

The watermint toolbox does not collect any information to third-party servers.

The watermint toolbox is for interacting with the services like Dropbox with your account. There is no third-party account involved. The Commands stores API token, logs, files, or reports on your PC storage.

## Sensitive data

Most sensitive data, such as API token, are saved on your PC storage in obfuscated & made restricted access. However, it's your responsibility to keep those data secret. 
Please do not share files, especially the `secrets` folder under toolbox work path (`C:\Users\<your user name>\.toolbox`, or `$HOME/.toolbox` by default).


