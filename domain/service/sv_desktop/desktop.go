package sv_desktop

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/model/mo_desktop"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

type Desktop interface {
	Lookup() (personal *mo_desktop.Desktop, business *mo_desktop.Desktop, err error)
}

func New() Desktop {
	return &desktopImpl{}
}

type desktopImpl struct {
}

func (z *desktopImpl) findInfo() (gjson.Result, error) {
	l := app_root.Log()

	findEnvHome := func(envName string) string {
		em := ut_runtime.EnvMap()
		if e, ok := em[envName]; ok {
			l.Debug("Home folder found", zap.String("envName", envName), zap.String("path", e))
			return e
		}
		l.Debug("Home folder not found")
		return ""
	}
	findInfoFile := func(path string) (gjson.Result, error) {
		ll := l.With(zap.String("path", path))
		b, err := ioutil.ReadFile(path)
		if err != nil {
			ll.Debug("Unable to read info.json", zap.Error(err))
			return gjson.Parse("{}"), err
		}
		if !gjson.ValidBytes(b) {
			ll.Debug("Invalid JSON format", zap.Error(err))
			return gjson.Parse("{}"), err
		}
		j := gjson.ParseBytes(b)
		return j, nil
	}

	switch runtime.GOOS {
	case "windows":
		if eh := findEnvHome("APPDATA"); eh != "" {
			if j, e := findInfoFile(filepath.Join(eh, "Dropbox", "info.json")); e == nil {
				return j, nil
			}
		}
		if eh := findEnvHome("LOCALAPPDATA"); eh != "" {
			if j, e := findInfoFile(filepath.Join(eh, "Dropbox", "info.json")); e == nil {
				return j, nil
			}
		}

	default:
		if eh := findEnvHome("HOME"); eh != "" {
			if j, e := findInfoFile(filepath.Join(eh, ".dropbox", "info.json")); e == nil {
				return j, nil
			}
		}
	}

	return gjson.Parse("{}"), errors.New("valid info.json not found")
}

func (z *desktopImpl) Lookup() (personal *mo_desktop.Desktop, business *mo_desktop.Desktop, err error) {
	l := app_root.Log()

	info, err := z.findInfo()
	if err != nil {
		l.Debug("info.json not found or invalid", zap.Error(err))
		return nil, nil, err
	}
	personal = &mo_desktop.Desktop{}
	business = &mo_desktop.Desktop{}

	var lastErr error
	if err := api_parser.ParseModel(personal, info.Get(mo_desktop.TypePersonal)); err != nil || personal.Path == "" {
		l.Debug("personal Dropbox not found or invalid", zap.Error(err))
		personal = nil
		lastErr = err
	}
	if err := api_parser.ParseModel(business, info.Get(mo_desktop.TypeBusiness)); err != nil || business.Path == "" {
		l.Debug("business Dropbox not found or invalid", zap.Error(err))
		business = nil
		lastErr = err
	}
	if personal == nil && business == nil {
		return nil, nil, lastErr
	}
	return personal, business, nil
}
