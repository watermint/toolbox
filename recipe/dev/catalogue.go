package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"go.uber.org/zap"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"text/template"
)

type Catalogue struct {
}

func (z *Catalogue) Preset() {
}

func (z *Catalogue) scanRecipe(path string, c app_control.Control) {
	l := c.Log()
	root, err := os.Getwd()
	if err != nil {
		l.Warn("Unable to retrieve current work dir", zap.Error(err))
		return
	}
	basePath := root

	l.Debug("scanning", zap.String("path", path))
	fset := token.NewFileSet()
	allPkgs := make([]*ast.Package, 0)

	var parseDir func(relPath string)
	parseDir = func(relPath string) {
		path0 := filepath.Join(basePath, relPath)
		l.Debug("Scanning", zap.String("path", path0))
		pkgs, err := parser.ParseDir(fset, path0, nil, 0)
		if err != nil {
			l.Error("Parse error", zap.Error(err))
			return
		}
		for _, pkg := range pkgs {
			allPkgs = append(allPkgs, pkg)
		}
		entries, err := ioutil.ReadDir(path0)
		for _, entry := range entries {
			if entry.IsDir() && !strings.HasPrefix(entry.Name(), ".") {
				parseDir(filepath.Join(relPath, entry.Name()))
			}
		}
	}
	parseDir("")

	cfg := &types.Config{
		Error: func(err error) {
			l.Debug("error", zap.Error(err))
		},
		Importer: importer.Default(),
	}

	recipeType := reflect.TypeOf((*rc_recipe.Recipe)(nil)).Elem()
	var recipeAstType *types.Interface
	l.Debug("Recipe type", zap.String("name", recipeType.Name()), zap.String("pkg", recipeType.PkgPath()))
	for _, pkg := range allPkgs {
		for f0, f := range pkg.Files {
			l.Debug("scan files", zap.String("f0", f0))
			if r := f.Scope.Lookup(recipeType.Name()); r != nil {
				l.Debug("finding recipe", zap.String("r", r.Name))
				info := types.Info{
					Types: make(map[ast.Expr]types.TypeAndValue),
					Defs:  make(map[*ast.Ident]types.Object),
					Uses:  make(map[*ast.Ident]types.Object),
				}
				q, err := cfg.Check(basePath, fset, []*ast.File{f}, &info)
				l.Debug("check error", zap.Error(err))
				ro := q.Scope().Lookup(recipeType.Name())
				if rat, ok := ro.Type().Underlying().(*types.Interface); ok {
					recipeAstType = rat
					break
				}
			}
		}
	}
	if recipeAstType == nil {
		l.Warn("Recipe interface not found")
		return
	}

	type RecipeImpl struct {
		Package string
		Name    string
	}

	recipePackages := make(map[string]bool)
	recipePkgAliases := make(map[string]string)
	recipeImpls := make([]*RecipeImpl, 0)

	for _, pkg := range allPkgs {
		for f0, f := range pkg.Files {
			info := types.Info{
				Types: make(map[ast.Expr]types.TypeAndValue),
				Defs:  make(map[*ast.Ident]types.Object),
				Uses:  make(map[*ast.Ident]types.Object),
			}

			q, err := cfg.Check(basePath, fset, []*ast.File{f}, &info)
			if err != nil {
				l.Debug("unable to check", zap.Error(err))
			}
			for _, n := range q.Scope().Names() {
				l.Debug("name", zap.String("name", n))
				obj := q.Scope().Lookup(n)
				if strings.HasSuffix(obj.Pkg().Name(), "_test") {
					continue
				}
				if _, ok := obj.Type().Underlying().(*types.Struct); !ok {
					continue
				}
				ptr := types.NewPointer(obj.Type())
				ut := ptr.Underlying()
				impl := types.Implements(ut, recipeAstType)
				if impl {
					rel, _ := filepath.Rel(obj.Pkg().Path(), filepath.Dir(f0))
					l.Debug("underlying", zap.String("on", obj.Name()), zap.String("pkg", obj.Pkg().Name()), zap.String("f0", f0), zap.String("rel", rel))
					recipePkg := app.Pkg + "/" + rel
					recipePkgAlias := strings.ReplaceAll(rel, "/", "")
					recipePackages[recipePkg] = true
					recipeImpl := &RecipeImpl{
						Package: recipePkgAlias,
						Name:    obj.Name(),
					}
					recipeImpls = append(recipeImpls, recipeImpl)
					recipePkgAliases[recipePkg] = recipePkgAlias
				}
			}
		}
	}

	recipePkgSorted := make([]string, 0)
	for p := range recipePackages {
		recipePkgSorted = append(recipePkgSorted, p)
	}
	sort.Strings(recipePkgSorted)

	genPrefixes := []string{"recipe", "ingredient"}

	for _, gp := range genPrefixes {
		outName := fmt.Sprintf("catalogue/%s.go", gp)
		gtName := fmt.Sprintf("catalogue_%s.go.tmpl", gp)
		l.Debug("generate", zap.String("template", gtName))
		t0, err := c.Resource(gtName)
		if err != nil {
			l.Error("err", zap.Error(err))
			return
		}
		t1, err := template.New("catalogue").Parse(string(t0))
		if err != nil {
			l.Error("err", zap.Error(err))
			return
		}

		gImports := make([]string, 0)
		gRecipes := make([]*RecipeImpl, 0)
		for _, p := range recipePkgSorted {
			if strings.HasPrefix(p, app.Pkg+"/"+gp) {
				gImports = append(gImports, p)
			}
		}
		for _, r := range recipeImpls {
			if strings.HasPrefix(r.Package, gp) {
				gRecipes = append(gRecipes, r)
			}
		}

		outPath := filepath.Join(basePath, outName)
		l.Info("Generating catalogue", zap.String("path", outPath))
		f, err := os.Create(outPath)
		if err != nil {
			l.Error("Failed to generate catalogue", zap.Error(err))
			continue
		}
		err = t1.Execute(f, map[string]interface{}{
			"ImportAliases": recipePkgAliases,
			"Imports":       gImports,
			"Recipes":       gRecipes,
		})
		if err != nil {
			l.Error("Unable to execute template", zap.Error(err))
		}
		f.Close()
	}
}

func (z *Catalogue) Exec(c app_control.Control) error {
	z.scanRecipe("recipe", c)
	return nil
}

func (z *Catalogue) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Catalogue{}, rc_recipe.NoCustomValues)
}
