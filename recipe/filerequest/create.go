package filerequest

import (
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/domain/service/sv_filerequest"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"sync"
	"time"
)

type Create struct {
	Title            string
	Path             mo_path.DropboxPath
	Deadline         mo_time.TimeOptional
	AllowLateUploads string
	Peer             rc_conn.ConnUserFile
	FileRequest      rp_model.RowReport
}

func (z *Create) Preset() {
	z.FileRequest.SetModel(&mo_filerequest.FileRequest{})
}

func (z *Create) Exec(c app_control.Control) error {
	opts := make([]sv_filerequest.CreateOpt, 0)
	if z.Deadline.Ok() {
		opts = append(opts, sv_filerequest.OptDeadline(z.Deadline.String()))
		if z.AllowLateUploads != "" {
			opts = append(opts, sv_filerequest.OptAllowLateUploads(z.AllowLateUploads))
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
	wg := sync.WaitGroup{}
	create := func(i int) {
		wg.Add(1)
		rc_exec.Exec(c, z, func(r rc_recipe.Recipe) {
			m := r.(*Create)
			m.Title = fmt.Sprintf("watermint toolbox [%d] %s", i, time.Now().String())
			m.Path = qt_recipe.NewTestDropboxFolderPath("file-request")
			//m.Deadline = mo_time.NewOptional(time.Now().Add(5 * time.Hour))
		})
		wg.Done()
	}
	for i := 0; i < 5000; i++ {
		go create(i)
	}
	wg.Wait()

	return rc_exec.Exec(c, z, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.Title = "watermint toolbox " + time.Now().String()
		m.Path = qt_recipe.NewTestDropboxFolderPath("file-request")
		m.Deadline = mo_time.NewOptional(time.Now().Add(24 * time.Hour))
	})
}
