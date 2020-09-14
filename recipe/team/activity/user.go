package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_activity"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/strings/es_mailaddr"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"time"
)

type MsgUser struct {
	ProgressScanUser app_msg.Message
}

var (
	MUser = app_msg.Apply(&MsgUser{}).(*MsgUser)
)

type UserSummary struct {
	User           string `json:"user"`
	Logins         int64  `json:"logins"`
	Devices        int64  `json:"devices"`
	Sharing        int64  `json:"sharing"`
	FileOperations int64  `json:"file_operations"`
	Paper          int64  `json:"paper"`
	Others         int64  `json:"others"`
}

type UserIn struct {
	User string `json:"user"`
}

type UserWorker struct {
	ctl        app_control.Control
	ctx        dbx_context.Context
	reps       rp_model.RowReport
	repSummary rp_model.TransactionReport
	user       *mo_member.Member
	StartTime  string
	EndTime    string
	Category   string
}

func (z *UserWorker) Exec() error {
	userIn := &UserIn{User: z.user.Email}
	ui := z.ctl.UI()
	l := z.ctl.Log().With(esl.Any("userIn", userIn))

	rep, err := z.reps.OpenNew(
		rp_model.Suffix("-"+es_mailaddr.EscapeSpecial(z.user.Email, "_")),
		rp_model.NoConsoleOutput(),
	)
	if err != nil {
		l.Debug("unable to create report", esl.Error(err))
		z.repSummary.Failure(err, userIn)
		return err
	}
	defer rep.Close()

	ui.Progress(MUser.ProgressScanUser.With("User", userIn.User))

	summary := &UserSummary{
		User: z.user.Email,
	}

	handler := func(event *mo_activity.Event) error {
		rep.Row(event.Compatible())
		switch event.EventCategory {
		case "logins":
			summary.Logins++
		case "devices":
			summary.Devices++
		case "sharing":
			summary.Sharing++
		case "file_operations":
			summary.FileOperations++
		case "paper":
			summary.Paper++
		default:
			summary.Others++
		}
		app_ui.ShowProgress(ui)

		return nil
	}

	err = sv_activity.New(z.ctx).List(handler,
		sv_activity.StartTime(z.StartTime),
		sv_activity.EndTime(z.EndTime),
		sv_activity.Category(z.Category),
		sv_activity.AccountId(z.user.AccountId),
	)
	if err != nil {
		z.repSummary.Failure(err, userIn)
		return err
	}

	z.repSummary.Success(userIn, summary)
	return nil
}

type User struct {
	Peer        dbx_conn.ConnBusinessAudit
	StartTime   mo_time.TimeOptional
	EndTime     mo_time.TimeOptional
	Category    mo_string.OptionalString
	User        rp_model.RowReport
	UserSummary rp_model.TransactionReport
}

func (z *User) Preset() {
	z.User.SetModel(&mo_activity.Compatible{})
	z.UserSummary.SetModel(&UserIn{}, &UserSummary{}, rp_model.HiddenColumns(
		"result.user", // duplicated to `input.user`
	))
}

func (z *User) Exec(c app_control.Control) error {
	if err := z.UserSummary.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	q := c.NewLegacyQueue()
	for _, member := range members {
		q.Enqueue(&UserWorker{
			ctl:        c,
			ctx:        z.Peer.Context(),
			reps:       z.User,
			repSummary: z.UserSummary,
			user:       member,
			StartTime:  z.StartTime.Iso8601(),
			EndTime:    z.EndTime.Iso8601(),
			Category:   z.Category.Value(),
		})
	}
	q.Wait()

	return nil
}

func (z *User) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &User{}, func(r rc_recipe.Recipe) {
		rc := r.(*User)
		if t, ok := rc.StartTime.(*mo_time.TimeImpl); ok {
			t.UpdateTime(time.Now().Add(-10 * time.Minute).Format(dbx_util.DateTimeFormat))
		}
	})
	if err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "user_summary", func(cols map[string]string) error {
		if _, ok := cols["result.logins"]; !ok {
			return errors.New("`logins` is not found")
		}
		return nil
	})
}
