package export

import (
	"errors"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/service/sv_file"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Node struct {
	Peer            fg_conn.ConnFigmaApi
	Key             string
	Id              string
	Scale           mo_int.RangeInt
	Format          mo_string.SelectString
	Path            mo_path.ExistingFileSystemPath
	TryDownloadUrl  app_msg.Message
	TryDownloadPath app_msg.Message
	NodeNotFound    app_msg.Message
}

func (z *Node) Preset() {
	z.Scale.SetRange(sv_file.ImageScaleMin, sv_file.ImageScaleMax, sv_file.ImageScaleDefault)
	z.Format.SetOptions("pdf", sv_file.ImageFormats...)
}

func (z *Node) Exec(c app_control.Control) error {
	svf := sv_file.New(z.Peer.Client())
	doc, err := svf.Info(z.Key)
	if err != nil {
		return err
	}

	node, found := doc.FindById(z.Id)
	if !found {
		c.UI().Error(z.NodeNotFound.With("Id", z.Id))
		return errors.New("node not found")
	}

	urls, err := svf.Image(z.Key, z.Id, z.Scale.Value(), z.Format.Value())
	if err != nil {
		return err
	}
	url, ok := urls[z.Id]
	if !ok || url == "" {
		return errors.New("no image was generated on the Figma side")
	}

	path := filepath.Join(z.Path.Path(), es_filepath.Escape(node.Name)+"."+z.Format.Value())

	c.UI().Progress(z.TryDownloadUrl.With("Url", url))
	c.UI().Progress(z.TryDownloadPath.With("Path", path))
	return es_download.Download(c.Log(), url, path)
}

func (z *Node) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("node", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()

	return rc_exec.ExecMock(c, &Node{}, func(r rc_recipe.Recipe) {
		m := r.(*Node)
		m.Key = "abc"
		m.Id = "0:1"
		m.Path = mo_path.NewExistingFileSystemPath(p)
	})
}
