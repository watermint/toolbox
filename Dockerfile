FROM techknowlogick/xgo:latest

ENV PROJECT_ROOT=/source
ENV BUILD_FILE=$PROJECT_ROOT/build.sh
ENV mkdir -p $PROJECT_ROOT
ADD . $PROJECT_ROOT
WORKDIR $PROJECT_ROOT

ENTRYPOINT ["bash", "/source/build.sh"]

