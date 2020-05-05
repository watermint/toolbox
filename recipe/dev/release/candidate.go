package release

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/recipe/dev"
	"github.com/watermint/toolbox/recipe/dev/ci/auth"
	"github.com/watermint/toolbox/recipe/dev/test"
	"os"
)

const (
	defaultTestResource = "test/dev/resource.json"
)

type Candidate struct {
	rc_recipe.RemarkSecret
	TestResource string
	Auth         *auth.Connect
}

func (z *Candidate) Preset() {
	z.TestResource = defaultTestResource
}

func (z *Candidate) verifyMessages(c app_control.Control) error {
	enMessagesRaw, err := app_resource.Bundle().Messages().Bytes("messages.json")
	if err != nil {
		return err
	}
	enMessages := make(map[string]string)
	if err := json.Unmarshal(enMessagesRaw, &enMessages); err != nil {
		return err
	}

	l := c.Log()
	for _, la := range lang.Supported {
		if la.IsDefault() {
			continue
		}
		code := la.CodeString()
		suffix := la.Suffix()

		ll := l.With(es_log.String("Language", code))
		ll.Info("Verify messages for language")

		msgRaw, err := app_resource.Bundle().Messages().Bytes(fmt.Sprintf("messages%s.json", suffix))
		if err != nil {
			ll.Error("Unable to load message resource", es_log.Error(err))
			return err
		}
		msgs := make(map[string]string)
		if err := json.Unmarshal(msgRaw, &msgs); err != nil {
			return err
		}

		missing := false
		for k, v := range enMessages {
			if _, ok := msgs[k]; !ok {
				ll.Warn("Missing key", es_log.String("key", k), es_log.String("message", v))
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
	if !dbx_conn_impl.IsEndToEndTokenAllAvailable(c) {
		l.Error("At least one of end to end resource is not available.")
		return errors.New("end to end resource is not available")
	}

	l.Info("Testing all end to end test")
	err = rc_exec.Exec(c, &test.Recipe{}, func(r rc_recipe.Recipe) {
		m := r.(*test.Recipe)
		m.All = true
		_, err := os.Lstat(z.TestResource)
		if err == nil {
			m.Resource = mo_string.NewOptional(z.TestResource)
		} else {
			l.Warn("Unable to read test resource", es_log.String("path", z.TestResource), es_log.Error(err))
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
	c.Log().Debug("Verify message result", es_log.Error(err))

	return qt_errors.ErrorNoTestRequired
}
