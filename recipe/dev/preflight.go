package dev

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"github.com/watermint/toolbox/recipe/dev/spec"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	minimumSpecDocVersion = 59
)

type Preflight struct {
	TestMode bool
}

func (z *Preflight) Preset() {
	z.TestMode = false
}

func (z *Preflight) Test(c app_control.Control) error {
	z.TestMode = true
	return z.Exec(c)
}

func (z *Preflight) sortMessages(c app_control.Control, filename string) error {
	l := c.Log().With(zap.String("filename", filename))
	p := filepath.Join("resources", filename)
	content, err := ioutil.ReadFile(p)
	if err != nil {
		l.Info("SKIP: Unable to open resource file", zap.Error(err))
		return nil
	}
	messages := make(map[string]string)
	if err = json.Unmarshal(content, &messages); err != nil {
		l.Warn("Unable to unmarshal message file", zap.Error(err))
		return err
	}
	buf := &bytes.Buffer{}
	je := json.NewEncoder(buf)
	je.SetEscapeHTML(false)
	je.SetIndent("", "  ")
	if err = je.Encode(messages); err != nil {
		l.Warn("Unable to create sorted image", zap.Error(err))
		return err
	}
	if err := ioutil.WriteFile(p, buf.Bytes(), 0644); err != nil {
		l.Warn("Unable to update message", zap.Error(err))
		return err
	}
	return nil
}

func (z *Preflight) Exec(c app_control.Control) error {
	l := c.Log()

	release := 0
	if !z.TestMode {
		v, err := ioutil.ReadFile("version")
		if err != nil {
			l.Error("Unable to read version file", zap.Error(err))
			return err
		}
		release, err = strconv.Atoi(string(v))
		if err != nil {
			l.Error("Unable to parse version number", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Generating English documents")
		err := rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
			rr := r.(*Doc)
			rr.TestMode = z.TestMode
			rr.Badge = true
			rr.MarkdownReadme = true
			rr.Lang = "en"
			rr.Filename = "README.md"
			rr.CommandPath = "doc/generated/"
		})
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}
	{
		l.Info("Generating Japanese documents")
		err := rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
			rr := r.(*Doc)
			rr.TestMode = z.TestMode
			rr.Badge = true
			rr.MarkdownReadme = true
			rr.Lang = "ja"
			rr.Filename = "README_ja.md"
			rr.CommandPath = "doc/generated_ja/"
		})
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}

	if !z.TestMode {
		l.Info("Generating Spec document (English)")
		err := rc_exec.Exec(c, &spec.Doc{}, func(r rc_recipe.Recipe) {
			rr := r.(*spec.Doc)
			rr.Lang = "en"
			rr.FilePath = "doc/generated/spec.json"
		})
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}

	if !z.TestMode {
		l.Info("Generating Spec document (Japanese)")
		err := rc_exec.Exec(c, &spec.Doc{}, func(r rc_recipe.Recipe) {
			rr := r.(*spec.Doc)
			rr.Lang = "ja"
			rr.FilePath = "doc/generated_ja/spec.json"
		})
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}

	cloneSpec := func(path string) error {
		if z.TestMode {
			l.Debug("Skip for test")
			return nil
		}
		s, err := ioutil.ReadFile(filepath.Join(path, "spec.json"))
		if err != nil {
			l.Error("Unable to open current spec file", zap.Error(err))
			return err
		}
		err = ioutil.WriteFile(filepath.Join(path, fmt.Sprintf("spec_%d.json", release)), s, 0644)
		if err != nil {
			l.Error("Unable to create version spec file", zap.Error(err))
			return err
		}
		return nil
	}

	if err := cloneSpec("doc/generated"); err != nil {
		return err
	}
	if err := cloneSpec("doc/generated_ja"); err != nil {
		return err
	}
	if !z.TestMode && minimumSpecDocVersion < release {
		l.Info("Generating release notes")
		err := rc_exec.Exec(c, &spec.Diff{}, func(r rc_recipe.Recipe) {
			rr := r.(*spec.Diff)
			rr.Release1 = fmt.Sprintf("%d", release-1)
			rr.FilePath = "doc/generated/changes.md"
		})
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}

	{
		cl := c.(app_control_launcher.ControlLauncher)
		cat := cl.Catalogue()
		l.Info("Verify recipes")
		for _, r := range cat.Recipes() {
			spec := rc_spec.New(r)
			for _, m := range spec.Messages() {
				l.Debug("message", zap.String("key", m.Key()), zap.String("text", c.UI().Text(m)))
			}
		}

		l.Info("Verify ingredients")
		for _, r := range cat.Ingredients() {
			spec := rc_spec.New(r)
			for _, m := range spec.Messages() {
				l.Debug("message", zap.String("key", m.Key()), zap.String("text", c.UI().Text(m)))
			}
		}

		l.Info("Verify message objects")
		for _, m := range cat.Messages() {
			msgs := app_msg.Messages(m)
			for _, msg := range msgs {
				l.Debug("message", zap.String("key", msg.Key()), zap.String("text", c.UI().Text(msg)))
			}
		}
	}

	l.Info("Sorting message resources")
	if err := z.sortMessages(c, "messages.json"); err != nil {
		return err
	}
	if err := z.sortMessages(c, "messages_ja.json"); err != nil {
		return err
	}

	l.Info("Verify message resources")
	qm := c.Messages().(app_msg_container.Quality)
	missing := qm.MissingKeys()
	if len(missing) > 0 {
		suggested := make([]string, 0)
		for _, k := range missing {
			l.Error("Key missing", zap.String("key", k))
			suggested = append(suggested, "\""+k+"\":\"\",")
		}
		sort.Strings(suggested)
		fmt.Println(strings.Join(suggested, "\n"))
	}

	return qt_messages.VerifyMessages(c)
}
