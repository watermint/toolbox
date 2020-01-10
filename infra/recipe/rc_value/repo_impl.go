package rc_value

import (
	"errors"
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
	valueTypes = []rc_recipe.Value{
		newValueBool(),
		newValueInt(),
		newValueString(),
		newValueKvStorageStorage(""),
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

	ErrorMissingRequiredOption = errors.New("missing required option")
)

// Find value of type.
// Returns nil when the value type is not supported
func valueOfType(t reflect.Type, r interface{}, name string) rc_recipe.Value {
	for _, vt := range valueTypes {
		if v := vt.Accept(t, r, name); v != nil {
			return v
		}
	}
	return nil
}

// Returns nil if the given rcp is not supported type.
func NewRepository(scr interface{}) rc_recipe.Repository {
	l := app_root.Log()

	// Create a new recipe instance
	srt := reflect.ValueOf(scr).Elem().Type()
	rcp := reflect.New(srt).Interface()

	// Type
	rt := reflect.ValueOf(rcp).Elem().Type()
	rv := reflect.ValueOf(rcp).Elem()

	if rt.Kind() != reflect.Struct {
		l.Error("Recipe is not a struct", zap.String("name", rt.Name()), zap.String("pkg", rt.PkgPath()))
		return nil
	}

	vals := make(map[string]rc_recipe.Value)
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

	if cr, ok := rcp.(rc_recipe.SelfContainedRecipe); ok {
		cr.Preset()
	}

	// Apply preset values
	for k, v := range vals {
		f := fieldValue[k]
		v.ApplyPreset(f.Interface())
	}

	return &repositoryImpl{
		values:     vals,
		rcp:        rcp,
		rcpName:    rcpName,
		fieldValue: fieldValue,
	}
}

type repositoryImpl struct {
	rcp        interface{}
	rcpName    string
	values     map[string]rc_recipe.Value
	fieldValue map[string]reflect.Value
}

func (z *repositoryImpl) FieldValue(name string) rc_recipe.Value {
	return z.values[name]
}

func (z *repositoryImpl) Messages() []app_msg.Message {
	msgs := make([]app_msg.Message, 0)
	for k, v := range z.values {
		if vm, ok := v.(rc_recipe.ValueMessage); ok {
			if m, ok := vm.Message(); ok {
				msgs = append(msgs, m)
			}
		}
		if _, ok := v.(rc_recipe.ValueConn); ok {
			if k != "Peer" {
				msgs = append(msgs, app_msg.ObjMessage(z.rcp, "conn."+strcase.ToSnake(k)))
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
	if cv, ok := v.(rc_recipe.ValueCustomValueText); ok {
		return cv.ValueText()
	} else {
		av := v.Apply()
		return fmt.Sprintf("%v", av)
	}
}

func (z *repositoryImpl) Conns() map[string]rc_conn.ConnDropboxApi {
	conns := make(map[string]rc_conn.ConnDropboxApi)
	for k, v := range z.values {
		if vc, ok := v.(rc_recipe.ValueConn); ok {
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
		if vf, ok := v.(rc_recipe.ValueFeed); ok {
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
		if vf, ok := v.(rc_recipe.ValueFeed); ok {
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
		if vr, ok := v.(rc_recipe.ValueReport); ok {
			if rep, ok := vr.Report(); ok {
				reps[k] = rep
			}
		}
		if vr, ok := v.(rc_recipe.ValueReports); ok {
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
	if r, ok := z.rcp.(rc_recipe.Recipe); ok {
		return r
	} else {
		return nil
	}
}

func (z *repositoryImpl) SpinUp(ctl app_control.Control) (rc_recipe.Recipe, error) {
	l := ctl.Log()
	ui := ctl.UI()

	valKeys := make([]string, 0)
	for k := range z.values {
		valKeys = append(valKeys, k)
	}
	sort.Strings(valKeys)

	for _, k := range valKeys {
		v := z.values[k]
		promot := false
		if _, ok := v.(rc_recipe.ValueConn); ok {
			if k != "Peer" {
				ui.Header(app_msg.ObjMessage(z.rcp, "conn."+strcase.ToSnake(k)))
				promot = true
			}
		}

		err := v.SpinUp(ctl)
		switch err {
		case nil:
			if promot {
				ui.Info(MRepository.ProgressDoneValueInitialization)
			}
			continue

		case ErrorMissingRequiredOption:
			ui.Error(MRepository.ErrorMissingRequiredOption.With("Key", strcase.ToSnake(k)))
			return nil, err

		default:
			// TODO: replace with UI message
			l.Error("Invalid argument, or unable to spin up", zap.String("key", k), zap.Error(err))
			return nil, err
		}
	}

	if r, ok := z.rcp.(rc_recipe.Recipe); ok {
		return r, nil
	} else {
		return nil, nil
	}
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
				f.BoolVar(bv, flagName, *bv, ui.TextK(flagDesc.Key()))
			case *int64:
				f.Int64Var(bv, flagName, *bv, ui.TextK(flagDesc.Key()))
			case *string:
				f.StringVar(bv, flagName, *bv, ui.TextK(flagDesc.Key()))
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
