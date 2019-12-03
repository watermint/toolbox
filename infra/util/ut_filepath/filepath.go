package ut_filepath

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/domain/service/sv_desktop"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"os/user"
	"runtime"
	"strings"
	"text/template"
	"unicode"
)

var (
	isWindows = runtime.GOOS == "windows"
)

func Rel(basePath, targetPath string) (rel string, err error) {
	l := app_root.Log()

	isSeparator := func(c rune) bool {
		switch {
		case c == '/', c == '\\':
			return true
		case c == ':' && isWindows:
			return true
		default:
			return false
		}
	}

	bpr := []rune(basePath)
	tpr := []rune(targetPath)

	bl := len(bpr)
	tl := len(tpr)

	l = l.With(zap.Int("basePathLen", bl), zap.Int("targetPathLen", tl))

	if bl < 1 || tl < 1 {
		l.Debug("Empty path")
		return "", errors.New("empty path")
	}

	if isSeparator(bpr[bl-1]) {
		bpr = bpr[:bl-1]
		bl = len(bpr)
	}
	if isSeparator(tpr[tl-1]) {
		tpr = tpr[:tl-1]
		tl = len(tpr)
	}

	if tl == bl {
		same := true
		for i := 0; i < tl; i++ {
			if unicode.ToLower(bpr[i]) != unicode.ToLower(tpr[i]) {
				same = false
			}
		}
		if same {
			return ".", nil
		}
	}
	if tl <= bl {
		l.Debug("Target path is shorter or than base path, or same length")
		return "", errors.New("target path is shorter than base path")
	}

	errMsg := "target path does not have same base path"

	for i := 0; i < bl; i++ {
		if unicode.ToLower(bpr[i]) != unicode.ToLower(tpr[i]) {
			return "", errors.New(errMsg)
		}
	}
	if isSeparator(bpr[bl-1]) {
		return string(tpr[bl:]), nil
	}
	if isSeparator(tpr[bl]) {
		return string(tpr[bl+1:]), nil
	}
	return "", errors.New(errMsg)
}

type FormatError struct {
	Reason string
	Key    string
}

func (z *FormatError) Error() string {
	return z.String()
}
func (z *FormatError) String() string {
	return "{{." + z.Key + "}}: " + z.Reason
}

// Format path if a path contains pattern like `{{.DropboxPersonal}}`.
func FormatPathWithPredefinedVariables(path string) (string, error) {
	predefined := make(map[string]func() (string, error))
	predefined["DropboxPersonal"] = func() (s string, e error) {
		p, _, _ := sv_desktop.New().Lookup()
		if p != nil {
			return p.Path, nil
		}
		return "", errors.New("personal dropbox desktop folder not found")
	}
	predefined["DropboxBusiness"] = func() (s string, e error) {
		_, p, _ := sv_desktop.New().Lookup()
		if p != nil {
			return p.Path, nil
		}
		return "", errors.New("business dropbox desktop folder not found")
	}
	predefined["Home"] = func() (s string, e error) {
		u, err := user.Current()
		if err == nil {
			return u.HomeDir, nil
		}
		return "", errors.New("unable to retrieve current user home")
	}
	predefined["AlwaysErrorForTest"] = func() (s string, e error) {
		return "", errors.New("always error")
	}
	data := make(map[string]string)

	for k, vf := range predefined {
		ptn := "{{." + k + "}}"
		if strings.Index(path, ptn) >= 0 {
			v, err := vf()
			if err != nil {
				return "", &FormatError{
					Reason: err.Error(),
					Key:    k,
				}
			}
			data[k] = v
		}
	}

	var buf bytes.Buffer
	pathTmpl, err := template.New("path").Parse(path)
	if err != nil {
		return "", err
	}

	err = pathTmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
