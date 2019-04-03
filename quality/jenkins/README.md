# Jenkins setup instruction

* Build docker image
* Install Go plugin
* Configure Go from 'Global Tool Configuration' with name 'Go <version>'
* Configure credential for coveralls with credential id 'github.com-watermint-toolbox-goverall-token'
* Copy test tokens into `/var/jenkins_home/.toolbox/secrets`, then change owner to `jenkins:jenkins`
