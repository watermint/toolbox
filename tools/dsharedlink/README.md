# dsharedlink

Shared link utility.

# Usage

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

## How to build

* Copy `credentials.sample` with name `credentials.secret`.
* Update `ApiKey` and `ApiSecret` for your Application ID.
* Build entire project using Dockerfile on top of the project.
