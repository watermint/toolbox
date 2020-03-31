package filerequest

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_filerequest"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"strings"
)

type Clone struct {
	File         fd_file.RowFeed
	Peer         rc_conn.ConnBusinessFile
	OperationLog rp_model.TransactionReport
}

func (z *Clone) Preset() {
	z.File.SetModel(&mo_filerequest.MemberFileRequest{})
	z.OperationLog.SetModel(&mo_filerequest.MemberFileRequest{}, &mo_filerequest.MemberFileRequest{})
}

func (z *Clone) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	emailToMember := mo_member.MapByEmail(members)

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		fm := m.(*mo_filerequest.MemberFileRequest)
		if fm.Email == "" || fm.Destination == "" || fm.Title == "" {
			z.OperationLog.Failure(errors.New("invalid data"), fm)
			return nil
		}
		member, ok := emailToMember[strings.ToLower(fm.Email)]
		if !ok {
			z.OperationLog.Failure(errors.New("entry not found for the id"), fm)
			return nil
		}

		opts := make([]sv_filerequest.CreateOpt, 0)
		if fm.Deadline != "" {
			opts = append(opts, sv_filerequest.OptDeadline(fm.Deadline))
		}
		if fm.DeadlineAllowLateUploads != "" {
			opts = append(opts, sv_filerequest.OptAllowLateUploads(fm.DeadlineAllowLateUploads))
		}
		req, err := sv_filerequest.New(z.Peer.Context().AsMemberId(member.TeamMemberId)).Create(
			fm.Title,
			mo_path.NewDropboxPath(fm.Destination),
			opts...,
		)
		if err != nil {
			z.OperationLog.Failure(err, fm)
		} else {
			z.OperationLog.Success(fm, req)
		}
		return nil
	})
}

func (z *Clone) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Clone{}, func(r rc_recipe.Recipe) {
		f, err := qt_file.MakeTestFile("filerequest-clone", `account_id,team_member_id,email,status,surname,given_name,file_request_id,url,title,created,is_open,file_count,destination,deadline,deadline_allow_late_uploads
dbid:xxxxxxxxxxxxxxx-xxxxxxxxxxxxxxxxxxx,dbmid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx,xxx+xxx@xxxxxxxxx.xxx,active,xx,xx,xxxxxxxxxxxxxxxxxxxx,https://www.dropbox.com/request/xxxxxxxxxxxxxxxxxxxx,xxxxxx,2017-10-16T03:08:21Z,false,1,/xxxxxxxxxx,2017-10-23T03:00:00Z,two_days
`)
		if err != nil {
			return
		}
		m := r.(*Clone)
		m.File.SetFilePath(f)
	})
	if e, _ := qt_recipe.RecipeError(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}
