package activity

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_activity"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_activity"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_mailaddr"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
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

type UserWorker struct {
	k          rc_kitchen.Kitchen
	ctx        api_context.Context
	reps       rp_model.RowReport
	repSummary rp_model.TransactionReport
	user       *mo_member.Member
	StartTime  string
	EndTime    string
	Category   string
}

func (z *UserWorker) Exec() error {
	userIn := &UserIn{User: z.user.Email}
	ui := z.k.UI()
	l := z.k.Log().With(zap.Any("userIn", userIn))

	rep, err := z.reps.OpenNew(rp_model.Suffix("-" + ut_mailaddr.EscapeSpecial(z.user.Email, "_")))
	if err != nil {
		l.Debug("unable to create report", zap.Error(err))
		z.repSummary.Failure(err, userIn)
		return err
	}
	defer rep.Close()

	ui.Info("recipe.team.activity.user.progress.scan_user", app_msg.P{
		"User": userIn.User,
	})

	summary := &UserSummary{
		User: z.user.Email,
	}

	handler := func(event *mo_activity.Event) error {
		rep.Row(event)
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
	Peer        rc_conn.ConnBusinessAudit
	StartTime   string
	EndTime     string
	Category    string
	User        rp_model.RowReport
	UserSummary rp_model.TransactionReport
}

func (z *User) Preset() {
	z.User.SetModel(&mo_activity.Event{})
	z.UserSummary.SetModel(&UserIn{}, &UserSummary{}, rp_model.HiddenColumns(
		"result.user", // duplicated to `input.user`
	))
}

func (z *User) Console() {
}

func (z *User) Exec(k rc_kitchen.Kitchen) error {
	l := k.Log()

	if z.StartTime != "" {
		if t, ok := ut_time.ParseTimestamp(z.StartTime); ok {
			l.Debug("Rebase StartTime", zap.String("startTime", z.StartTime))
			z.StartTime = api_util.RebaseAsString(t)
			l.Debug("Rebased StartTime", zap.String("startTime", z.StartTime))
		} else {
			return errors.New("invalid date/time format for -start-date")
		}
	}
	if z.EndTime != "" {
		if t, ok := ut_time.ParseTimestamp(z.EndTime); ok {
			l.Debug("Rebase EndTime", zap.String("endTime", z.StartTime))
			z.StartTime = api_util.RebaseAsString(t)
			l.Debug("Rebased EndTime", zap.String("endTime", z.StartTime))
		} else {
			return errors.New("invalid date/time format for -end-date")
		}
	}

	if err := z.UserSummary.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	q := k.NewQueue()
	for _, member := range members {
		q.Enqueue(&UserWorker{
			k:          k,
			ctx:        z.Peer.Context(),
			reps:       z.User,
			repSummary: z.UserSummary,
			user:       member,
			StartTime:  z.StartTime,
			EndTime:    z.EndTime,
			Category:   z.Category,
		})
	}
	q.Wait()

	return nil
}

func (z *User) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &User{}, func(r rc_recipe.Recipe) {
		rc := r.(*User)
		rc.StartTime = api_util.RebaseAsString(time.Now().Add(-10 * time.Minute))
	})
	if err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "user_summary", func(cols map[string]string) error {
		if _, ok := cols["result.logins"]; !ok {
			return errors.New("`logins` is not found")
		}
		return nil
	})
}
