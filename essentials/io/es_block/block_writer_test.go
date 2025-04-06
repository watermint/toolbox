package es_block

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

func TestBlockWriter(t *testing.T) {
	wd, err := qt_file.MakeTestFolder("block_writer", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(wd)
	}()
	l := esl.Default()

	// small chunks
	{
		path := filepath.Join(wd, "4_10.bin")
		content := "https://github.com/watermint/toolbox"
		bwf := NewWriterFactory(l, 4, 10)
		bw := bwf.Open(
			path,
			int64(len(content)),
			func(w BlockWriter, offset, blockSize int64) {
				x := int(offset)
				y := x + int(blockSize)
				w.WriteBlock([]byte(content[x:y]), int64(x))
			},
			func(w BlockWriter, size int64) {
				if size != int64(len(content)) {
					t.Error(size, len(content))
				}
			},
			func(w BlockWriter, offset int64, err error) {
				t.Error(offset, err)
			},
		)
		bw.Wait()

		fileContent, err := os.ReadFile(path)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(fileContent, []byte(content)) {
			t.Error("mismatch", fileContent, content)
		}
	}

	// large block size (larger than content size)
	{
		path := filepath.Join(wd, "4_10.bin")
		content := "https://github.com/watermint/toolbox"
		bwf := NewWriterFactory(l, 4, len(content)+1)
		bw := bwf.Open(
			path,
			int64(len(content)),
			func(w BlockWriter, offset, blockSize int64) {
				x := int(offset)
				y := x + int(blockSize)
				w.WriteBlock([]byte(content[x:y]), int64(x))
			},
			func(w BlockWriter, size int64) {
				if size != int64(len(content)) {
					t.Error(size, len(content))
				}
			},
			func(w BlockWriter, offset int64, err error) {
				t.Error(offset, err)
			},
		)
		bw.Wait()

		fileContent, err := os.ReadFile(path)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(fileContent, []byte(content)) {
			t.Error("mismatch", fileContent, content)
		}
	}
}
