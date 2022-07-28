FROM golang:1.18

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y zip git curl jq

ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin
RUN mkdir /dist
ENV PROJECT_ROOT=/app

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
