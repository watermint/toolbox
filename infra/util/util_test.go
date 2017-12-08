package util

import "testing"

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
