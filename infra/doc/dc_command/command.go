package dc_command

import (
	"sort"

	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

// Define section types for documentation organization
type SectionType int

const (
	SectionTypeHeader SectionType = iota + 1 // Use 1-based index for priority
	SectionTypeCommandSecurity
	SectionTypeCommandAuth
	SectionTypeInstall
	SectionTypeUsage
	SectionTypeFeed
	SectionTypeReport
	SectionTypeGridDataInput
	SectionTypeGridDataOutput
	SectionTypeTextInput
)

// Priority returns the priority of the section type
func (t SectionType) Priority() int {
	return int(t)
}

// TypedSection represents a section with its type information
type TypedSection struct {
	sectionType SectionType
	section     dc_section.Section
}

// CommonSections represents common sections that are shared across commands
type CommonSections struct {
	media dc_index.MediaType
}

// NewCommon creates a new instance of CommonSections
func NewCommon(media dc_index.MediaType) *CommonSections {
	return &CommonSections{
		media: media,
	}
}

// GetSections returns common sections based on media type
func (z *CommonSections) GetSections() []TypedSection {
	sections := make([]TypedSection, 0)

	// Installation instructions are common for web media
	if z.media == dc_index.MediaWeb {
		sections = append(sections, TypedSection{
			sectionType: SectionTypeInstall,
			section:     app_msg.Apply(NewInstall()).(dc_section.Section),
		})
	}

	return sections
}

// CommandSections generates command-specific sections
func generateCommandSections(media dc_index.MediaType, spec rc_recipe.Spec) []TypedSection {
	sections := make([]TypedSection, 0)
	sections = append(sections, TypedSection{
		sectionType: SectionTypeHeader,
		section:     app_msg.Apply(NewHeader(spec)).(dc_section.Section),
	})

	// Command specific security and auth sections
	if media != dc_index.MediaKnowledge {
		if 0 < len(spec.ConnScopes()) {
			sections = append(sections, TypedSection{
				sectionType: SectionTypeCommandSecurity,
				section:     app_msg.Apply(NewSecurity(spec)).(dc_section.Section),
			})
			sections = append(sections, TypedSection{
				sectionType: SectionTypeCommandAuth,
				section:     app_msg.Apply(NewAuth(spec)).(dc_section.Section),
			})
		}
	}

	// Command specific usage
	sections = append(sections, TypedSection{
		sectionType: SectionTypeUsage,
		section:     app_msg.Apply(NewUsage(media, spec)).(dc_section.Section),
	})

	// Data related sections
	if 0 < len(spec.Feeds()) {
		sections = append(sections, TypedSection{
			sectionType: SectionTypeFeed,
			section:     app_msg.Apply(NewFeed(spec)).(dc_section.Section),
		})
	}
	if 0 < len(spec.Reports()) {
		sections = append(sections, TypedSection{
			sectionType: SectionTypeReport,
			section:     app_msg.Apply(NewReport(media, spec)).(dc_section.Section),
		})
	}
	if 0 < len(spec.GridDataInput()) {
		sections = append(sections, TypedSection{
			sectionType: SectionTypeGridDataInput,
			section:     app_msg.Apply(NewGridDataInput(spec)).(dc_section.Section),
		})
	}
	if 0 < len(spec.GridDataOutput()) {
		sections = append(sections, TypedSection{
			sectionType: SectionTypeGridDataOutput,
			section:     app_msg.Apply(NewGridDataOutput(spec)).(dc_section.Section),
		})
	}
	if 0 < len(spec.TextInput()) {
		sections = append(sections, TypedSection{
			sectionType: SectionTypeTextInput,
			section:     app_msg.Apply(NewTextInput(spec)).(dc_section.Section),
		})
	}

	return sections
}

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
	return z.Desc.With("CliPath", z.spec.CliPath()).With("ToolboxName", app_definitions.Name)
}

// Sort sections by their defined priority
func sortSectionsByPriority(sections []TypedSection) []dc_section.Section {
	// Sort by priority
	sort.Slice(sections, func(i, j int) bool {
		return sections[i].sectionType.Priority() < sections[j].sectionType.Priority()
	})

	// Convert to dc_section.Section array
	result := make([]dc_section.Section, len(sections))
	for i, s := range sections {
		result[i] = s.section
	}
	return result
}

func (z DocCommand) Sections() []dc_section.Section {
	// Get command-specific sections
	commandSections := generateCommandSections(z.media, z.spec)

	// Get common sections
	commonSections := NewCommon(z.media).GetSections()

	// Combine all sections into a single array
	allTypedSections := make([]TypedSection, 0, len(commandSections)+len(commonSections))
	allTypedSections = append(allTypedSections, commandSections...)
	allTypedSections = append(allTypedSections, commonSections...)

	// Sort sections by priority and return
	return sortSectionsByPriority(allTypedSections)
}
