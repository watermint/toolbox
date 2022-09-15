package benchmark

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/io/es_file_random"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"io/ioutil"
	"net/http"
	"time"
)

type Uploadlink struct {
	rc_recipe.RemarkSecret
	Peer   dbx_conn.ConnScopedIndividual
	Path   mo_path.DropboxPath
	SizeKb int
}

func (z *Uploadlink) Preset() {
	z.SizeKb = 1024
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
}

func (z *Uploadlink) Exec(c app_control.Control) error {
	l := c.Log()

	link, err := sv_file.NewFiles(z.Peer.Client()).UploadLink(z.Path)
	if err != nil || link == "" {
		l.Error("Unable to create an upload link", esl.Error(err))
		return err
	}

	l.Info("Upload link", esl.String("link", link))

	timeStart := time.Now()
	res, err := http.Post(link, "application/octet-stream", es_file_random.NewReader(uint64(z.SizeKb*1024)))
	timeEnd := time.Now()

	l.Info("The request finished", esl.Duration("Duration", timeEnd.Sub(timeStart)))
	if err != nil {
		l.Error("Unable to upload a file", esl.Error(err))
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		l.Warn("Unable to read a body", esl.Error(err))
	}

	l.Info("Successfully uploaded",
		esl.Int("StatusCode", res.StatusCode),
		esl.Int64("ContentLength", res.ContentLength),
		esl.ByteString("Body", body),
		esl.Any("Header", res.Header),
	)
	return nil
}

func (z *Uploadlink) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Uploadlink{}, func(r rc_recipe.Recipe) {
		m := r.(*Uploadlink)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("upload-link")
	})
}
