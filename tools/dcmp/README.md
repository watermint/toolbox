# dcmp

File comparison utility. Generate report differences between local files and files on Dropbox.
Comparing files by not only filename/size but also a content hash.

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
