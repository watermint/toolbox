package oper

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/poc/oper/oper_api"
	"github.com/watermint/toolbox/poc/oper/oper_auth"
	"github.com/watermint/toolbox/poc/oper/oper_msg"
	"go.uber.org/zap"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	LogFieldName     = "Logger"
	ContextFieldName = "Ctx"
)

type Operator struct {
	Context       Context
	Resource      *Resource
	Authenticator oper_auth.Authenticator
	Op            interface{}
}

func (z *Operator) Init() {
	z.Resource = z.LocateResource()
}

func (z *Operator) Title() string {
	return z.Resource.Title
}

func (z *Operator) Desc() string {
	return z.Resource.Desc
}

func (z *Operator) Tag() string {
	zt := reflect.TypeOf(z.Op)
	if zt.Kind() == reflect.Ptr {
		zt = zt.Elem()
	}
	return zt.Name()
}

func (z *Operator) InjectLog() {
	xt := reflect.TypeOf(z.Op)
	xv := reflect.ValueOf(z.Op)
	if xt.Kind() == reflect.Ptr {
		xt = xt.Elem()
		xv = xv.Elem()
	}
	if _, ok := xt.FieldByName(LogFieldName); ok {
		zvf := xv.FieldByName(LogFieldName)
		if zvf.Type().String() == "*zap.Logger" {
			zvf.Set(reflect.ValueOf(z.Context.Log()))
		}
	}
}

func (z *Operator) ConfigFlag(f *flag.FlagSet) {

}

func (z *Operator) inject(xv reflect.Value, fieldName string, v interface{}) {
	xt := xv.Type()
	if xt.Kind() == reflect.Ptr {
		xt = xt.Elem()
		xv = xv.Elem()
	}

	for i := xt.NumField() - 1; i >= 0; i-- {
		xtf := xt.Field(i)
		xvf := xv.Field(i)
		if xtf.Type.Kind() == reflect.Struct {
			z.inject(xvf, fieldName, v)
		} else if xtf.Name == fieldName {
			xvf.Set(reflect.ValueOf(v))
		}
	}
}

func (z *Operator) InjectOptDropboxAuthToken() {
	xv := reflect.ValueOf(z.Op)
	xt := xv.Type()
	if xt.Kind() != reflect.Ptr {
		return
	}
	xv = xv.Elem()
	xt = xt.Elem()
	var datType *oper_api.DropboxApiToken
	dat := reflect.TypeOf(datType).Elem()

	for i := xt.NumField() - 1; i >= 0; i-- {
		xtf := xt.Field(i)
		xvf := xv.Field(i)
		if !strings.HasPrefix(xtf.Name, "Opt") {
			continue
		}
		if !xtf.Type.Implements(dat) {
			continue
		}
		aPtr := reflect.New(xtf.Type)
		a0 := aPtr.Elem().Interface().(oper_api.DropboxApiToken)

		a1, err := z.Authenticator.Auth(a0)
		if err != nil {
			z.Log().Error(
				"Unable to inject auth",
				zap.String("field", xtf.Name),
				zap.String("tag", a0.Tag()),
				zap.String("app_key", a0.ApiKey()),
				zap.Error(err),
			)
			continue // TODO: just ignore for now
		}

		xvf.Set(reflect.ValueOf(a1))
	}
}

func (z *Operator) InjectContext() {
	z.inject(
		reflect.ValueOf(z.Op),
		ContextFieldName,
		z.Context.WithMessages(
			oper_msg.NewMessageMap(z.Resource.Messages, z.Log()),
		),
	)
}

func (z *Operator) Log() *zap.Logger {
	return z.Context.Log()
}

func (z *Operator) IsExecutable() bool {
	return z.Executable() != nil
}

func (z *Operator) Executable() Executable {
	switch e := z.Op.(type) {
	case Executable:
		return e
	}
	return nil
}

func (z *Operator) IsGroup() bool {
	return z.Group() != nil
}

func (z *Operator) Group() Group {
	switch g := z.Op.(type) {
	case Group:
		return g
	}
	return nil
}

func (z *Operator) SubOperators() []Operator {
	ops := z.Group().Operations()
	opr := make([]Operator, len(ops))
	for i, op := range ops {
		opr[i] = Operator{
			Context: z.Context,
			Op:      op,
		}
	}
	return opr
}

func (z *Operator) LocateResource() *Resource {
	xt := reflect.TypeOf(z.Op)
	if xt.Kind() == reflect.Ptr {
		xt = xt.Elem()
	}

	selfPath := reflect.TypeOf(z).Elem().PkgPath()
	rel, err := filepath.Rel(selfPath, xt.PkgPath())
	if err != nil {
		z.Log().Debug("Unable to identify rel path", zap.Error(err))
		return nil
	}
	loc := filepath.Join(rel, xt.Name()+".json")

	z.Log().Debug("Locate resource",
		zap.String("self", selfPath),
		zap.String("pkg", xt.PkgPath()),
		zap.String("name", xt.Name()),
		zap.String("rel", rel),
		zap.String("resLoc", loc),
	)

	resBytes, err := z.Context.Box().Bytes(loc)
	if err != nil {
		z.Log().Debug("Unable to find resource", zap.Error(err))
		return nil
	}

	res := &Resource{}
	err = json.Unmarshal(resBytes, res)
	if err != nil {
		z.Log().Debug("Unable to unmarshal resource", zap.Error(err))
		return nil
	}

	z.Log().Info("Loaded resource", zap.Any("res", res))

	resBytes, err = z.Context.Box().Bytes("sys_messages.json")
	if err != nil {
		z.Log().Debug("Unable to find resource `sys_messages.json`", zap.Error(err))
		return nil
	}
	msgs := make(map[string]string)
	err = json.Unmarshal(resBytes, &msgs)
	if err != nil {
		z.Log().Debug("Unable to unmarshal sys messages", zap.Error(err))
		return nil
	}
	if res.Messages == nil {
		res.Messages = msgs
	} else {
		for k, v := range msgs {
			res.Messages[k] = v
		}
	}

	return res
}
