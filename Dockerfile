FROM watermint/xgo-glide

RUN mkdir /dist
ENV PROJECT_ROOT=$GOPATH/src/github.com/watermint/toolbox
RUN mkdir -p $PROJECT_ROOT

ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT
RUN glide install

ENTRYPOINT $PROJECT_ROOT/build/build_on_docker.sh
