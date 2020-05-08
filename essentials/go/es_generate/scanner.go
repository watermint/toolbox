package es_generate

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
)

var (
	ErrorTargetInterfaceNotFound = errors.New("the target interface is not found in the source tree")
)

type Scanner interface {
	// Find struct which implements the given interface.
	// Please pass the type `t` as like reflect.TypeOf((*INTERFACE_PKG.INTERFACE)(nil)).Elem()
	FindStructImplements(refType reflect.Type) ([]*StructType, error)

	// Find struct which has name starts from prefix.
	FindStructHasPrefix(prefix string) ([]*StructType, error)

	// Exclude `*_test.go` files.
	ExcludeTest() Scanner

	// Limit path with prefix
	PathFilterPrefix(prefix string) Scanner
}

func NewScanner(c app_control.Control, path string) (Scanner, error) {
	s := &scannerImpl{
		c:           c,
		path:        path,
		excludeTest: false,
	}
	if err := s.load(); err != nil {
		return nil, err
	}
	return s, nil
}

type scannerImpl struct {
	c           app_control.Control
	path        string
	excludeTest bool
	fst         *token.FileSet
	allPkg      []*ast.Package
	pathPrefix  string
}

func (z *scannerImpl) PathFilterPrefix(prefix string) Scanner {
	return &scannerImpl{
		c:           z.c,
		path:        z.path,
		excludeTest: z.excludeTest,
		fst:         z.fst,
		allPkg:      z.allPkg,
		pathPrefix:  prefix,
	}
}

func (z *scannerImpl) ExcludeTest() Scanner {
	return &scannerImpl{
		c:           z.c,
		path:        z.path,
		excludeTest: true,
		fst:         z.fst,
		allPkg:      z.allPkg,
		pathPrefix:  z.pathPrefix,
	}
}

func (z *scannerImpl) log() esl.Logger {
	return z.c.Log().With(esl.String("path", z.path), esl.Bool("excludeTest", z.excludeTest))
}

func (z *scannerImpl) load() error {
	l := z.c.Log()
	l.Debug("Loading")
	z.fst = token.NewFileSet()
	z.allPkg = make([]*ast.Package, 0)

	var parseDir func(relPath string) error
	parseDir = func(relPath string) error {
		path0 := filepath.Join(z.path, relPath)
		l.Debug("Scanning", esl.String("path", path0))
		pkgs, err := parser.ParseDir(z.fst, path0, nil, 0)
		if err != nil {
			l.Error("Parse error", esl.Error(err))
			return err
		}
		for _, pkg := range pkgs {
			z.allPkg = append(z.allPkg, pkg)
		}
		entries, err := ioutil.ReadDir(path0)
		for _, entry := range entries {
			if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
				if err := parseDir(filepath.Join(relPath, entry.Name())); err != nil {
					return err
				}
			}
		}
		return nil
	}
	return parseDir("")
}

func (z *scannerImpl) typesConfig() *types.Config {
	return &types.Config{
		Error: func(err error) {
			z.log().Debug("error", esl.Error(err))
		},
		Importer: importer.Default(),
	}
}

func (z *scannerImpl) findAstInterface(refType reflect.Type) (astType *types.Interface, err error) {
	l := z.log()

	cfg := z.typesConfig()
	l.Debug("Recipe type", esl.String("name", refType.Name()), esl.String("pkg", refType.PkgPath()))
	for _, pkg := range z.allPkg {
		for f0, f := range pkg.Files {
			l.Debug("scan files", esl.String("f0", f0))
			if r := f.Scope.Lookup(refType.Name()); r != nil {
				l.Debug("finding recipe", esl.String("r", r.Name))
				info := types.Info{
					Types: make(map[ast.Expr]types.TypeAndValue),
					Defs:  make(map[*ast.Ident]types.Object),
					Uses:  make(map[*ast.Ident]types.Object),
				}
				q, err := cfg.Check(z.path, z.fst, []*ast.File{f}, &info)
				l.Debug("check error", esl.Error(err))
				ro := q.Scope().Lookup(refType.Name())
				if rat, ok := ro.Type().Underlying().(*types.Interface); ok {
					astType = rat
					break
				}
			}
		}
	}
	if astType == nil {
		l.Debug("Recipe interface not found")
		return nil, ErrorTargetInterfaceNotFound
	}
	return astType, nil
}

