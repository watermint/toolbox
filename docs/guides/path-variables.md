---
layout: page
title: Path variables
---

# Path variables

Path variables are predefined variables which will be replaced on runtime. For example, if you specify a path with the variable like `{% raw %}{{{% endraw %}.DropboxPersonal{% raw %}}}{% endraw %}/Pictures`, then the path will be replaced with actual path to Personal Dropbox's folder. But the tool does not guarantee the existence or accuracy.

| Path variable                  | Description                                                                                |
|--------------------------------|--------------------------------------------------------------------------------------------|
| {% raw %}{{{% endraw %}.DropboxPersonal{% raw %}}}{% endraw %}           | Path to Dropbox Personal account root folder.                                              |
| {% raw %}{{{% endraw %}.DropboxBusiness{% raw %}}}{% endraw %}           | Path to Dropbox Business account root folder.                                              |
| {% raw %}{{{% endraw %}.DropboxBusinessOrPersonal{% raw %}}}{% endraw %} | Path to Dropbox Business account root folder, or Personal Dropbox account if it not found. |
| {% raw %}{{{% endraw %}.DropboxPersonalOrBusiness{% raw %}}}{% endraw %} | Path to Dropbox Personal account root folder, or Business Dropbox account if it not found. |
| {% raw %}{{{% endraw %}.Home{% raw %}}}{% endraw %}                      | The home folder of the current user.                                                       |
| {% raw %}{{{% endraw %}.Username{% raw %}}}{% endraw %}                  | The name of the current user.                                                              |
| {% raw %}{{{% endraw %}.Hostname{% raw %}}}{% endraw %}                  | The host name of the current computer.                                                     |
| {% raw %}{{{% endraw %}.ExecPath{% raw %}}}{% endraw %}                  | Path to this program.                                                                      |
| {% raw %}{{{% endraw %}.Rand8{% raw %}}}{% endraw %}                     | Randomized 8 digit number leading with 0.                                                  |
| {% raw %}{{{% endraw %}.Year{% raw %}}}{% endraw %}                      | Current local year with format 'yyyy' like 2021.                                           |
| {% raw %}{{{% endraw %}.Month{% raw %}}}{% endraw %}                     | Current local month with format 'mm' like 01.                                              |
| {% raw %}{{{% endraw %}.Day{% raw %}}}{% endraw %}                       | Current local day with format 'dd' like 05.                                                |
| {% raw %}{{{% endraw %}.Date{% raw %}}}{% endraw %}                      | Current local date with format yyyy-mm-dd.                                                 |
| {% raw %}{{{% endraw %}.Time{% raw %}}}{% endraw %}                      | Current local time with format HH-MM-SS.                                                   |
| {% raw %}{{{% endraw %}.DateUTC{% raw %}}}{% endraw %}                   | Current UTC date with format yyyy-mm-dd.                                                   |
| {% raw %}{{{% endraw %}.TimeUTC{% raw %}}}{% endraw %}                   | Current UTC time with format HH-MM-SS.                                                     |


