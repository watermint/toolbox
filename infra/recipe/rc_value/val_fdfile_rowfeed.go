package rc_value

import (
	"compress/gzip"
	"github.com/watermint/essentials/estring/ecase"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/feed/fd_file_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

const (
	FeedBackupFilePrefix = "backup_feed_"
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

func (z *ValueFdFileRowFeed) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
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
	l := esl.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.path)
	if err != nil {
		p = z.path
		l.Debug("Unable to format", esl.String("path", z.path), esl.Error(err))
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

func (z *ValueFdFileRowFeed) captureReader(c app_control.Control, r io.ReadCloser, useBackup bool) (v interface{}, err error) {
	l := c.Log()
	backupId := sc_random.MustGetSecureRandomString(8)
	backupName := FeedBackupFilePrefix + backupId + ".gz"
	backupPath := filepath.Join(c.Workspace().Log(), backupName)
	l.Debug("Create backup", esl.String("backupId", backupId), esl.String("backupPath", backupPath))
	backup, err := os.Create(backupPath)
	if err != nil {
		l.Debug("Unable to create the backup", esl.Error(err))
		return nil, err
	}
	defer func() {
		ioErr := backup.Close()
		l.Debug("Backup closed", esl.Error(ioErr))
	}()

	backupStream := gzip.NewWriter(backup)
	size, ioErr := io.Copy(backupStream, r)
	if ioErr != nil {
		l.Debug("Unable to copy", esl.Error(ioErr))
		return nil, ioErr
	}
	ioErr = backupStream.Close()
	l.Debug("Backup completed", esl.Int64("size", size), esl.Error(ioErr))

	if useBackup {
		l.Debug("Use backup as an input file", esl.String("path", backupPath))
	}

	return &CapturedData{
		BackupId:   backupId,
		BackupName: backupName,
		BackupPath: backupPath,
		SourcePath: z.rf.FilePath(),
		SourceExt:  filepath.Ext(z.rf.FilePath()),
	}, nil
}

func (z *ValueFdFileRowFeed) Capture(ctl app_control.Control) (v interface{}, err error) {
	filePath := z.path

	if z.path == "" {
		filePath = z.rf.FilePath()
	}
	return captureFile(ctl, filePath, func(path string) {
		z.path = path
		z.rf.SetFilePath(path)
	})
}

func (z *ValueFdFileRowFeed) Restore(v es_json.Json, ctl app_control.Control) error {
	return restoreFile(v, ctl, func(path string) {
		z.rf.SetFilePath(path)
	})
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
		ui.Header(MValFdFileRowFeed.HeadFeed.With("Name", ecase.ToLowerSnakeCase(z.rf.Spec().Name())))
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
