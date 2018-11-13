package oper

import (
	"encoding/json"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"path/filepath"
	"reflect"
)

const (
	LogFieldName = "Logger"
)

var (
	SupportedLanguages = []language.Tag{
		language.English,
		language.Japanese,
	}
	DefaultLanguage              = language.MustParse("en_US")
	DefaultLanguageBaseTag, _, _ = DefaultLanguage.Raw()
	DefaultLanguageBase          = DefaultLanguageBaseTag.String()
)

type Operator struct {
	Context   *Context
	Resources map[string]*Resource
	Op        interface{}
}

func (z *Operator) Init() {
	z.Resources = make(map[string]*Resource)
	z.LocateResource(DefaultLanguage)
	if z.Context.Lang != DefaultLanguage {
		z.LocateResource(z.Context.Lang)
	}
}

func (z *Operator) fetchText(f func(r *Resource) string, alt string) string {
	if r, ok := z.Resources[z.Context.LangBase()]; ok {
		s := f(r)
		if s != "" {
			return s
		}
	}
	if r, ok := z.Resources[DefaultLanguageBase]; ok {
		s := f(r)
		if s != "" {
			return s
		}
	}
	return alt
}

func (z *Operator) Title() string {
	return z.fetchText(
		func(r *Resource) string {
			return r.Title
		},
		"<no title>",
	)
}

func (z *Operator) Desc() string {
	return z.fetchText(
		func(r *Resource) string {
			return r.Desc
		},
		"<no title>",
	)
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
			zvf.Set(reflect.ValueOf(z.Context.Logger))
		}
	}
}

func (z *Operator) Log() *zap.Logger {
	return z.Context.Logger
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

func (z *Operator) LocateResource(lang language.Tag) *Resource {
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
	langPart := ""
	langBaseTag, _, _ := lang.Raw()
	langBase := langBaseTag.String()
	if DefaultLanguage != lang {
		langPart = "_" + langBase
	}

	loc := filepath.Join(rel, xt.Name()+langPart+".json")

	z.Log().Debug("Locate resource",
		zap.String("self", selfPath),
		zap.String("pkg", xt.PkgPath()),
		zap.String("name", xt.Name()),
		zap.String("rel", rel),
		zap.String("resLoc", loc),
		zap.Any("lang", lang),
	)

	resBytes, err := z.Context.Box.Bytes(loc)
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

	z.Resources[langBase] = res

	return res
}
