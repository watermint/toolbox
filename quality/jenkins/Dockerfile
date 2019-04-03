FROM jenkins/jenkins:lts

USER root
ENV GO_BINARY=go1.12.1.linux-amd64.tar.gz

RUN apt-get update && apt-get install -y gcc
RUN wget https://dl.google.com/go/$GO_BINARY
RUN tar -C /usr/local -xzf $GO_BINARY

USER jenkins
RUN export PATH=$PATH:/usr/local/go/bin
