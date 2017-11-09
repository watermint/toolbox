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

# Enforce shared link expiration

## Update expiration of shared links at +7 days if expiration not set

```sh
$ ./dsharedlink expire -team -days 7
```

## Update and overwrite expiration date

```sh
$ ./dsharedlink expire -team -days 7 -overwrite
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

