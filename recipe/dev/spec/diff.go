package spec

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_doc"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type Diff struct {
	Lang                string
	Release1            string
	Release2            string
	FilePath            string
	ReleaseCurrent      app_msg.Message
	ReleaseVersion      app_msg.Message
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

func (z *Diff) loadSpec(c app_control.Control, relName string) (r map[string]*rc_doc.Recipe, err error) {
	fn := "spec.json"
	if relName != "" {
		fn = "spec_" + relName + ".json"
	}
	p := filepath.Join("doc/generated", fn)
	l := c.Log().With(zap.String("path", p))
	l.Debug("Loading")

	j, err := ioutil.ReadFile(p)
	if err != nil {
		l.Error("unable to read spec", zap.Error(err))
		return nil, err
	}
	r = make(map[string]*rc_doc.Recipe)
	if err = json.Unmarshal(j, &r); err != nil {
		l.Error("Unable to unmarshal spec", zap.Error(err))
		return nil, err
	}
	return r, nil
}

type diffRowsContext struct {
	r1               []rc_doc.DocRows
	r2               []rc_doc.DocRows
	changeHeader     *sync.Once
	changeHeaderFunc func()
	msgAdd           app_msg.Message
	msgDel           app_msg.Message
	msgChanged       app_msg.Message
}

func (z *Diff) diffRows(mui app_ui.UI, dc *diffRowsContext) {
	r1 := make(map[string]rc_doc.DocRows)
	r2 := make(map[string]rc_doc.DocRows)
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

func (z *Diff) diffSpec(mui app_ui.UI, s1, s2 *rc_doc.Recipe) {
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

	rr1 := make([]rc_doc.DocRows, 0)
	for _, x := range r1 {
		rr1 = append(rr1, x)
	}
	rr2 := make([]rc_doc.DocRows, 0)
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
	r1, err := z.loadSpec(c, z.Release1)
	if err != nil {
		return nil
	}
	r2, err := z.loadSpec(c, z.Release2)
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
	if z.FilePath == "" {
		w = ut_io.NewDefaultOut(c.IsTest())
	} else {
		var err error
		w, err = os.Create(z.FilePath)
		if err != nil {
			l.Error("Unable to create file", zap.Error(err), zap.String("path", z.FilePath))
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

	mui := app_ui.NewMarkdown(c.Messages(), w, false)
	mui.Header(z.DocHeader.With("Release1", relName(z.Release1)).With("Release2", relName(z.Release2)))

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
	for r, _ := range r1 {
		if _, ok := r2[r]; ok {
			changed = append(changed, r)
		}
	}
	sort.Strings(changed)
	for _, r := range changed {
		z.diffSpec(mui, r1[r], r2[r])
	}

	return nil
}

func (z *Diff) Exec(c app_control.Control) error {
	if cl, ok := app_control_launcher.ControlWithLang(z.Lang, c); ok {
		return z.makeDiff(cl)
	} else {
		return z.makeDiff(c)
	}
}

func (z *Diff) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}
