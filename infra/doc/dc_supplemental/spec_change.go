package dc_supplemental

import (
	"strings"

	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_compatibility"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgSpecChange struct {
	DocDesc app_msg.Message
}

var (
	MSpecChange = app_msg.Apply(&MsgSpecChange{}).(*MsgSpecChange)
)

type DocSpecChange struct {
	Desc app_msg.Message
}

func NewDocSpecChange() *DocSpecChange {
	return &DocSpecChange{
		Desc: MSpecChange.DocDesc,
	}
}

func (z DocSpecChange) DocId() dc_index.DocId {
	return dc_index.DocSupplementalSpecChange
}

func (z DocSpecChange) DocDesc() app_msg.Message {
	return z.Desc
}

func (z DocSpecChange) Sections() []dc_section.Section {
	return []dc_section.Section{
		&SecSpecChange{},
	}
}

type SecSpecChange struct {
	TitleSection app_msg.Message

	HeadPathChange                           app_msg.Message
	BodyPathChange                           app_msg.Message
	TableHeaderPathChangeCliPathFrom         app_msg.Message
	TableHeaderPathChangeCliPathTo           app_msg.Message
	TableHeaderPathChangeDesc                app_msg.Message
	TableHeaderPathChangePruneAfterBuildDate app_msg.Message

	HeadPrune                           app_msg.Message
	BodyPrune                           app_msg.Message
	TableHeaderPruneCliPath             app_msg.Message
	TableHeaderPruneDesc                app_msg.Message
	TableHeaderPrunePruneAfterBuildDate app_msg.Message
}

func (z SecSpecChange) Title() app_msg.Message {
	return z.TitleSection
}

func (z SecSpecChange) Body(ui app_ui.UI) {
	scheduledPrune := make([]rc_recipe.Spec, 0)
	scheduledPruneDefinition := make(map[string]rc_compatibility.PruneDefinition, 0)
	for _, p := range rc_compatibility.Definitions.ListPlannedPrune() {
		spec := app_catalogue.Current().RecipeSpec(strings.Join(append(p.Current.Path, p.Current.Name), " "))
		if !spec.IsSecret() {
			scheduledPrune = append(scheduledPrune, spec)
			scheduledPruneDefinition[spec.CliPath()] = p
		}
	}
	scheduledPathChange := make([]rc_recipe.Spec, 0)
	scheduledPathChangeDefinition := make(map[string]rc_compatibility.PathChangeDefinition, 0)
	for _, p := range rc_compatibility.Definitions.ListAlivePathChange() {
		spec := app_catalogue.Current().RecipeSpec(strings.Join(append(p.Current.Path, p.Current.Name), " "))
		if !spec.IsSecret() {
			scheduledPathChange = append(scheduledPathChange, spec)
			scheduledPathChangeDefinition[spec.CliPath()] = p
		}
	}

	ui.Header(z.HeadPathChange)
	ui.Break()
	ui.Info(z.BodyPathChange)
	ui.WithTable("PathChange", func(t app_ui.Table) {
		t.Header(z.TableHeaderPathChangeCliPathFrom, z.TableHeaderPathChangeCliPathTo, z.TableHeaderPathChangeDesc, z.TableHeaderPathChangePruneAfterBuildDate)
		for _, spec := range scheduledPathChange {
			p := scheduledPathChangeDefinition[spec.CliPath()]
			for _, fp := range p.FormerPaths {
				if p.Announcement == "" {
					t.Row(
						app_msg.Raw(strings.Join(append(fp.Path, fp.Name), " ")),
						app_msg.Raw(spec.CliPath()),
						spec.Title(),
						app_msg.Raw(p.PruneAfterBuildDate),
					)
				} else {
					t.Row(
						app_msg.Raw("["+strings.Join(append(fp.Path, fp.Name), " ")+"]("+p.Announcement+")"),
						app_msg.Raw(spec.CliPath()),
						spec.Title(),
						app_msg.Raw(p.PruneAfterBuildDate),
					)
				}
			}
		}
	})

	ui.Break()
	ui.Header(z.HeadPrune)
	ui.Break()
	ui.Info(z.BodyPrune)
	ui.WithTable("Prune", func(t app_ui.Table) {
		t.Header(z.TableHeaderPruneCliPath, z.TableHeaderPruneDesc, z.TableHeaderPrunePruneAfterBuildDate)
		for _, spec := range scheduledPrune {
			p := scheduledPruneDefinition[spec.CliPath()]
			if p.Announcement == "" {
				t.Row(
					app_msg.Raw(spec.CliPath()),
					spec.Title(),
					app_msg.Raw(p.PruneAfterBuildDate),
				)
			} else {
				t.Row(
					app_msg.Raw("["+spec.CliPath()+"]("+p.Announcement+")"),
					spec.Title(),
					app_msg.Raw(p.PruneAfterBuildDate),
				)
			}
		}
	})
}
