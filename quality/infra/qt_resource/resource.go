package qt_resource

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"io/ioutil"
	"os"
	"strings"
)

var (
	ErrorInvalidResource = errors.New("invalid resource")
)

func WithResource(r rc_recipe.Recipe, f func(j es_json.Json) error) error {
	l := es_log.Default()

	spec := rc_spec.New(r)
	rcPath, rcName := spec.Path()
	jPath := "recipe." + strings.Join(rcPath, ".") + "." + rcName
	ll := l.With(es_log.String("Recipe", spec.CliPath()))

	resPath := os.Getenv(app.EnvNameTestResource)
	if resPath == "" {
		ll.Debug("Resource path was not specified")
		return qt_errors.ErrorNotEnoughResource
	}

	ll.Debug("Loading resource", es_log.String("path", resPath))
	b, err := ioutil.ReadFile(resPath)
	if err != nil {
		return err
	}
	if !gjson.ValidBytes(b) {
		l.Debug("invalid json sequence", es_log.Error(err), es_log.String("path", resPath))
		return ErrorInvalidResource
	}

	j, err := es_json.Parse(b)
	if err != nil {
		ll.Debug("Invalid json sequence", es_log.Error(err))
		return err
	}

	ll.Debug("Looking for resource for the recipe", es_log.String("jPath", jPath))
	if jr, found := j.Find(jPath); found {
		ll.Debug("Resource found. Execute test")
		return f(jr)
	}

	ll.Debug("Resource not found")
	return qt_errors.ErrorNotEnoughResource
}
