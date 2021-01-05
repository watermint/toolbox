# Path variables

Path variables are predefined variables which will be replaced on runtime. For example, if you specify a path with the variable like `{{.DropboxPersonal}}/Pictures`, then the path will be replaced with actual path to Personal Dropbox's folder. But the tool does not guarantee the existence or accuracy.

| Path variable                  | Description                                                                                |
|--------------------------------|--------------------------------------------------------------------------------------------|
| {{.DropboxPersonal}}           | Path to Dropbox Personal account root folder.                                              |
| {{.DropboxBusiness}}           | Path to Dropbox Business account root folder.                                              |
| {{.DropboxBusinessOrPersonal}} | Path to Dropbox Business account root folder, or Personal Dropbox account if it not found. |
| {{.DropboxPersonalOrBusiness}} | Path to Dropbox Personal account root folder, or Business Dropbox account if it not found. |
| {{.Home}}                      | The home folder of the current user.                                                       |
| {{.Username}}                  | The name of the current user.                                                              |
| {{.Hostname}}                  | The host name of the current computer.                                                     |
| {{.ExecPath}}                  | Path to this program.                                                                      |
| {{.Rand8}}                     | Randomized 8 digit number leading with 0.                                                  |
| {{.Year}}                      | Current local year with format 'yyyy' like 2021.                                           |
| {{.Month}}                     | Current local month with format 'mm' like 01.                                              |
| {{.Day}}                       | Current local day with format 'dd' like 05.                                                |
| {{.Date}}                      | Current local date with format yyyy-mm-dd.                                                 |
| {{.Time}}                      | Current local time with format HH-MM-SS.                                                   |
| {{.DateUTC}}                   | Current UTC date with format yyyy-mm-dd.                                                   |
| {{.TimeUTC}}                   | Current UTC time with format HH-MM-SS.                                                     |

