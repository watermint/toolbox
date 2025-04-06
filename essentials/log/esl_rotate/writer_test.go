package esl_rotate

import (
	"os"
	"testing"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

func TestWriterImpl_Open(t *testing.T) {
	qt_file.TestWithTestFolder(t, "writer", false, func(path string) {
		Startup()
		w := NewWriter(path, "writer")
		if err := w.Open(ChunkSize(10), NumBackup(3)); err != nil {
			t.Error(err)
		}
		if n, err := w.Write([]byte("Hello")); err != nil {
			t.Error(n, err)
		}
		if err := w.Close(); err != nil {
			t.Error(err)
		}
		Shutdown()
	})
}

func TestRotate(t *testing.T) {
	qt_file.TestWithTestFolder(t, "rotate", false, func(path string) {
		l := esl.ConsoleOnly()
		numPurge := 0
		{
			hook := func(path string) {
				l.Info("Path", esl.String("path", path))
				numPurge++
			}

			Shutdown()
			w := NewWriter(path, "w1")
			if err := w.Open(
				ChunkSize(16),
				NumBackup(3),
				HookBeforeDelete(hook)); err != nil {
				t.Error(err)
			}

			if n, err := w.Write([]byte("01234567890123456789")); err != nil {
				t.Error(n, err)
			}

			// Should rotate
			if n, err := w.Write([]byte("01234567890123456789")); err != nil {
				t.Error(n, err)
			}

			// Should rotate
			if n, err := w.Write([]byte("01234567890123456789")); err != nil {
				t.Error(n, err)
			}

			// Should rotate & delete
			if n, err := w.Write([]byte("01234567890123456789")); err != nil {
				t.Error(n, err)
			}

			// Should rotate & delete
			if n, err := w.Write([]byte("01234567890123456789")); err != nil {
				t.Error(n, err)
			}

			// Should rotate & delete
			if n, err := w.Write([]byte("01234567890123456789")); err != nil {
				t.Error(n, err)
			}

			if err := w.Close(); err != nil {
				t.Error(err)
			}
		}
		Shutdown()
		if numPurge != 4 {
			t.Error(numPurge)
		}

		entries, err := os.ReadDir(path)
		if err != nil || len(entries) != 3 {
			t.Error(err, entries)
		}
	})
}
