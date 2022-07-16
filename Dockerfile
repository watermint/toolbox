FROM golang:1.18

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y zip git curl jq

ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go get golang.org/x/tools/cmd/cover
RUN go get github.com/jstemmer/go-junit-report
RUN go get github.com/google/go-licenses
RUN mkdir /dist
ENV PROJECT_ROOT=/app

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
