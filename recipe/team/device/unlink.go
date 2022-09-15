package device

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_device"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_device"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Unlink struct {
	rc_recipe.RemarkIrreversible
	DeleteOnUnlink bool
	File           fd_file.RowFeed
	Peer           dbx_conn.ConnScopedTeam
	OperationLog   rp_model.TransactionReport
}

func (z *Unlink) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSessionsModify,
	)
	z.File.SetModel(&mo_device.MemberSession{})
	z.OperationLog.SetModel(&mo_device.MemberSession{}, nil,
		rp_model.HiddenColumns(
			"input.familiar_name",
			"input.abbreviated_name",
			"input.member_folder_id",
			"input.external_id",
			"input.account_id",
			"input.persistent_id",
		),
	)
}

func (z *Unlink) unlink(session *mo_device.MemberSession) error {
	s := &mo_device.Metadata{
		Tag:          session.DeviceTag,
		TeamMemberId: session.TeamMemberId,
		Id:           session.Id,
	}
	err := sv_device.New(z.Peer.Client()).Revoke(s)
	if err != nil {
		z.OperationLog.Failure(err, session)
		return err
	}
	z.OperationLog.Success(session, nil)
	return nil
}

func (z *Unlink) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("unlink", z.unlink)
		q := s.Get("unlink")
		lastErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	})

	return lastErr
}

func (z *Unlink) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Unlink{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("session-unlink",
			`team_member_id,email,status,given_name,surname,familiar_name,display_name,abbreviated_name,external_id,account_id,device_tag,id,user_agent,os,browser,ip_address,country,created,updated,expires,host_name,client_type,client_version,platform,is_delete_on_unlink_supported,device_name,os_version,last_carrier
dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,xxx+xxxx@xxxxxxxxx.xxx,active,xx,xxxxx,xxxxx xx,xxxxx xx,xx,xxx xxx+xxxx@xxxxxxxxx.xxx xxxx-xx-xxxxx-xx-xx,dbid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,desktop_client,dbdsid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,,,,xx.xxx.x.xxx,United States,2019-09-20T23:47:33Z,2019-10-25T04:42:16Z,,xxxxxxxxxx,windows,83.4.152,Windows 10 1903,true,,,
`)
		if err != nil {
			return
		}
		m := r.(*Unlink)
		m.File.SetFilePath(f)
		m.DeleteOnUnlink = true
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
