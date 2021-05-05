package spec

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"text/template"
)

const (
	changesHeader = `---
layout: release
title: {{.Title}}
lang: {{.Lang}}
---

`
)

type Diff struct {
	rc_recipe.RemarkSecret
	DocLang             mo_string.OptionalString
	Release1            mo_string.OptionalString
	Release2            mo_string.OptionalString
	FilePath            mo_string.OptionalString
	ReleaseCurrent      app_msg.Message
	ReleaseVersion      app_msg.Message
	DocTitle            app_msg.Message
	DocHeader           app_msg.Message
	SpecAdded           app_msg.Message
	SpecDeleted         app_msg.Message
	SpecChanged         app_msg.Message
	SpecChangedRecipe   app_msg.Message
	ChangeRecipeConfig  app_msg.Message
	ChangeReportAdded   app_msg.Message
	ChangeReportDeleted app_msg.Message
	ChangeReportChanged app_msg.Message
	ChangeFeedAdded     app_msg.Message
	ChangeFeedDeleted   app_msg.Message
	ChangeFeedChanged   app_msg.Message
	TableHeaderName     app_msg.Message
	TableHeaderDesc     app_msg.Message
	TableHeaderPath     app_msg.Message
	TableHeaderTitle    app_msg.Message
}

func (z *Diff) Preset() {
}

