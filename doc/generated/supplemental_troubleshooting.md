# Firewall or proxy server settings

The tool automatically detects proxy configuration from the system. However, that may fail or cause misconfiguration. In those cases, please use the `-proxy` option to specify proxy server hostname and port number like `-proxy 192.168.1.1:8080` (for proxy server 192.168.1.1, and the port number 8080). 

Note: This tool does not support proxy servers with any authentication such as Basic authentication or NTLM.

# Performance issue

If the command feels slow or stalled, please try re-run with an option `-verbose`. That will show more detailed progress. But in most cases, the cause is simply you have a larger data to process. Otherwise, you may hit a rate limit from API servers. If you want to see rate limit status, please see capture logs and debug for more details. 

The tool automatically adjusts concurrency to avoid additional limitation from API servers. If you want to see current concurrency, please run the command like below. That will show a current window size (maximum concurrency) per endpoint. The debug message "WaiterStatus" reports current concurrency and window sizes. The map "runners" is for operations currently waiting for a result from API servers. The map "window" is for window size for each endpoint. The map "concurrency" is for window sizes for current running operations. From the below example, for the endpoint "https://api.dropboxapi.com/2/file_requests/create" the tool does not allow call that endpoint grater than one concurrency. That means it requires operation one by one, and there is no easy workaround to speed up operations.
```
tbx job log last -quiet | jq 'select(.msg == "WaiterStatus")' 
{
  "level": "DEBUG",
  "time": "2020-11-10T14:55:57.501+0900",
  "name": "z951.z960.z112064",
  "caller": "nw_congestion/congestion.go:310",
  "msg": "WaiterStatus",
  "goroutine": "gr:284877",
  "runners": {
    "gr:1": {
      "key": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create",
      "go_routine_id": "gr:1",
      "running_since": "2020-11-10T14:55:56.124899+09:00"
    }
  },
  "numRunners": 1,
  "waiters": [],
  "numWaiters": 0,
  "window": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url/check_job_status": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search/continue_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/users/get_current_account": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_reference/get": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/delete_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/get_metadata": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/sharing/list_mountable_folders": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/download": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/export": 4
  },
  "concurrency": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 1
  }
}
```

# Garbled output

If the tool output garbled, please stop the tool with Ctrl+C. This issue usually happens when your console does not have a font to display it. Then, please try to change the font that supports your language. Or, please try the option `-lang en` to overwrite language setting of the tool to English.

In PowerShell, you can change the font with (1) right-click on the title bar, (2) click properties, (3) then choose the font tab, (4) then change to an appropriate font like "MS Gothic."

# Log files

By default, log files are stored under the path "%USERPROFILE%\.toolbox\jobs" (e.g. `C:\Users\USERNAME\.toolbox\jobs`) on windows, or "$HOME\.toolbox\jobs" in Linux or macOS (e.g. `/Users/USERNAME/.toolbox/jobs`). Log files contain information such as (1) Runtime information, e.g. OS type/version/environment variables, (2) Runtime options to the tool (including a copy of input data files), (3) Account information of services such as Dropbox, (4) Request and response data to API servers, (5) Data in services such as file name, metadata, id, URL etc. (depends on the command).

Those logs do not contain password, credentials, or API token. But API tokens are stored under the path "%USERPROFILE%\.toolbox\secrets" (e.g. `C:\Users\USERNAME\.toolbox\secrets`) on windows, or "$HOME\.toolbox\secrets" in Linux or macOS (e.g. `/Users/USERNAME/.toolbox/secrets`). These secrets folder files are obfuscated but please do not share these files to anyone including a service provider support such as Dropbox support.

## Log format

There are several folders and files stored under the `jobs` folder. First, the job folder will be created every run with a name (internally called Job Id) with the format "yyyyMMdd-HHmmSS.xxx". The first part "yyyyMMdd-HHmmSS" is for local date/time of the command start. The second part ".xxx" is the sequential or random three-character ID to avoid conflict with a concurrent run.

Under the job folder, there are subfolders (1) `logs`: runtime logs including request/response data, parameters, or debug information, (2) `reports`: reports folder is for manage generated reports, (3) `kvs`: KVS folder is for runtime database folder. 

On troubleshooting, files under `logs` are essential to understand what happened in runtime. The tool generates several types of logs. Those logs are JSON Lines format. Note: JSON Lines is a format that separate data with line separators. Please see [JSON Lines](https://jsonlines.org/) for more detail about the specification.

Some logs are compressed with gzip format. If a log compressed, then the file has a suffix '.gz'. Additionally, logs such as capture logs and toolbox logs are divided by certain size. If you want to analyze logs, please consider using `job log` commands. For example, `job log last -quiet` will report toolbox logs of the latest job with decompressed and concatenated.

## Debug logs

The tool will record all debug information into debug logs that have a prefix "toolbox". All records have a source code file name and line at the operation. If you find some suspicious error, then go to the source code and debug it. Some troubleshooting requires statistical analysis such as for performance tuning or out of memory. It is better to work with a tool such as `grep` or [jq](https://stedolan.github.io/jq/). 

If you want to see heap size data in time series, please run the command like below. Then you can see time + heap size in CSV format.
```
tbx job log last -quiet | jq -r 'select(.msg == "Heap stats") | [.time, .HeapInuse] | @csv'
"2020-11-10T14:55:45.725+0900",18604032
"2020-11-10T14:55:50.725+0900",15130624
"2020-11-10T14:55:55.725+0900",17408000
"2020-11-10T14:56:00.725+0900",17014784
"2020-11-10T14:56:05.726+0900",19193856
"2020-11-10T14:56:10.725+0900",19136512
"2020-11-10T14:56:15.726+0900",16637952
"2020-11-10T14:56:20.725+0900",16678912
"2020-11-10T14:56:25.727+0900",16678912
"2020-11-10T14:56:30.730+0900",16678912
"2020-11-10T14:56:35.726+0900",16678912
```
## API transaction logs

The toll will record API requests and responses into capture logs that have a prefix "capture". This capture logs do not contain requests and responses of OAuth. Additionally, API token strings are replaced with `<secret>`.

