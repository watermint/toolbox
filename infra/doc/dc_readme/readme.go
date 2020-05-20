package dc_readme

import (
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func New(forPublish bool, commandPath string) []dc_section.Section {
	sections := make([]dc_section.Section, 0)
	sections = append(sections, NewHeader(forPublish))
	sections = append(sections, NewLicense())
	if forPublish {
		sections = append(sections, NewRelease())
	}
	sections = append(sections, NewUsage())
	if forPublish {
		sections = append(sections, NewCommand(forPublish, commandPath))
	}

	for i := 0; i < len(sections); i++ {
		sections[i] = app_msg.Apply(sections[i]).(dc_section.Section)
	}
	return sections
}
