---
layout: page
title: Path variables
---

# Path variables

Path variables are predefined variables which will be replaced on runtime. For example, if you specify a path with the variable like `{% raw %}{{.{% endraw %}DropboxPersonal}}/Pictures`, then the path will be replaced with actual path to Personal Dropbox's folder. But the tool does not guarantee the existence or accuracy.

| Path variable                  | Description                                                                                |
|--------------------------------|--------------------------------------------------------------------------------------------|
| {% raw %}{{.{% endraw %}DropboxPersonal}}           | Path to Dropbox Personal account root folder.                                              |
| {% raw %}{{.{% endraw %}DropboxBusiness}}           | Path to Dropbox Business account root folder.                                              |
| {% raw %}{{.{% endraw %}DropboxBusinessOrPersonal}} | Path to Dropbox Business account root folder, or Personal Dropbox account if it not found. |
| {% raw %}{{.{% endraw %}DropboxPersonalOrBusiness}} | Path to Dropbox Personal account root folder, or Business Dropbox account if it not found. |
| {% raw %}{{.{% endraw %}Home}}                      | The home folder of the current user.                                                       |
| {% raw %}{{.{% endraw %}Username}}                  | The name of the current user.                                                              |
| {% raw %}{{.{% endraw %}Hostname}}                  | The host name of the current computer.                                                     |
| {% raw %}{{.{% endraw %}ExecPath}}                  | Path to this program.                                                                      |
| {% raw %}{{.{% endraw %}Rand8}}                     | Randomized 8 digit number leading with 0.                                                  |
| {% raw %}{{.{% endraw %}Year}}                      | Current local year with format 'yyyy' like 2021.                                           |
| {% raw %}{{.{% endraw %}Month}}                     | Current local month with format 'mm' like 01.                                              |
| {% raw %}{{.{% endraw %}Day}}                       | Current local day with format 'dd' like 05.                                                |
| {% raw %}{{.{% endraw %}Date}}                      | Current local date with format yyyy-mm-dd.                                                 |
| {% raw %}{{.{% endraw %}Time}}                      | Current local time with format HH-MM-SS.                                                   |
| {% raw %}{{.{% endraw %}DateUTC}}                   | Current UTC date with format yyyy-mm-dd.                                                   |
| {% raw %}{{.{% endraw %}TimeUTC}}                   | Current UTC time with format HH-MM-SS.                                                     |


