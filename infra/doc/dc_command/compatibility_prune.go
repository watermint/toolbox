package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_compatibility"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewCompatibilityPrune(media dc_index.MediaType, spec rc_recipe.Spec, prune rc_compatibility.PruneDefinition) dc_section.Document {
	return &DocCompatibilityPrune{
		media: media,
		spec:  spec,
		prune: prune,
	}
}

type DocCompatibilityPrune struct {
	media dc_index.MediaType
	spec  rc_recipe.Spec
	prune rc_compatibility.PruneDefinition
	Desc  app_msg.Message
}

func (z DocCompatibilityPrune) DocId() dc_index.DocId {
	return dc_index.DocManualCommand
}

func (z DocCompatibilityPrune) DocDesc() app_msg.Message {
	return z.Desc.With("CliPath", z.spec.CliPath())
}

func (z DocCompatibilityPrune) Sections() []dc_section.Section {
	return []dc_section.Section{
		&SectionCompatibilityPrune{
			spec:  z.spec,
			prune: z.prune,
		},
	}
}

type SectionCompatibilityPrune struct {
	spec                 rc_recipe.Spec
	prune                rc_compatibility.PruneDefinition
	TitlePrune           app_msg.Message
	DescPrune            app_msg.Message
	DescTransitionPeriod app_msg.Message
	DescAnnouncement     app_msg.Message
}

func (z SectionCompatibilityPrune) Title() app_msg.Message {
	return z.TitlePrune.With("CliPath", z.spec.CliPath())
}

func (z SectionCompatibilityPrune) Body(ui app_ui.UI) {
	ui.Info(z.DescPrune.With("Path", z.spec.CliPath()).With("Title", ui.TextOrEmpty(z.spec.Title())))

	if z.prune.PruneAfterBuildDate != "" {
		ui.Info(z.DescTransitionPeriod.With("PruneAfterBuildDate", z.prune.PruneAfterBuildDate))
	}

	if z.prune.Announcement != "" {
		ui.Break()
		ui.Info(z.DescAnnouncement.With("URL", z.prune.Announcement))
	}
}
