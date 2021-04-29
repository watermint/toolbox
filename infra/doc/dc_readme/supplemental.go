package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/doc/dc_supplemental"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewSupplemental() dc_section.Section {
	return &Supplemental{}
}

type Supplemental struct {
	SupplementalTitle app_msg.Message
	DocumentName      app_msg.Message
	DocumentDesc      app_msg.Message
}

func (z Supplemental) Title() app_msg.Message {
	return z.SupplementalTitle
}

func (z Supplemental) Body(ui app_ui.UI) {
	ui.WithTable("supplemental documents", func(t app_ui.Table) {
		t.Header(z.DocumentName, z.DocumentDesc)
		for _, s := range dc_supplemental.Docs(dc_index.MediaRepository) {
			name := dc_index.DocName(dc_index.MediaRepository, s.DocId(), ui.Messages().Lang())
			link := dc_index.DocName(dc_index.MediaRepository, s.DocId(), ui.Messages().Lang())
			t.Row(app_msg.Raw("["+name+"]("+link+")"), s.DocDesc())
		}
	})
}
