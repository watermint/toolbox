package release

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_lang"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/recipe/dev"
	"github.com/watermint/toolbox/recipe/dev/ci"
	"github.com/watermint/toolbox/recipe/dev/test"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"os"
)

const (
	defaultTestResource = "test/dev/resource.json"
)

type Candidate struct {
	TestResource string
	Auth         *ci.Auth
}

func (z *Candidate) Preset() {
	z.TestResource = defaultTestResource
}

func (z *Candidate) verifyMessages(c app_control.Control) error {
	enMessagesRaw, err := c.Resource("messages.json")
	if err != nil {
		return err
	}
	enMessages := make(map[string]string)
	if err := json.Unmarshal(enMessagesRaw, &enMessages); err != nil {
		return err
	}

	l := c.Log()
	for _, lang := range app_lang.SupportedLanguages {
		if lang == language.English {
			continue
		}
		code := app_lang.Base(lang)
		suffix := app_lang.PathSuffix(lang)

		ll := l.With(zap.String("Language", code))
		ll.Info("Verify messages for language")

		msgRaw, err := c.Resource(fmt.Sprintf("messages%s.json", suffix))
		if err != nil {
			ll.Error("Unable to load message resource", zap.Error(err))
			return err
		}
		msgs := make(map[string]string)
		if err := json.Unmarshal(msgRaw, &msgs); err != nil {
			return err
		}

		missing := false
		for k, v := range enMessages {
			if _, ok := msgs[k]; !ok {
				ll.Warn("Missing key", zap.String("key", k), zap.String("message", v))
				missing = true
			}
		}
		if missing {
			ll.Error("One or more missing key found")
			return errors.New("one or more missing key found")
		}
	}
	return nil
}

func (z *Candidate) Exec(c app_control.Control) error {
	l := c.Log()

	l.Info("Verify translations")
	if err := z.verifyMessages(c); err != nil {
		return err
	}

	l.Info("Preview process")
	err := rc_exec.Exec(c, &dev.Preflight{}, rc_recipe.NoCustomValues)
	if err != nil {
		return err
	}

	l.Info("Ensure end to end resource availability")
	if !rc_conn_impl.IsEndToEndTokenAllAvailable(c) {
		l.Error("At least one of end to end resource is not available.")
		return errors.New("end to end resource is not available")
	}

	l.Info("Testing all end to end test")
	err = rc_exec.Exec(c, &test.Recipe{}, func(r rc_recipe.Recipe) {
		m := r.(*test.Recipe)
		m.All = true
		_, err := os.Lstat(z.TestResource)
		if err == nil {
			m.Resource = z.TestResource
		} else {
			l.Warn("Unable to read test resource", zap.String("path", z.TestResource), zap.Error(err))
		}
	})
	if err != nil {
		return err
	}

	l.Info("This release candidate looks good to go")
	return nil
}

func (z *Candidate) Test(c app_control.Control) error {
	err := z.verifyMessages(c)
	c.Log().Debug("Verify message result", zap.Error(err))

	return qt_errors.ErrorNoTestRequired
}
