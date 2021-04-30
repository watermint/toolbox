package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func New(media dc_index.MediaType, spec rc_recipe.Spec) dc_section.Document {
	return &DocCommand{
		media: media,
		spec:  spec,
	}
}

type DocCommand struct {
	media dc_index.MediaType
	spec  rc_recipe.Spec
	Desc  app_msg.Message
}

func (z DocCommand) DocId() dc_index.DocId {
	return dc_index.DocManualCommand
}

func (z DocCommand) DocDesc() app_msg.Message {
	return z.Desc
}

func (z DocCommand) Sections() []dc_section.Section {
	sections := make([]dc_section.Section, 0)
	sections = append(sections, NewHeader(z.spec))
	if 0 < len(z.spec.ConnScopes()) {
		sections = append(sections, NewSecurity(z.spec))
		sections = append(sections, NewAuth(z.spec))
	}
	if z.media == dc_index.MediaWeb {
		sections = append(sections, NewInstall())
	}
	sections = append(sections, NewUsage(z.spec))
	if 0 < len(z.spec.Feeds()) {
		sections = append(sections, NewFeed(z.spec))
	}
	if 0 < len(z.spec.Reports()) {
		sections = append(sections, NewReport(z.spec))
	}
	if 0 < len(z.spec.GridDataInput()) {
		sections = append(sections, NewGridDataInput(z.spec))
	}
	if 0 < len(z.spec.GridDataOutput()) {
		sections = append(sections, NewGridDataOutput(z.spec))
	}
	if 0 < len(z.spec.TextInput()) {
		sections = append(sections, NewTextInput(z.spec))
	}
	sections = append(sections, NewProxy(z.spec))

	for i := 0; i < len(sections); i++ {
		sections[i] = app_msg.Apply(sections[i]).(dc_section.Section)
	}

	return sections
}
