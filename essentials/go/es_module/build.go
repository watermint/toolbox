package es_module

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"os"
	"path/filepath"
	"regexp"
	"runtime/debug"
	"strings"
)

var (
	ErrorReadBuildInfo = errors.New("could not read build info")
	ErrorNoGoPath      = errors.New("go path not defined")
)

type Build interface {
	// Modules of the build
	Modules() []Module

	GoVersion() string
}

type buildImpl struct {
	modules   []Module
	goVersion string
}

func (z buildImpl) GoVersion() string {
	return z.goVersion
}

func (z buildImpl) Modules() []Module {
	return z.modules
}

func NewBuild(modules []Module) Build {
	return buildImpl{
		modules: modules,
	}
}

// ScanBuild build information around modules.
func ScanBuild() (b Build, err error) {
	l := esl.Default()
	dbi, ok := debug.ReadBuildInfo()
	if !ok {
		return nil, ErrorReadBuildInfo
	}

	goPathEnv := os.Getenv("GOPATH")
	if goPathEnv == "" {
		return nil, ErrorNoGoPath
	}
	l.Debug("GOPATH", esl.String("Path", goPathEnv))

	modules := make([]Module, 0)
	l.Debug("Dependencies", esl.Int("NumModules", len(dbi.Deps)))
	for _, dep := range dbi.Deps {
		// Skip the main module
		// https://github.com/golang/go/issues/29228
		if dep.Version == "(devel)" {
			continue
		}
		l.Debug("Loading module", esl.Any("module", dep))

		depRegex := regexp.MustCompile(`([A-Z])`)
		depPath := depRegex.ReplaceAllStringFunc(dep.Path, func(s string) string {
			return "!" + strings.ToLower(s)
		})

		modPath := filepath.Join(goPathEnv, "pkg", "mod", depPath+"@"+dep.Version)
		l.Debug("Looking for module root", esl.String("Path", modPath))

		modPathInfo, err := os.Lstat(modPath)
		if os.IsNotExist(err) {
			l.Warn("Module not found", esl.Any("module", dep), esl.String("path", modPath), esl.Error(err))
			continue
		}
		if !modPathInfo.IsDir() {
			l.Warn("Module path is not a folder", esl.Any("module", dep), esl.String("path", modPath))
			continue
		}

		modFs := os.DirFS(modPath)
		licenses, err := ScanLicense(modFs)
		if err != nil {
			l.Error("Unable to scan licenses", esl.String("path", modPath), esl.Error(err))
			return nil, err
		}

		modules = append(modules, NewModule(dep, licenses))
	}
	return NewBuild(modules), nil
}
