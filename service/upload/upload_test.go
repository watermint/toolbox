package upload

import (
	"fmt"
	"github.com/watermint/toolbox/infra"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"
)

func TestUploadAndCompare(t *testing.T) {
	infraOpts := infra.InfraOpts{}
	err := infraOpts.Startup()
	if err != nil {
		t.Skip("Skip")
		return
	}
	defer infraOpts.Shutdown()
	token := os.Getenv("TEST_TOKEN_DROPBOX_FULL")
	if token == "" {
		t.Skip("No token for test.")
		return
	}

	base := os.TempDir()
	tmpd, err := ioutil.TempDir(base, "traverse")
	if err != nil {
		t.Error(err)
	}
	println(tmpd)
	tmpf, err := ioutil.TempFile(tmpd, "localfile")
	if err != nil {
		t.Error(err)
	}
	tmpf.WriteString("Hello")
	tmpf.Close()

	dbxTestPath := "/test"
	dbxTestSession := fmt.Sprintf("%x", time.Now().Unix())
	dbxBasePath := path.Join(dbxTestPath, dbxTestSession)

	uc := &UploadContext{
		LocalPaths:         []string{tmpd},
		LocalRecursive:     true,
		LocalFollowSymlink: false,
		DropboxBasePath:    dbxBasePath,
		DropboxToken:       token,
		DeleteAfterUpload:  true,
		BandwidthLimit:     0,
		Concurrency:        1,
	}
	uc.Upload()
}
