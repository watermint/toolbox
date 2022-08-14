package rc_value

import (
	"compress/gzip"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"io"
	"os"
	"path/filepath"
)

type CapturedData struct {
	BackupId   string `path:"backup_id" json:"backup_id"`
	BackupName string `path:"backup_name" json:"backup_name"`
	BackupPath string `path:"backup_path" json:"backup_path"`
	SourcePath string `path:"source_path" json:"source_path"`
	SourceExt  string `path:"source_ext" json:"source_ext"`
}

func captureFile(c app_control.Control, origPath string, callback func(path string)) (v *CapturedData, err error) {
	l := c.Log().With(esl.String("path", origPath))
	switch origPath {
	case "":
		l.Debug("No file path")
		return nil, nil

	case "-":
		l.Debug("Capture from STDIN")
		return captureReader(c, os.Stdin, true, origPath, callback)

	default:
		if c.Feature().IsTransient() {
			callback(origPath)
			return nil, nil
		}
		l.Debug("Capture from a file")
		var formattedPath string
		formattedPath, err = es_filepath.FormatPathWithPredefinedVariables(origPath)
		if err != nil {
			l.Debug("Unable to format file path", esl.Error(err))
			return nil, err
		}
		l.Debug("Formatted path", esl.String("formattedPath", formattedPath))

		f, err := os.Open(formattedPath)
		if err != nil {
			l.Debug("Unable to open the feed file", esl.Error(err))
			return nil, err
		}
		defer func() {
			ioErr := f.Close()
			l.Debug("Source closed", esl.Error(ioErr))
		}()

		return captureReader(c, f, false, origPath, callback)
	}
}

func captureReader(c app_control.Control, r io.ReadCloser, useBackup bool, origPath string, callback func(path string)) (v *CapturedData, err error) {
	l := c.Log()
	if c.Feature().IsTransient() {
		if err := os.MkdirAll(c.Workspace().Log(), 0755); err != nil {
			l.Debug("Unable to create log path", esl.Error(err))
			return nil, err
		}
	}
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
		callback(backupPath)
	} else {
		l.Debug("Use origFile as an input file", esl.String("path", backupPath))
		callback(origPath)
	}

	return &CapturedData{
		BackupId:   backupId,
		BackupName: backupName,
		BackupPath: backupPath,
		SourcePath: origPath,
		SourceExt:  filepath.Ext(origPath),
	}, nil
}

func restoreFile(v es_json.Json, c app_control.Control, callback func(path string)) error {
	l := c.Log()
	rfd := &CapturedData{}
	if err := v.Model(rfd); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return err
	}

	backupPath := filepath.Join(c.Workspace().Log(), rfd.BackupName)
	restorePath := filepath.Join(c.Workspace().Job(), rfd.BackupId+rfd.SourceExt)

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
	callback(restorePath)

	return nil
}
