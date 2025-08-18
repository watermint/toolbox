package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewKtloAnnouncement() dc_section.Section {
	return &KtloAnnouncement{}
}

type KtloAnnouncement struct {
	AnnouncementTitle app_msg.Message
	AnnouncementBody  app_msg.Message
}

func (z KtloAnnouncement) Title() app_msg.Message {
	return z.AnnouncementTitle
}

func (z KtloAnnouncement) Body(ui app_ui.UI) {
	ui.Info(z.AnnouncementBody)
}
