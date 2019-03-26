node {
    def root = tool name: 'Go 1.12.1', type: 'go'

    withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
        stages {
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
}
