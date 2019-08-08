package app_vo_impl

import (
	"flag"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/experimental/app_conn"
	"github.com/watermint/toolbox/experimental/app_conn_impl"
	"github.com/watermint/toolbox/experimental/app_file"
	"github.com/watermint/toolbox/experimental/app_file_impl"
	"github.com/watermint/toolbox/experimental/app_root"
	"go.uber.org/zap"
	"reflect"
	"runtime"
	"strings"
)

type ValueContainer struct {
	PkgName string
	Values  map[string]interface{}
}

func NewValueContainer(vo interface{}) *ValueContainer {
	vc := &ValueContainer{
		Values: make(map[string]interface{}),
	}
	vc.From(vo)
	return vc
}

func (z *ValueContainer) From(vo interface{}) {
	l := app_root.Log()
	vot := reflect.TypeOf(vo)
	vov := reflect.ValueOf(vo)
	if vot.Kind() == reflect.Ptr {
		vot = reflect.ValueOf(vo).Elem().Type()
		vov = reflect.ValueOf(vo).Elem()
	}
	z.PkgName = vot.PkgPath()

	if vot.Kind() != reflect.Struct {
		l.Error("ValueObject is not a struct", zap.String("name", vot.Name()), zap.String("pkg", vot.PkgPath()))
		return
	}

	nf := vot.NumField()
	for i := 0; i < nf; i++ {
		vof := vot.Field(i)
		vvf := vov.Field(i)
		kn := vof.Name
		ll := l.With(zap.String("key", kn))

		switch vof.Type.Kind() {
		case reflect.Bool:
			v := vvf.Bool()
			z.Values[kn] = &v
		case reflect.Int:
			v := vvf.Int()
			z.Values[kn] = &v
		case reflect.String:
			v := vvf.String()
			z.Values[kn] = &v
		case reflect.Interface:
			switch {
			case vof.Type.Implements(reflect.TypeOf((*app_file.RowDataFile)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = &app_file_impl.Factory{}
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessMgmt)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = app_conn_impl.NewConnBusinessMgmt()
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessInfo)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = app_conn_impl.NewConnBusinessInfo()
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessAudit)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = app_conn_impl.NewConnBusinessAudit()
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessFile)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = app_conn_impl.NewConnBusinessFile()
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnUserFile)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = app_conn_impl.NewConnUserFile()
				}

			default:
				ll.Warn("Unsupported type", zap.Any("kind", vof.Type.Kind()))
			}

		default:
			ll.Warn("Unsupported type", zap.Any("kind", vof.Type.Kind()))
		}
	}
}

func (z *ValueContainer) Apply(vo interface{}) {
	l := app_root.Log()
	defer func() {
		if r := recover(); r != nil {
			switch r0 := r.(type) {
			case *runtime.TypeAssertionError:
				l.Debug("Unable to convert type", zap.Error(r0))
			default:
				l.Debug("Unexpected error", zap.Any("r", r))
			}
		}
	}()

	vot := reflect.TypeOf(vo)
	vov := reflect.ValueOf(vo)

	// follow pointer
	if vot.Kind() == reflect.Ptr {
		vot = reflect.ValueOf(vo).Elem().Type()
		vov = reflect.ValueOf(vo).Elem()
	}

	if vot.Kind() != reflect.Struct {
		l.Error("ValueObject is not a struct", zap.String("name", vot.Name()), zap.String("pkg", vot.PkgPath()))
		return
	}

	nf := vot.NumField()
	for i := 0; i < nf; i++ {
		vof := vot.Field(i)
		vvf := vov.Field(i)
		kn := vof.Name
		ll := l.With(zap.String("key", kn))

		switch vof.Type.Kind() {
		case reflect.Bool:
			if v, e := z.Values[kn]; e {
				vvf.SetBool(*v.(*bool))
			} else {
				ll.Debug("Unable to find value")
			}
		case reflect.Int:
			if v, e := z.Values[kn]; e {
				vvf.SetInt(*v.(*int64))
			} else {
				ll.Debug("Unable to find value")
			}
		case reflect.String:
			if v, e := z.Values[kn]; e {
				vvf.SetString(*v.(*string))
			} else {
				ll.Debug("Unable to find value")
			}
		case reflect.Interface:
			switch {
			case vof.Type.Implements(reflect.TypeOf((*app_file.RowDataFile)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("Unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessMgmt)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessInfo)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessFile)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnBusinessAudit)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*app_conn.ConnUserFile)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			default:
				ll.Warn("Unsupported type", zap.Any("kind", vof.Type.Kind()))
			}
		default:
			ll.Debug("Not supported type", zap.Any("kind", vof.Type.Kind()))
		}
	}
}

func (z *ValueContainer) MessageKey(name string) string {
	pkg := z.PkgName
	pkg = strings.ReplaceAll(pkg, "github.com/watermint/toolbox/experimental/", "")
	pkg = strings.ReplaceAll(pkg, "/", ".")
	return pkg + ".flag." + strcase.ToSnake(name)
}

func (z *ValueContainer) MakeFlagSet(f *flag.FlagSet) {
	for n, d := range z.Values {
		kf := strcase.ToKebab(n)
		desc := z.MessageKey(n)

		switch dv := d.(type) {
		case *bool:
			f.BoolVar(dv, kf, *dv, desc)
		case *int64:
			f.Int64Var(dv, kf, *dv, desc)
		case *string:
			f.StringVar(dv, kf, *dv, desc)
		case *app_file_impl.Factory:
			f.StringVar(&dv.FilePath, kf, dv.FilePath, desc)
		case *app_conn_impl.ConnBusinessMgmt:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *app_conn_impl.ConnBusinessInfo:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *app_conn_impl.ConnBusinessAudit:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *app_conn_impl.ConnBusinessFile:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *app_conn_impl.ConnUserFile:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		}
	}
}
