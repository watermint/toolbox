package ut_filepath

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"golang.org/x/tools/go/ssa/interp/testdata/src/runtime"
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
