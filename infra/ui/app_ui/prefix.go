package app_ui

import (
	"github.com/watermint/toolbox/infra/app"
	"reflect"
	"strings"
)

func prefixFor(p interface{}) string {
	if p == nil {
		return ""
	}
	path := make([]string, 0)

	rt := reflect.ValueOf(p).Elem().Type()
	pkg := rt.PkgPath()
	pkg = strings.ReplaceAll(pkg, app.Pkg, "")
	if strings.HasPrefix(pkg, "/") {
		pkg = pkg[1:]
	}
	if pkg != "" {
		path = append(path, strings.Split(pkg, "/")...)
	}
	path = append(path, strings.ToLower(rt.Name()))
	return strings.Join(path, ".")
}
