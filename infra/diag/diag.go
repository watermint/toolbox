package diag

import (
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/util"
	"net/http"
	"os"
	"os/user"
	"runtime"
)

const ()

type DiagnosticsInfra struct {
	OS        string
	Arch      string
	GoVersion string
	NumCpu    int
}

func NewDiagnosticsInfra() DiagnosticsInfra {
	return DiagnosticsInfra{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		NumCpu:    runtime.NumCPU(),
	}
}

type DiagnosticsRuntime struct {
	Args             []string
	Hostname         string
	Environment      []string
	WorkingDirectory string
	PID              int
	ExecutorUID      int
	CurrentUID       string
	UserName         string
}

func NewDiagnosticsRuntime() DiagnosticsRuntime {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	usr, _ := user.Current()

	return DiagnosticsRuntime{
		Args:             os.Args,
		Hostname:         hostname,
		Environment:      os.Environ(),
		WorkingDirectory: wd,
		PID:              os.Getpid(),
		ExecutorUID:      os.Geteuid(),
		CurrentUID:       usr.Uid,
		UserName:         usr.Username,
	}
}

type DiagnosticsNetwork struct {
	TestUrl      string
	StatusCode   int
	ErrorSummary string
}

func NewDiagnosticsNetwork(testUrl string) DiagnosticsNetwork {
	seelog.Tracef("Start diagnostics of network: [%s]", testUrl)
	resp, err := http.Head(testUrl)

	if err != nil {
		return DiagnosticsNetwork{
			TestUrl:      testUrl,
			ErrorSummary: err.Error(),
		}
	}

	return DiagnosticsNetwork{
		TestUrl:    testUrl,
		StatusCode: resp.StatusCode,
	}
}

func DigestDiagnosticsInfra() string {
	d := NewDiagnosticsInfra()
	return util.MarshalObjectToString(d)
}

func DigestDiagnosticsRuntime() string {
	d := NewDiagnosticsRuntime()
	return util.MarshalObjectToString(d)
}

func DigestDiagnosticsNetwork(testUrl string) string {
	d := NewDiagnosticsNetwork(testUrl)
	return util.MarshalObjectToString(d)
}

func LogDiagnostics() {
	seelog.Infof("Diagnostics(Infra): %s", DigestDiagnosticsInfra())
	seelog.Infof("Diagnostics(Runtime): %s", DigestDiagnosticsRuntime())
	seelog.Flush()
}

func LogNetworkDiagnostics() {
	testUrls := []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}
	for _, u := range testUrls {
		seelog.Infof("Diagnostics(Network:%s): %s", u, DigestDiagnosticsNetwork(u))
	}
}
