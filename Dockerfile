FROM karalabe/xgo-1.8.x

RUN apt-get update -y
RUN apt-get upgrade -y
RUN apt-get install -y zip git curl

ENV GOBIN=/usr/local/go/bin
ENV PATH=$PATH:/usr/local/go/bin
RUN curl https://glide.sh/get | sh

RUN mkdir /dist
ENV PROJECT_ROOT=$GOPATH/src/github.com/watermint/toolbox
RUN mkdir -p $PROJECT_ROOT

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
RUN glide install

ENTRYPOINT $PROJECT_ROOT/build/build_on_docker.sh
