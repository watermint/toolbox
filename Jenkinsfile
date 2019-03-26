node {
    dir('src/github.com/watermint/toolbox') {
        checkout scm
    }
    def root = tool name: 'Go 1.12.1', type: 'go'

    withEnv(["GOROOT=${root}", "GOPATH=${WORKSPACE}", "PATH+GO=${root}/bin"]) {
        env.PATH="${GOPATH}/bin:$PATH"
        stage('Prepare') {
            sh 'go get golang.org/x/tools/cmd/cover'
            sh 'go get github.com/modocache/gover'
            sh 'go get github.com/mattn/goveralls'
            sh 'go get github.com/Masterminds/glide'
        }

        stage('Test') {
            dir('src/github.com/watermint/toolbox') {
                sh 'glide install'
                sh 'go list -f \'{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}\' $(glide novendor) | xargs -L 1 sh -c'
                sh 'gover'
            }
        }

        stage('Report') {
            environment {
                COVERALLS_TOKEN = credentials('COVERALLS_TOKEN')
            }
            dir('src/github.com/watermint/toolbox') {
                sh 'goveralls -coverprofile=gover.coverprofile -service=travis-ci -repotoken $COVERALLS_TOKEN'
            }
        }
    }
}
