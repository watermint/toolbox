package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func New(spec rc_recipe.Spec) []dc_section.Section {
	sections := make([]dc_section.Section, 0)
	sections = append(sections, NewHeader(spec))
	if 0 < len(spec.ConnScopes()) {
		sections = append(sections, NewSecurity(spec))
		sections = append(sections, NewAuth(spec))
	}
	sections = append(sections, NewUsage(spec))
	if 0 < len(spec.Feeds()) {
		sections = append(sections, NewFeed(spec))
	}
	if 0 < len(spec.Reports()) {
		sections = append(sections, NewReport(spec))
	}
	if 0 < len(spec.GridDataInput()) {
		sections = append(sections, NewGridDataInput(spec))
	}
	if 0 < len(spec.GridDataOutput()) {
		sections = append(sections, NewGridDataOutput(spec))
	}
	sections = append(sections, NewProxy(spec))

	for i := 0; i < len(sections); i++ {
		sections[i] = app_msg.Apply(sections[i]).(dc_section.Section)
	}

	return sections
}
