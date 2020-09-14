package filerequest

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_filerequest"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"time"
)

type Create struct {
	rc_recipe.RemarkIrreversible
	Title            string
	Path             mo_path.DropboxPath
	Deadline         mo_time.TimeOptional
	AllowLateUploads mo_string.OptionalString
	Peer             dbx_conn.ConnUserFile
	FileRequest      rp_model.RowReport
}

func (z *Create) Preset() {
	z.FileRequest.SetModel(&mo_filerequest.FileRequest{})
}

func (z *Create) Exec(c app_control.Control) error {
	opts := make([]sv_filerequest.CreateOpt, 0)
	if z.Deadline.Ok() {
		opts = append(opts, sv_filerequest.OptDeadline(z.Deadline.Value()))
		if z.AllowLateUploads.IsExists() {
			opts = append(opts, sv_filerequest.OptAllowLateUploads(z.AllowLateUploads.Value()))
		}
	}
	if err := z.FileRequest.Open(); err != nil {
		return err
	}
	fr, err := sv_filerequest.New(z.Peer.Context()).Create(z.Title, z.Path, opts...)
	if err != nil {
		return err
	}
	z.FileRequest.Row(fr)
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, z, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.Title = "watermint toolbox " + time.Now().String()
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("file-request")
		m.Deadline = mo_time.NewOptional(time.Now().Add(24 * time.Hour))
	})
	ers := dbx_error.NewErrors(err)
	if ers != nil && ers.Endpoint() != nil && ers.Endpoint().IsRateLimit() {
		// In case of the account has 4,000 > file requests
		c.Log().Info("The test account has more than 4,000 file requests")
		return nil
	}
	return err
}
