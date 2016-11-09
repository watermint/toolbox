# toolbox

Tools for Dropbox and Dropbox Business

## Build

```sh
$ docker build -t toolbox . && rm -fr /tmp/dist && docker run -v /tmp/dist:/dist:rw --rm toolbox
```
