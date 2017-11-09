# dfm: Dropbox File/folder Management util

# Move

File/folder move utility.
This command moves files/folders into the specified path within Dropbox.

## Usage

```bash
$ ./dfm move [OPTION]... SRC DEST
```


# Restore

File/folder restore utility. This command restores deleted files/folders under the specified path inside Dropbox.

## Usage

```bash
$ ./dfm restore [OPTION]... PATH
```


# Compare

File comparison command. Generate report differences between local files and files on Dropbox.
Comparing files by not only filename/size but also a content hash.

## Usage

```bash
$ ./dfm compare LOCAL_PATH [DROPBOX_PATH]
```

## Generate report

Generate report by `.xlsx` format.

```bash
$ ./dfm compare -xlsx REPORTNAME.xlsx LOCAL_PATH [DROPBOX_PATH]
```

Generate report by CSV

```bash
$ ./dfm compare -csv-dir CSV_DIR LOCAL_PATH [DROPBOX_PATH]
```

# Upload

Bulk file and/or directory upload command.

## Usage

```sh
$ ./dfm upload [OPTION]... SRC [SRC]... DEST
```

## Options

```
Usage:
dfm upload [OPTION]... SRC [SRC]... DEST
  -L	Follow symlinks
  -bwlimit int
    	Limit upload bandwidth; KBytes per second (not kbps)
  -c int
    	Upload concurrency (default 1)
  -concurrency int
    	Upload concurrency (default 1)
  -follow-symlink
    	Follow symlinks
  -proxy string
    	HTTP/HTTPS proxy (hostname:port)
  -r	Recurse into directories (default true)
  -recursive
    	Recurse into directories (default true)
  -revoke-token
    	Revoke token on exit
  -work string
    	Work directory
```
