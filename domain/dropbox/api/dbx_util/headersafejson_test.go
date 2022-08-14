package dbx_util

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestHeaderSafeJson(t *testing.T) {
	// test 1
	{
		type Data struct {
			Title   string `json:"path"`
			Num     int    `json:"num"`
			Enabled bool   `json:"enabled"`
			Message string `json:"message"`
			Data    []byte `json:"data"`
		}

		p := &Data{
			Title:   "\t<h1>\\こんにちは、\"世界\"。\n</h1>",
			Num:     1,
			Enabled: false,
			Message: "☺️",
			Data:    []byte{0x7f, 0x80, 0xee, 0xf0},
		}
		j, err := HeaderSafeJson(p)
		if err != nil {
			t.Error(err)
		}

		for i, j0 := range j {
			if 0x7f < j0 {
				t.Error("contains 8bit char", i)
			}
		}

		t.Log(j)
		q := &Data{}
		if err := json.Unmarshal([]byte(j), q); err != nil {
			t.Error(err)
		}

		if q.Title != p.Title {
			t.Error("title broken")
		}
		if q.Num != p.Num {
			t.Error("num broken")
		}
		if q.Enabled != p.Enabled {
			t.Error("enabled broken")
		}
		if q.Message != p.Message {
			t.Error("message broken")
		}
		if bytes.Compare(q.Data, p.Data) != 0 {
			t.Error("data broken")
		}
	}

	// test 2 : should fail on Unicode plane 1 >= (U+10000) char found
	{
		type Data struct {
			Title   string `json:"path"`
			Num     int    `json:"num"`
			Enabled bool   `json:"enabled"`
			Message string `json:"message"`
			Data    string `json:"data"`
		}

		p := &Data{
			Title:   "\t<h1>\\こんにちは、\"世界\"。\n</h1>",
			Num:     1,
			Enabled: false,
			Message: "🍜️", // U+
			Data:    "\x7f",
		}
		_, err := HeaderSafeJson(p)
		if err == nil {
			t.Error("should fail")
		}
	}
}
