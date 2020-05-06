package rp_artifact

import (
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestArtifacts(t *testing.T) {
	qt_file.TestWithTestFolder(t, "artifact", false, func(path string) {
		ws, err := app_workspace.NewWorkspace(path)
		if err != nil {
			t.Error(err)
			return
		}
		err = ioutil.WriteFile(filepath.Join(ws.Report(), "test.csv"), []byte("1,2,3"), 0644)
		if err != nil {
			t.Error(err)
			return
		}
		err = ioutil.WriteFile(filepath.Join(ws.Report(), "test.zip"), []byte("PK"), 0644)
		if err != nil {
			t.Error(err)
			return
		}

		artifacts := Artifacts(ws)
		if len(artifacts) != 1 {
			t.Error(artifacts)
		}
		if artifacts[0].Name() != "test.csv" {
			t.Error(artifacts[0].Name())
		}
	})
}
