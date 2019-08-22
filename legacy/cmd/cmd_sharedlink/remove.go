package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type CmdSharedLinkRemove struct {
	*cmd2.SimpleCommandlet
	report       app_report.Factory
	optPath      string
	optRecursive bool
}

func (z *CmdSharedLinkRemove) Name() string {
	return "remove"
}

func (z *CmdSharedLinkRemove) Desc() string {
	return "cmd.sharedlink.remove.desc"
}

func (z *CmdSharedLinkRemove) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdSharedLinkRemove) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descPath := z.ExecContext.Msg("cmd.sharedlink.remove.flag.path").T()
	f.StringVar(&z.optPath, "path", "", descPath)

	descRecursive := z.ExecContext.Msg("cmd.sharedlink.remove.flag.recursive").T()
	f.BoolVar(&z.optRecursive, "recursive", false, descRecursive)
}

func (z *CmdSharedLinkRemove) Exec(args []string) {
	if z.optPath == "" {
		z.ExecContext.Msg("cmd.sharedlink.remove.err.not_enough_arguments").TellError()
		return
	}

	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	if z.optRecursive {
		z.removeRecursive(ctx, mo_path.NewPath(z.optPath))
	} else {
		z.removePathAt(ctx, mo_path.NewPath(z.optPath))
	}
}

func (z *CmdSharedLinkRemove) removePathAt(ctx api_context.Context, path mo_path.Path) {
	svc := sv_sharedlink.New(ctx)
	links, err := svc.ListByPath(mo_path.NewPath(z.optPath))
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	if len(links) < 1 {
		z.ExecContext.Msg("cmd.sharedlink.remove.err.no_link_found").TellError()
		return
	}
	for _, link := range links {
		log := z.ExecContext.Log().With(zap.String("linkPath", link.LinkPathLower()))
		log.Debug("Removing")
		err := svc.Remove(link)
		if err != nil {
			log.Debug("Failed", zap.Error(err))
			api_util.UIMsgFromError(err).TellError()
			continue
		}
		z.report.Report(link)
	}
}

func (z *CmdSharedLinkRemove) removeRecursive(ctx api_context.Context, path mo_path.Path) {
	svc := sv_sharedlink.New(ctx)
	links, err := svc.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	if len(links) < 1 {
		z.ExecContext.Msg("cmd.sharedlink.remove.err.no_link_found").TellError()
		return
	}
	for _, link := range links {
		log := z.ExecContext.Log().With(zap.String("linkPath", link.LinkPathLower()))
		rel, err := filepath.Rel(strings.ToLower(path.Path()), link.LinkPathLower())
		if err != nil {
			log.Debug("Skip", zap.Error(err))
			continue
		}
		if strings.HasPrefix(rel, "..") {
			log.Debug("Skip")
			continue
		}
		log.Debug("Removing")
		err = svc.Remove(link)
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			continue
		}
		log.Debug("Removed")
		z.report.Report(link)
	}
}
