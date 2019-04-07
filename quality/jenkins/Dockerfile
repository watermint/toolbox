FROM jenkins/jenkins:lts

USER root
ENV GO_BINARY=go1.12.1.linux-amd64.tar.gz

RUN apt-get update && apt-get install -y gcc
RUN wget https://dl.google.com/go/$GO_BINARY
RUN tar -C /usr/local -xzf $GO_BINARY
ENV CERT_PATH /var/jenkins_cert
RUN mkdir -p $CERT_PATH
ADD https.jks $CERT_PATH/https.jks

USER jenkins

# --- If you do not require https, comment out following lines
ARG HTTPS_KEYSTORE_PASSWORD
ENV JENKINS_OPTS --httpPort=-1 --httpsPort=8443 --httpsKeyStore=$CERT_PATH/https.jks --httpsKeyStorePassword=$HTTPS_KEYSTORE_PASSWORD
EXPOSE 8443
# ---

RUN export PATH=$PATH:/usr/local/go/bin

