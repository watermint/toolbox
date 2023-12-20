package es_reflect

import (
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"math"
	"reflect"
	"strings"
)

func PathFromMultiBase(bases []string, r interface{}) (path []string, name string) {
	minDepth := math.MaxInt
	for _, base := range bases {
		p, n := Path(base, r)
		if len(p) < minDepth {
			minDepth = len(p)
			path = p
			name = n
		}
	}
	return
}

func Path(base string, r interface{}) (path []string, name string) {
	path = make([]string, 0)

	rt := reflect.ValueOf(r).Elem().Type()
	pkg := rt.PkgPath()
	pkg = strings.ReplaceAll(pkg, base, "")
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	if pkg != "" {
		path = append(path, strings.Split(pkg, "/")...)
	}
	//	return path, strings.ToLower(rt.Name())
	return path, es_case.ToLowerSnakeCase(rt.Name())
}

var (
	KeyBasePackages = []string{
		app_definitions.Pkg,
	}
)

func Key(r interface{}) string {
	path, name := PathFromMultiBase(KeyBasePackages, r)
	return strings.Join(append(path, name), ".")
}

func KeyWithBase(base string, r interface{}) string {
	path, name := Path(base, r)
	return strings.Join(append(path, name), ".")
}

func NewInstance(v interface{}) interface{} {
	t := reflect.ValueOf(v).Elem().Type()
	return reflect.New(t).Interface()
}
