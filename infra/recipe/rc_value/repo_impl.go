package rc_value

import (
	"errors"
	"flag"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_multi"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/data/da_json"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"reflect"
	"sort"
	"strings"
)

var (
	valueTypes = []rc_recipe.Value{
		newValueAppMsgMessage("", app_msg.Raw("")),
		newValueAsConnAsana(dbx_conn_impl.DefaultPeerName),
		newValueBool(),
		newValueDbxConnScopedIndividual(dbx_conn_impl.DefaultPeerName),
		newValueDbxConnScopedTeam(dbx_conn_impl.DefaultPeerName),
		newValueDbxConnUserFile(dbx_conn_impl.DefaultPeerName),
		newValueFdFileRowFeed(""),
		newValueGhConnGithubPublic(),
		newValueGhConnGithubRepo(dbx_conn_impl.DefaultPeerName),
		newValueGoogConnMail(dbx_conn_impl.DefaultPeerName),
		newValueGoogConnSheets(dbx_conn_impl.DefaultPeerName),
		newValueDaGridDataInput(nil, ""),
		newValueDaGridDataOutput(nil, ""),
		newValueDaTextInput(nil, ""),
		newValueDaJsonInput(nil, ""),
		newValueInt(),
		newValueKvStorageStorage(""),
		newValueMoFilter(""),
		newValueMoPathDropboxPath(""),
		newValueMoPathFileSystemPath(""),
		newValueMoTimeTime(""),
		newValueMoUrlUrl(""),
		newValueOptionalString(),
		newValueRangeInt(),
		newValueRcRecipeRecipe("", nil),
		newValueRpModelRowReport(""),
		newValueRpModelTransactionReport(""),
		newValueSelectString(),
		newValueSlack(dbx_conn_impl.DefaultPeerName),
		newValueString(),
	}

	ErrorMissingRequiredOption = errors.New("missing required option")
	ErrorInvalidValue          = errors.New("invalid value")
)

// Find value of type.
// Returns nil when the value type is not supported
func valueOfType(recipe interface{}, t reflect.Type, r interface{}, name string) rc_recipe.Value {
	for _, vt := range valueTypes {
		if v := vt.Accept(recipe, t, r, name); v != nil {
			return v
		}
	}
	return nil
}

