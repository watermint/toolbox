package rc_vo_impl

import (
	"flag"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"reflect"
	"runtime"
	"strings"
)

// Deprecated: use rc_value.ValueRepository
type ValueContainer struct {
	PkgName string
	Values  map[string]interface{}
}

// Deprecated: use rc_value.NewValueRepository
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
	z.PkgName = vot.PkgPath() + "." + strcase.ToSnake(vot.Name())

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
			case vof.Type.Implements(reflect.TypeOf((*fd_file.ModelFile)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = fd_file_impl.NewData()
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessMgmt)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = rc_conn_impl.NewConnBusinessMgmt()
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessInfo)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = rc_conn_impl.NewConnBusinessInfo()
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessAudit)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = rc_conn_impl.NewConnBusinessAudit()
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessFile)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = rc_conn_impl.NewConnBusinessFile()
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnUserFile)(nil)).Elem()):
				if !vvf.IsNil() {
					z.Values[kn] = vvf.Interface()
				} else {
					z.Values[kn] = rc_conn_impl.NewConnUserFile()
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
			case vof.Type.Implements(reflect.TypeOf((*fd_file.ModelFile)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("Unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessMgmt)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessInfo)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessFile)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessAudit)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vvf.Set(reflect.ValueOf(v))
				} else {
					ll.Debug("unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnUserFile)(nil)).Elem()):
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
	pkg = strings.ReplaceAll(pkg, app.Pkg+"/", "")
	pkg = strings.ReplaceAll(pkg, "/", ".")
	return pkg + ".flag." + strcase.ToSnake(name)
}

func (z *ValueContainer) MakeFlagSet(f *flag.FlagSet, ui app_ui.UI) {
	for n, d := range z.Values {
		kf := strcase.ToKebab(n)
		desc := ui.Text(z.MessageKey(n))

		switch dv := d.(type) {
		case *bool:
			f.BoolVar(dv, kf, *dv, desc)
		case *int64:
			f.Int64Var(dv, kf, *dv, desc)
		case *string:
			f.StringVar(dv, kf, *dv, desc)
		case *fd_file_impl.CsvData:
			f.StringVar(&dv.FilePath, kf, dv.FilePath, desc)
		case *rc_conn_impl.ConnBusinessMgmt:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *rc_conn_impl.ConnBusinessInfo:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *rc_conn_impl.ConnBusinessAudit:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *rc_conn_impl.ConnBusinessFile:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		case *rc_conn_impl.ConnUserFile:
			f.StringVar(&dv.PeerName, kf, dv.PeerName, desc)
		}
	}
}

func (z ValueContainer) Serialize() map[string]interface{} {
	s := make(map[string]interface{})
	for n, d := range z.Values {
		switch dv := d.(type) {
		case *bool:
			s[n] = *dv
		case *int64:
			s[n] = *dv
		case *string:
			s[n] = *dv
		case *fd_file_impl.CsvData:
			s[n] = dv.FilePath
		case *rc_conn_impl.ConnBusinessMgmt:
			s[n] = dv.PeerName
		case *rc_conn_impl.ConnBusinessInfo:
			s[n] = dv.PeerName
		case *rc_conn_impl.ConnBusinessAudit:
			s[n] = dv.PeerName
		case *rc_conn_impl.ConnBusinessFile:
			s[n] = dv.PeerName
		case *rc_conn_impl.ConnUserFile:
			s[n] = dv.PeerName
		}
	}
	return s
}
