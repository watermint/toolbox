package build

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_array"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"github.com/watermint/toolbox/quality/infra/qt_msgusage"
	"github.com/watermint/toolbox/recipe/dev/spec"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	minimumSpecDocVersion = 59
)

type Preflight struct {
	rc_recipe.RemarkConsole
	rc_recipe.RemarkSecret
}

func (z *Preflight) Preset() {
}

func (z *Preflight) Test(c app_control.Control) error {
	return z.Exec(c)
}

func (z *Preflight) sortMessages(c app_control.Control, filename string) error {
	l := c.Log().With(esl.String("filename", filename))
	p := filepath.Join("resources/messages", filename)
	content, err := ioutil.ReadFile(p)
	if err != nil {
		l.Info("SKIP: Unable to open resource file", esl.Error(err))
		return nil
	}
	messages := make(map[string]string)
	if err = json.Unmarshal(content, &messages); err != nil {
		l.Warn("Unable to unmarshal message file", esl.Error(err))
		return err
	}

	touchedKeys := qt_msgusage.Record().Used()
	definedKeys := make([]string, 0)
	for k := range messages {
		definedKeys = append(definedKeys, k)
	}
	unusedKeys := es_array.NewByString(definedKeys...).Diff(es_array.NewByString(touchedKeys...)).AsStringArray()
	sort.Strings(unusedKeys)
	for _, k := range unusedKeys {
		l.Warn("Unused key found, removing it", esl.String("key", k))
		delete(messages, k)
	}

	buf := &bytes.Buffer{}
	je := json.NewEncoder(buf)
	je.SetEscapeHTML(false)
	je.SetIndent("", "  ")
	if err = je.Encode(messages); err != nil {
		l.Warn("Unable to create sorted image", esl.Error(err))
		return err
	}
	if err := ioutil.WriteFile(p, buf.Bytes(), 0644); err != nil {
		l.Warn("Unable to update message", esl.Error(err))
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
		case strings.HasPrefix(nameLower, "spec") && strings.HasSuffix(nameLower, ".json.gz"):
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
				l.Error("Unable to remove file", esl.Error(err), esl.String("path", p))
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
	src := filepath.Join(path, "spec.json.gz")
	dst := filepath.Join(path, fmt.Sprintf("spec_%d.json.gz", release))

	s, err := os.Open(src)
	if err != nil {
		l.Error("Unable to open current spec file", esl.Error(err))
		return err
	}
	defer func() {
		_ = s.Close()
	}()
	d, err := os.Create(dst)
	if err != nil {
		l.Error("Unable to create version spec file", esl.Error(err))
		return err
	}
	defer func() {
		_ = d.Close()
	}()

	_, err = io.Copy(d, s)
	return err
}

func (z *Preflight) Exec(c app_control.Control) error {
	l := c.Log()

	release := 0
	if !c.Feature().IsTest() {
		v, err := ioutil.ReadFile("version")
		if err != nil {
			l.Error("Unable to read version file", esl.Error(err))
			return err
		}
		release, err = strconv.Atoi(string(v))
		if err != nil {
			l.Error("Unable to parse version number", esl.Error(err))
			return err
		}
	}

	for _, la := range lang.Supported {
		langCode := la.CodeString()
		suffix := la.Suffix()

		path := fmt.Sprintf("doc/generated%s/", suffix)
		ll := l.With(esl.String("lang", langCode), esl.String("suffix", suffix))

		if !c.Feature().IsTest() {
			ll.Info("Clean up generated document folder")
			if err := z.deleteOldGeneratedFiles(c, path); err != nil {
				return err
			}

			ll.Info("Generating README & command documents")
			err := rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
				rr := r.(*Doc)
				rr.Badge = true
				rr.DocLang = mo_string.NewOptional(langCode)
			})
			if err != nil {
				l.Error("Failed to generate documents", esl.Error(err))
				return err
			}

			ll.Info("Generating Spec document")
			err = rc_exec.Exec(c, &spec.Doc{}, func(r rc_recipe.Recipe) {
				rr := r.(*spec.Doc)
				rr.Lang = mo_string.NewOptional(langCode)
				rr.FilePath = mo_string.NewOptional(filepath.Join(path, "spec.json.gz"))
			})
			if err != nil {
				l.Error("Failed to generate documents", esl.Error(err))
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

		if !c.Feature().IsTest() && minimumSpecDocVersion < release {
			ll.Info("Generating release changes")
			err := rc_exec.Exec(c, &spec.Diff{}, func(r rc_recipe.Recipe) {
				rr := r.(*spec.Diff)
				rr.DocLang = mo_string.NewOptional(langCode)
				rr.Release1 = mo_string.NewOptional(fmt.Sprintf("%d", release-1))
				rr.Release2 = mo_string.NewOptional(fmt.Sprintf("%d", release))
				rr.FilePath = mo_string.NewOptional(filepath.Join(path, "changes.md"))
			})
			if err != nil {
				l.Error("Failed to generate documents", esl.Error(err))
				return err
			}
		}
	}

	{
		cat := app_catalogue.Current()
		l.Info("Verify recipes")
		for _, r := range cat.Recipes() {
			spec := rc_spec.New(r)
			for _, m := range spec.Messages() {
				qt_msgusage.Record().Touch(m.Key())
				l.Debug("message", esl.String("key", m.Key()), esl.String("text", c.UI().Text(m)))
			}
		}

		l.Info("Verify ingredients")
		for _, r := range cat.Ingredients() {
			spec := rc_spec.New(r)
			for _, m := range spec.Messages() {
				qt_msgusage.Record().Touch(m.Key())
				l.Debug("message", esl.String("key", m.Key()), esl.String("text", c.UI().Text(m)))
			}
		}

		l.Info("Verify message objects")
		for _, m := range cat.Messages() {
			m1 := app_msg.Apply(m)
			msgs := app_msg.Messages(m1)
			for _, msg := range msgs {
				qt_msgusage.Record().Touch(msg.Key())
				l.Debug("message", esl.String("key", msg.Key()), esl.String("text", c.UI().Text(msg)))
			}
		}

		l.Info("Verify features")
		for _, f := range cat.Features() {
			key := app_feature.OptInName(f)
			qt_msgusage.Record().Touch(key)
			ll := l.With(esl.String("key", key))
			ll.Debug("feature disclaimer", esl.String("msg", c.UI().Text(app_feature.OptInDisclaimer(f))))
			ll.Debug("feature agreement", esl.String("msg", c.UI().Text(app_feature.OptInAgreement(f))))
			ll.Debug("feature desc", esl.String("msg", c.UI().Text(app_feature.OptInDescription(f))))
		}
	}

	l.Info("Verify message resources")
	verifyErr := qt_messages.VerifyMessages(c)
	if verifyErr != nil {
		return verifyErr
	}

	for _, la := range lang.Supported {
		suffix := la.Suffix()
		l.Info("Sorting message resources", esl.String("suffix", suffix))
		if err := z.sortMessages(c, fmt.Sprintf("messages%s.json", suffix)); err != nil {
			return err
		}
	}

	return nil
}
