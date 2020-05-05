package rc_value

import (
	"errors"
	"github.com/iancoleman/strcase"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/log/es_log"
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

func (z *ValueMoPathFileSystemPath) Accept(t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
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
	l := es_log.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.filePath)
	if err != nil {
		p = z.filePath
		l.Debug("Unable to format", es_log.String("path", z.filePath), es_log.Error(err))
	}
	z.path = mo_path2.NewFileSystemPath(p)
	return z.path
}

func (z *ValueMoPathFileSystemPath) Debug() interface{} {
	return map[string]string{
		"path": z.filePath,
	}
}

func (z *ValueMoPathFileSystemPath) SpinUp(ctl app_control.Control) error {
	l := ctl.Log().With(es_log.String("path", z.filePath))
	ui := ctl.UI()

	if z.filePath == "" {
		return ErrorMissingRequiredOption
	}

	if e, ok := z.path.(mo_path2.ExistingFileSystemPath); ok && e.ShouldExist() {
		l.Debug("verify the given file")
		ls, err := os.Lstat(z.filePath)
		if err != nil {
			ui.Error(MRepository.ErrorMoPathFsPathNotFound.With("Path", z.filePath).With("Key", strcase.ToKebab(z.name)))
			l.Debug("The file is not found", es_log.Error(err))
			return ErrorPathNotFound
		}
		l.Debug("The file found", es_log.Int64("size", ls.Size()), es_log.Bool("isDir", ls.IsDir()))
	}
	return nil
}

func (z *ValueMoPathFileSystemPath) SpinDown(ctl app_control.Control) error {
	return nil
}
