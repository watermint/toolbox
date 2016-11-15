# dupload

Bulk file and/or directory Uploader.

## Usage

```sh
$ ./dupload -localPath YOUR_LOCAL_PATH -dropboxPath DROPBOX_PATH
```

## How to build

* Copy `credentials.sample` with name `credentials.secret`.
* Update `ApiKey` and `ApiSecret` for your Application ID.
* Build entire project using Dockerfile on top of the project.
