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
lang: {{.Lang}}
---

{{.Body}}
`
)

type LayoutType int

const (
	LayoutPage LayoutType = iota
	LayoutHome
	LayoutCommand
	LayoutContributor
)

func Generate(media dc_index.MediaType, layout LayoutType, mc app_msg_container.Container, doc Document) string {
	sections := doc.Sections()
	body := app_ui.MakeMarkdown(mc, func(ui app_ui.UI) {
		for _, s := range sections {
			sec := app_msg.Apply(s).(Section)
			ui.Header(sec.Title())
			sec.Body(ui)
			ui.Break()
		}
	})

	compiledDoc := app_msg.Apply(doc).(Document)
	title := mc.Compile(compiledDoc.DocDesc())

	// #522 : markdownify look like not work on GitHub Pages
	if layout == LayoutHome {
		return body
	}

	switch media {
	case dc_index.MediaRepository, dc_index.MediaKnowledge:
		return body
	case dc_index.MediaWeb:
		tmpl, err := template.New("web").Parse(WebHeader)
		if err != nil {
			panic(err)
		}
		buf := bytes.Buffer{}
		var layoutName string
		switch layout {
		case LayoutHome:
			layoutName = "home"
		case LayoutCommand:
			layoutName = "command"
		case LayoutContributor:
			layoutName = "contributor"
		default:
			layoutName = "page"
		}
		err = tmpl.Execute(&buf, map[string]string{
			"Title":  title,
			"Layout": layoutName,
			"Lang":   mc.Lang().CodeString(),
			"Body":   strings.ReplaceAll(body, "{{.", "{% raw %}{{.{% endraw %}"),
		})
		if err != nil {
			panic(err)
		}
		return buf.String()

	default:
		panic("Undefined media type")
	}
}
