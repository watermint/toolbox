# Jenkins setup instruction

* Prepare https key store (https.jks), otherwise comment update Dockerfile
* Copy ~/.toolbox/secrets/*.tokens to this build directory
* Build docker image
* Install Go plugin
* Configure Go from 'Global Tool Configuration' with name 'Go <version>'
* Configure credential for coveralls with credential id 'github.com-watermint-toolbox-goverall-token'

## HTTPS

To build Jenkins docker with https, please sepcify keystore password like below

```
docker build --build-arg HTTPS_KEYSTORE_PASSWORD=<password> -t toolbox-jenkins .
```

Then, run like below

```
docker run -d -p 443:8443 toolbox-jenkins
```

