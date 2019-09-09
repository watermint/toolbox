package app_util

import (
	"bytes"
	"text/template"
)

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
