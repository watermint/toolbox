package app_run

import (
	"encoding/json"
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	RunBookName     = "tbx.runbook"
	RunBookTestName = "tbx.runbook.test"
)

type RunEntry struct {
	Args []string `json:"args"`
}

type RunBook struct {
	Entry []*RunEntry `json:"entry"`
}

func (z *RunBook) Exec(bx, web *rice.Box) {
	l := app_root.Log()
	for _, e := range z.Entry {
		l.Info("Run", zap.Strings("Args", e.Args))
		Run(e.Args, bx, web)
	}
}

func NewRunBook(path string) (rb *RunBook, found bool) {
	l := app_root.Log()
	_, err := os.Lstat(path)
	if err != nil {
		return nil, false
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		l.Error("Unable to read the runbook file", zap.Error(err))
		return nil, false
	}

	rb = &RunBook{}
	if err = json.Unmarshal(content, rb); err != nil {
		l.Error("Unable to unmarshal the runbook file", zap.Error(err))
		return nil, false
	}

	return rb, true
}

func DefaultRunBook(forTest bool) (rb *RunBook, found bool) {
	name := RunBookName
	if forTest {
		name = RunBookTestName
	}
	path := filepath.Join(filepath.Dir(os.Args[0]), name)
	return NewRunBook(path)
}
