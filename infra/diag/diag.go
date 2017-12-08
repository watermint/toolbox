package diag

import (
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/util"
	"net/http"
	"os"
	"os/user"
	"runtime"
)

type InfraDiag struct {
	OS        string
	Arch      string
	GoVersion string
	NumCpu    int
}

func NewInfraDiag() InfraDiag {
	return InfraDiag{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		GoVersion: runtime.Version(),
		NumCpu:    runtime.NumCPU(),
	}
}

type RuntimeDiag struct {
	Args             []string
	Hostname         string
	Environment      []string
	WorkingDirectory string
	PID              int
	ExecutorUID      int
	CurrentUID       string
	UserName         string
}

func NewRuntimeDiag() RuntimeDiag {
	hostname, _ := os.Hostname()
	wd, _ := os.Getwd()
	usr, _ := user.Current()

	return RuntimeDiag{
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

type NetworkDiag struct {
	TestUrl      string
	StatusCode   int
	ErrorSummary string
}

func NewNetworkDiag(testUrl string) NetworkDiag {
	seelog.Tracef("Start diagnostics of network: [%s]", testUrl)
	resp, err := http.Head(testUrl)

	if err != nil {
		return NetworkDiag{
			TestUrl:      testUrl,
			ErrorSummary: err.Error(),
		}
	}

	return NetworkDiag{
		TestUrl:    testUrl,
		StatusCode: resp.StatusCode,
	}
}

func DigestDiagnosticsInfra() string {
	d := NewInfraDiag()
	return util.MarshalObjectToString(d)
}

func DigestDiagnosticsRuntime() string {
	d := NewRuntimeDiag()
	return util.MarshalObjectToString(d)
}

func DigestDiagnosticsNetwork(testUrl string) string {
	d := NewNetworkDiag(testUrl)
	return util.MarshalObjectToString(d)
}

func LogDiagnostics() {
	seelog.Tracef("Diagnostics(Infra): %s", DigestDiagnosticsInfra())
	seelog.Tracef("Diagnostics(Runtime): %s", DigestDiagnosticsRuntime())
	seelog.Flush()
}

func QuickNetworkDiagnostics() error {
	resp, err := http.Head("https://www.dropbox.com")
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		seelog.Warnf("Server error: status code[%d]", resp.StatusCode)
		return errors.New(fmt.Sprintf("Server error response code(%d)", resp.StatusCode))
	}
	return nil
}

func LogNetworkDiagnostics() {
	testUrls := []string{
		"https://www.dropbox.com",
		"https://api.dropboxapi.com",
	}
	for _, u := range testUrls {
		seelog.Tracef("Diagnostics(Network:%s): %s", u, DigestDiagnosticsNetwork(u))
	}
	seelog.Flush()
}
