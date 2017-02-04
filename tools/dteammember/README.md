# dteammember

Team member management utility.

# Usage

## Detach (delete account from Team, but keep account as Dropbox Basic account)

```sh
$ ./dteammember detach -user someone@example.com
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

## How to build

* Copy `credentials.sample` with name `credentials.secret`.
* Update `ApiKey` and `ApiSecret` for your Application ID.
* Build entire project using Dockerfile on top of the project.
