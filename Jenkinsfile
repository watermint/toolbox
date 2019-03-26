pipeline {
    agent any
    stages {
        stage('Env') {
            environment {
                GO_BINARY = "go1.12.1.linux-amd64.tar.gz"
            }
            steps {
                sh 'export GOPATH=$PWD/go'
                sh 'export PATH=$PATH:$GOPATH'
                sh 'if [ ! -d go ]; then mkdir go fi'
                sh 'if [ ! -e $GO_BINARY ]; then wget https://dl.google.com/go/$GO_BINARY fi'
                sh 'tar -C $GOPATH -xzf $GO_BINARY'
            }
        }
        stage('Prepare') {
            steps {
                sh 'go get golang.org/x/tools/cmd/cover'
                sh 'go get github.com/modocache/gover'
                sh 'go get github.com/mattn/goveralls'
                sh 'go get github.com/Masterminds/glide'
                sh 'glide install'
            }
        }
        stage('Test') {
            steps {
                sh 'go list -f \'{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}\' $(glide novendor) | xargs -L 1 sh -c'
                sh 'gover'
            }
        }
        stage('Report') {
            environment {
                GOVERALLS_TOKEN = credentials('GOVERALLS_TOKEN')
            }
            steps {
                sh 'goveralls -coverprofile=gover.coverprofile -service=travis-ci -repotoken $COVERALLS_TOKEN'
            }
        }
    }
}