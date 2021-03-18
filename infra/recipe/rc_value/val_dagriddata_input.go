package rc_value

import (
	"compress/gzip"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_griddata"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"io"
	"os"
	"path/filepath"
	"reflect"
)

func newValueDaGridDataInput(recipe interface{}, name string) rc_recipe.Value {
	return &ValueDaGridDataInput{
		recipe:  recipe,
		gdInput: da_griddata.NewInput(recipe, name),
		name:    name,
		path:    "",
	}
}

type ValueDaGridDataInput struct {
	recipe  interface{}
	gdInput da_griddata.GridDataInput
	name    string
	path    string
}

func (z *ValueDaGridDataInput) ValueText() string {
	return ""
}

func (z *ValueDaGridDataInput) Accept(recipe interface{}, t reflect.Type, v0 interface{}, name string) rc_recipe.Value {
	if t.Implements(reflect.TypeOf((*da_griddata.GridDataInput)(nil)).Elem()) {
		return newValueDaGridDataInput(recipe, name)
	}
	return nil
}

func (z *ValueDaGridDataInput) Bind() interface{} {
	return &z.path
}

func (z *ValueDaGridDataInput) Init() (v interface{}) {
	return z.gdInput
}

func (z *ValueDaGridDataInput) ApplyPreset(v0 interface{}) {
	z.gdInput = v0.(da_griddata.GridDataInput)
	if z.gdInput.FilePath() != "" {
		z.path = z.gdInput.FilePath()
	}
}

func (z *ValueDaGridDataInput) Apply() (v interface{}) {
	l := esl.Default()
	p, err := es_filepath.FormatPathWithPredefinedVariables(z.path)
	if err != nil {
		p = z.path
		l.Debug("Unable to format", esl.String("path", z.path), esl.Error(err))
	}

	if p != "" {
		z.gdInput.SetFilePath(p)
	}
	return z.gdInput
}

func (z *ValueDaGridDataInput) Debug() interface{} {
	return z.gdInput.Debug()
}

func (z *ValueDaGridDataInput) captureReader(c app_control.Control, r io.ReadCloser, useBackup bool) (v interface{}, err error) {
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
		z.path = backupPath
		z.gdInput.SetFilePath(backupPath)
	}

	return &FileRowFeedData{
		BackupId:   backupId,
		BackupName: backupName,
		BackupPath: backupPath,
		SourcePath: z.gdInput.FilePath(),
		SourceExt:  filepath.Ext(z.gdInput.FilePath()),
	}, nil
}

func (z *ValueDaGridDataInput) Capture(ctl app_control.Control) (v interface{}, err error) {
	filePath := z.path

	if z.path == "" {
		filePath = z.gdInput.FilePath()
	}
	l := ctl.Log().With(esl.String("filePath", filePath))

	switch filePath {
	case "":
		l.Debug("No file path")
		return nil, nil

	case "-":
		l.Debug("Capture from STDIN")
		return z.captureReader(ctl, os.Stdin, true)

	default:
		l.Debug("Capture from a file")
		filePath, err = es_filepath.FormatPathWithPredefinedVariables(filePath)
		if err != nil {
			l.Debug("Unable to format file path", esl.Error(err))
			return nil, err
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

		return z.captureReader(ctl, f, false)
	}
}

func (z *ValueDaGridDataInput) Restore(v es_json.Json, ctl app_control.Control) error {
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
	z.gdInput.SetFilePath(restorePath)

	return nil
}

func (z *ValueDaGridDataInput) SpinUp(ctl app_control.Control) (err error) {
	if z.gdInput.FilePath() == "" {
		return ErrorMissingRequiredOption
	} else {
		return z.gdInput.Open(ctl)
	}
}

func (z *ValueDaGridDataInput) SpinDown(ctl app_control.Control) error {
	return nil
}

func (z *ValueDaGridDataInput) Spec() (typeName string, typeAttr interface{}) {
	return z.gdInput.Spec().Name(), nil
}

func (z *ValueDaGridDataInput) GridDataInput() (gd da_griddata.GridDataInput, valid bool) {
	return z.gdInput, true
}
