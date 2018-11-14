package app_util

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"text/template"
)

// size: length of the string
func GenerateRandomString(size int) (string, error) {
	if size < 1 {
		return "", errors.New(fmt.Sprintf("Size must greater than 1, given size was %d", size))
	}
	seq := make([]byte, size)
	_, err := rand.Read(seq)
	if err != nil {
		return "", err
	}
	encoded := base64.URLEncoding.EncodeToString(seq)
	return encoded[:size], nil
}

func CompileTemplate(tmpl string, data interface{}) (string, error) {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer

	if err := t.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

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
