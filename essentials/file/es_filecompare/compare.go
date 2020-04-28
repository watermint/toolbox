package es_filecompare

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"go.uber.org/zap"
	"os"
)

type Comparator interface {
	Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error)
}

type SizeComparator struct {
	l *zap.Logger
}

func (z SizeComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(zap.String("localPath", localPath), zap.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		if f.Size == localFile.Size() {
			l.Debug("Same file size", zap.Int64("size", localFile.Size()))
			return true, nil
		}
		l.Debug("Size diff found", zap.Int64("localFileSize", localFile.Size()), zap.Int64("dbxFileSize", f.Size))
		return true, nil
	}
	l.Debug("Fallback")
	return false, nil
}

type TimeComparator struct {
	l *zap.Logger
}

func (z TimeComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(zap.String("localPath", localPath), zap.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lt := dbx_util.RebaseTime(localFile.ModTime())
		dt, err := dbx_util.Parse(f.ClientModified)
		if err != nil {
			l.Debug("Unable to parse client modified", zap.Error(err))
			return false, err
		}
		if lt.Equal(dt) {
			l.Debug("Same modified time", zap.String("clientModified", dt.String()))
			return true, nil
		}
		l.Debug("Modified time diff found",
			zap.String("localModTime", lt.String()),
			zap.String("dbxModTime", dt.String()),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

type HashComparator struct {
	l *zap.Logger
}

func (z HashComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(zap.String("localPath", localPath), zap.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lch, err := dbx_util.ContentHash(localPath)
		if err != nil {
			l.Debug("Unable to calc local file content hash", zap.Error(err))
			return false, err
		}
		if lch == f.ContentHash {
			l.Debug("Same content hash", zap.String("hash", f.ContentHash))
			return true, nil
		}

		l.Debug("Content hash diff found",
			zap.String("localFileHash", lch),
			zap.String("dbxFileHash", f.ContentHash),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

// Returns true if it determined as same file
func Compare(l *zap.Logger, localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	sc := &SizeComparator{l: l}
	tc := &TimeComparator{l: l}
	hc := &HashComparator{l: l}

	eq, err := sc.Compare(localPath, localFile, dbxEntry)
	if err != nil || !eq {
		return eq, err
	}
	eq, err = tc.Compare(localPath, localFile, dbxEntry)
	if err != nil {
		return eq, err
	}
	// determine as true, if same size & time
	if eq {
		return true, nil
	}

	// otherwise, compare content hash
	eq, err = hc.Compare(localPath, localFile, dbxEntry)
	return eq, err
}
