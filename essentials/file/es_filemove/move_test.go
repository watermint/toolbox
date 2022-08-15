package es_filemove

import (
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCopyThenDelete(t *testing.T) {
	createFile := func(path, content string, mode os.FileMode, modTime time.Time) error {
		err := os.WriteFile(path, []byte(content), mode)
		if err != nil {
			t.Error(err)
			return err
		}
		err = os.Chtimes(path, modTime, modTime)
		if err != nil {
			t.Error(err)
			return err
		}
		return nil
	}

	qt_file.TestWithTestFolder(t, "copy-and-delete", false, func(path string) {
		// Test file prep
		target1Path := filepath.Join(path, "target1.txt")
		target1Mode := os.FileMode(0641)
		target1Time := time.Now().Add(-1 * time.Hour)
		target1Text := "target1-content"

		if err := createFile(target1Path, target1Text, target1Mode, target1Time); err != nil {
			return
		}

		// Success case
		success1Path := filepath.Join(path, "success1.txt")
		if err := CopyThenDelete(target1Path, success1Path); err != nil {
			t.Error(err)
			return
		}
		success1Info, err := os.Lstat(success1Path)
		if err != nil {
			t.Error(err)
		}
		if success1Info.Mode() != target1Mode {
			t.Error("mode mismatch", target1Mode, success1Info.Mode())
		}
		if !success1Info.ModTime().Equal(target1Time) {
			t.Error("time mismatch", success1Info.ModTime(), target1Time)
		}

		// Error case not found
		if err := CopyThenDelete(target1Path, success1Path); err == nil {
			t.Error("not found was not reported")
			return
		}
	})

	qt_file.TestWithTestFolder(t, "copy-and-delete", false, func(path string) {
		// Test file prep
		target1Path := filepath.Join(path, "target1.txt")
		target1Mode := os.FileMode(0660)
		target1Time := time.Now().Add(-1 * time.Hour)
		target1Text := "target1-content"

		if err := createFile(target1Path, target1Text, target1Mode, target1Time); err != nil {
			return
		}

		// Error no folder in dst path
		failure1Path := filepath.Join(path, "no_existent", "failure1.txt")
		if err := CopyThenDelete(target1Path, failure1Path); err == nil {
			t.Error("not failure was not reported")
			return
		}
	})
}

func TestMove(t *testing.T) {
	createFile := func(path, content string, mode os.FileMode, modTime time.Time) error {
		err := os.WriteFile(path, []byte(content), mode)
		if err != nil {
			t.Error(err)
			return err
		}
		err = os.Chtimes(path, modTime, modTime)
		if err != nil {
			t.Error(err)
			return err
		}
		return nil
	}

	qt_file.TestWithTestFolder(t, "move", false, func(path string) {
		// Test file prep
		target1Path := filepath.Join(path, "target1.txt")
		target1Mode := os.FileMode(0641)
		target1Time := time.Now().Add(-1 * time.Hour)
		target1Text := "target1-content"

		if err := createFile(target1Path, target1Text, target1Mode, target1Time); err != nil {
			return
		}

		// Success case
		success1Path := filepath.Join(path, "success1.txt")
		if err := Move(target1Path, success1Path); err != nil {
			t.Error(err)
			return
		}
		success1Info, err := os.Lstat(success1Path)
		if err != nil {
			t.Error(err)
		}
		if success1Info.Mode() != target1Mode {
			t.Error("mode mismatch", target1Mode, success1Info.Mode())
		}
		if !success1Info.ModTime().Equal(target1Time) {
			t.Error("time mismatch", success1Info.ModTime(), target1Time)
		}

		// Error case not found
		if err := Move(target1Path, success1Path); err == nil {
			t.Error("not found was not reported")
			return
		}
	})

	qt_file.TestWithTestFolder(t, "move", false, func(path string) {
		// Test file prep
		target1Path := filepath.Join(path, "target1.txt")
		target1Mode := os.FileMode(0660)
		target1Time := time.Now().Add(-1 * time.Hour)
		target1Text := "target1-content"

		if err := createFile(target1Path, target1Text, target1Mode, target1Time); err != nil {
			return
		}

		// Error no folder in dst path
		failure1Path := filepath.Join(path, "no_existent", "failure1.txt")
		if err := Move(target1Path, failure1Path); err == nil {
			t.Error("not failure was not reported")
			return
		}
	})
}
