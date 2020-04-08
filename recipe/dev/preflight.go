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
	"github.com/watermint/toolbox/infra/ui/app_lang"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"github.com/watermint/toolbox/recipe/dev/spec"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	minimumSpecDocVersion = 59
)

type Preflight struct {
}

func (z *Preflight) Preset() {
}

func (z *Preflight) Test(c app_control.Control) error {
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

func (z *Preflight) deleteOldGeneratedFiles(c app_control.Control, path string) error {
	l := c.Log()
	whiteList := func(name string) bool {
		nameLower := strings.ToLower(name)
		switch {
		case nameLower == "changes.md":
			return true
		case strings.HasPrefix(nameLower, "spec") && strings.HasSuffix(nameLower, ".json"):
			return true
		default:
			return false
		}
	}

	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, e := range entries {
		if !whiteList(e.Name()) {
			p := filepath.Join(path, e.Name())
			if c.Feature().IsTest() {
				continue
			}
			err := os.Remove(p)
			if err != nil {
				l.Error("Unable to remove file", zap.Error(err), zap.String("path", p))
				return nil
			}
		}
	}
	return nil
}

func (z *Preflight) cloneSpec(c app_control.Control, path string, release int) error {
	l := c.Log()
	if c.Feature().IsTest() {
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

func (z *Preflight) Exec(c app_control.Control) error {
	l := c.Log()

	release := 0
	if !c.Feature().IsTest() {
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

	for _, lang := range app_lang.SupportedLanguages {
		langCode := app_lang.Base(lang)
		suffix := app_lang.PathSuffix(lang)

		path := fmt.Sprintf("doc/generated%s/", suffix)
		ll := l.With(zap.String("lang", langCode), zap.String("suffix", suffix))

		if !c.Feature().IsTest() {
			ll.Info("Clean up generated document folder")
			if err := z.deleteOldGeneratedFiles(c, path); err != nil {
				return err
			}

			ll.Info("Generating README & command documents")
			err := rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
				rr := r.(*Doc)
				rr.Badge = true
				rr.MarkdownReadme = true
				rr.Lang = langCode
				rr.Filename = fmt.Sprintf("README%s.md", suffix)
				rr.CommandPath = path
			})
			if err != nil {
				l.Error("Failed to generate documents", zap.Error(err))
				return err
			}

			ll.Info("Generating Spec document")
			err = rc_exec.Exec(c, &spec.Doc{}, func(r rc_recipe.Recipe) {
				rr := r.(*spec.Doc)
				rr.Lang = langCode
				rr.FilePath = filepath.Join(path, "spec.json")
			})
			if err != nil {
				l.Error("Failed to generate documents", zap.Error(err))
				return err
			}
			l.Info("Verify message resources")
			if err := qt_messages.VerifyMessages(c); err != nil {
				return err
			}
		}

		ll.Info("Clone spec")
		if err := z.cloneSpec(c, path, release); err != nil {
			return err
		}

		ll.Info("Sorting message resources")
		if err := z.sortMessages(c, fmt.Sprintf("messages%s.json", suffix)); err != nil {
			return err
		}

		if !c.Feature().IsTest() && minimumSpecDocVersion < release {
			ll.Info("Generating release notes")
			err := rc_exec.Exec(c, &spec.Diff{}, func(r rc_recipe.Recipe) {
				rr := r.(*spec.Diff)
				rr.Lang = langCode
				rr.Release1 = fmt.Sprintf("%d", release-1)
				rr.Release2 = fmt.Sprintf("%d", release)
				rr.FilePath = filepath.Join(path, "changes.md")
			})
			if err != nil {
				l.Error("Failed to generate documents", zap.Error(err))
				return err
			}
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

	l.Info("Verify message resources")
	return qt_messages.VerifyMessages(c)
}
