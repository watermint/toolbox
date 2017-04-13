package compare

import (
	"fmt"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/service/report"
	"github.com/watermint/toolbox/service/upload"
	"io/ioutil"
	"os"
	"path"
	"testing"
	"time"
)

func TestContentHash(t *testing.T) {
	tmpf, err := ioutil.TempFile("", "hash_test")
	if err != nil {
		t.Error("Unable to create temp file", err)
	}

	h, err := ContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != HASH_FOR_EMPTY {
		t.Errorf("Hash not matched expected[%s] actual[%s]", HASH_FOR_EMPTY, h)
	}

	if _, err = tmpf.WriteString("Hello"); err != nil {
		t.Error("Unabble to write content", err)
	}
	hashHello := "70bc18bef5ae66b72d1995f8db90a583a60d77b4066e4653f1cead613025861c"
	h, err = ContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != hashHello {
		t.Errorf("Hash not matched expected[%s] actual[%s]", hashHello, h)
	}
	expectedChunked := "4fdd596c2141f63b577aa7be1e73bb78fbd3966a1d8b46d56650e26b07acea16"
	for i := 0; i < BLOCK_SIZE/2; i++ {
		_, err = tmpf.WriteString("COMPARE")
	}
	h, err = ContentHash(tmpf.Name())
	if err != nil {
		t.Error("Unable to read file", err)
	}
	if h != expectedChunked {
		t.Errorf("Hash not matched expected[%s] actual[%s]", expectedChunked, h)
	}
}

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
	wd, err := os.Getwd()
	if err != nil {
		t.Error("Could not acquire wd", err)
		return
	}

	localBasePath := path.Join(wd, "infra")
	dbxTestPath := "/test"
	dbxTestSession := fmt.Sprintf("%x", time.Now().Unix())
	dbxBasePath := path.Join(dbxTestPath, dbxTestSession)

	uc := &upload.UploadContext{
		LocalPaths:         []string{localBasePath},
		LocalRecursive:     true,
		LocalFollowSymlink: false,
		DropboxBasePath:    dbxBasePath,
		DropboxToken:       token,
		BandwidthLimit:     0,
		Concurrency:        1,
	}
	uc.Upload()

	co := CompareOpts{
		InfraOpts:       &infraOpts,
		ReportOpts:      &report.MultiReportOpts{},
		DropboxToken:    token,
		DropboxBasePath: dbxBasePath,
		LocalBasePath:   localBasePath,
	}
	match, err := Compare(&co)
	if err != nil {
		t.Error(err)
		return
	}
	if !match {
		t.Error("Contents are not matched")
	}
}
