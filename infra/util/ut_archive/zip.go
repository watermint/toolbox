package ut_archive

import (
	"archive/zip"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Create(arcPath, targetPath, arcComment string) error {
	l := app_root.Log()
	arcFile, err := os.Create(arcPath)
	if err != nil {
		l.Error("Unable to create archive", zap.Error(err))
		return err
	}

	arc := zip.NewWriter(arcFile)
	arc.SetComment(arcComment)

	var archive func(b, d string) error
	archive = func(b, d string) error {
		entries, err := ioutil.ReadDir(filepath.Join(b, d))
		if err != nil {
			l.Error("Unable to read dir", zap.Error(err))
			return err
		}
		for _, e := range entries {
			if e.IsDir() {
				if err = archive(b, filepath.Join(d, e.Name())); err != nil {
					return err
				}
			} else {
				ep := filepath.Join(d, e.Name())
				l.Debug("Add file into the archive", zap.String("EntryPath", ep))
				w, err := arc.CreateHeader(&zip.FileHeader{
					Name:     filepath.ToSlash(ep),
					Comment:  e.Name(),
					Modified: e.ModTime(),
				})
				if err != nil {
					l.Debug("Unable to add file to the archive", zap.Error(err))
					continue
				}

				rp := filepath.Join(b, d, e.Name())
				r, err := os.Open(rp)
				if err != nil {
					l.Debug("Unable to read file", zap.Error(err))
					continue
				}

				_, err = io.Copy(w, r)
				if err != nil {
					l.Debug("Unable to insert file into the archive", zap.Error(err))
					continue
				}

				r.Close()
				l.Debug("Try remove the log file", zap.String("path", rp))
				err = os.Remove(rp)
				l.Debug("Removed", zap.String("path", rp), zap.Error(err))
			}
		}
		return nil
	}

	err = archive(targetPath, "")
	arc.Flush()
	arc.Close()
	if err != nil {
		return err
	}

	return nil
}
