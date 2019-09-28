package cap

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"go.uber.org/zap"
	"time"
)

type ExpiryVO struct {
	Peer app_conn.ConnBusinessFile
	Days int
	At   string
}

type Expiry struct{}

func (z Expiry) Console() {
}

func (z *Expiry) Requirement() app_vo.ValueObject {
	return &ExpiryVO{}
}

func (z *Expiry) Exec(k app_kitchen.Kitchen) error {
	ui := k.UI()
	l := k.Log()
	evo := k.Value().(*ExpiryVO)
	var newExpiry time.Time
	if evo.Days > 0 && evo.At != "" {
		l.Debug("Both Days/At specified", zap.Int("evo.Days", evo.Days), zap.String("evo.At", evo.At))
		ui.Error("recipe.team.sharedlink.cap.expiry.err.please_specify_days_or_at")
		return errors.New("please specify days or at")
	}
	if evo.Days < 0 {
		l.Debug("Days options should not be negative", zap.Int("evo.Days", evo.Days))
		ui.Error("recipe.team.sharedlink.cap.expiry.err.days_should_not_negative")
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
			ui.Error("recipe.team.sharedlink.cap.expiry.err.invalid_date_time_format_for_at_option")
			return errors.New("invalid date/time format for `at`")
		}
	}

	l = l.With(zap.String("newExpiry", newExpiry.String()))

	rep, err := k.Report("updated_sharedlink", &mo_sharedlink.SharedLinkMember{})
	if err != nil {
		return err
	}
	defer rep.Close()

	conn, err := evo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	members, err := sv_member.New(conn).List()
	if err != nil {
		return err
	}

	// TODO: move below logic to uc_* package
	for _, member := range members {
		l.Debug("Scanning member shared links", zap.Any("member", member))

		connMember := conn.AsMemberId(member.TeamMemberId)
		links, err := sv_sharedlink.New(connMember).List()
		if err != nil {
			return err
		}

		for _, link := range links {
			ll := l.With(zap.Any("link", link))
			if link.LinkVisibility() != "public" {
				ll.Debug("Skip non public link")
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

				if le.IsZero() || le.After(newExpiry) {
					ll.Debug("The link have long or no expiration")
					update = true
				}
			}

			if !update {
				ll.Debug("Skip")
				continue
			}

			updated, err := sv_sharedlink.New(connMember).Update(link, sv_sharedlink.Expires(newExpiry))
			if err != nil {
				ll.Warn("Unable to update expiration")
				continue
			}
			ll.Debug("Updated", zap.Any("updated", updated))
			rep.Row(mo_sharedlink.NewSharedLinkMember(updated, member))
		}
	}
	return nil
}

func (z *Expiry) Test(c app_control.Control) error {
	// should fail
	if err := z.Exec(app_kitchen.NewKitchen(c, &ExpiryVO{Days: 1, At: "2019-09-05T01:02:03Z"})); err == nil {
		return errors.New("days and at should not be accepted same time")
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, &ExpiryVO{Days: -1})); err == nil {
		return errors.New("negative days should not be accepted")
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, &ExpiryVO{At: "Invalid time format"})); err == nil {
		return errors.New("invalid time format should not be accepted")
	}
	return nil
}
