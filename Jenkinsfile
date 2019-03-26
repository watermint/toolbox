node {
    dir('src/github.com/watermint/toolbox') {
        checkout scm
    }
    def root = tool name: 'Go 1.12.1', type: 'go'

    withEnv(["GOROOT=${root}", "GOPATH=${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}/", "PATH+GO=${root}/bin"]) {
        env.PATH="${GOPATH}/bin:$PATH"
        stage 'Prepare'
        sh 'go get golang.org/x/tools/cmd/cover'
        sh 'go get github.com/modocache/gover'
        sh 'go get github.com/mattn/goveralls'
        sh 'go get github.com/Masterminds/glide'
        sh 'glide install'

        stage('Test')
        sh 'cd src/github.com/watermint/toolbox'
        sh 'go list -f \'{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}\' $(glide novendor) | xargs -L 1 sh -c'
        sh 'gover'
    }
}
