package device

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_device"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_device"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type UnlinkVO struct {
}

type UnlinkWorker struct {
	session *mo_device.MemberSession
	rep     rp_model.TransactionReport
	ctx     api_context.DropboxApiContext
	ctl     app_control.Control
}

func (z *UnlinkWorker) Exec() error {
	ui := z.ctl.UI()
	ui.InfoK("recipe.team.device.unlink.progress", app_msg.P{
		"Member":      z.session.Email,
		"SessionType": z.session.DeviceTag,
		"SessionId":   z.session.Id,
	})

	s := &mo_device.Metadata{
		Tag:          z.session.DeviceTag,
		TeamMemberId: z.session.TeamMemberId,
		Id:           z.session.Id,
	}
	err := sv_device.New(z.ctx).Revoke(s)
	if err != nil {
		z.rep.Failure(err, z.session)
		return err
	}
	z.rep.Success(z.session, nil)
	return nil
}

type Unlink struct {
	DeleteOnUnlink bool
	File           fd_file.RowFeed
	Peer           rc_conn.ConnBusinessFile
	OperationLog   rp_model.TransactionReport
}

func (z *Unlink) Preset() {
	z.File.SetModel(&mo_device.MemberSession{})
	z.OperationLog.SetModel(&mo_device.MemberSession{}, nil)
}

func (z *Unlink) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	q := c.NewQueue()
	err := z.File.EachRow(func(m interface{}, rowIndex int) error {
		q.Enqueue(&UnlinkWorker{
			session: m.(*mo_device.MemberSession),
			rep:     z.OperationLog,
			ctx:     z.Peer.Context(),
			ctl:     c,
		})
		return nil
	})
	q.Wait()
	return err
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
	if e, _ := qt_recipe.RecipeError(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
