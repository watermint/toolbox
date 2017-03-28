# dcmp

File comparison utility.

# Usage

```sh
$ ./dcmp LOCAL_PATH [DROPBOX_PATH]
```

## Generate report

Generate report by `.xlsx` format.

```sh
$ ./dcmp -xlsx REPORTNAME.xlsx LOCAL_PATH [DROPBOX_PATH]
```

Generate report by CSV

```sh
$ ./dcmp -csv-dir CSV_DIR LOCAL_PATH [DROPBOX_PATH]
```

## Options

```
  -proxy string
    	HTTP/HTTPS proxy (hostname:port)
  -revoke-token
    	Revoke token on exit
  -work string
    	Work directory
```
