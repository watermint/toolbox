package rc_value

import (
	"flag"
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
	"strings"
)

var (
	valueTypes = []Value{
		newValueBool(),
		newValueInt(),
		newValueString(),
		newValueMoTimeTime(""),
		newValueMoPathPath(""),
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
func valueOfType(t reflect.Type, name string) Value {
	for _, vt := range valueTypes {
		if v := vt.Accept(t, name); v != nil {
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

	// Apply messages
	app_msg.Apply(rcp)

	vals := make(map[string]Value)
	fieldValue := make(map[string]reflect.Value)
	rcpName := rt.PkgPath() + "." + strcase.ToSnake(rt.Name())

	numField := rt.NumField()
	for i := 0; i < numField; i++ {
		var rtf reflect.StructField = rt.Field(i)
		var rvf reflect.Value = rv.Field(i)
		fn := rtf.Name
		ll := l.With(zap.String("fieldName", fn))

		vot := valueOfType(rtf.Type, fn)
		if vot != nil {
			ll.Debug("Set value", zap.Any("debug", vot.Debug()))
			vals[fn] = vot
			fieldValue[fn] = rvf

			vi := vot.Init()
			if vi != nil {
				ll.Debug("Set initial value", zap.Any("valueInstance", vi))
				rvf.Set(reflect.ValueOf(vi))
			}
		} else {
			ll.Debug("Non value type")
		}
	}

	rcp.Preset()

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

func (z *repositoryImpl) FieldNames() []string {
	names := make([]string, 0)
	for k, v := range z.values {
		if v.Bind() != nil {
			names = append(names, k)
		}
	}
	return names
}

func (z *repositoryImpl) FieldValue(name string) interface{} {
	return z.values[name].Apply()
}

func (z *repositoryImpl) Conns() map[string]rc_conn.ConnDropboxApi {
	conns := make(map[string]rc_conn.ConnDropboxApi)
	for k, v := range z.values {
		if conn, ok := v.IsConn(); ok {
			conns[k] = conn
		}
	}
	return conns
}

func (z *repositoryImpl) Feeds() map[string]fd_file.RowFeed {
	feeds := make(map[string]fd_file.RowFeed)
	for k, v := range z.values {
		if feed, ok := v.IsFeed(); ok {
			feeds[k] = feed
		}
	}
	return feeds
}

func (z *repositoryImpl) FeedSpecs() map[string]fd_file.Spec {
	feeds := make(map[string]fd_file.Spec)
	for k, v := range z.values {
		if feed, ok := v.IsFeed(); ok {
			feeds[k] = feed.Spec()
		}
	}
	return feeds
}

func (z *repositoryImpl) Reports() map[string]rp_model.Report {
	reps := make(map[string]rp_model.Report)
	for k, v := range z.values {
		if rep, ok := v.IsReport(); ok {
			reps[k] = rep
		}
	}
	return reps
}

func (z *repositoryImpl) ReportSpecs() map[string]rp_model.Spec {
	reps := make(map[string]rp_model.Spec)
	for k, v := range z.values {
		if rep, ok := v.IsReport(); ok {
			reps[k] = rep.Spec()
		}
	}
	return reps
}

func (z *repositoryImpl) Fork(ctl app_control.Control) Repository {
	rep := NewRepository(z.rcp).(*repositoryImpl)
	for k, v := range z.values {
		rep.values[k] = v.Fork(ctl)
	}
	rep.Apply()
	return rep
}

func (z *repositoryImpl) Apply() rc_recipe.Recipe {
	for k, v := range z.values {
		av := v.Apply()
		if fv, ok := z.fieldValue[k]; ok && av != nil {
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
