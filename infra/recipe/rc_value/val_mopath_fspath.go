package rc_value

import (
	"errors"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/islet/estring/ecase"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"os"
	"reflect"
)

var (
	ErrorPathNotFound = errors.New("path not found")
)

func newValueMoPathFileSystemPath(name string) rc_recipe.Value {
	v := &ValueMoPathFileSystemPath{name: name}
	v.path = mo_path2.NewFileSystemPath("")
	return v
}

type ValueMoPathFileSystemPath struct {
	name     string
	filePath string
	path     mo_path2.FileSystemPath
}

func (z *ValueMoPathFileSystemPath) Spec() (typeName string, typeAttr interface{}) {
	se := false
	if efs, ok := z.path.(mo_path2.ExistingFileSystemPath); ok {
		se = efs.ShouldExist()
	}
	return es_reflect.Key(app.Pkg, z.path), map[string]bool{
		"shouldExist": se,
	}
}

func (z *ValueMoPathFileSystemPath) ValueText() string {
	return z.filePath
}

func (z *ValueMoPathFileSystemPath) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*mo_path2.FileSystemPath)(nil)).Elem()) {
		return newValueMoPathFileSystemPath(name)
	}
	return nil
}

func (z *ValueMoPathFileSystemPath) Bind() interface{} {
	return &z.filePath
}

func (z *ValueMoPathFileSystemPath) Init() (v interface{}) {
	return z.path
}

func (z *ValueMoPathFileSystemPath) ApplyPreset(v0 interface{}) {
	z.path = v0.(mo_path2.FileSystemPath)
	z.filePath = z.path.Path()
}

func (z *ValueMoPathFileSystemPath) Apply() (v interface{}) {
	l := esl.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.filePath)
	if err != nil {
		p = z.filePath
		l.Debug("Unable to format", esl.String("path", z.filePath), esl.Error(err))
	}
	z.path = mo_path2.NewFileSystemPath(p)
	return z.path
}

func (z *ValueMoPathFileSystemPath) Debug() interface{} {
	return map[string]string{
		"path": z.filePath,
	}
}

func (z *ValueMoPathFileSystemPath) Capture(ctl app_control.Control) (v interface{}, err error) {
	return z.filePath, nil
}

func (z *ValueMoPathFileSystemPath) Restore(v es_json.Json, ctl app_control.Control) error {
	if w, found := v.String(); found {
		z.filePath = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *ValueMoPathFileSystemPath) SpinUp(ctl app_control.Control) error {
	l := ctl.Log().With(esl.String("path", z.filePath))
	ui := ctl.UI()

	if z.filePath == "" {
		return ErrorMissingRequiredOption
	}

	if e, ok := z.path.(mo_path2.ExistingFileSystemPath); ok && e.ShouldExist() {
		l.Debug("verify the given file")
		ls, err := os.Lstat(z.filePath)
		if err != nil {
			ui.Error(MRepository.ErrorMoPathFsPathNotFound.With("Path", z.filePath).With("Key", ecase.ToLowerKebabCase(z.name)))
			l.Debug("The file is not found", esl.Error(err))
			return ErrorPathNotFound
		}
		l.Debug("The file found", esl.Int64("size", ls.Size()), esl.Bool("isDir", ls.IsDir()))
	}
	return nil
}

func (z *ValueMoPathFileSystemPath) SpinDown(ctl app_control.Control) error {
	return nil
}
