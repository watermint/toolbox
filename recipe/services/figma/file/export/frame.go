package export

import (
	"errors"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/model/mo_file"
	"github.com/watermint/toolbox/domain/figma/service/sv_file"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Frame struct {
	Peer      fg_conn.ConnFigmaApi
	Key       string
	Scale     mo_int.RangeInt
	Format    mo_string.SelectString
	Path      mo_path.ExistingFileSystemPath
	FrameInfo app_msg.Message
}

func (z *Frame) Preset() {
	z.Scale.SetRange(sv_file.ImageScaleMin, sv_file.ImageScaleMax, sv_file.ImageScaleDefault)
	z.Format.SetOptions("pdf", sv_file.ImageFormats...)
}

func (z *Frame) download(page *mo_file.NodeWithPath, doc *mo_file.Document, c app_control.Control) error {
	svf := sv_file.New(z.Peer.Client())
	urls, err := svf.Image(z.Key, page.Node.Id, z.Scale.Value(), z.Format.Value())
	if err != nil {
		return err
	}
	url, ok := urls[page.Node.Id]
	if !ok || url == "" {
		return errors.New("no image was generated on the Figma side")
	}

	path := filepath.Join(z.Path.Path(), es_filepath.Escape(doc.Name+"-"+page.Path(" ", "_"))+"."+z.Format.Value())

	return es_download.Download(c.Log(), url, path)
}

func (z *Frame) Exec(c app_control.Control) error {
	svf := sv_file.New(z.Peer.Client())
	doc, err := svf.Info(z.Key)
	if err != nil {
		return err
	}

	frames := doc.NodesWithPathByType("FRAME")

	c.UI().Info(z.FrameInfo.With("Frame", len(frames)))
	var lastErr error
	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("frame", z.download, doc, c)
		q := s.Get("frame")
		for _, page := range frames {
			q.Enqueue(&page)
		}
	}, eq_sequence.SingleThread(),
		eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
			lastErr = err
		}),
	)

	return lastErr
}

func (z *Frame) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("frame", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()

	return rc_exec.ExecMock(c, &Frame{}, func(r rc_recipe.Recipe) {
		m := r.(*Frame)
		m.Key = "abc"
		m.Path = mo_path.NewExistingFileSystemPath(p)
	})
}
