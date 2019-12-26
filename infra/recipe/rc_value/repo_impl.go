package rc_value

import (
	"flag"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"reflect"
	"sort"
	"strings"
)

var (
	valueTypes = []Value{
		newValueBool(),
		newValueInt(),
		newValueString(),
		newValueAppMsgMessage("", app_msg.Raw("")),
		newValueMoTimeTime(""),
		newValueMoPathDropboxPath(""),
		newValueMoPathFileSystemPath(""),
		newValueRcRecipeRecipe("", nil),
		newValueRcConnBusinessInfo(rc_conn_impl.DefaultPeerName),
		newValueRcConnBusinessMgmt(rc_conn_impl.DefaultPeerName),
		newValueRcConnBusinessFile(rc_conn_impl.DefaultPeerName),
		newValueRcConnBusinessAudit(rc_conn_impl.DefaultPeerName),
		newValueRcConnUserFile(rc_conn_impl.DefaultPeerName),
		newValueRpModelRowReport(""),
		newValueRpModelTransactionReport(""),
		newValueFdFileRowFeed(""),
	}
)

// Find value of type.
// Returns nil when the value type is not supported
func valueOfType(t reflect.Type, r rc_recipe.Recipe, name string) Value {
	for _, vt := range valueTypes {
		if v := vt.Accept(t, r, name); v != nil {
			return v
		}
	}
	return nil
}

// Returns nil if the given rcp is not supported type.
func NewRepository(scr rc_recipe.Recipe) Repository {
	l := app_root.Log()

	// Create a new recipe instance
	srt := reflect.ValueOf(scr).Elem().Type()
	rcp := reflect.New(srt).Interface().(rc_recipe.SelfContainedRecipe)

	// Type
	rt := reflect.ValueOf(rcp).Elem().Type()
	rv := reflect.ValueOf(rcp).Elem()

	if rt.Kind() != reflect.Struct {
		l.Error("Recipe is not a struct", zap.String("name", rt.Name()), zap.String("pkg", rt.PkgPath()))
		return nil
	}

	vals := make(map[string]Value)
	fieldValue := make(map[string]reflect.Value)
	rcpName := rt.PkgPath() + "." + strcase.ToSnake(rt.Name())

	numField := rt.NumField()
	for i := 0; i < numField; i++ {
		var rtf reflect.StructField = rt.Field(i)
		var rvf reflect.Value = rv.Field(i)
		fn := rtf.Name
		ll := l.With(zap.String("fieldName", fn))

		vot := valueOfType(rtf.Type, rcp, fn)
		if vot != nil {
			ll.Debug("Set value", zap.Any("debug", vot.Debug()))
			vals[fn] = vot
			fieldValue[fn] = rvf

			vi := vot.Init()
			if vi != nil {
				ll.Debug("Set initial value", zap.Any("valueInstance", vi))
				switch rtf.Type.Kind() {
				case reflect.Bool:
					rvf.SetBool(vi.(bool))
				case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
					rvf.SetInt(vi.(int64))
				case reflect.String:
					rvf.SetString(vi.(string))
				default:
					rvf.Set(reflect.ValueOf(vi))
				}
			}
		} else {
			ll.Debug("Non value type")
		}
	}

	rcp.Preset()

	// Apply preset values
	for k, v := range vals {
		f := fieldValue[k]
		v.ApplyPreset(f.Interface())
	}

	// TODO: require write back to repo vals at this point

	return &repositoryImpl{
		values:     vals,
		rcp:        rcp,
		rcpName:    rcpName,
		fieldValue: fieldValue,
	}
}

type repositoryImpl struct {
	rcp        rc_recipe.Recipe
	rcpName    string
	values     map[string]Value
	fieldValue map[string]reflect.Value
}

func (z *repositoryImpl) Messages() []app_msg.Message {
	msgs := make([]app_msg.Message, 0)
	for _, v := range z.values {
		if vm, ok := v.(ValueMessage); ok {
			if m, ok := vm.Message(); ok {
				msgs = append(msgs, m)
			}
		}
	}
	return msgs
}

func (z *repositoryImpl) FieldNames() []string {
	names := make([]string, 0)
	for k, v := range z.values {
		if v.Bind() != nil {
			names = append(names, k)
		}
	}
	sort.Strings(names)
	return names
}

func (z *repositoryImpl) FieldValueText(name string) string {
	v := z.values[name]
	if cv, ok := v.(ValueCustomValueText); ok {
		return cv.ValueText()
	} else {
		av := v.Apply()
		return fmt.Sprintf("%v", av)
	}
}

