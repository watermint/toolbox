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
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
)

type Preflight struct {
	rc_recipe.RemarkConsole
	rc_recipe.RemarkSecret
	Quick bool
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

	entries, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, e := range entries {
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
	return nil
}

func (z *Preflight) Exec(c app_control.Control) error {
	l := c.Log()

	var targetLanguages = make([]lang.Lang, 0)
	if z.Quick {
		targetLanguages = append(targetLanguages, lang.Default)
	} else {
		targetLanguages = lang.Supported
	}

	for _, la := range targetLanguages {
		langCode := la.CodeString()
		suffix := la.Suffix()
		webLangPath := la.CodeString() + "/"
		if la.IsDefault() {
			webLangPath = ""
		}

		ll := l.With(esl.String("lang", langCode), esl.String("suffix", suffix))

		if !c.Feature().IsTest() {
			ll.Info("Clean up docs/{lang}/commands folder")
			if err := z.deleteOldGeneratedFiles(c, fmt.Sprintf("docs/%scommands", webLangPath)); err != nil {
				return err
			}
			ll.Info("Clean up docs/{lang}/guide folder")
			if err := z.deleteOldGeneratedFiles(c, fmt.Sprintf("docs/%sguides", webLangPath)); err != nil {
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

			l.Info("Verify message resources")
			if err := qt_messages.VerifyMessages(c); err != nil {
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
