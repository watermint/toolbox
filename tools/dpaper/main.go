package main

import (
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"os"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

List document ids
{{.Command}} list
{{.Command}} list -filter created

Download
{{.Command}} download -filter created -format markdown -out-path ./
{{.Command}} download -docid zO1E7coc54sE8IuMdUoxz -format markdown -out zO1E7coc54sE8IuMdUoxz.md

Archive
{{.Command}} archive -docid zO1E7coc54sE8IuMdUoxz

Unshare
{{.Command}} unshare -docid zO1E7coc54sE8IuMdUoxz
`

	data := struct {
		AppName    string
		AppVersion string
		AppHash    string
		Command    string
	}{
		AppName:    knowledge.AppName,
		AppVersion: knowledge.AppVersion,
		AppHash:    knowledge.AppHash,
		Command:    os.Args[0],
	}
	infra.ShowUsage(tmpl, data)
}

func main() {

}
