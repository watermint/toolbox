# dteammember

Team member management utility.

# Usage

## Detach (delete account from Team, but keep account as Dropbox Basic account)

```sh
$ ./dteammember detach -user someone@example.com
```

## List team member(s)

```shh
$ ./dteammember list -status invited
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
