package util

import "testing"

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
