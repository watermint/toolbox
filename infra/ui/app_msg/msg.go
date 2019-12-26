package app_msg

import (
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/app"
	"reflect"
	"strings"
)

type P map[string]interface{}

type Message interface {
	Key() string
	Params() []P
	With(key string, value interface{}) Message
	AsOptional() MessageOptional
}

type MessageOptional interface {
	Message
	Optional()
}

type messageImpl struct {
	K string
	P []P
}

func (z *messageImpl) With(key string, value interface{}) Message {
	np := make([]P, 0)
	np = append(np, P{key: value})
	np = append(np, z.P...)
	return &messageImpl{
		K: z.K,
		P: np,
	}
}

func (z *messageImpl) AsOptional() MessageOptional {
	return &messageOptionalImpl{
		K: z.K,
		P: z.P,
	}
}

func (z *messageImpl) Key() string {
	return z.K
}

func (z *messageImpl) Params() []P {
	return z.P
}

type messageOptionalImpl struct {
	K string
	P []P
}

func (z *messageOptionalImpl) Optional() {
}

func (z *messageOptionalImpl) AsOptional() MessageOptional {
	return &messageOptionalImpl{
		K: z.K,
		P: z.P,
	}
}

func (z *messageOptionalImpl) With(key string, value interface{}) Message {
	np := make([]P, 0)
	np = append(np, P{key: value})
	np = append(np, z.P...)
	return &messageOptionalImpl{
		K: z.K,
		P: np,
	}
}

func (z *messageOptionalImpl) Key() string {
	return z.K
}

func (z *messageOptionalImpl) Params() []P {
	return z.P
}

func M(key string, p ...P) Message {
	return &messageImpl{
		K: key,
		P: p,
	}
}

func Raw(text string) Message {
	return &messageImpl{
		K: "raw",
		P: []P{
			{
				"Raw": text,
			},
		},
	}
}

func applyReflect(mot reflect.Type, mov reflect.Value) {
	base := mot.PkgPath() + "." + strcase.ToSnake(mot.Name())
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
				K: base + "." + strcase.ToSnake(kn),
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
