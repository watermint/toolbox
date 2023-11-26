package es_reflect

import (
	"github.com/watermint/toolbox/essentials/islet/estring/ecase"
	"reflect"
	"strings"
)

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
	return path, ecase.ToLowerSnakeCase(rt.Name())
}

func Key(base string, r interface{}) string {
	path, name := Path(base, r)
	return strings.Join(append(path, name), ".")
}

func NewInstance(v interface{}) interface{} {
	t := reflect.ValueOf(v).Elem().Type()
	return reflect.New(t).Interface()
}