func (z *Diff) loadSpec(c app_control.Control, relName string) (r map[string]*dc_recipe.Recipe, err error) {
	fn := "spec.json.gz"
	if relName != "" {
		fn = "spec_" + relName + ".json.gz"
	}
	p := filepath.Join("resources/release/"+c.Messages().Lang().CodeString(), fn)
	l := c.Log().With(esl.String("path", p))
	l.Debug("Loading")

	f, err := os.Open(p)
	if err != nil {
		l.Error("unable to read spec", esl.Error(err))
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	g, err := gzip.NewReader(f)
	if err != nil {
		l.Error("unable to read spec", esl.Error(err))
		return nil, err
	}
	defer func() {
		_ = g.Close()
	}()
	j, err := ioutil.ReadAll(g)
	if err != nil {
		l.Error("unable to read spec", esl.Error(err))
		return nil, err
	}

	r = make(map[string]*dc_recipe.Recipe)
	if err = json.Unmarshal(j, &r); err != nil {
		l.Error("Unable to unmarshal spec", esl.Error(err))
		return nil, err
	}
	return r, nil
}

type diffRowsContext struct {
	r1               []dc_recipe.DocRows
	r2               []dc_recipe.DocRows
	changeHeader     *sync.Once
	changeHeaderFunc func()
	msgAdd           app_msg.Message
	msgDel           app_msg.Message
	msgChanged       app_msg.Message
}

func (z *Diff) diffRows(mui app_ui.UI, dc *diffRowsContext) {
	r1 := make(map[string]dc_recipe.DocRows)
	r2 := make(map[string]dc_recipe.DocRows)
	added := make([]string, 0)
	deleted := make([]string, 0)
	changed := make([]string, 0)

	for _, x := range dc.r1 {
		r1[x.RowsName()] = x
	}
	for _, x := range dc.r2 {
		r2[x.RowsName()] = x
	}

	for r := range r1 {
		if _, ok := r2[r]; !ok {
			deleted = append(deleted, r)
		}
	}
	for r := range r2 {
		if _, ok := r1[r]; !ok {
			added = append(added, r)
		}
	}
	for x := range r1 {
		if _, ok := r2[x]; ok {
			xd := cmp.Diff(r1[x], r2[x])
			if xd != "" {
				changed = append(changed, x)
			}
		}
	}

	if len(added) > 0 {
		dc.changeHeader.Do(dc.changeHeaderFunc)
		sort.Strings(added)
		mui.SubHeader(dc.msgAdd)
		mt := mui.InfoTable("added")
		mt.Header(z.TableHeaderName, z.TableHeaderDesc)
		for _, x := range added {
			mt.Row(app_msg.Raw(x), app_msg.Raw(r2[x].RowsDesc()))
		}
		mt.Flush()
	}

	if len(deleted) > 0 {
		dc.changeHeader.Do(dc.changeHeaderFunc)
		sort.Strings(deleted)
		mui.SubHeader(dc.msgDel)
		mt := mui.InfoTable("deleted")
		mt.Header(z.TableHeaderName, z.TableHeaderDesc)
		for _, x := range deleted {
			mt.Row(app_msg.Raw(x), app_msg.Raw(r1[x].RowsDesc()))
		}
		mt.Flush()
	}

	if len(changed) > 0 {
		dc.changeHeader.Do(dc.changeHeaderFunc)
		sort.Strings(changed)

		for _, c := range changed {
			mui.SubHeader(dc.msgChanged.With("RowsName", c))
			mui.Code(cmp.Diff(r1[c], r2[c]))
		}
	}
}

func (z *Diff) diffSpec(mui app_ui.UI, s1, s2 *dc_recipe.Recipe) {
	wholeDiff := cmp.Diff(s1, s2)
	// no diff
	if wholeDiff == "" {
		return
	}

	// keep reports & feeds
	r1 := s1.Reports
	r2 := s2.Reports
	f1 := s1.Feeds
	f2 := s2.Feeds

	// compare only for direct attributes
	s1.Reports = nil
	s2.Reports = nil
	s1.Feeds = nil
	s2.Feeds = nil

	changeHeader := &sync.Once{}
	changeHeaderFunc := func() {
		mui.Header(z.SpecChangedRecipe.With("Recipe", s1.Path))
		mui.Break()
	}

	wholeDiff = cmp.Diff(s1, s2)
	s1.Reports = r1
	s2.Reports = r2
	s1.Feeds = f1
	s2.Feeds = f2

	if wholeDiff != "" {
		changeHeader.Do(changeHeaderFunc)
		mui.SubHeader(z.ChangeRecipeConfig)
		mui.Break()
		mui.Code(wholeDiff)
	}

	rr1 := make([]dc_recipe.DocRows, 0)
	for _, x := range r1 {
		rr1 = append(rr1, x)
	}
	rr2 := make([]dc_recipe.DocRows, 0)
	for _, x := range r2 {
		rr2 = append(rr2, x)
	}
	z.diffRows(mui, &diffRowsContext{
		r1:               rr1,
		r2:               rr2,
		msgAdd:           z.ChangeReportAdded,
		msgDel:           z.ChangeReportDeleted,
		msgChanged:       z.ChangeReportChanged,
		changeHeader:     changeHeader,
		changeHeaderFunc: changeHeaderFunc,
	})
}

func (z *Diff) makeDiff(c app_control.Control) error {
	l := c.Log()
	r1, err := z.loadSpec(c, z.Release1.Value())
	if err != nil {
		return nil
	}
	r2, err := z.loadSpec(c, z.Release2.Value())
	if err != nil {
		return nil
	}

	added := make([]string, 0)
	addedTitle := make(map[string]string)
	deleted := make([]string, 0)
	deletedTitle := make(map[string]string)
	for r, s := range r1 {
		if _, ok := r2[r]; !ok {
			deleted = append(deleted, r)
			deletedTitle[r] = s.Title
		}
	}
	for r, s := range r2 {
		if _, ok := r1[r]; !ok {
			added = append(added, r)
			addedTitle[r] = s.Title
		}
	}

	var w io.WriteCloser
	shouldClose := false
	if !z.FilePath.IsExists() {
		w = es_stdout.NewDefaultOut(c.Feature())
	} else {
		var err error
		w, err = os.Create(z.FilePath.Value())
		if err != nil {
			l.Error("Unable to create file", esl.Error(err), esl.String("path", z.FilePath.Value()))
			return err
		}
		shouldClose = true
	}
	defer func() {
		if shouldClose {
			w.Close()
		}
	}()

	relName := func(x string) string {
		if x == "" {
			return c.UI().Text(z.ReleaseCurrent)
		}
		return c.UI().Text(z.ReleaseVersion.With("Release", strings.Replace(x, "_", "", 1)))
	}

	buf := bytes.Buffer{}

	changeTmpl, err := template.New("release_header").Parse(changesHeader)
	if err != nil {
		l.Debug("Unable to compile the template", esl.Error(err))
		return err
	}
	err = changeTmpl.Execute(&buf, map[string]interface{}{
		"Title": c.UI().Text(z.DocTitle.With("Release", z.Release1.Value())),
		"Lang":  c.Messages().Lang().CodeString(),
	})
	if err != nil {
		l.Debug("Unable to exec the template", esl.Error(err))
		return err
	}

	mui := app_ui.NewMarkdown(c.Messages(), c.Log(), &buf, es_dialogue.DenyAll())
	mui.Header(z.DocHeader.With("Release1", relName(z.Release1.Value())).With("Release2", relName(z.Release2.Value())))

	if len(added) > 0 {
		mui.Header(z.SpecAdded)
		sort.Strings(added)
		mt := mui.InfoTable("added")
		mt.Header(z.TableHeaderPath, z.TableHeaderTitle)
		for _, r := range added {
			mt.Row(app_msg.Raw(r), app_msg.Raw(addedTitle[r]))
		}
		mt.Flush()
		mui.Break()
	}

	if len(deleted) > 0 {
		mui.Header(z.SpecDeleted)
		sort.Strings(deleted)
		mt := mui.InfoTable("deleted")
		mt.Header(z.TableHeaderPath, z.TableHeaderTitle)
		for _, r := range deleted {
			mt.Row(app_msg.Raw(r), app_msg.Raw(deletedTitle[r]))
		}
		mt.Flush()
		mui.Break()
	}

	// Search for changed
	changed := make([]string, 0)
	for r := range r1 {
		if _, ok := r2[r]; ok {
			changed = append(changed, r)
		}
	}
	sort.Strings(changed)
	for _, r := range changed {
		z.diffSpec(mui, r1[r], r2[r])
	}

	content := strings.ReplaceAll(buf.String(), "{{.", "{% raw %}{{.{% endraw %}")
	if _, err = w.Write([]byte(content)); err != nil {
		return err
	}

	return nil
}

func (z *Diff) Exec(c app_control.Control) error {
	if z.DocLang.IsExists() {
		return z.makeDiff(c.WithLang(z.DocLang.Value()))
	} else {
		return z.makeDiff(c)
	}
}

func (z *Diff) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
