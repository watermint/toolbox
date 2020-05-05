package es_filecompare

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"os"
)

type Comparator interface {
	Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error)
}

type SizeComparator struct {
	l es_log.Logger
}

func (z SizeComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(es_log.String("localPath", localPath), es_log.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		if f.Size == localFile.Size() {
			l.Debug("Same file size", es_log.Int64("size", localFile.Size()))
			return true, nil
		}
		l.Debug("Size diff found", es_log.Int64("localFileSize", localFile.Size()), es_log.Int64("dbxFileSize", f.Size))
		return true, nil
	}
	l.Debug("Fallback")
	return false, nil
}

type TimeComparator struct {
	l es_log.Logger
}

func (z TimeComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(es_log.String("localPath", localPath), es_log.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lt := dbx_util.RebaseTime(localFile.ModTime())
		dt, err := dbx_util.Parse(f.ClientModified)
		if err != nil {
			l.Debug("Unable to parse client modified", es_log.Error(err))
			return false, err
		}
		if lt.Equal(dt) {
			l.Debug("Same modified time", es_log.String("clientModified", dt.String()))
			return true, nil
		}
		l.Debug("Modified time diff found",
			es_log.String("localModTime", lt.String()),
			es_log.String("dbxModTime", dt.String()),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

type HashComparator struct {
	l es_log.Logger
}

func (z HashComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(es_log.String("localPath", localPath), es_log.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lch, err := dbx_util.ContentHash(localPath)
		if err != nil {
			l.Debug("Unable to calc local file content hash", es_log.Error(err))
			return false, err
		}
		if lch == f.ContentHash {
			l.Debug("Same content hash", es_log.String("hash", f.ContentHash))
			return true, nil
		}

		l.Debug("Content hash diff found",
			es_log.String("localFileHash", lch),
			es_log.String("dbxFileHash", f.ContentHash),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

// Returns true if it determined as same file
func Compare(l es_log.Logger, localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
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
