package app_msg

import (
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/islet/estring/ecase"
	"github.com/watermint/toolbox/infra/app"
	"reflect"
	"strings"
)

func applyReflect(mot reflect.Type, mov reflect.Value) {
	base := mot.PkgPath() + "." + ecase.ToLowerSnakeCase(mot.Name())
	base = strings.ReplaceAll(base, app.Pkg+"/", "")
	base = strings.ReplaceAll(base, "/", ".")

	nf := mot.NumField()
	for i := 0; i < nf; i++ {
		mof := mot.Field(i)
		mvf := mov.Field(i)
		kn := mof.Name

		switch {
		case mof.Type.Implements(reflect.TypeOf((*Message)(nil)).Elem()):
			mvf.Set(reflect.ValueOf(&messageImpl{
				K: base + "." + ecase.ToLowerSnakeCase(kn),
			}))
		}
	}
}

func Apply(mo interface{}) interface{} {
	mot := reflect.TypeOf(mo)
	mov := reflect.ValueOf(mo)
	if mot.Kind() == reflect.Ptr {
		mot = reflect.ValueOf(mo).Elem().Type()
		mov = reflect.ValueOf(mo).Elem()
	}

	applyReflect(mot, mov)
	return mo
}

func Messages(mo interface{}) []Message {
	msgs := make([]Message, 0)
	mot := reflect.TypeOf(mo)
	mov := reflect.ValueOf(mo)
	if mot.Kind() == reflect.Ptr {
		mot = reflect.ValueOf(mo).Elem().Type()
		mov = reflect.ValueOf(mo).Elem()
	}

	nf := mot.NumField()
	for i := 0; i < nf; i++ {
		mof := mot.Field(i)
		mvf := mov.Field(i)

		switch {
		case mof.Type.Implements(reflect.TypeOf((*Message)(nil)).Elem()):
			msg := mvf.Interface().(Message)
			msgs = append(msgs, msg)
		}
	}
	return msgs
}

func ObjMessage(r interface{}, suffix string) Message {
	return CreateMessage(es_reflect.Key(app.Pkg, r) + "." + suffix)
}
