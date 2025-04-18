package sv_desktop

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"

	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_desktop"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/log/esl"
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
	l := esl.Default()

	findEnvHome := func(envName string) string {
		home := os.Getenv(envName)
		if home != "" {
			l.Debug("home folder found", esl.String("envName", envName), esl.String("path", home))
		} else {
			l.Debug("home folder not found", esl.String("envName", envName))
		}
		return home
	}
	findInfoFile := func(path string) (gjson.Result, error) {
		ll := l.With(esl.String("path", path))
		b, err := os.ReadFile(path)
		if err != nil {
			ll.Debug("Unable to read info.json", esl.Error(err))
			return gjson.Parse("{}"), err
		}
		if !gjson.ValidBytes(b) {
			ll.Debug("Invalid JSON format", esl.Error(err))
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
	l := esl.Default()

	info, err := z.findInfo()
	if err != nil {
		l.Debug("info.json not found or invalid", esl.Error(err))
		return nil, nil, err
	}
	personal = &mo_desktop.Desktop{}
	business = &mo_desktop.Desktop{}

	var lastErr error
	if err := api_parser.ParseModel(personal, info.Get(mo_desktop.TypePersonal)); err != nil || personal.Path == "" {
		l.Debug("personal Dropbox not found or invalid", esl.Error(err))
		personal = nil
		lastErr = err
	}
	if err := api_parser.ParseModel(business, info.Get(mo_desktop.TypeBusiness)); err != nil || business.Path == "" {
		l.Debug("business Dropbox not found or invalid", esl.Error(err))
		business = nil
		lastErr = err
	}
	if personal == nil && business == nil {
		return nil, nil, lastErr
	}
	return personal, business, nil
}
