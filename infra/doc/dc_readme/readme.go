package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
)

func New(media dc_index.MediaType, mc app_msg_container.Container, forPublish bool) dc_section.Document {
	return &Readme{
		media:      media,
		mc:         mc,
		forPublish: forPublish,
	}
}

type Readme struct {
	media      dc_index.MediaType
	mc         app_msg_container.Container
	forPublish bool
	Desc       app_msg.Message
}

func (z Readme) DocId() dc_index.DocId {
	return dc_index.DocRootReadme
}

func (z Readme) DocDesc() app_msg.Message {
	return z.Desc
}

func (z Readme) Sections() []dc_section.Section {
	sections := make([]dc_section.Section, 0)
	sections = append(sections, NewHeader(z.forPublish))
	sections = append(sections, NewLicense())
	if z.forPublish {
		sections = append(sections, NewRelease())
		sections = append(sections, NewAnnouncements())
		sections = append(sections, NewLifecycle())
	}
	sections = append(sections, NewSecuritySection())
	sections = append(sections, NewUsage())
	if z.forPublish {
		sections = append(sections, NewCommand(z.forPublish, z.media, z.mc))
	}

	for i := 0; i < len(sections); i++ {
		sections[i] = app_msg.Apply(sections[i]).(dc_section.Section)
	}

	return sections
}
