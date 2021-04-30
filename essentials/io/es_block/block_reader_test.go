package es_block

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"testing"
)

func TestNewPlainFileSystem(t *testing.T) {
	f, err := qt_file.MakeTestFile("pfs", "0123456789ABCDEFGHIJ") // 20 bytes
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.Remove(f)
	}()

	{
		pfs := NewPlainReader(esl.Default(), 10)
		offsets, err := pfs.FileBlocks(f)
		if len(offsets) != 2 || err != nil {
			t.Error(offsets, err)
		}
		if offsets[0] != 0 || offsets[1] != 10 {
			t.Error(offsets[0], offsets[1])
		}

		// block 0
		{
			d, l, err := pfs.ReadBlock(f, 0)
			if err != nil || l || string(d) != "0123456789" {
				t.Error(err, l, string(d))
			}
		}

		// block 1
		{
			d, l, err := pfs.ReadBlock(f, 10)
			if err != nil || !l || string(d) != "ABCDEFGHIJ" {
				t.Error(err, l, string(d))
			}
		}
	}

	{
		pfs := NewPlainReader(esl.Default(), 15)
		offsets, err := pfs.FileBlocks(f)
		if len(offsets) != 2 || err != nil {
			t.Error(offsets, err)
		}
		if offsets[0] != 0 || offsets[1] != 15 {
			t.Error(offsets[0], offsets[1])
		}

		// block 0
		{
			d, l, err := pfs.ReadBlock(f, 0)
			if err != nil || l || string(d) != "0123456789ABCDE" {
				t.Error(err, l, string(d))
			}
		}

		// block 1
		{
			d, l, err := pfs.ReadBlock(f, 15)
			if err != nil || !l || string(d) != "FGHIJ" {
				t.Error(err, l, string(d))
			}
		}
	}
}
