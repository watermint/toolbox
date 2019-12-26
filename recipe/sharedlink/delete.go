package sharedlink

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"go.uber.org/zap"
	"path/filepath"
	"strings"
)

type Delete struct {
	Peer       rc_conn.ConnUserFile
	Path       mo_path.DropboxPath
	Recursive  bool
	SharedLink rp_model.TransactionReport
}

func (z *Delete) Preset() {
	z.SharedLink.SetModel(&mo_sharedlink.Metadata{}, nil)
}

func (z *Delete) Console() {
}

func (z *Delete) Exec(k rc_kitchen.Kitchen) error {
	if err := z.SharedLink.Open(); err != nil {
		return err
	}

	if z.Recursive {
		return z.removeRecursive(k)
	} else {
		return z.removePathAt(k)
	}
}

func (z *Delete) removePathAt(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	l := k.Log()
	links, err := sv_sharedlink.New(z.Peer.Context()).ListByPath(z.Path)
	if err != nil {
		return err
	}
	if len(links) < 1 {
		ui.Info("recipe.sharedlink.delete.info.no_links_at_the_path", app_msg.P{
			"Path": z.Path.Path(),
		})
		return nil
	}

	var lastErr error
	for _, link := range links {
		ui.Info("recipe.sharedlink.delete.progress", app_msg.P{
			"Url":  link.LinkUrl(),
			"Path": link.LinkPathLower(),
		})
		err = sv_sharedlink.New(z.Peer.Context()).Remove(link)
		if err != nil {
			l.Debug("Unable to remove link", zap.Error(err), zap.Any("link", link))
			z.SharedLink.Failure(err, link)
			lastErr = err
		} else {
			z.SharedLink.Success(link, nil)
		}
	}
	return lastErr
}

func (z *Delete) removeRecursive(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	l := k.Log().With(zap.String("path", z.Path.Path()))
	links, err := sv_sharedlink.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	if len(links) < 1 {
		ui.Info("recipe.sharedlink.delete.info.no_links_at_the_path", app_msg.P{
			"Path": z.Path.Path(),
		})
		return nil
	}

	var lastErr error
	for _, link := range links {
		l = l.With(zap.String("linkPath", link.LinkPathLower()))
		rel, err := filepath.Rel(strings.ToLower(z.Path.Path()), link.LinkPathLower())
		if err != nil {
			l.Debug("Skip due to path calc error", zap.Error(err))
			continue
		}
		if strings.HasPrefix(rel, "..") {
			l.Debug("Skip due to non related path")
			continue
		}

		ui.Info("recipe.sharedlink.delete.progress", app_msg.P{
			"Url":  link.LinkUrl(),
			"Path": link.LinkPathLower(),
		})
		err = sv_sharedlink.New(z.Peer.Context()).Remove(link)
		if err != nil {
			l.Debug("Unable to remove link", zap.Error(err), zap.Any("link", link))
			z.SharedLink.Failure(err, link)
			lastErr = err
		} else {
			z.SharedLink.Success(link, nil)
		}
	}
	return lastErr
}

func (z *Delete) Test(c app_control.Control) error {
	return qt_endtoend.ImplementMe()
}
