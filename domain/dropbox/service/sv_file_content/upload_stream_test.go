package sv_file_content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
	"time"
)

func TestStreamImpl_Upload(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(client dbx_client.Client) {
		sv := NewUploadStream(client, false, false)
		_, err := sv.Upload(
			qtr_endtoend.NewTestDropboxFolderPath("test.txt"),
			es_rewinder.NewReadRewinderOnMemory([]byte("hello")),
			time.Now(),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