func (z *scannerImpl) FindStructImplements(refType reflect.Type) (sts []*StructType, err error) {
	l := z.log()

	astType, err := z.findAstInterface(refType)
	if err != nil {
		return nil, err
	}

	sts = make([]*StructType, 0)
	cfg := z.typesConfig()

	for _, pkg := range z.allPkg {
		for f0, f := range pkg.Files {
			info := types.Info{
				Types: make(map[ast.Expr]types.TypeAndValue),
				Defs:  make(map[*ast.Ident]types.Object),
				Uses:  make(map[*ast.Ident]types.Object),
			}

			q, err := cfg.Check(z.path, z.fst, []*ast.File{f}, &info)
			if err != nil {
				l.Debug("unable to check", esl.Error(err))
			}
			for _, n := range q.Scope().Names() {
				l.Debug("name", esl.String("name", n))
				obj := q.Scope().Lookup(n)
				pkgPath := obj.Pkg().Path()
				pkgName := obj.Pkg().Name()
				rel, _ := filepath.Rel(pkgPath, filepath.Dir(f0))
				if z.pathPrefix != "" && !strings.HasPrefix(rel, z.pathPrefix) {
					continue
				}
				if z.excludeTest && strings.HasSuffix(pkgName, "_test") {
					continue
				}
				if _, ok := obj.Type().Underlying().(*types.Struct); !ok {
					continue
				}
				ptr := types.NewPointer(obj.Type())
				ut := ptr.Underlying()
				impl := types.Implements(ut, astType)
				if impl {
					l.Debug("underlying", esl.String("on", obj.Name()), esl.String("pkg", obj.Pkg().Name()), esl.String("f0", f0), esl.String("rel", rel))
					st := &StructType{
						Package: rel,
						Name:    obj.Name(),
					}
					sts = append(sts, st)
				}
			}
		}
	}
	return sts, nil
}

func (z *scannerImpl) FindStructHasPrefix(prefix string) (sts []*StructType, err error) {
	l := z.log()

	sts = make([]*StructType, 0)
	cfg := z.typesConfig()
	for _, pkg := range z.allPkg {
		for f0, f := range pkg.Files {
			l.Debug("scan files", esl.String("f0", f0))
			for _, r := range f.Scope.Objects {

				l.Debug("finding message object", esl.String("r", r.Name))
				info := types.Info{
					Types: make(map[ast.Expr]types.TypeAndValue),
					Defs:  make(map[*ast.Ident]types.Object),
					Uses:  make(map[*ast.Ident]types.Object),
				}
				q, err := cfg.Check(z.path, z.fst, []*ast.File{f}, &info)
				l.Debug("check error", esl.Error(err))
				for _, qn := range q.Scope().Names() {
					ro := q.Scope().Lookup(qn)
					pkgPath := ro.Pkg().Path()
					rel, _ := filepath.Rel(pkgPath, filepath.Dir(f0))
					if z.pathPrefix != "" && !strings.HasPrefix(rel, z.pathPrefix) {
						continue
					}
					if z.excludeTest && strings.HasSuffix(ro.Pkg().Name(), "_test") {
						continue
					}
					if !strings.HasPrefix(qn, prefix) {
						continue
					}
					if _, ok := ro.Type().Underlying().(*types.Struct); !ok {
						continue
					}
					st := &StructType{
						Package: rel,
						Name:    qn,
					}
					sts = append(sts, st)
				}
			}
		}
	}
	return sts, nil
}
