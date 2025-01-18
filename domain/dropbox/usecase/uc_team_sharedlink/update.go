package uc_team_sharedlink

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

// Return true to include target
type UpdateFilter func(target *Target) bool

// Create update options for the target
type UpdateCreateOpts func(target *Target) (opts []sv_sharedlink.LinkOpt)

// Callback on skip
type UpdateOnSkip func(target *Target)

type UpdateOnFailure func(target *Target, err error)

type UpdateOnSuccess func(target *Target, updated mo_sharedlink.SharedLink)

type UpdateOpts struct {
	Filter    UpdateFilter
	Opts      UpdateCreateOpts
	OnSkip    UpdateOnSkip
	OnMissing SelectorOnMissing
	OnSuccess UpdateOnSuccess
	OnFailure UpdateOnFailure
}

func Update(target *Target, c app_control.Control, ctx dbx_client.Client, sel Selector, opts UpdateOpts, baseNamespace dbx_filesystem.BaseNamespaceType) error {
	l := c.Log().With(esl.String("member", target.Member.Email), esl.String("url", target.Entry.Url))
	mc := ctx.AsMemberId(target.Member.TeamMemberId, baseNamespace)

	defer func() {
		_ = sel.Processed(target.Entry.Url)
	}()

	if !opts.Filter(target) {
		l.Debug("Skipped")
		opts.OnSkip(target)
		return nil
	}

	updated, err := sv_sharedlink.New(mc).Update(target.Entry.SharedLink(), opts.Opts(target)...)
	if err != nil {
		l.Debug("Unable to update visibility of the link", esl.Error(err))
		opts.OnFailure(target, err)
		return err
	}

	l.Debug("Updated", esl.Any("updated", updated))
	opts.OnSuccess(target, updated)
	return nil
}
