package ut_encoding

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
)

func NewBomAwareCsvReader(r io.Reader) *csv.Reader {
	var (
		bomUtf8    = []byte{0xef, 0xbb, 0xbf}
		bomUtf16BE = []byte{0xfe, 0xff}
		bomUtf16LE = []byte{0xff, 0xfe}
	)
	br := bufio.NewReader(r)
	mark, err := br.Peek(3)
	if err != nil {
		panic(err)
	}
	dec := unicode.UTF8.NewDecoder()

	if bytes.HasPrefix(mark, bomUtf8) {
		br.Discard(len(bomUtf8))
	} else if bytes.HasPrefix(mark, bomUtf16BE) {
		dec = unicode.UTF16(unicode.BigEndian, unicode.UseBOM).NewDecoder()
	} else if bytes.HasPrefix(mark, bomUtf16LE) {
		dec = unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewDecoder()
	}

	return csv.NewReader(transform.NewReader(br, dec))
}

func NewBomAawareCsvWriter(w io.Writer) *csv.Writer {
	bomUtf8 := []byte{0xef, 0xbb, 0xbf}
	w.Write(bomUtf8)
	return csv.NewWriter(w)
}
