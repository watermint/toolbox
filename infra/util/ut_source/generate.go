package ut_source

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"io"
	"strings"
	"text/template"
)

type Generator interface {
	// Generate source with templateName
	Generate(tmplName string, out io.Writer) error
}

func NewStructTypeGenerator(ctl app_control.Control, sts []*StructType) Generator {
	return &structTypeGenerator{
		ctl: ctl,
		sts: sts,
	}
}

type structTypeGenerator struct {
	ctl app_control.Control
	sts []*StructType
}

func (z *structTypeGenerator) Generate(tmplName string, out io.Writer) error {
	makeAlias := func(pkg string) string {
		return strings.ReplaceAll(pkg, "/", "")
	}
	aliases := make(map[string]string)
	packages := UniqSortedPackages(z.sts)
	for _, pkg := range packages {
		aliases[pkg] = makeAlias(pkg)
	}
	objects := SortedStructTypes(z.sts)
	aliasObjects := make([]*StructType, 0)
	for _, st := range objects {
		sta := &StructType{
			Package: makeAlias(st.Package),
			Name:    st.Name,
		}
		aliasObjects = append(aliasObjects, sta)
	}

	t0, err := z.ctl.Resource(tmplName)
	if err != nil {
		return err
	}
	t1, err := template.New(tmplName).Parse(string(t0))
	if err != nil {
		return err
	}
	err = t1.Execute(out, map[string]interface{}{
		"ImportAliases": aliases,
		"Imports":       packages,
		"Objects":       aliasObjects,
	})
	return err
}
