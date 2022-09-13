package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_activity"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/essentials/strings/es_mailaddr"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"time"
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

type User struct {
	Peer        dbx_conn.ConnScopedTeam
	StartTime   mo_time.TimeOptional
	EndTime     mo_time.TimeOptional
	Category    mo_string.OptionalString
	User        rp_model.RowReport
	UserSummary rp_model.TransactionReport
}

func (z *User) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeEventsRead,
		dbx_auth.ScopeMembersRead,
	)
	z.User.SetModel(&mo_activity.Compatible{})
	z.UserSummary.SetModel(&UserIn{}, &UserSummary{}, rp_model.HiddenColumns(
		"result.user", // duplicated to `input.user`
	))
}

func (z *User) activity(member *mo_member.Member, c app_control.Control) error {
	userIn := &UserIn{User: member.Email}
	l := c.Log().With(esl.Any("userIn", userIn))

	rep, err := z.User.OpenNew(
		rp_model.Suffix("-"+es_mailaddr.EscapeSpecial(member.Email, "_")),
		rp_model.NoConsoleOutput(),
	)
	if err != nil {
		l.Debug("unable to create report", esl.Error(err))
		z.UserSummary.Failure(err, userIn)
		return err
	}
	defer rep.Close()

	summary := &UserSummary{
		User: member.Email,
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

		return nil
	}

	err = sv_activity.New(z.Peer.Client()).List(handler,
		sv_activity.StartTime(z.StartTime.Iso8601()),
		sv_activity.EndTime(z.EndTime.Iso8601()),
		sv_activity.Category(z.Category.Value()),
		sv_activity.AccountId(member.AccountId),
	)
	if err != nil {
		z.UserSummary.Failure(err, userIn)
		return err
	}

	z.UserSummary.Success(userIn, summary)
	return nil
}

func (z *User) Exec(c app_control.Control) error {
	if err := z.UserSummary.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("activity", z.activity, c)
		q := s.Get("activity")
		for _, member := range members {
			q.Enqueue(member)
		}
	})
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
