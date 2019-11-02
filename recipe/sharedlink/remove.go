package sharedlink

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type RemoveVO struct {
	Peer      app_conn.ConnUserFile
	Path      string
	Recursive bool
}

const (
	reportRemove = "link"
)

type Remove struct {
}

func (z *Remove) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportRemove, rp_model.TransactionHeader(&mo_sharedlink.Metadata{}, nil)),
	}
}

func (z *Remove) Console() {
}

func (z *Remove) Requirement() app_vo.ValueObject {
	return &RemoveVO{}
}

func (z *Remove) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*RemoveVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	if vo.Recursive {
		return z.removeRecursive(k, ctx, vo.Path)
	} else {
		return z.removePathAt(k, ctx, vo.Path)
	}
}

func (z *Remove) removePathAt(k app_kitchen.Kitchen, ctx api_context.Context, path string) error {
	ui := k.UI()
	l := k.Log()
	links, err := sv_sharedlink.New(ctx).ListByPath(mo_path.NewPath(path))
	if err != nil {
		return err
	}
	if len(links) < 1 {
		ui.Info("recipe.sharedlink.remove.info.no_links_at_the_path", app_msg.P{
			"Path": path,
		})
		return nil
	}
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportRemove)
	if err != nil {
		return err
	}
	defer rep.Close()

	var lastErr error
	for _, link := range links {
		ui.Info("recipe.sharedlink.remove.progress", app_msg.P{
			"Url":  link.LinkUrl(),
			"Path": link.LinkPathLower(),
		})
		err = sv_sharedlink.New(ctx).Remove(link)
		if err != nil {
			l.Debug("Unable to remove link", zap.Error(err), zap.Any("link", link))
			msg := app_msg.M("recipe.sharedlink.remove.err.unable_to_remove", app_msg.P{
				"Error": err.Error(),
			})
			rep.Failure(msg, link, nil)
			lastErr = err
		} else {
			rep.Success(link, nil)
		}
	}
	return lastErr
}

func (z *Remove) removeRecursive(k app_kitchen.Kitchen, ctx api_context.Context, path string) error {
	ui := k.UI()
	l := k.Log().With(zap.String("path", path))
	links, err := sv_sharedlink.New(ctx).List()
	if err != nil {
		return err
	}
	if len(links) < 1 {
		ui.Info("recipe.sharedlink.remove.info.no_links_at_the_path", app_msg.P{
			"Path": path,
		})
		return nil
	}
	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportRemove)
	if err != nil {
		return err
	}
	defer rep.Close()

	var lastErr error
	for _, link := range links {
		l = l.With(zap.String("linkPath", link.LinkPathLower()))
		rel, err := filepath.Rel(strings.ToLower(path), link.LinkPathLower())
		if err != nil {
			l.Debug("Skip due to path calc error", zap.Error(err))
			continue
		}
		if strings.HasPrefix(rel, "..") {
			l.Debug("Skip due to non related path")
			continue
		}

		ui.Info("recipe.sharedlink.remove.progress", app_msg.P{
			"Url":  link.LinkUrl(),
			"Path": link.LinkPathLower(),
		})
		err = sv_sharedlink.New(ctx).Remove(link)
		if err != nil {
			l.Debug("Unable to remove link", zap.Error(err), zap.Any("link", link))
			msg := app_msg.M("recipe.sharedlink.remove.err.unable_to_remove", app_msg.P{
				"Error": err.Error(),
			})
			rep.Failure(msg, link, nil)
			lastErr = err
		} else {
			rep.Success(link, nil)
		}
	}
	return lastErr
}

func (z *Remove) Test(c app_control.Control) error {
	return qt_test.ImplementMe()
}
