FROM golang:1.19

ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin
RUN mkdir /dist
ENV PROJECT_ROOT=/app

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
