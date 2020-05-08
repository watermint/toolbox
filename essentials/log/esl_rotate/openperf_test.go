package esl_rotate

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func BenchmarkOpen(b *testing.B) {
	f, err := ioutil.TempFile("", "open")
	if err != nil {
		b.Error(err)
		return
	}
	written := 0
	for i := 0; i < b.N; i++ {
		n, err := f.WriteString(fmt.Sprintf("%d\n", i))
		if err != nil {
			b.Error(err)
			return
		}
		written += n
	}
	if err := f.Close(); err != nil {
		b.Error(err)
	}
}

func BenchmarkAppend(b *testing.B) {
	f, err := ioutil.TempFile("", "open")
	if err != nil {
		b.Error(err)
		return
	}
	fs, err := f.Stat()
	if err != nil {
		b.Error(err)
		return
	}
	if err := f.Close(); err != nil {
		b.Error(err)
		return
	}
	path := f.Name()

	written := 0
	for i := 0; i < b.N; i++ {
		g, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, fs.Mode())
		if err != nil {
			b.Error(err)
			return
		}

		n, err := g.WriteString(fmt.Sprintf("%d\n", i))
		if err != nil {
			b.Error(err)
			return
		}
		written += n

		if err := g.Close(); err != nil {
			b.Error(err)
			return
		}
	}
}
