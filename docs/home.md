# watermint toolbox

![watermint toolbox]({{ site.baseurl }}/images/logo.png){: width="160" }

The multi-purpose utility command-line tool for web services including Dropbox, Figma, Google, GitHub, etc.

# Do more with the watermint toolbox

The watermint toolbox has 281 commands to solve your daily tasks. For example, if you are an admin of Dropbox for teams and need to manage a group. You can bulk create groups or bulk add members to groups via group commands.

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

* [#836 Remove binaries that are more than six months old after release](https://github.com/watermint/toolbox/discussions/836)
* [#835 Google commands deprecation](https://github.com/watermint/toolbox/discussions/835)
* [#813 License change : MIT License to Apache License, Version 2.0](https://github.com/watermint/toolbox/discussions/813)
* [#815 Lifecycle: Availability period for each release](https://github.com/watermint/toolbox/discussions/815)
* [#799 Path change: Dropbox and Dropbox for teams commands have been  moved to under `dropbox`](https://github.com/watermint/toolbox/discussions/799)
* [#797 Path change: commands under `services` have been moved to a new location](https://github.com/watermint/toolbox/discussions/797)
* [#796 Deprecation: Dropbox Team space commands will be removed](https://github.com/watermint/toolbox/discussions/796)

# Security and privacy

The watermint toolbox is designed to simplify the use of cloud service APIs. It will not use the data in any way that is contrary to your intentions.

The watermint toolbox does not store the data it retrieves via the linked cloud service API on a separate server, contrary to the intent of the specified command.

For example, if you use the watermint toolbox to retrieve data from a cloud service, those data will only be stored on your PC. Furthermore, in the case of commands that upload files or data to a cloud service, they will only be stored in the location specified by your account.

## Data protection

When you use the watermint toolbox to retrieve data from the cloud service API, your data is stored on your PC as report data or log data. More sensitive information, such as the authentication token for the cloud service API, is also stored on your PC.

It is your responsibility to keep these data stored on your PC secure.

Important information such as authentication tokens are obfuscated so that their contents cannot be easily read. However, this obfuscation is not intended to enhance security, but to prevent unintentional operational errors. If a malicious third party copies your token information to another PC, they may be able to access cloud services that you did not intend.

## Use

As previously stated, the watermint toolbox is designed to store data on your PC or in your cloud account. Processes other than your intended operation include data retrieval for release lifecycle management, as outlined below.

The watermint toolbox has the capability to deactivate specific releases that have critical bugs or security issues. This is achieved by retrieving data from a repository hosted on GitHub approximately every 30 days to assess the status of a release. This access does not collect any private data (such as your cloud account information, local files, token, etc.). It merely checks the release status, but as a side effect, your IP address is sent to GitHub when downloading data.

Please be aware that this access information (date, time and IP address) may be used in the future to estimate the usage of each release.

## Sharing

The watermint toolbox project does not currently manage or obtain data including IP addresses, information that only GitHub, the company that hosts the project, has the possibility to access. However, the project may in the future make this information available, and may disclose anonymised release-by-release usage to project members if deemed necessary for the operation of the project.

Any such changes will be announced on the announcement page and this security & privacy policy page at least 30 days before the change takes effect.

