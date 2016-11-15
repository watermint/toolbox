package diag

import (
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"net/http"
	"os"
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
}

func NewDiagnosticsRuntime() DiagnosticsRuntime {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()

	return DiagnosticsRuntime{
		Args:             os.Args,
		Hostname:         hostname,
		Environment:      os.Environ(),
		WorkingDirectory: wd,
		PID:              os.Getpid(),
		ExecutorUID:      os.Geteuid(),
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
	j, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("%v", d)
	}
	return string(j)
}

func DigestDiagnosticsRuntime() string {
	d := NewDiagnosticsRuntime()
	j, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("%v", d)
	}
	return string(j)
}

func DigestDiagnosticsNetwork(testUrl string) string {
	d := NewDiagnosticsNetwork(testUrl)
	j, err := json.Marshal(d)
	if err != nil {
		return fmt.Sprintf("%v", d)
	}
	return string(j)
}

func LogDiagnostics() {
	testUrls := []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}

	seelog.Infof("Diagnostics(Infra): %s", DigestDiagnosticsInfra())
	seelog.Infof("Diagnostics(Runtime): %s", DigestDiagnosticsRuntime())
	for _, u := range testUrls {
		seelog.Infof("Diagnostics(Network:%s): %s", u, DigestDiagnosticsNetwork(u))
	}
	seelog.Flush()
}
