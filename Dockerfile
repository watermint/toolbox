FROM golang:1.7

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y zip git

RUN curl https://glide.sh/get | sh

RUN cd $GOPATH

RUN mkdir /dist

ENV PROJECT_ROOT=$GOPATH/src/github.com/watermint/toolbox
RUN mkdir -p $PROJECT_ROOT

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
RUN glide install

ENTRYPOINT $PROJECT_ROOT/build/build_on_docker.sh
