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
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"time"
)

type UserVO struct {
	Peer      app_conn.ConnBusinessAudit
	StartTime string
	EndTime   string
	Category  string
}

const (
	reportEventUser        = "user_activity"
	reportEventUserSummary = "user_summary"
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
	k          app_kitchen.Kitchen
	ctx        api_context.Context
	reps       *rp_spec_impl.Specs
	repSummary rp_model.Report
	user       *mo_member.Member
	vo         *UserVO
}

func (z *UserWorker) Exec() error {
	userIn := &UserIn{User: z.user.Email}
	ui := z.k.UI()
	l := z.k.Log().With(zap.Any("userIn", userIn))
	rep, err := z.reps.Open(reportEventUser, rp_model.Suffix("-"+z.user.Email))
	if err != nil {
		l.Debug("unable to create report", zap.Error(err))
		z.repSummary.Failure(err, userIn)
		return err
	}

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
		sv_activity.StartTime(z.vo.StartTime),
		sv_activity.EndTime(z.vo.EndTime),
		sv_activity.Category(z.vo.Category),
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
}

func (z *User) Console() {
}

func (z *User) Requirement() app_vo.ValueObject {
	return &UserVO{}
}

func (z *User) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*UserVO)

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	reps := rp_spec_impl.New(z, k.Control())
	repSummary, err := reps.Open(reportEventUserSummary)
	if err != nil {
		return err
	}
	defer repSummary.Close()

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}

	q := k.NewQueue()
	for _, member := range members {
		q.Enqueue(&UserWorker{
			k:          k,
			ctx:        ctx,
			reps:       reps,
			repSummary: repSummary,
			user:       member,
			vo:         vo,
		})
	}
	q.Wait()

	return nil
}

func (z *User) Test(c app_control.Control) error {
	lvo := &UserVO{
		StartTime: api_util.RebaseAsString(time.Now().Add(-10 * time.Minute)),
	}
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, reportEventUserSummary, func(cols map[string]string) error {
		if _, ok := cols["logins"]; !ok {
			return errors.New("`logins` is not found")
		}
		return nil
	})
}

func (z *User) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(
			reportEventUser,
			&mo_activity.Event{},
		),
		rp_spec_impl.Spec(
			reportEventUserSummary,
			rp_model.TransactionHeader(
				&UserIn{},
				&UserSummary{},
			),
			rp_model.HiddenColumns(
				"result.user", // duplicated to `input.user`
			),
		),
	}
}
