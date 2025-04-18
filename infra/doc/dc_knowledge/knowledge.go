package dc_knowledge

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/terminal/es_dialogue"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/doc/dc_command"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_options"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

// MsgKnowledge defines messages for knowledge base documentation
type MsgKnowledge struct {
	DocDesc             app_msg.Message
	Title               app_msg.Message
	Description         app_msg.Message
	CommonOptionsTitle  app_msg.Message
	CommandsTitle       app_msg.Message
	AdditionalDocsTitle app_msg.Message
}

var (
	MKnowledge = app_msg.Apply(&MsgKnowledge{}).(*MsgKnowledge)
)

// BaseURL is the base URL for the documentation website
var BaseURL = app_definitions.LandingPage

// KnowledgeGenerator is an interface for generating text data for LLM training
type KnowledgeGenerator interface {
	// GenerateKnowledge generates knowledge data from command specifications and additional documents
	GenerateKnowledge(specs []rc_recipe.Spec, additionalDocs []dc_section.Document) string
}

// KnowledgeDoc is a struct representing a document for LLM training
type KnowledgeDoc struct {
	media dc_index.MediaType
}

// New creates a new instance of KnowledgeGenerator
func New(media dc_index.MediaType) KnowledgeGenerator {
	return &KnowledgeDoc{
		media: media,
	}
}

// DocId returns the document ID
func (z *KnowledgeDoc) DocId() dc_index.DocId {
	return dc_index.DocManualCommand
}

// DocDesc returns the document description
func (z *KnowledgeDoc) DocDesc() app_msg.Message {
	return MKnowledge.DocDesc
}

// buildURL builds a full URL from the base URL and path
func (z *KnowledgeDoc) buildURL(path string) string {
	return fmt.Sprintf("%s/%s", BaseURL, strings.TrimPrefix(path, "docs/"))
}

// writeMultiMarkdownHeader writes the multi-markdown header with metadata
func (z *KnowledgeDoc) writeMultiMarkdownHeader(builder *strings.Builder, title string, docId dc_index.DocId, commandName string) {
	builder.WriteString("---\n")
	builder.WriteString(fmt.Sprintf("Title: %s\n", title))
	path := dc_index.DocName(dc_index.MediaWeb, docId, es_lang.English, dc_index.CommandName(commandName), dc_index.RefPath(false))
	builder.WriteString(fmt.Sprintf("URL: %s\n", z.buildURL(path)))
	builder.WriteString("---\n\n")
}

// generateCommonOptionsDoc generates documentation for common options
func (z *KnowledgeDoc) generateCommonOptionsDoc(spec rc_recipe.Spec, mc app_msg_container.Container, logger esl.Logger) string {
	var docText strings.Builder
	z.writeMultiMarkdownHeader(&docText, "Common Options for All Commands", dc_index.DocManualCommand, "common-options")

	var buf bytes.Buffer
	ui := app_ui.NewMarkdown(mc, logger, &buf, es_dialogue.DenyAll())

	dc_command.GenerateCommonOptionsDetail(ui,
		MKnowledge.CommonOptionsTitle,
		dc_options.MDoc.HeaderOption,
		dc_options.MDoc.HeaderDescription,
		dc_options.MDoc.HeaderDefault)

	docText.WriteString(buf.String())
	docText.WriteString("\n\n")

	return docText.String()
}

// generateCommandDoc generates documentation for a single command
func (z *KnowledgeDoc) generateCommandDoc(spec rc_recipe.Spec, mc app_msg_container.Container, logger esl.Logger) string {
	if spec.IsSecret() {
		return ""
	}
	var docText strings.Builder
	cliPath := spec.CliPath()
	urlPath := strings.ReplaceAll(cliPath, " ", "/")

	// Write header for this command
	z.writeMultiMarkdownHeader(&docText, cliPath, dc_index.DocManualCommand, urlPath)

	// Generate command documentation
	cmdDoc := dc_command.New(z.media, spec)
	sections := cmdDoc.Sections()

	// Add each section
	var buf bytes.Buffer
	ui := app_ui.NewMarkdown(mc, logger, &buf, es_dialogue.DenyAll())
	for _, section := range sections {
		sec := app_msg.Apply(section).(dc_section.Section)
		ui.Header(sec.Title())
		sec.Body(ui)
		ui.Break()
	}

	docText.WriteString(buf.String())
	docText.WriteString("\n\n")

	return docText.String()
}

