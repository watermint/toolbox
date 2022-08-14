package es_zip

import (
	"archive/zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CompressPath creates zip archive `arcPath` for `targetPath`
func CompressPath(arcPath, targetPath, arcComment string) error {
	l := esl.Default()
	arcFile, err := os.Create(arcPath)
	if err != nil {
		l.Error("Unable to create archive", esl.Error(err))
		return err
	}

	arc := zip.NewWriter(arcFile)
	arc.SetComment(arcComment)

	var archive func(b, d string) error
	archive = func(b, d string) error {
		entries, err := ioutil.ReadDir(filepath.Join(b, d))
		if err != nil {
			l.Error("Unable to read dir", esl.Error(err))
			return err
		}
		for _, e := range entries {
			if e.IsDir() {
				if err = archive(b, filepath.Join(d, e.Name())); err != nil {
					return err
				}
			} else {
				ep := filepath.Join(d, e.Name())
				l.Debug("Add file into the archive", esl.String("EntryPath", ep))
				w, err := arc.CreateHeader(&zip.FileHeader{
					Name:     filepath.ToSlash(ep),
					Comment:  e.Name(),
					Method:   zip.Deflate,
					Modified: e.ModTime(),
				})
				if err != nil {
					l.Debug("Unable to add file to the archive", esl.Error(err))
					continue
				}

				rp := filepath.Join(b, d, e.Name())
				r, err := os.Open(rp)
				if err != nil {
					l.Debug("Unable to read file", esl.Error(err))
					continue
				}

				_, err = io.Copy(w, r)
				if err != nil {
					l.Debug("Unable to insert file into the archive", esl.Error(err))
					continue
				}

				_ = r.Close()
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
