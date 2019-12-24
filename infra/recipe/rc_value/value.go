package rc_value

import (
	"errors"
	"flag"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_model_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

func NewValueRepository() *ValueRepository {
	vc := &ValueRepository{
		Values:  make(map[string]interface{}),
		Reports: make(map[string]rp_model.Report),
	}
	return vc
}

type ValueTime struct {
	Time string
}

type ValueDropboxPath struct {
	Path string
}

type ValueRepository struct {
	PkgName string
	Values  map[string]interface{}
	Reports map[string]rp_model.Report
}

func (z *ValueRepository) Feeds() map[string]fd_file.RowFeed {
	feeds := make(map[string]fd_file.RowFeed)
	for _, v := range z.Values {
		switch vv := v.(type) {
		case *fd_file_impl.RowFeed:
			feeds[vv.Spec().Name()] = vv
		}
	}
	return feeds
}

func (z *ValueRepository) FeedSpecs() map[string]fd_file.Spec {
	feeds := make(map[string]fd_file.Spec)
	for _, v := range z.Values {
		switch vv := v.(type) {
		case *fd_file_impl.RowFeed:
			spec := vv.Spec()
			feeds[spec.Name()] = spec
		}
	}
	return feeds
}

func (z *ValueRepository) Fork(ctl app_control.Control) *ValueRepository {
	vals := make(map[string]interface{})
	reps := make(map[string]rp_model.Report)
	for k, v := range z.Values {
		switch vv := v.(type) {
		case *ValueTime:
			vals[k] = &ValueTime{Time: vv.Time}
		case *ValueDropboxPath:
			vals[k] = &ValueDropboxPath{Path: vv.Path}
		case *fd_file_impl.RowFeed:
			vals[k] = vv.Fork()
		default:
			vals[k] = v
		}
	}
	for k, v := range z.Reports {
		switch vv := v.(type) {
		case *rp_model_impl.RowReport:
			reps[k] = vv.Fork(ctl)
		case *rp_model_impl.TransactionReport:
			reps[k] = vv.Fork(ctl)
		}
	}
	return &ValueRepository{
		PkgName: z.PkgName,
		Values:  vals,
		Reports: reps,
	}
}

func (z *ValueRepository) Init(vo interface{}) error {
	l := app_root.Log()

	vot := reflect.TypeOf(vo)
	vov := reflect.ValueOf(vo)

	// follow pointer
	if vot.Kind() == reflect.Ptr {
		vot = reflect.ValueOf(vo).Elem().Type()
		vov = reflect.ValueOf(vo).Elem()
	}

	if vot.Kind() != reflect.Struct {
		l.Error("ValueObject is not a struct", zap.String("name", vot.Name()), zap.String("pkg", vot.PkgPath()))
		return errors.New("given obj is not a struct")
	}
	z.PkgName = vot.PkgPath() + "." + strcase.ToSnake(vot.Name())

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
			case vof.Type.Implements(reflect.TypeOf((*mo_time.Time)(nil)).Elem()):
				ll.Debug("init mo_time.Time instance")
				vvf.Set(reflect.ValueOf(mo_time.Zero()))
				z.Values[kn] = &ValueTime{}

			case vof.Type.Implements(reflect.TypeOf((*mo_path.DropboxPath)(nil)).Elem()):
				ll.Debug("init mo_path.DropboxPath instance")
				vvf.Set(reflect.ValueOf(mo_path.NewDropboxPath("")))
				z.Values[kn] = &ValueDropboxPath{}

			case vof.Type.Implements(reflect.TypeOf((*fd_file.RowFeed)(nil)).Elem()):
				ll.Debug("init fd_file.RowFeed instance")
				fd := fd_file_impl.NewRowFeed(strcase.ToSnake(kn))
				vvf.Set(reflect.ValueOf(fd))
				z.Values[kn] = fd

			case vof.Type.Implements(reflect.TypeOf((*rp_model.RowReport)(nil)).Elem()):
				ll.Debug("init rp_model.RowReport instance")
				rr := rp_model_impl.NewRowReport(strcase.ToSnake(kn))
				vvf.Set(reflect.ValueOf(rr))
				z.Reports[kn] = rr

			case vof.Type.Implements(reflect.TypeOf((*rp_model.TransactionReport)(nil)).Elem()):
				ll.Debug("init rp_model.TransactionReport instance")
				rr := rp_model_impl.NewTransactionReport(strcase.ToSnake(kn))
				vvf.Set(reflect.ValueOf(rr))
				z.Reports[kn] = rr

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessMgmt)(nil)).Elem()):
				z.Values[kn] = rc_conn_impl.NewConnBusinessMgmt()

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessInfo)(nil)).Elem()):
				z.Values[kn] = rc_conn_impl.NewConnBusinessInfo()

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessAudit)(nil)).Elem()):
				z.Values[kn] = rc_conn_impl.NewConnBusinessAudit()

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnBusinessFile)(nil)).Elem()):
				z.Values[kn] = rc_conn_impl.NewConnBusinessFile()

			case vof.Type.Implements(reflect.TypeOf((*rc_conn.ConnUserFile)(nil)).Elem()):
				z.Values[kn] = rc_conn_impl.NewConnUserFile()
			}
		}
	}
	return nil
}