// Returns nil if the given rcp is not supported type.
func NewRepository(scr interface{}) rc_recipe.Repository {
	l := esl.Default()

	// Create a new recipe instance
	srt := reflect.ValueOf(scr).Elem().Type()
	rcp := reflect.New(srt).Interface()

	// Type
	rt := reflect.ValueOf(rcp).Elem().Type()
	rv := reflect.ValueOf(rcp).Elem()

	if rt.Kind() != reflect.Struct {
		l.Error("Recipe is not a struct", esl.String("name", rt.Name()), esl.String("pkg", rt.PkgPath()))
		return nil
	}

	vals := make(map[string]rc_recipe.Value)
	fieldValue := make(map[string]reflect.Value)
	rcpName := rt.PkgPath() + "." + strcase.ToSnake(rt.Name())

	numField := rt.NumField()
	for i := 0; i < numField; i++ {
		var rtf = rt.Field(i)
		var rvf = rv.Field(i)
		fn := rtf.Name
		ll := l.With(esl.String("fieldName", fn))

		vot := valueOfType(rcp, rtf.Type, rcp, fn)
		if vot != nil {
			ll.Debug("Set value", esl.Any("debug", vot.Debug()))
			vals[fn] = vot
			fieldValue[fn] = rvf

			vi := vot.Init()
			if vi != nil {
				ll.Debug("Set initial value", esl.Any("valueInstance", vi))
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

	if scr, ok := rcp.(rc_recipe.Preset); ok {
		l.Debug("Call preset")
		scr.Preset()
	}

	ri := &RepositoryImpl{
		values:     vals,
		rcp:        rcp,
		rcpName:    rcpName,
		fieldValue: fieldValue,
	}
	ri.ApplyCustom()

	return ri
}

type RepositoryImpl struct {
	rcp        interface{}
	rcpName    string
	values     map[string]rc_recipe.Value
	fieldValue map[string]reflect.Value
}

func (z *RepositoryImpl) ApplyCustom() {
	l := esl.Default()
	for k, v := range z.values {
		f := z.fieldValue[k]
		l.Debug("Apply preset", esl.String("k", k), esl.Any("v", f.Interface()))
		v.ApplyPreset(f.Interface())
	}
}

func (z *RepositoryImpl) Current() interface{} {
	return z.rcp
}

func (z *RepositoryImpl) FieldValue(name string) rc_recipe.Value {
	return z.values[name]
}

func (z *RepositoryImpl) Messages() []app_msg.Message {
	msgs := make([]app_msg.Message, 0)
	for k, v := range z.values {
		if vm, ok := v.(rc_recipe.ValueMessage); ok {
			if m, ok := vm.Message(); ok {
				msgs = append(msgs, m)
			}
		}
		if vm, ok := v.(rc_recipe.ValueMessages); ok {
			if m, ok := vm.Messages(); ok {
				msgs = append(msgs, m...)
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

func (z *RepositoryImpl) FieldNames() []string {
	names := make([]string, 0)
	for k, v := range z.values {
		switch b := v.Bind().(type) {
		case nil:
			continue
		case mo_multi.MultiValue:
			names = append(names, b.Fields()...)
		default:
			names = append(names, k)
		}
	}
	sort.Strings(names)
	return names
}

func (z *RepositoryImpl) FieldValueText(name string) string {
	if v, ok := z.values[name]; !ok {
		return "" // Might be nested value in MultiValue
	} else if cv, ok := v.(rc_recipe.ValueCustomValueText); ok {
		return cv.ValueText()
	} else {
		av := v.Apply()
		return fmt.Sprintf("%v", av)
	}
}

func (z *RepositoryImpl) Conns() map[string]api_conn.Connection {
	conns := make(map[string]api_conn.Connection)
	for k, v := range z.values {
		if vc, ok := v.(rc_recipe.ValueConn); ok {
			if conn, ok := vc.Conn(); ok {
				conns[k] = conn
			}
		}
		if vc, ok := v.(rc_recipe.ValueConns); ok {
			for k, c := range vc.Conns() {
				conns[k] = c
			}
		}
	}
	return conns
}

func (z *RepositoryImpl) GridDataInputSpecs() (specs map[string]da_griddata.GridDataInputSpec) {
	specs = make(map[string]da_griddata.GridDataInputSpec)
	for k, v := range z.values {
		if vf, ok := v.(rc_recipe.ValueGridDataInput); ok {
			if gd, ok := vf.GridDataInput(); ok {
				specs[k] = gd.Spec()
			}
		}
	}
	return specs
}

func (z *RepositoryImpl) GridDataOutputSpecs() (specs map[string]da_griddata.GridDataOutputSpec) {
	specs = make(map[string]da_griddata.GridDataOutputSpec)
	for k, v := range z.values {
		if vf, ok := v.(rc_recipe.ValueGridDataOutput); ok {
			if gd, ok := vf.GridDataOutput(); ok {
				specs[k] = gd.Spec()
			}
		}
	}
	return specs
}

func (z *RepositoryImpl) TextInputSpecs() map[string]da_text.TextInputSpec {
	specs := make(map[string]da_text.TextInputSpec)
	for k, v := range z.values {
		if vf, ok := v.(rc_recipe.ValueTextInput); ok {
			if v0, ok := vf.TextInput(); ok {
				specs[k] = v0.Spec()
			}
		}
	}
	return specs
}

func (z *RepositoryImpl) JsonInputSpecs() map[string]da_json.JsonInputSpec {
	specs := make(map[string]da_json.JsonInputSpec)
	for k, v := range z.values {
		if vf, ok := v.(rc_recipe.ValueJsonInput); ok {
			if v0, ok := vf.JsonInput(); ok {
				specs[k] = v0.Spec()
			}
		}
	}
	return specs
}

func (z *RepositoryImpl) Feeds() map[string]fd_file.RowFeed {
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

func (z *RepositoryImpl) FeedSpecs() map[string]fd_file.Spec {
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

func (z *RepositoryImpl) Reports() map[string]rp_model.Report {
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

func (z *RepositoryImpl) ReportSpecs() map[string]rp_model.Spec {
	reps := make(map[string]rp_model.Spec)
	for k, r := range z.Reports() {
		reps[k] = r.Spec()
	}
	return reps
}

func (z *RepositoryImpl) Apply() rc_recipe.Recipe {
	l := esl.Default()
	for k, v := range z.values {
		fv, ok := z.fieldValue[k]
		if !ok {
			continue
		}
		av := v.Apply()
		switch fv.Type().Kind() {
		case reflect.Bool:
			l.Debug("apply bool", esl.String("k", k), esl.Any("av", av))
			fv.SetBool(av.(bool))
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			l.Debug("apply int", esl.String("k", k), esl.Any("av", av))
			fv.SetInt(av.(int64))
		case reflect.String:
			l.Debug("apply string", esl.String("k", k), esl.Any("av", av))
			fv.SetString(av.(string))
		default:
			l.Debug("apply interface", esl.String("k", k), esl.Any("av", av))
			fv.Set(reflect.ValueOf(av))
		}
	}
	if r, ok := z.rcp.(rc_recipe.Recipe); ok {
		return r
	} else {
		return nil
	}
}

func (z *RepositoryImpl) Capture(ctl app_control.Control) (v interface{}, err error) {
	l := ctl.Log()
	vals := make(map[string]interface{})

	for k, val := range z.values {
		if vc, err := val.Capture(ctl); err != nil {
			l.Debug("Unable to capture value", esl.String("key", k), esl.Error(err))
			return nil, err
		} else {
			vals[k] = vc
		}
	}
	return vals, nil
}

func (z *RepositoryImpl) Restore(j es_json.Json, ctl app_control.Control) error {
	l := ctl.Log()
	if w, found := j.Object(); found {
		for k, val := range z.values {
			ll := l.With(esl.String("key", k))
			if x, ok := w[k]; ok {
				ll.Debug("Restore value")
				if err := val.Restore(x, ctl); err != nil {
					ll.Debug("Unable to restore value for the key", esl.Error(err))
					return err
				}
			} else {
				ll.Debug("Restore value not found")
				return rc_recipe.ErrorValueRestoreFailed
			}
		}
	} else {
		l.Debug("Restore value not found (not an object)")
		return rc_recipe.ErrorValueRestoreFailed
	}
	l.Debug("Restore completed")
	return nil
}

func (z *RepositoryImpl) SpinUp(ctl app_control.Control) (rc_recipe.Recipe, error) {
	l := ctl.Log()
	ui := ctl.UI()

	valKeys := make([]string, 0)
	for k := range z.values {
		valKeys = append(valKeys, k)
	}
	sort.Strings(valKeys)

	var lastErr error
	for _, k := range valKeys {
		v := z.values[k]
		prompt := false
		if _, ok := v.(rc_recipe.ValueConn); ok {
			if k != "Peer" {
				ui.Header(app_msg.ObjMessage(z.rcp, "conn."+strcase.ToSnake(k)))
				prompt = true
			}
		}
		l.Debug("spin up", esl.String("k", k), esl.Any("v.debug", v.Debug()))

		err := v.SpinUp(ctl)
		switch err {
		case nil:
			if prompt {
				ui.Info(MRepository.ProgressDoneValueInitialization)
			}
			continue

		case ErrorInvalidValue:
			ui.Error(MRepository.ErrorInvalidValue.With("Key", strcase.ToKebab(k)))
			lastErr = err

		case ErrorMissingRequiredOption:
			ui.Error(MRepository.ErrorMissingRequiredOption.With("Key", strcase.ToKebab(k)))
			lastErr = err

		default:
			// TODO: replace with UI message
			l.Error("Invalid argument, or unable to spin up", esl.String("key", k), esl.Error(err))
			lastErr = err
		}
	}
	if lastErr != nil {
		l.Debug("fail spin up")
		return nil, lastErr
	}

	if r, ok := z.rcp.(rc_recipe.Recipe); ok {
		return r, nil
	} else {
		return nil, nil
	}
}

func (z *RepositoryImpl) SpinDown(ctl app_control.Control) error {
	l := ctl.Log()
	var lastErr error
	for k, v := range z.values {
		err := v.SpinDown(ctl)
		if err != nil {
			lastErr = err
			// TODO: replace with UI message
			l.Error("Unable to spin down", esl.String("key", k), esl.Error(err))
		}
	}
	if lastErr != nil {
		return lastErr
	}
	return nil
}

func (z *RepositoryImpl) ApplyFlags(f *flag.FlagSet, ui app_ui.UI) {
	for k, v := range z.values {
		flagName := strcase.ToKebab(k)
		flagDesc := z.FieldDesc(k)

		b := v.Bind()
		if b != nil {
			switch bv := b.(type) {
			case *bool:
				f.BoolVar(bv, flagName, *bv, ui.Text(flagDesc))
			case *int64:
				f.Int64Var(bv, flagName, *bv, ui.Text(flagDesc))
			case *string:
				f.StringVar(bv, flagName, *bv, ui.Text(flagDesc))
			case mo_multi.MultiValue:
				bv.ApplyFlags(f, flagDesc, ui)
			}
		}
	}
}

func (z *RepositoryImpl) fieldMessageKeyBase() string {
	key := z.rcpName
	key = strings.ReplaceAll(key, app.Pkg+"/", "")
	key = strings.ReplaceAll(key, "/", ".")
	return key + ".flag."
}

func (z *RepositoryImpl) fieldMessageKey(name string) string {
	return z.fieldMessageKeyBase() + strcase.ToSnake(name)
}

func (z *RepositoryImpl) FieldCustomDefault(name string) app_msg.MessageOptional {
	return app_msg.CreateMessage(z.fieldMessageKey(name) + ".default").AsOptional()
}

func (z *RepositoryImpl) FieldDesc(name string) app_msg.Message {
	msg := app_msg.CreateMessage(z.fieldMessageKey(name))
	for k, v := range z.values {
		if k == name {
			return msg
		}
		switch v0 := v.Bind().(type) {
		case mo_multi.MultiValue:
			if strings.HasPrefix(name, v0.Name()) {
				base := app_msg.CreateMessage(z.fieldMessageKeyBase() + strcase.ToSnake(v0.Name()))
				return v0.FieldDesc(base, name)
			}
		}
	}
	return msg
}

func (z *RepositoryImpl) Debug() map[string]interface{} {
	dbg := make(map[string]interface{})
	for k, v := range z.values {
		dbg[k] = v.Debug()
	}
	return dbg
}
