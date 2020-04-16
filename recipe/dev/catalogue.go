package dev

import (
	"errors"
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

var (
	ErrorUnableToDetermineSourceRoot = errors.New("unable to determine source root path")
)

type Catalogue struct {
}

func (z *Catalogue) Preset() {
}

func (z *Catalogue) determineSourceRoot(c app_control.Control) (basePath string, err error) {
	l := c.Log()
	basePath, err = os.Getwd()
	if err != nil {
		l.Warn("Unable to retrieve current work dir", zap.Error(err))
		return "", ErrorUnableToDetermineSourceRoot
	}
	return basePath, nil
}

func (z *Catalogue) scanSourceTree(c app_control.Control) (basePath string, fset *token.FileSet, allPkgs []*ast.Package, err error) {
	l := c.Log()
	basePath, err = z.determineSourceRoot(c)
	if err != nil {
		return "", nil, nil, err
	}

	l.Debug("scanning", zap.String("path", basePath))
	fset = token.NewFileSet()
	allPkgs = make([]*ast.Package, 0)

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

	return basePath, fset, allPkgs, nil
}

func (z *Catalogue) scanRecipe(c app_control.Control) {
	basePath, fset, allPkgs, err := z.scanSourceTree(c)
	if err != nil {
		return
	}
	l := c.Log().With(zap.String("path", basePath))

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

func (z *Catalogue) scanMessages(c app_control.Control) {
	basePath, fset, allPkgs, err := z.scanSourceTree(c)
	if err != nil {
		return
	}
	l := c.Log().With(zap.String("path", basePath))

	cfg := &types.Config{
		Error: func(err error) {
			l.Debug("error", zap.Error(err))
		},
		Importer: importer.Default(),
	}

	type MsgDef struct {
		Package string
		Name    string
	}
	msgDefs := make([]*MsgDef, 0)
	msgPkgs := make(map[string]bool)

	for _, pkg := range allPkgs {
		for f0, f := range pkg.Files {
			l.Debug("scan files", zap.String("f0", f0))
			for _, r := range f.Scope.Objects {
				l.Debug("finding message object", zap.String("r", r.Name))
				info := types.Info{
					Types: make(map[ast.Expr]types.TypeAndValue),
					Defs:  make(map[*ast.Ident]types.Object),
					Uses:  make(map[*ast.Ident]types.Object),
				}
				q, err := cfg.Check(basePath, fset, []*ast.File{f}, &info)
				l.Debug("check error", zap.Error(err))
				for _, qn := range q.Scope().Names() {
					ro := q.Scope().Lookup(qn)
					if !strings.HasPrefix(qn, "Msg") {
						continue
					}
					if _, ok := ro.Type().Underlying().(*types.Struct); !ok {
						continue
					}
					rel, _ := filepath.Rel(ro.Pkg().Path(), filepath.Dir(f0))

					msgDefs = append(msgDefs, &MsgDef{
						Package: rel,
						Name:    qn,
					})
					msgPkgs[rel] = true
				}
			}
		}
	}

	msgPkgSorted := make([]string, 0)
	msgPkgAlias := make(map[string]string)
	msgObjects := make(map[string]bool)
	msgObjectSorted := make([]string, 0)

	for mp := range msgPkgs {
		msgPkgSorted = append(msgPkgSorted, mp)
		msgPkgAlias[mp] = strings.ReplaceAll(mp, "/", "_")
	}
	sort.Strings(msgPkgSorted)

	for _, mo := range msgDefs {
		mod := msgPkgAlias[mo.Package] + "." + mo.Name
		msgObjects[mod] = true
	}
	for mo := range msgObjects {
		msgObjectSorted = append(msgObjectSorted, mo)
	}
	sort.Strings(msgObjectSorted)

	l.Info("Generate test cases for message object")
	t0, err := c.Resource("catalogue_message.go.tmpl")
	if err != nil {
		l.Error("err", zap.Error(err))
		return
	}
	t1, err := template.New("catalogue").Parse(string(t0))
	if err != nil {
		l.Error("err", zap.Error(err))
		return
	}

	outPath := filepath.Join(basePath, "catalogue/message.go")
	l.Info("Generating catalogue", zap.String("path", outPath))
	f, err := os.Create(outPath)
	if err != nil {
		l.Error("Failed to generate catalogue", zap.Error(err))
		return
	}
	err = t1.Execute(f, map[string]interface{}{
		"ImportAliases": msgPkgAlias,
		"Imports":       msgPkgSorted,
		"Objects":       msgObjectSorted,
	})
	if err != nil {
		l.Error("Unable to execute template", zap.Error(err))
	}
	f.Close()
}

func (z *Catalogue) Exec(c app_control.Control) error {
	z.scanRecipe(c)
	z.scanMessages(c)
	return nil
}

func (z *Catalogue) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Catalogue{}, rc_recipe.NoCustomValues)
}