func (z *ValueRepository) Apply(vo interface{}, ctl app_control.Control) error {
	l := app_root.Log()
	ui := ctl.UI()
	//defer func() {
	//	if r := recover(); r != nil {
	//		switch r0 := r.(type) {
	//		case *runtime.TypeAssertionError:
	//			l.Debug("Unable to convert type", zap.Error(r0))
	//		default:
	//			l.Debug("Unexpected error", zap.Any("r", r))
	//		}
	//	}
	//}()

	vot := reflect.TypeOf(vo)
	vov := reflect.ValueOf(vo)

	// follow pointer
	if vot.Kind() == reflect.Ptr {
		vot = reflect.ValueOf(vo).Elem().Type()
		vov = reflect.ValueOf(vo).Elem()
	}

	if vot.Kind() != reflect.Struct {
		l.Error("ValueObject is not a struct", zap.String("name", vot.Name()), zap.String("pkg", vot.PkgPath()))
		return errors.New("given obj is not a struct")
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
			case vof.Type.Implements(reflect.TypeOf((*rp_model.RowReport)(nil)).Elem()):
				rr := z.Reports[kn].(*rp_model_impl.RowReport)
				vvf.Set(reflect.ValueOf(rr.Fork(ctl)))

			case vof.Type.Implements(reflect.TypeOf((*rp_model.TransactionReport)(nil)).Elem()):
				rr := z.Reports[kn].(*rp_model_impl.TransactionReport)
				vvf.Set(reflect.ValueOf(rr.Fork(ctl)))

			case vof.Type.Implements(reflect.TypeOf((*app_msg.Message)(nil)).Elem()):
				l.Debug("Message", zap.String("name", vof.Name))

			case vof.Type.Implements(reflect.TypeOf((*mo_time.Time)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					vt := v.(*ValueTime)
					if vt.Time == "" {
						ui.Error("infra.recipe.rc_value.value.error.mo_time.empty_time", app_msg.P{
							"Key": kn,
						})
						return errors.New("please specify date/time")
					}
					t, err := mo_time.New(vt.Time)
					if err != nil {
						ui.Error("infra.recipe.rc_value.value.error.mo_time.invalid_time_format", app_msg.P{
							"Key":  kn,
							"Time": vt.Time,
						})
						return err
					}
					vvf.Set(reflect.ValueOf(t))
				} else {
					ll.Debug("Unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*mo_path.DropboxPath)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					dbxPath := v.(*ValueDropboxPath)
					vvf.Set(reflect.ValueOf(mo_path.NewDropboxPath(dbxPath.Path)))
				} else {
					ll.Debug("Unable to find value")
				}

			case vof.Type.Implements(reflect.TypeOf((*fd_file.RowFeed)(nil)).Elem()):
				if v, e := z.Values[kn]; e {
					rf := v.(fd_file.RowFeed)
					if err := rf.ApplyModel(ctl); err != nil {
						return err
					}
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
	return nil
}

func (z *ValueRepository) MessageKey(name string) string {
	pkg := z.PkgName
	pkg = strings.ReplaceAll(pkg, app.Pkg+"/", "")
	pkg = strings.ReplaceAll(pkg, "/", ".")
	return pkg + ".flag." + strcase.ToSnake(name)
}

func (z *ValueRepository) MakeFlagSet(f *flag.FlagSet, ui app_ui.UI) {
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
		case *ValueTime:
			f.StringVar(&dv.Time, kf, dv.Time, desc)
		case *ValueDropboxPath:
			f.StringVar(&dv.Path, kf, dv.Path, desc)
		case *fd_file_impl.RowFeed:
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

func (z ValueRepository) Serialize() map[string]interface{} {
	s := make(map[string]interface{})
	for n, d := range z.Values {
		switch dv := d.(type) {
		case *bool:
			s[n] = *dv
		case *int64:
			s[n] = *dv
		case *string:
			s[n] = *dv
		case *ValueTime:
			s[n] = dv.Time
		case *ValueDropboxPath:
			s[n] = dv.Path
		case *fd_file_impl.RowFeed:
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
