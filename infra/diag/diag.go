package diag

import (
	"encoding/json"
	"fmt"
	"github.com/cihub/seelog"
	"os"
	"runtime"
)

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

func LogDiagnostics() {
	seelog.Infof("Diagnostics(Infra): %s", DigestDiagnosticsInfra())
	seelog.Infof("Diagnostics(Runtime): %s", DigestDiagnosticsRuntime())
	seelog.Flush()
}
