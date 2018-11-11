package util

import (
	"encoding/csv"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestGenerateRandomString(t *testing.T) {
	for l := -1; l < 32; l++ {
		s, e := GenerateRandomString(l)
		if l < 1 && e == nil {
			t.Errorf("Should fail with size (%d)", l)
		}
		if l >= 1 && (e != nil || len(s) != l) {
			t.Errorf("Error or invalid length (%d): generated (%s)", l, s)
		}
	}
}

func TestCompileTemplate(t *testing.T) {
	tmpl := "Hello {{.Name}}"
	data1 := struct {
		Name string
	}{
		Name: "World",
	}

	c, err := CompileTemplate(tmpl, data1)
	if err != nil {
		t.Error(err)
	}
	if c != "Hello World" {
		t.Errorf("Unxpected text: [%s]", c)
	}
	data2 := make(map[string]string)
	data2["Name"] = "Go"

	c, err = CompileTemplate(tmpl, data2)
	if err != nil {
		t.Error(err)
	}
	if c != "Hello Go" {
		t.Errorf("Unexpected text: [%s]", c)
	}
}

func TestNewBomAwareCsvReader(t *testing.T) {
	forig, err := os.Open("util_test_bomawarecsv_orig.txt")
	if err != nil {
		t.Error(err)
	}
	defer forig.Close()
	futf8, err := os.Open("util_test_bomawarecsv_utf8.txt")
	if err != nil {
		t.Error(err)
	}
	defer futf8.Close()
	futf8bom, err := os.Open("util_test_bomawarecsv_utf8bom.txt")
	if err != nil {
		t.Error(err)
	}
	defer futf8bom.Close()
	futf16le, err := os.Open("util_test_bomawarecsv_utf16le.txt")
	if err != nil {
		t.Error(err)
	}
	defer futf16le.Close()
	futf16lebom, err := os.Open("util_test_bomawarecsv_utf16lebom.txt")
	if err != nil {
		t.Error(err)
	}
	defer futf16lebom.Close()
	futf16be, err := os.Open("util_test_bomawarecsv_utf16be.txt")
	if err != nil {
		t.Error(err)
	}
	defer futf16be.Close()
	futf16bebom, err := os.Open("util_test_bomawarecsv_utf16bebom.txt")
	if err != nil {
		t.Error(err)
	}
	defer futf16bebom.Close()

	orig := csv.NewReader(forig)
	cutf8 := NewBomAwareCsvReader(futf8)
	cutf8bom := NewBomAwareCsvReader(futf8bom)
	cutf16le := NewBomAwareCsvReader(futf16le)
	cutf16lebom := NewBomAwareCsvReader(futf16lebom)
	cutf16be := NewBomAwareCsvReader(futf16be)
	cutf16bebom := NewBomAwareCsvReader(futf16bebom)

	for {
		lorig, err := orig.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Error(err)
		}
		lutf8, err := cutf8.Read()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(lorig, lutf8) {
			t.Error("utf8 not match")
		}

		lutf8bom, err := cutf8bom.Read()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(lorig, lutf8bom) {
			t.Error("utf8bom not match")
		}

		lutf16le, err := cutf16le.Read()
		if err != nil {
			t.Error(err)
		}
		if reflect.DeepEqual(lorig, lutf16le) {
			t.Error("utf16le should not match")
		}

		lutf16lebom, err := cutf16lebom.Read()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(lorig, lutf16lebom) {
			t.Error("utf16lebom not match")
		}

		lutf16be, err := cutf16be.Read()
		if err != nil {
			t.Error(err)
		}
		if reflect.DeepEqual(lorig, lutf16be) {
			t.Error("utf16be should not match")
		}

		lutf16bebom, err := cutf16bebom.Read()
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(lorig, lutf16bebom) {
			t.Error("utf16bebom not match")
		}
	}

	//
	//srcUtf8 := bytes.NewReader(textUtf8)
	//csvUtf8 := NewBomAwareCsvReader(srcUtf8)
	//
	//srcUtf16Le := bytes.NewReader(textUtf16LeSeq)
	//csvUtf16Le := NewBomAwareCsvReader(srcUtf16Le)
	//
	//srcUtf16Be := bytes.NewReader(textUtf16BeSeq)
	//csvUtf16Be := NewBomAwareCsvReader(srcUtf16Be)
	//
	//for i := 0; i < numRows; i++ {
	//	l8, err := csvUtf8.Read()
	//	if err != nil {
	//		t.Error(err)
	//	}
	//	l16le, err := csvUtf16Le.Read()
	//	if err != nil {
	//		t.Error(err)
	//	}
	//	l16be, err := csvUtf16Be.Read()
	//	if err != nil {
	//		t.Error(err)
	//	}
	//
	//	colsOfRow := make([]string, len(cols))
	//	for j, c := range cols {
	//		colsOfRow[j] = fmt.Sprintf("%s%d", c, i)
	//	}
	//	if !reflect.DeepEqual(l8, colsOfRow) {
	//		t.Error("Utf8 not matched", l8)
	//	}
	//	if !reflect.DeepEqual(l16le, colsOfRow) {
	//		t.Error("Utf16le not matched", l16le)
	//	}
	//	if !reflect.DeepEqual(l16be, colsOfRow) {
	//		t.Error("Utf16be not matched", l16be)
	//	}
	//}
}
