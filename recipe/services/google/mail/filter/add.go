package filter

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_filter"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_filter"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"strings"
)

type Add struct {
	Peer                   goog_conn.ConnGoogleMail
	Filter                 rp_model.RowReport
	UserId                 string
	CriteriaFrom           mo_string.OptionalString
	CriteriaTo             mo_string.OptionalString
	CriteriaQuery          mo_string.OptionalString
	CriteriaNegatedQuery   mo_string.OptionalString
	CriteriaHasAttachment  bool
	CriteriaNoAttachment   bool
	CriteriaExcludeChats   bool
	CriteriaSize           int
	CriteriaSizeComparison mo_string.OptionalString
	AddLabels              mo_string.OptionalString
	RemoveLabels           mo_string.OptionalString
	Forward                mo_string.OptionalString
	AddLabelIfNotExist     bool
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailModify,
		goog_auth.ScopeGmailSettingsBasic,
	)
	z.Filter.SetModel(&mo_filter.Filter{})
	z.UserId = "me"
}

func (z *Add) Exec(c app_control.Control) error {
	opts := make([]sv_filter.Opt, 0)
	if z.CriteriaFrom.IsExists() {
		opts = append(opts, sv_filter.From(z.CriteriaFrom.Value()))
	}
	if z.CriteriaTo.IsExists() {
		opts = append(opts, sv_filter.To(z.CriteriaTo.Value()))
	}
	if z.CriteriaQuery.IsExists() {
		opts = append(opts, sv_filter.Query(z.CriteriaQuery.Value()))
	}
	if z.CriteriaNegatedQuery.IsExists() {
		opts = append(opts, sv_filter.NegatedQuery(z.CriteriaNegatedQuery.Value()))
	}
	if z.CriteriaHasAttachment {
		opts = append(opts, sv_filter.HasAttachment(true))
	}
	if z.CriteriaNoAttachment {
		opts = append(opts, sv_filter.HasAttachment(false))
	}
	if z.CriteriaSize > 0 {
		opts = append(opts, sv_filter.Size(z.CriteriaSize))
	}
	if z.CriteriaSizeComparison.IsExists() {
		opts = append(opts, sv_filter.SizeComparison(z.CriteriaSizeComparison.Value()))
	}
	if z.AddLabels.IsExists() {
		labelNames := strings.Split(z.AddLabels.Value(), ",")
		var labelIds []string
		var err error
		if z.AddLabelIfNotExist {
			labelIds, err = sv_label.FindOrAddLabelIdsByNames(z.Peer.Context(), c.UI(), z.UserId, labelNames)
		} else {
			labelIds, err = sv_label.FindLabelIdsByNames(z.Peer.Context(), c.UI(), z.UserId, labelNames)
		}
		if err != nil {
			return err
		}
		opts = append(opts, sv_filter.AddLabelIds(labelIds))
	}
	if z.RemoveLabels.IsExists() {
		labelNames := strings.Split(z.RemoveLabels.Value(), ",")
		var labelIds []string
		var err error
		if z.AddLabelIfNotExist {
			labelIds, err = sv_label.FindOrAddLabelIdsByNames(z.Peer.Context(), c.UI(), z.UserId, labelNames)
		} else {
			labelIds, err = sv_label.FindLabelIdsByNames(z.Peer.Context(), c.UI(), z.UserId, labelNames)
		}
		if err != nil {
			return err
		}
		opts = append(opts, sv_filter.RemoveLabelIds(labelIds))
	}

	if err := z.Filter.Open(); err != nil {
		return nil
	}

	label, err := sv_filter.New(z.Peer.Context(), z.UserId).Add(opts...)
	if err != nil {
		return err
	}
	z.Filter.Row(label)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &Add{}, "recipe-services-google-mail-filter-add.json.gz", func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.AddLabels = mo_string.NewOptional("xxxxx")
		m.CriteriaQuery = mo_string.NewOptional("from:@google.com")
	})
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.CriteriaFrom = mo_string.NewOptional("@gmail.com")
		m.Forward = mo_string.NewOptional("toolbox@example.com")
	})
}
