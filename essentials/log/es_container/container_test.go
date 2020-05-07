package es_container

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewToolboxCaller(t *testing.T) {
	qt_file.TestWithTestFolder(t, "toolbox", false, func(path string) {
		l, err := NewToolbox(path, app_budget.BudgetUnlimited, es_log.ConsoleDefaultLevel())
		if err != nil {
			t.Error(err)
			return
		}
		lg := l.Logger()

		err = es_log.EnsureCallerSkip(lg, "msg", "caller", func() string {
			entries, err := ioutil.ReadDir(path)
			if err != nil {
				t.Error(err)
				return ""
			}
			for _, entry := range entries {
				if strings.HasPrefix(entry.Name(), app.LogToolbox) {
					logPath := filepath.Join(path, entry.Name())
					content, err := ioutil.ReadFile(logPath)
					if err != nil {
						t.Error(err)
					} else {
						return string(content)
					}
				}
			}
			t.Error("log not found")
			return ""
		})
		if err != nil {
			t.Error(err)
		}
	})
}
