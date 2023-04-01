package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/figma/api/fg_conn"
	"github.com/watermint/toolbox/domain/figma/model/mo_file"
	"github.com/watermint/toolbox/domain/figma/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Info struct {
	Peer     fg_conn.ConnFigmaApi
	Key      string
	AllNodes bool
	Document rp_model.RowReport
	Node     rp_model.RowReport
}

func (z *Info) Preset() {
	z.Document.SetModel(
		&mo_file.Document{},
		rp_model.HiddenColumns(
			"thumbnailUrl",
			"document",
		),
	)
	z.Node.SetModel(
		&mo_file.Node{},
		rp_model.HiddenColumns(
			"thumbnailUrl",
			"children",
		),
	)
}

func (z *Info) reportNode(node *mo_file.Node) {
	switch node.Type {
	case "CANVAS", "FRAME":
		z.Node.Row(node)
	default:
		if z.AllNodes {
			z.Node.Row(node)
		}
	}
	for _, c := range node.Children {
		z.reportNode(c)
	}
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.Document.Open(); err != nil {
		return err
	}
	if err := z.Node.Open(); err != nil {
		return err
	}
	if r, m := sv_file.VerifyFileKey(z.Key); r != sv_file.VerifyFileKeyLooksOkay {
		c.UI().Error(m)
		return errors.New(c.UI().Text(m))
	}
	doc, err := sv_file.New(z.Peer.Client()).Info(z.Key)
	if err != nil {
		return err
	}

	z.Document.Row(doc)

	z.reportNode(doc.Document)

	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, func(r rc_recipe.Recipe) {
		m := r.(*Info)
		m.Key = "abc123"
	})
}
