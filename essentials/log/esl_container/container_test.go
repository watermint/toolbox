package esl_container

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewToolboxCaller(t *testing.T) {
	qt_file.TestWithTestFolder(t, "toolbox", false, func(path string) {
		l, err := NewToolbox(path, app_budget.BudgetUnlimited, esl.ConsoleDefaultLevel())
		if err != nil {
			t.Error(err)
			return
		}
		lg := l.Logger()

		err = esl.EnsureCallerSkip(lg, "msg", "caller", func() string {
			entries, err := ioutil.ReadDir(path)
			if err != nil {
				t.Error(err)
				return ""
			}
			for _, entry := range entries {
				if strings.HasPrefix(entry.Name(), app_definitions.LogToolbox) {
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
