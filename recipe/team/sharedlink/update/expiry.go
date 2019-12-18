package update

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/rc_conn"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"time"
)

type ExpiryVO struct {
	Peer       rc_conn.ConnBusinessFile
	Days       int
	At         string
	Visibility string
}

type ExpiryScanWorker struct {
	k          rc_kitchen.Kitchen
	ctl        app_control.Control
	ctx        api_context.Context
	rep        rp_model.Report
	repSkipped rp_model.Report
	member     *mo_member.Member
	newExpiry  time.Time
	visibility string
}

func (z *ExpiryScanWorker) Exec() error {
	ui := z.ctl.UI()
	l := z.ctl.Log().With(zap.Any("member", z.member))

	l.Debug("Scanning member shared links")
	ui.Info("recipe.team.sharedlink.update.expiry.scan", app_msg.P{"MemberEmail": z.member.Email})

	ctxMember := z.ctx.AsMemberId(z.member.TeamMemberId)
	links, err := sv_sharedlink.New(ctxMember).List()
	if err != nil {
		l.Debug("Unable to scan shared link", zap.Error(err))
		ui.ErrorM(api_util.MsgFromError(err))
		return err
	}

	q := z.k.NewQueue()

	for _, link := range links {
		ll := l.With(zap.Any("link", link))
		if link.LinkVisibility() != z.visibility {
			ll.Debug("Skip link", zap.String("targetVisibility", z.visibility))
			z.repSkipped.Row(mo_sharedlink.NewSharedLinkMember(link, z.member))
			continue
		}

		update := false

		switch {
		case link.LinkExpires() == "":
			ll.Debug("The link doesn't have expiration")
			update = true

		default:
			le, v := ut_time.ParseTimestamp(link.LinkExpires())
			if !v {
				ll.Warn("Invalid timestamp format from API response")
				continue
			}

			if le.IsZero() || le.After(z.newExpiry) {
				ll.Debug("The link have long or no expiration")
				update = true
			}
		}

		if !update {
			z.repSkipped.Row(mo_sharedlink.NewSharedLinkMember(link, z.member))
			ll.Debug("Skip")
			continue
		}

		q.Enqueue(&ExpiryWorker{
			ctl:       z.ctl,
			ctx:       ctxMember,
			rep:       z.rep,
			member:    z.member,
			link:      link,
			newExpiry: z.newExpiry,
		})
	}
	q.Wait()

	return nil
}

type ExpiryWorker struct {
	ctl       app_control.Control
	ctx       api_context.Context
	rep       rp_model.Report
	member    *mo_member.Member
	link      mo_sharedlink.SharedLink
	newExpiry time.Time
}

func (z *ExpiryWorker) Exec() error {
	ui := z.ctl.UI()
	l := z.ctl.Log().With(zap.Any("link", z.link.Metadata()))

	ui.Info("recipe.team.sharedlink.update.expiry.updating", app_msg.P{
		"MemberEmail":   z.member.Email,
		"Url":           z.link.LinkUrl(),
		"CurrentExpiry": z.link.LinkExpires(),
		"NewExpiry":     api_util.RebaseAsString(z.newExpiry),
	})

	updated, err := sv_sharedlink.New(z.ctx).Update(z.link, sv_sharedlink.Expires(z.newExpiry))
	if err != nil {
		l.Debug("Unable to update expiration")
		z.rep.Failure(err, mo_sharedlink.NewSharedLinkMember(z.link, z.member))
		return err
	}

	l.Debug("Updated", zap.Any("updated", updated))
	z.rep.Success(
		mo_sharedlink.NewSharedLinkMember(z.link, z.member),
		updated,
	)

	return nil
}

const (
	reportExpiryUpdated = "updated_sharedlink"
	reportExpirySkipped = "skipped_sharedlink"
)

type Expiry struct {
}

func (z Expiry) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportExpirySkipped,
			&mo_sharedlink.SharedLinkMember{},
			rp_model.HiddenColumns(
				"shared_link_id",
				"account_id",
				"team_member_id",
			),
		),
		rp_spec_impl.Spec(reportExpiryUpdated,
			rp_model.TransactionHeader(
				&mo_sharedlink.SharedLinkMember{},
				&mo_sharedlink.Metadata{},
			),
			rp_model.HiddenColumns(
				"input.shared_link_id",
				"input.account_id",
				"input.team_member_id",
				"result.tag",
				"result.id",
				"result.url",
				"result.name",
				"result.path_lower",
				"result.visibility",
			),
		),
	}
}

func (z Expiry) Console() {
}

func (z *Expiry) Requirement() rc_vo.ValueObject {
	return &ExpiryVO{
		Visibility: "public",
	}
}

func (z *Expiry) Exec(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	l := k.Log()
	evo := k.Value().(*ExpiryVO)
	var newExpiry time.Time
	if evo.Days > 0 && evo.At != "" {
		l.Debug("Both Days/At specified", zap.Int("evo.Days", evo.Days), zap.String("evo.At", evo.At))
		ui.Error("recipe.team.sharedlink.update.expiry.err.please_specify_days_or_at")
		return errors.New("please specify days or at")
	}
	if evo.Days < 0 {
		l.Debug("Days options should not be negative", zap.Int("evo.Days", evo.Days))
		ui.Error("recipe.team.sharedlink.update.expiry.err.days_should_not_negative")
		return errors.New("days should not be negative")
	}

	switch {
	case evo.Days > 0:
		newExpiry = api_util.RebaseTime(time.Now().Add(time.Duration(evo.Days*24) * time.Hour))
		l.Debug("New expiry", zap.Int("evo.Days", evo.Days), zap.String("newExpiry", newExpiry.String()))

	default:
		var valid bool
		if newExpiry, valid = ut_time.ParseTimestamp(evo.At); !valid {
			l.Debug("Invalid date/time format for at option", zap.String("evo.At", evo.At))
			ui.Error("recipe.team.sharedlink.update.expiry.err.invalid_date_time_format_for_at_option")
			return errors.New("invalid date/time format for `at`")
		}
	}

	l = l.With(zap.String("newExpiry", newExpiry.String()))

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportExpiryUpdated)
	if err != nil {
		return err
	}
	defer rep.Close()
	repSkipped, err := rp_spec_impl.New(z, k.Control()).Open(reportExpirySkipped)
	if err != nil {
		return err
	}
	defer repSkipped.Close()

	ctx, err := evo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(ctx).List()
	if err != nil {
		return err
	}

	q := k.NewQueue()

	for _, member := range members {
		q.Enqueue(&ExpiryScanWorker{
			k:          k,
			ctl:        k.Control(),
			ctx:        ctx,
			rep:        rep,
			repSkipped: repSkipped,
			member:     member,
			newExpiry:  newExpiry,
			visibility: evo.Visibility,
		})
	}
	q.Wait()

	return nil
}

func (z *Expiry) Test(c app_control.Control) error {
	// should fail
	if err := z.Exec(rc_kitchen.NewKitchen(c, &ExpiryVO{Days: 1, At: "2019-09-05T01:02:03Z"})); err == nil {
		return errors.New("days and at should not be accepted same time")
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, &ExpiryVO{Days: -1})); err == nil {
		return errors.New("negative days should not be accepted")
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, &ExpiryVO{At: "Invalid time format"})); err == nil {
		return errors.New("invalid time format should not be accepted")
	}
	return qt_recipe.ImplementMe()
}
