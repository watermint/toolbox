package rc_value

import (
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

func newValueFdFileRowFeed(name string) rc_recipe.Value {
	v := &ValueFdFileRowFeed{name: name}
	v.rf = fd_file_impl.NewRowFeed(name)
	return v
}

type ValueFdFileRowFeed struct {
	name string
	rf   fd_file.RowFeed
	path string
}

func (z *ValueFdFileRowFeed) Spec() (typeName string, typeAttr interface{}) {
	return es_reflect.Key(app.Pkg, z.rf), nil
}

func (z *ValueFdFileRowFeed) ValueText() string {
	return z.path
}

func (z *ValueFdFileRowFeed) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*fd_file.RowFeed)(nil)).Elem()) {
		return newValueFdFileRowFeed(name)
	}
	return nil
}

func (z *ValueFdFileRowFeed) Bind() interface{} {
	return &z.path
}

func (z *ValueFdFileRowFeed) Init() (v interface{}) {
	return z.rf
}

func (z *ValueFdFileRowFeed) ApplyPreset(v0 interface{}) {
	z.rf = v0.(fd_file.RowFeed)
	if z.rf.FilePath() != "" {
		z.path = z.rf.FilePath()
	}
}

func (z *ValueFdFileRowFeed) Apply() (v interface{}) {
	l := app_root.Log()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.path)
	if err != nil {
		p = z.path
		l.Debug("Unable to format", zap.String("path", z.path), zap.Error(err))
	}

	if p != "" {
		z.rf.SetFilePath(p)
	}
	return z.rf
}

func (z *ValueFdFileRowFeed) Debug() interface{} {
	return map[string]string{
		"path": z.path,
	}
}

func (z *ValueFdFileRowFeed) SpinUp(ctl app_control.Control) (err error) {
	if z.rf.FilePath() == "" {
		err = ErrorMissingRequiredOption
	} else {
		err = z.rf.Open(ctl)
	}
	if err != nil {
		ui := ctl.UI()
		ui.Break()
		ui.Header(MValFdFileRowFeed.HeadFeed.With("Name", strcase.ToSnake(z.rf.Spec().Name())))
		ui.Info(MValFdFileRowFeed.FeedDesc)

		FeedSpec(z.rf.Spec(), ctl.UI())
		return err
	}
	return nil
}

func (z *ValueFdFileRowFeed) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueFdFileRowFeed) Feed() (feed fd_file.RowFeed, valid bool) {
	return z.rf, true
}

func FeedSpec(spec fd_file.Spec, ui app_ui.UI) {
	cols := spec.Columns()
	sampleCols := make([]string, 0)
	for _, col := range cols {
		sampleCols = append(sampleCols, ui.Text(spec.ColumnExample(col)))
	}
	ui.Info(MValFdFileRowFeed.FeedSample.
		With("Header", strings.Join(cols, ",")).
		With("Body", strings.Join(sampleCols, ",")))
	ui.Break()

	t := ui.InfoTable(spec.Name())

	t.Header(
		MValFdFileRowFeed.HeadColName,
		MValFdFileRowFeed.HeadColDesc,
		MValFdFileRowFeed.HeadColExample,
	)
	for _, col := range cols {
		t.Row(
			app_msg.Raw(col),
			spec.ColumnDesc(col),
			spec.ColumnExample(col),
		)
	}
	t.Flush()
	ui.Break()
}
