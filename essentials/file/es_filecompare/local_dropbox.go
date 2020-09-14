package es_filecompare

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/essentials/file/es_fileentry"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type LocalDropboxComparator interface {
	Compare(localFile es_fileentry.LocalEntry, dbxEntry mo_file.Entry) (bool, error)
}

type SizeComparator struct {
	l esl.Logger
}

func (z SizeComparator) Compare(localFile es_fileentry.LocalEntry, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(esl.Any("local", localFile), esl.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		if f.Size == localFile.Size {
			l.Debug("Same file size", esl.Int64("size", localFile.Size))
			return true, nil
		}
		l.Debug("Size diff found", esl.Int64("localFileSize", localFile.Size), esl.Int64("dbxFileSize", f.Size))
		return true, nil
	}
	l.Debug("Fallback")
	return false, nil
}

type TimeComparator struct {
	l esl.Logger
}

func (z TimeComparator) Compare(localFile es_fileentry.LocalEntry, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(esl.Any("localFile", localFile), esl.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lt := dbx_util.RebaseTime(localFile.ModTime)
		dt, err := dbx_util.Parse(f.ClientModified)
		if err != nil {
			l.Debug("Unable to parse client modified", esl.Error(err))
			return false, err
		}
		if lt.Equal(dt) {
			l.Debug("Same modified time", esl.String("clientModified", dt.String()))
			return true, nil
		}
		l.Debug("Modified time diff found",
			esl.String("localModTime", lt.String()),
			esl.String("dbxModTime", dt.String()),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

type HashComparator struct {
	l esl.Logger
}

func (z HashComparator) Compare(localFile es_fileentry.LocalEntry, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(esl.Any("localFile", localFile), esl.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lch, err := dbx_util.FileContentHash(localFile.Path)
		if err != nil {
			l.Debug("Unable to calc local file content hash", esl.Error(err))
			return false, err
		}
		if lch == f.ContentHash {
			l.Debug("Same content hash", esl.String("hash", f.ContentHash))
			return true, nil
		}

		l.Debug("Content hash diff found",
			esl.String("localFileHash", lch),
			esl.String("dbxFileHash", f.ContentHash),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

// Returns true if it determined as same file
func Compare(l esl.Logger, localFile es_fileentry.LocalEntry, dbxEntry mo_file.Entry) (bool, error) {
	sc := &SizeComparator{l: l}
	tc := &TimeComparator{l: l}
	hc := &HashComparator{l: l}

	eq, err := sc.Compare(localFile, dbxEntry)
	if err != nil || !eq {
		return eq, err
	}
	eq, err = tc.Compare(localFile, dbxEntry)
	if err != nil {
		return eq, err
	}
	// determine as true, if same size & time
	if eq {
		return true, nil
	}

	// otherwise, compare content hash
	eq, err = hc.Compare(localFile, dbxEntry)
	return eq, err
}
