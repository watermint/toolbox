package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewDropboxBusiness() dc_section.Section {
	return &DropboxBusiness{}
}

type DropboxBusiness struct {
	SectionTitle app_msg.Message
	SectionDesc  app_msg.Message
}

func (z DropboxBusiness) Title() app_msg.Message {
	return z.SectionTitle
}

func (z DropboxBusiness) Body(ui app_ui.UI) {
	path := dc_index.DocName(dc_index.DocSupplementalDropboxBusiness, ui.Messages().Lang()) + ".md"
	ui.Info(z.SectionDesc.With("Path", path))
}
