package doc

import (
	"fmt"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"strings"
)

type Markdown struct {
	rc_recipe.RemarkSecret
	Content da_text.TextInput
}

func (z *Markdown) Preset() {
}

func (z *Markdown) Exec(c app_control.Control) error {
	l := c.Log()
	content, err := z.Content.Content()
	if err != nil {
		return err
	}
	parserExtensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(parserExtensions)
	markdown := p.Parse(content)
	l.Info("Markdown parsed", esl.Any("markdown", markdown))

	contentsById := make(map[string]string)

	isContainer := func(node ast.Node) (*ast.Container, bool) {
		return node.AsContainer(), node.AsLeaf() == nil && node.AsContainer() != nil
	}
	containerId := func(container *ast.Container) string {
		if container == nil {
			return ""
		} else if container.Attribute == nil || container.ID == nil || len(container.ID) == 0 {
			return ""
		} else {
			return string(container.ID)
		}
	}
	leafId := func(leaf *ast.Leaf) string {
		if leaf == nil {
			return ""
		} else if leaf.Attribute == nil || leaf.ID == nil || len(leaf.ID) == 0 {
			return ""
		} else {
			return string(leaf.ID)
		}
	}
	isLeaf := func(node ast.Node) (*ast.Leaf, bool) {
		return node.AsLeaf(), node.AsLeaf() != nil
	}

	var walkMarkdown func(node ast.Node, path []string)
	walkMarkdown = func(node ast.Node, path []string) {
		if container, ok := isContainer(node); ok {
			cId := containerId(container)
			l.Debug("Container", esl.Any("path", path), esl.Any("container", container))
			for i, child := range container.Children {
				walkMarkdown(child, append(path, fmt.Sprintf("%s:%d", cId, i)))
			}
		} else if leaf, ok := isLeaf(node); ok {
			lId := strings.Join(append(path, leafId(leaf)), "/")
			l.Debug("Leaf", esl.Any("path", path), esl.Any("leaf", leaf))
			contentsById[lId] = string(leaf.Literal)
		}
	}

	walkMarkdown(markdown, []string{})
	l.Info("Contents by ID", esl.Any("contentsById", contentsById))

	return nil
}

func (z *Markdown) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Markdown{}, func(r rc_recipe.Recipe) {
		m := r.(*Markdown)
		m.Content.SetFilePath("markdown_test.md")
	})
}
