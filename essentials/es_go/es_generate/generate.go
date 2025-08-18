package es_generate

import (
	"bytes"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"go/format"
	"strings"
	"text/template"
)

type Generator interface {
	// Generate source with templateName
	Generate(tmplName string) (src []byte, err error)
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

func (z *structTypeGenerator) Generate(tmplName string) ([]byte, error) {
	l := esl.Default().With(esl.String("tmpl", tmplName))
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

	t0, err := app_resource.Bundle().Templates().Bytes(tmplName)
	if err != nil {
		l.Debug("Unable to load the template", esl.Error(err))
		return nil, err
	}
	t1, err := template.New(tmplName).Parse(string(t0))
	if err != nil {
		l.Debug("Unable to parse the template", esl.Error(err))
		return nil, err
	}
	var buf bytes.Buffer
	err = t1.Execute(&buf, map[string]interface{}{
		"ImportAliases": aliases,
		"Imports":       packages,
		"Objects":       aliasObjects,
	})
	if err != nil {
		l.Debug("Unable to execute the template", esl.Error(err))
		return nil, err
	}

	// Format
	src, err := format.Source(buf.Bytes())
	if err != nil {
		l.Warn("Unable to execute go format", esl.Error(err))
		return buf.Bytes(), nil
	}

	return src, nil
}
