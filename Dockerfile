FROM golang:1.14

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y zip git curl jq

ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin
RUN go get github.com/GeertJohan/go.rice
RUN go get github.com/GeertJohan/go.rice/rice
RUN go get github.com/derekparker/delve/cmd/dlv

RUN mkdir /dist
ENV PROJECT_ROOT=$GOPATH/src/github.com/watermint/toolbox
RUN mkdir -p $PROJECT_ROOT

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
