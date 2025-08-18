package es_generate

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"go/importer"
	"go/types"
	"golang.org/x/tools/go/packages"
)

var (
	ErrorNoPackageFound = errors.New("no package found")
)

func NewEnhancedImporter() types.Importer {
	return &importerWithModDependency{
		packages: make(map[string]*types.Package),
		primary:  importer.Default(),
	}
}

type importerWithModDependency struct {
	packages map[string]*types.Package
	primary  types.Importer
}

func (z *importerWithModDependency) Import(path string) (*types.Package, error) {
	l := esl.Default().With(esl.String("path", path))

	if pkg, ok := z.packages[path]; ok {
		l.Debug("Package found in cache", esl.String("name", pkg.Name()))
		return pkg, nil
	}

	p, err := z.primary.Import(path)
	if err == nil {
		l.Debug("Package found by primary", esl.String("name", p.Name()))
		z.packages[path] = p
		return p, nil
	}
	l.Debug("Package not found by primary", esl.Error(err))

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
	}

	pkgs, err := packages.Load(cfg, path)
	if err != nil {
		l.Debug("Unable to load package", esl.Error(err))
		return nil, err
	}
	l.Debug("Package loaded", esl.Int("numPkgs", len(pkgs)))
	if len(pkgs) == 0 {
		l.Debug("No package found")
		return nil, ErrorNoPackageFound
	}

	z.packages[path] = pkgs[0].Types
	l.Debug("Package found",
		esl.String("name", pkgs[0].Types.Name()),
		esl.String("path", pkgs[0].Types.Path()),
		esl.String("pkgPath", pkgs[0].PkgPath),
		esl.Strings("scopeNames", pkgs[0].Types.Scope().Names()))
	return pkgs[0].Types, nil
}
