package update

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"time"
)

type ExpiryScanWorker struct {
	ctl        app_control.Control
	ctx        dbx_context.Context
	rep        rp_model.TransactionReport
	repSkipped rp_model.RowReport
	member     *mo_member.Member
	newExpiry  time.Time
	visibility string
}

func (z *ExpiryScanWorker) Exec() error {
	ui := z.ctl.UI()
	l := z.ctl.Log().With(zap.Any("member", z.member))

	l.Debug("Scanning member shared links")
	ui.InfoK("recipe.team.sharedlink.update.expiry.scan", app_msg.P{"MemberEmail": z.member.Email})

	ctxMember := z.ctx.AsMemberId(z.member.TeamMemberId)
	links, err := sv_sharedlink.New(ctxMember).List()
	if err != nil {
		l.Debug("Unable to scan shared link", zap.Error(err))
		ui.Error(dbx_util.MsgFromError(err))
		return err
	}

	q := z.ctl.NewQueue()

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
	ctx       dbx_context.Context
	rep       rp_model.TransactionReport
	member    *mo_member.Member
	link      mo_sharedlink.SharedLink
	newExpiry time.Time
}

func (z *ExpiryWorker) Exec() error {
	ui := z.ctl.UI()
	l := z.ctl.Log().With(zap.Any("link", z.link.Metadata()))

	ui.InfoK("recipe.team.sharedlink.update.expiry.updating", app_msg.P{
		"MemberEmail":   z.member.Email,
		"Url":           z.link.LinkUrl(),
		"CurrentExpiry": z.link.LinkExpires(),
		"NewExpiry":     dbx_util.RebaseAsString(z.newExpiry),
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

type Expiry struct {
	Peer       dbx_conn.ConnBusinessFile
	Days       int
	At         mo_time.TimeOptional
	Visibility string
	Updated    rp_model.TransactionReport
	Skipped    rp_model.RowReport
}

func (z *Expiry) Preset() {
	z.Visibility = "public"
	z.Skipped.SetModel(&mo_sharedlink.SharedLinkMember{}, rp_model.HiddenColumns(
		"shared_link_id",
		"account_id",
		"team_member_id",
	))
	z.Updated.SetModel(&mo_sharedlink.SharedLinkMember{}, &mo_sharedlink.Metadata{}, rp_model.HiddenColumns(
		"input.shared_link_id",
		"input.account_id",
		"input.team_member_id",
		"result.tag",
		"result.id",
		"result.url",
		"result.name",
		"result.path_lower",
		"result.visibility",
	))
}

func (z *Expiry) Exec(c app_control.Control) error {
	ui := c.UI()
	l := c.Log()
	var newExpiry time.Time
	if z.Days > 0 && z.At.Ok() {
		l.Debug("Both Days/At specified", zap.Int("evo.Days", z.Days), zap.String("evo.At", z.At.String()))
		ui.ErrorK("recipe.team.sharedlink.update.expiry.err.please_specify_days_or_at")
		return errors.New("please specify one of `-days` or `-at`")
	}
	if z.Days < 0 {
		l.Debug("Days options should not be negative", zap.Int("evo.Days", z.Days))
		ui.ErrorK("recipe.team.sharedlink.update.expiry.err.days_should_not_negative")
		return errors.New("days should not be negative")
	}

	switch {
	case z.Days > 0:
		newExpiry = dbx_util.RebaseTime(time.Now().Add(time.Duration(z.Days*24) * time.Hour))
		l.Debug("New expiry", zap.Int("evo.Days", z.Days), zap.String("newExpiry", newExpiry.String()))

	default:
		if !z.At.Ok() {
			l.Debug("Invalid date/time format for at option", zap.String("evo.At", z.At.String()))
			ui.ErrorK("recipe.team.sharedlink.update.expiry.err.invalid_date_time_format_for_at_option")
			return errors.New("invalid date/time format for `at`")
		}
		newExpiry = z.At.Time()
	}

	l = l.With(zap.String("newExpiry", newExpiry.String()))

	if err := z.Updated.Open(); err != nil {
		return err
	}
	if err := z.Skipped.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	q := c.NewQueue()

	for _, member := range members {
		q.Enqueue(&ExpiryScanWorker{
			ctl:        c,
			ctx:        z.Peer.Context(),
			rep:        z.Updated,
			repSkipped: z.Skipped,
			member:     member,
			newExpiry:  newExpiry,
			visibility: z.Visibility,
		})
	}
	q.Wait()

	return nil
}

func (z *Expiry) Test(c app_control.Control) error {
	// should fail
	{
		err := rc_exec.Exec(c, &Expiry{}, func(r rc_recipe.Recipe) {
			rc := r.(*Expiry)
			rc.Days = 1
			rc.At = mo_time.NewOptional(time.Now().Add(1 * time.Second))
		})
		if err == nil {
			return errors.New("days and at should not be accepted same time")
		}
	}

	{
		err := rc_exec.Exec(c, &Expiry{}, func(r rc_recipe.Recipe) {
			rc := r.(*Expiry)
			rc.Days = -1
		})
		if err == nil {
			return errors.New("negative days should not be accepted")
		}
	}

	{
		err := rc_exec.ExecMock(c, &Expiry{}, func(r rc_recipe.Recipe) {
			m := r.(*Expiry)
			m.Days = 7
		})
		if e, _ := qt_recipe.RecipeError(c.Log(), err); e != nil {
			return e
		}
	}

	{
		err := rc_exec.ExecMock(c, &Expiry{}, func(r rc_recipe.Recipe) {
			m := r.(*Expiry)
			m.At = mo_time.NewOptional(time.Now().Add(1 * time.Second))
		})
		if e, _ := qt_recipe.RecipeError(c.Log(), err); e != nil {
			return e
		}
	}

	return nil
}
