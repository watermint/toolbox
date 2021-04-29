package dc_section

import (
	"bytes"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"strings"
	"text/template"
)

type Section interface {
	Title() app_msg.Message
	Body(ui app_ui.UI)
}

type Document interface {
	DocId() dc_index.DocId
	DocDesc() app_msg.Message
	Sections() []Section
}

const (
	WebHeader = `---
layout: {{.Layout}}
title: {{.Title}}
---

{{.Body}}
`
)

type LayoutType int

const (
	LayoutPage LayoutType = iota
	LayoutCommand
)

func Generate(media dc_index.MediaType, layout LayoutType, mc app_msg_container.Container, sections ...Section) string {
	body := app_ui.MakeMarkdown(mc, func(ui app_ui.UI) {
		for _, section := range sections {
			ui.Header(section.Title())
			section.Body(ui)
			ui.Break()
		}
	})

	title := ""
	if 0 < len(sections) {
		title = mc.Compile(sections[0].Title())
	}

	switch media {
	case dc_index.MediaRepository:
		return body
	case dc_index.MediaWeb:
		tmpl, err := template.New("web").Parse(WebHeader)
		if err != nil {
			panic(err)
		}
		buf := bytes.Buffer{}
		layoutName := "page"
		if layout == LayoutCommand {
			layoutName = "command"
		}
		err = tmpl.Execute(&buf, map[string]string{
			"Title":  title,
			"Layout": layoutName,
			"Body": strings.ReplaceAll(strings.ReplaceAll(body,
				"{{", "{% raw %}{{{% endraw %}"),
				"}}", "{% raw %}}}{% endraw %}"),
		})
		if err != nil {
			panic(err)
		}
		return buf.String()

	default:
		panic("Undefined media type")
	}
}
