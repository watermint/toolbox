package rc_value

import (
	"compress/gzip"
	"github.com/iancoleman/strcase"
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

type FileRowFeedData struct {
	BackupId   string `path:"backup_id" json:"backup_id"`
	BackupName string `path:"backup_name" json:"backup_name"`
	BackupPath string `path:"backup_path" json:"backup_path"`
	SourcePath string `path:"source_path" json:"source_path"`
	SourceExt  string `path:"source_ext" json:"source_ext"`
}

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

func (z *ValueFdFileRowFeed) Capture(ctl app_control.Control) (v interface{}, err error) {
	filePath := z.path
	if z.path == "" {
		filePath = z.rf.FilePath()
	}
	l := ctl.Log().With(esl.String("filePath", filePath))

	if filePath == "" {
		l.Debug("No file path")
		return nil, nil
	}

	f, err := os.Open(filePath)
	if err != nil {
		l.Debug("Unable to open the feed file", esl.Error(err))
		return nil, err
	}
	defer func() {
		ioErr := f.Close()
		l.Debug("Source closed", esl.Error(ioErr))
	}()

	backupId := sc_random.MustGenerateRandomString(8)
	backupName := FeedBackupFilePrefix + backupId + ".gz"
	backupPath := filepath.Join(ctl.Workspace().Log(), backupName)
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
	size, ioErr := io.Copy(backupStream, f)
	if ioErr != nil {
		l.Debug("Unable to copy", esl.Error(ioErr))
		return nil, ioErr
	}
	ioErr = backupStream.Close()
	l.Debug("Backup completed", esl.Int64("size", size), esl.Error(ioErr))

	return &FileRowFeedData{
		BackupId:   backupId,
		BackupName: backupName,
		BackupPath: backupPath,
		SourcePath: z.rf.FilePath(),
		SourceExt:  filepath.Ext(z.rf.FilePath()),
	}, nil
}

func (z *ValueFdFileRowFeed) Restore(v es_json.Json, ctl app_control.Control) error {
	l := ctl.Log()
	rfd := &FileRowFeedData{}
	if err := v.Model(rfd); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return err
	}

	backupPath := filepath.Join(ctl.Workspace().Log(), rfd.BackupName)
	restorePath := filepath.Join(ctl.Workspace().Job(), rfd.BackupId+rfd.SourceExt)

	l.Debug("Restore from the backup",
		esl.Any("data", rfd),
		esl.String("backupPath", backupPath),
		esl.String("restorePath", restorePath),
	)

	backupFile, err := os.Open(backupPath)
	if err != nil {
		l.Debug("Unable to open the backup", esl.Error(err))
		return err
	}
	defer func() {
		ioErr := backupFile.Close()
		l.Debug("backup file closed", esl.Error(ioErr))
	}()

	backupStream, err := gzip.NewReader(backupFile)
	if err != nil {
		l.Debug("Unable to read the archive", esl.Error(err))
		return err
	}

	restoreFile, err := os.Create(restorePath)
	if err != nil {
		l.Debug("Unable to create restore file", esl.Error(err))
		return err
	}
	defer func() {
		ioErr := restoreFile.Close()
		l.Debug("Restore file closed", esl.Error(ioErr))
	}()

	size, ioErr := io.Copy(restoreFile, backupStream)
	if ioErr != nil {
		l.Debug("Unable to copy", esl.Error(ioErr))
		return ioErr
	}

	l.Debug("Restore completed, file path now points to restore path", esl.Int64("size", size))
	z.rf.SetFilePath(restorePath)

	return nil
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
