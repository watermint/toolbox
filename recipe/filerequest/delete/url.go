package delete

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_filerequest"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Url struct {
	Url                    string
	Peer                   dbx_conn.ConnUserFile
	Deleted                rp_model.RowReport
	Force                  bool
	ProgressClose          app_msg.Message
	ErrorUnableToClose     app_msg.Message
	ErrorFileRequestIsOpen app_msg.Message
}

func (z *Url) Preset() {
	z.Deleted.SetModel(&mo_filerequest.FileRequest{})
}

func (z *Url) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()
	reqs, err := sv_filerequest.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	if err := z.Deleted.Open(); err != nil {
		return err
	}
	for _, r := range reqs {
		ll := l.With(es_log.Any("File request", r))
		if r.Url != z.Url {
			ll.Debug("skip unrelated file request")
			continue
		}
		switch {
		case z.Force && r.IsOpen:
			ll.Debug("File request is open and force option")
			ui.Info(z.ProgressClose.With("Title", r.Title))
			r.IsOpen = false
			r, err = sv_filerequest.New(z.Peer.Context()).Update(r)
			if err != nil {
				ui.Error(z.ErrorUnableToClose.With("Error", err))
				return err
			}
			ll.Debug("The request closed", es_log.Any("after", r))

		case r.IsOpen:
			ll.Debug("The file request is open")
			ui.Error(z.ErrorFileRequestIsOpen.With("Title", r.Title))
			return errors.New("the file request is open")
		}

		deleted, err := sv_filerequest.New(z.Peer.Context()).Delete(r.Id)
		if err != nil {
			return err
		}
		for _, d := range deleted {
			l.Debug("Deleted request", es_log.Any("deleted", d))
			z.Deleted.Row(d)
		}
	}
	return nil
}

func (z *Url) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Url{}, func(r rc_recipe.Recipe) {
		m := r.(*Url)
		m.Url = "https://www.dropbox.com/request/oaCAVmEyrqYnkZX9955Y"
	})
}
