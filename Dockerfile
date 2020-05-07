FROM golang:1.14

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y zip git curl jq

ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin
RUN go get github.com/GeertJohan/go.rice
RUN go get github.com/GeertJohan/go.rice/rice
RUN go get github.com/derekparker/delve/cmd/dlv
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/jstemmer/go-junit-report
RUN mkdir /dist
ENV PROJECT_ROOT=/app

WORKDIR $PROJECT_ROOT