// generateAdditionalDoc generates documentation for an additional document
func (z *KnowledgeDoc) generateAdditionalDoc(doc dc_section.Document, mc app_msg_container.Container) string {
	var docText strings.Builder
	docId := doc.DocId()
	docDesc := mc.Compile(doc.DocDesc())

	// Write header for this additional document
	z.writeMultiMarkdownHeader(&docText, docDesc, docId, "")

	// Generate document content
	docContent := dc_section.Generate(z.media, dc_section.LayoutPage, mc, doc)
	docText.WriteString(docContent)
	docText.WriteString("\n\n")

	return docText.String()
}

// GenerateKnowledge generates knowledge data from command specifications and additional documents
func (z *KnowledgeDoc) GenerateKnowledge(specs []rc_recipe.Spec, additionalDocs []dc_section.Document) string {
	var knowledgeText strings.Builder
	mc := app_msg_container_impl.NewContainer()
	l := esl.Default()

	l.Debug("Start generating knowledge base documentation", esl.Int("specs", len(specs)), esl.Int("additionalDocs", len(additionalDocs)))

	// Generate index header
	l.Debug("Generating index header", esl.String("title", mc.Compile(MKnowledge.Title)))
	z.writeMultiMarkdownHeader(&knowledgeText, mc.Compile(MKnowledge.Title), dc_index.DocWebHome, "")
	knowledgeText.WriteString(fmt.Sprintf("# %s\n\n", mc.Compile(MKnowledge.Title)))
	knowledgeText.WriteString(fmt.Sprintf("%s\n\n", mc.Compile(MKnowledge.Description)))

	// Generate common options documentation
	if len(specs) > 0 {
		l.Debug("Generating common options documentation", esl.String("firstCommand", specs[0].CliPath()))
		commonOptionsDoc := z.generateCommonOptionsDoc(specs[0], mc, l)
		knowledgeText.WriteString(commonOptionsDoc)
	}

	// Generate command documentation
	l.Debug("Start generating command documentation", esl.Int("commands", len(specs)))
	knowledgeText.WriteString(fmt.Sprintf("## %s\n\n", mc.Compile(MKnowledge.CommandsTitle)))
	for i, spec := range specs {
		if spec.IsSecret() {
			l.Debug("Skip secret command", esl.String("command", spec.CliPath()))
			continue
		}
		cliPath := spec.CliPath()
		l.Debug("Generating documentation for command",
			esl.String("command", cliPath),
			esl.Int("index", i+1),
			esl.Int("total", len(specs)))

		urlPath := strings.ReplaceAll(cliPath, " ", "/")
		path := dc_index.DocName(z.media, dc_index.DocManualCommand, es_lang.English, dc_index.CommandName(urlPath), dc_index.RefPath(false))
		l.Debug("Command URL path generated", esl.String("path", path))

		commandDoc := z.generateCommandDoc(spec, mc, l)
		knowledgeText.WriteString(commandDoc)
	}
	l.Debug("Completed generating command documentation")

	// Generate additional documentation
	l.Debug("Start generating additional documentation", esl.Int("documents", len(additionalDocs)))
	knowledgeText.WriteString(fmt.Sprintf("## %s\n\n", mc.Compile(MKnowledge.AdditionalDocsTitle)))
	for i, doc := range additionalDocs {
		docId := doc.DocId()
		docDescZ := doc.DocDesc()
		if docDescZ == nil {
			l.Debug("Skip additional document without description", esl.Int("docId", int(docId)))
			panic("additional document without description")
		}
		l.Debug("Generating additional document",
			esl.String("docDesc", doc.DocDesc().Key()),
			esl.Int("index", i+1),
			esl.Int("total", len(additionalDocs)))

		docDesc := mc.Compile(doc.DocDesc())
		path := dc_index.DocName(z.media, docId, es_lang.English, dc_index.RefPath(false))
		l.Debug("Document URL path generated", esl.String("path", path))

		knowledgeText.WriteString(fmt.Sprintf("- [%s](%s)\n", docDesc, z.buildURL(path)))
		additionalDoc := z.generateAdditionalDoc(doc, mc)
		knowledgeText.WriteString(additionalDoc)
	}
	l.Debug("Completed generating additional documentation")

	l.Debug("Completed generating knowledge base documentation")
	return knowledgeText.String()
}
