FROM golang:1.11

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y zip git curl

ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin
RUN curl https://glide.sh/get | sh
RUN go get github.com/GeertJohan/go.rice
RUN go get github.com/GeertJohan/go.rice/rice

RUN mkdir /dist
ENV PROJECT_ROOT=$GOPATH/src/github.com/watermint/toolbox
RUN mkdir -p $PROJECT_ROOT

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
RUN glide install

ENTRYPOINT $PROJECT_ROOT/build.sh
