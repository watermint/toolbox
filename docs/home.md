# watermint toolbox

![watermint toolbox]({{ site.baseurl }}/images/logo.png){: width="160" }

The multi-purpose utility command-line tool for web services including Dropbox, Figma, Google, GitHub, etc.

# Do more with the watermint toolbox

The watermint toolbox has 304 commands to solve your daily tasks. For example, if you are an admin of Dropbox for teams and need to manage a group. You can bulk create groups or bulk add members to groups via group commands.

![Demo]({{ site.baseurl }}/images/demo.gif)

The watermint toolbox runs on Windows, macOS (Darwin), and Linux without any additional libraries. You can run a command immediately after you download and extract the binary.

Please see command references for more detail.

| Reference                                                                        |
|----------------------------------------------------------------------------------|
| [Commands]({{ site.baseurl }}/commands/toc.html)                                 |
| [Commands of Dropbox for teams]({{ site.baseurl }}/guides/dropbox-business.html) |

# Built executable

Pre-compiled binaries can be found in [Latest Release](https://github.com/watermint/toolbox/releases/latest). If you are building directly from the source, please refer [BUILD.md](BUILD.md).

## Installing using Homebrew on macOS/Linux

First, you need to install Homebrew. Please refer the instruction on [the official site](https://brew.sh/). Then run following commands to install watermint toolbox.
```
brew tap watermint/toolbox
brew install toolbox
```

# Licensing & Disclaimers

watermint toolbox is licensed under the Apache License, Version 2.0.
Please see LICENSE.md or LICENSE.txt for more detail.

Please carefully note:
> Unless required by applicable law or agreed to in writing, Licensor provides the Work (and each Contributor provides its Contributions) on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.

# Announcements

* [#813 License change : MIT License to Apache License, Version 2.0](https://github.com/watermint/toolbox/discussions/813)
* [#815 Lifecycle: Availability period for each release](https://github.com/watermint/toolbox/discussions/815)
* [#793 Google commands require re-authentication on Release 130](https://github.com/watermint/toolbox/discussions/793)
* [#799 Path change: Dropbox and Dropbox for teams commands have been  moved to under `dropbox`](https://github.com/watermint/toolbox/discussions/799)
* [#797 Path change: commands under `services` have been moved to a new location](https://github.com/watermint/toolbox/discussions/797)
* [#796 Deprecation: Dropbox Team space commands will be removed](https://github.com/watermint/toolbox/discussions/796)

# Security and privacy

## Information Not Collected 

The watermint toolbox does not collect any information to third-party servers.

The watermint toolbox is for interacting with the services like Dropbox with your account. There is no third-party account involved. The Commands stores API token, logs, files, or reports on your PC storage.

## Sensitive data

Most sensitive data, such as API token, are saved on your PC storage in obfuscated & made restricted access. However, it's your responsibility to keep those data secret. 
Please do not share files, especially the `secrets` folder under toolbox work path (`C:\Users\<your user name>\.toolbox`, or `$HOME/.toolbox` by default).

## Internet access other than the service that is the subject of the command used

The watermint toolbox has the ability to deactivate certain releases that have critical bugs or security issues. It does this by retrieving data from a repository hosted on GitHub about once every 30 days to check the status of a release.
This access does not collect your private data (such as your Dropbox, Google, local files, token, etc). It just checks the release status, but as a side effect your IP address is sent to GitHub when downloading data. I know IP address is also a PII. But this is the same as visiting a general website and is not a special operation.
The watermint toolbox project administrator cannot even see how many times these files have been downloaded.

