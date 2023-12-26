package dc_command

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_compatibility"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"strings"
)

func NewCompatibilityNewPath(media dc_index.MediaType, spec rc_recipe.Spec, pathPair rc_compatibility.PathPair, newPathSpec rc_compatibility.PathChangeDefinition) dc_section.Document {
	return &DocCompatibilityNewPath{
		media:       media,
		spec:        spec,
		pathPair:    pathPair,
		newPathSpec: newPathSpec,
	}
}

type DocCompatibilityNewPath struct {
	media       dc_index.MediaType
	spec        rc_recipe.Spec
	pathPair    rc_compatibility.PathPair
	newPathSpec rc_compatibility.PathChangeDefinition
	Desc        app_msg.Message
}

func (z DocCompatibilityNewPath) DocId() dc_index.DocId {
	return dc_index.DocManualCommand
}

func (z DocCompatibilityNewPath) DocDesc() app_msg.Message {
	return z.Desc.With("CliPath", z.spec.CliPath())
}

func (z DocCompatibilityNewPath) Sections() []dc_section.Section {
	return []dc_section.Section{
		&SectionCompatibilityNewPath{
			spec:        z.spec,
			pathPair:    z.pathPair,
			newPathSpec: z.newPathSpec,
		},
	}
}

type SectionCompatibilityNewPath struct {
	spec                 rc_recipe.Spec
	pathPair             rc_compatibility.PathPair
	newPathSpec          rc_compatibility.PathChangeDefinition
	TitleMoved           app_msg.Message
	DescMoved            app_msg.Message
	DescTransitionPeriod app_msg.Message
	DescAnnouncement     app_msg.Message
}

func (z SectionCompatibilityNewPath) Title() app_msg.Message {
	return z.TitleMoved.With("CliPath", strings.Join(append(z.pathPair.Path, z.pathPair.Name), " "))
}

func (z SectionCompatibilityNewPath) Body(ui app_ui.UI) {
	ui.Info(z.DescMoved.With("CurrentPath", z.spec.CliPath()).
		With("FormerPath", strings.Join(append(z.pathPair.Path, z.pathPair.Name), " ")).
		With("Title", ui.TextOrEmpty(z.spec.Title())))

	if z.newPathSpec.PruneAfterBuildDate != "" {
		ui.Break()
		ui.Info(z.DescTransitionPeriod.With("PruneAfterBuildDate", z.newPathSpec.PruneAfterBuildDate))
	}

	if z.newPathSpec.Announcement != "" {
		ui.Break()
		ui.Info(z.DescAnnouncement.With("URL", z.newPathSpec.Announcement))
	}
}
