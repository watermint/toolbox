# dupload

Bulk file and/or directory Uploader.

## Usage

```sh
$ ./dupload [OPTION]... SRC [SRC]... DEST
```

## Options

```
Usage:
dupload [OPTION]... SRC [SRC]... DEST
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

## How to build

* Copy `credentials.sample` with name `credentials.secret`.
* Update `ApiKey` and `ApiSecret` for your Application ID.
* Build entire project using Dockerfile on top of the project.
