package es_project

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"path/filepath"
	"runtime"
)

var (
	ErrorUnableToResolveRepositoryRoot = errors.New("unable to resolve repository root")
)

func DetectRepositoryRoot() (string, error) {
	l := esl.Default()
	isRoot := func(p string) bool {
		ll := l.With(esl.String("path", p))
		rootFiles := map[string]bool{
			"LICENSE.md": false,
			"README.md":  false,
		}
		entries, err := os.ReadDir(p)
		if err != nil {
			ll.Debug("unable to read directory", esl.Error(err))
			return false
		}
		for _, entry := range entries {
			if _, ok := rootFiles[entry.Name()]; ok {
				rootFiles[entry.Name()] = true
			}
		}
		for _, t := range rootFiles {
			if !t {
				return false
			}
		}
		return true
	}
	traverse := func(p string) (q string, err2 error) {
		for {
			if isRoot(p) {
				return p, nil
			}
			p = filepath.ToSlash(filepath.Dir(p))
			if len(p) <= 1 {
				return "", ErrorUnableToResolveRepositoryRoot
			}
		}
	}

	wd, err := os.Getwd()
	if err != nil {
		_, file, _, _ := runtime.Caller(0)
		if p, err := traverse(file); err == nil {
			return p, nil
		}
	} else {
		if p, err := traverse(wd); err == nil {
			return p, nil
		}
	}

	l.Debug("unable to retrieve working directory", esl.Error(err))
	return "", ErrorUnableToResolveRepositoryRoot
}