func (z *repositoryImpl) Conns() map[string]rc_conn.ConnDropboxApi {
	conns := make(map[string]rc_conn.ConnDropboxApi)
	for k, v := range z.values {
		if vc, ok := v.(ValueConn); ok {
			if conn, ok := vc.Conn(); ok {
				conns[k] = conn
			}
		}
	}
	return conns
}

func (z *repositoryImpl) Feeds() map[string]fd_file.RowFeed {
	feeds := make(map[string]fd_file.RowFeed)
	for k, v := range z.values {
		if vf, ok := v.(ValueFeed); ok {
			if feed, ok := vf.Feed(); ok {
				feeds[k] = feed
			}
		}
	}
	return feeds
}

func (z *repositoryImpl) FeedSpecs() map[string]fd_file.Spec {
	feeds := make(map[string]fd_file.Spec)
	for k, v := range z.values {
		if vf, ok := v.(ValueFeed); ok {
			if feed, ok := vf.Feed(); ok {
				feeds[k] = feed.Spec()
			}
		}
	}
	return feeds
}

func (z *repositoryImpl) Reports() map[string]rp_model.Report {
	reps := make(map[string]rp_model.Report)
	for k, v := range z.values {
		if vr, ok := v.(ValueReport); ok {
			if rep, ok := vr.Report(); ok {
				reps[k] = rep
			}
		}
		if vr, ok := v.(ValueReports); ok {
			reps0 := vr.Reports()
			for k0, v0 := range reps0 {
				reps[k0] = v0
			}
		}
	}

	return reps
}

func (z *repositoryImpl) ReportSpecs() map[string]rp_model.Spec {
	reps := make(map[string]rp_model.Spec)
	for k, r := range z.Reports() {
		reps[k] = r.Spec()
	}
	return reps
}

func (z *repositoryImpl) Apply() rc_recipe.Recipe {
	for k, v := range z.values {
		fv, ok := z.fieldValue[k]
		if !ok {
			continue
		}
		av := v.Apply()
		switch fv.Type().Kind() {
		case reflect.Bool:
			fv.SetBool(av.(bool))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fv.SetInt(av.(int64))
		case reflect.String:
			fv.SetString(av.(string))
		default:
			fv.Set(reflect.ValueOf(av))
		}
	}
	return z.rcp
}

func (z *repositoryImpl) SpinUp(ctl app_control.Control) (rc_recipe.Recipe, error) {
	l := ctl.Log()
	var lastErr error
	for k, v := range z.values {
		err := v.SpinUp(ctl)
		if err != nil {
			lastErr = err
			// TODO: replace with UI message
			l.Error("Invalid argument, or unable to spin up", zap.String("key", k), zap.Error(err))
		}
	}
	if lastErr != nil {
		return nil, lastErr
	}
	return z.rcp, nil
}

func (z *repositoryImpl) SpinDown(ctl app_control.Control) error {
	l := ctl.Log()
	var lastErr error
	for k, v := range z.values {
		err := v.SpinDown(ctl)
		if err != nil {
			lastErr = err
			// TODO: replace with UI message
			l.Error("Unable to spin down", zap.String("key", k), zap.Error(err))
		}
	}
	if lastErr != nil {
		return lastErr
	}
	return nil
}

func (z *repositoryImpl) ApplyFlags(f *flag.FlagSet, ui app_ui.UI) {
	for k, v := range z.values {
		flagName := strcase.ToKebab(k)
		flagDesc := z.FieldDesc(k)

		b := v.Bind()
		if b != nil {
			switch bv := b.(type) {
			case *bool:
				f.BoolVar(bv, flagName, *bv, ui.Text(flagDesc.Key()))
			case *int64:
				f.Int64Var(bv, flagName, *bv, ui.Text(flagDesc.Key()))
			case *string:
				f.StringVar(bv, flagName, *bv, ui.Text(flagDesc.Key()))
			}
		}
	}
}

func (z *repositoryImpl) fieldMessageKey(name string) string {
	key := z.rcpName
	key = strings.ReplaceAll(key, app.Pkg+"/", "")
	key = strings.ReplaceAll(key, "/", ".")
	return key + ".flag." + strcase.ToSnake(name)
}

func (z *repositoryImpl) FieldCustomDefault(name string) app_msg.MessageOptional {
	return app_msg.M(z.fieldMessageKey(name) + ".default").AsOptional()
}

func (z *repositoryImpl) FieldDesc(name string) app_msg.Message {
	return app_msg.M(z.fieldMessageKey(name))
}

func (z *repositoryImpl) Debug() map[string]interface{} {
	dbg := make(map[string]interface{})
	for k, v := range z.values {
		dbg[k] = v.Debug()
	}
	return dbg
}
