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

