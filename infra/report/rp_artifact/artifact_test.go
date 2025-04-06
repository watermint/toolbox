package rp_artifact

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

func TestArtifacts(t *testing.T) {
	qt_file.TestWithTestFolder(t, "artifact", false, func(path string) {
		ws, err := app_workspace.NewWorkspace(path, false)
		if err != nil {
			t.Error(err)
			return
		}
		err = os.WriteFile(filepath.Join(ws.Report(), "test.csv"), []byte("1,2,3"), 0644)
		if err != nil {
			t.Error(err)
			return
		}
		err = os.WriteFile(filepath.Join(ws.Report(), "test.zip"), []byte("PK"), 0644)
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
