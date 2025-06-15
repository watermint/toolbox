package doc

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_knowledge"
	"github.com/watermint/toolbox/infra/doc/dc_supplemental"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os"
	"path/filepath"
)

type MsgKnowledge struct {
	ProgressGenerate app_msg.Message
	SuccessGenerate  app_msg.Message
}

var (
	MKnowledge = app_msg.Apply(&MsgKnowledge{}).(*MsgKnowledge)
)

type Knowledge struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkExperimental
	CommandLimit int
	OutputPath   string
}

func (z *Knowledge) Preset() {
	z.CommandLimit = 50
}

func (z *Knowledge) Exec(c app_control.Control) error {
	l := c.Log()

	// Get all recipes and convert them to specifications
	recipes := app_catalogue.Current().Recipes()
	specs := make([]rc_recipe.Spec, 0, len(recipes))

	count := 0
	for _, r := range recipes {
		if count >= z.CommandLimit && z.CommandLimit > 0 {
			break
		}
		spec := rc_spec.New(r)
		if !spec.IsSecret() {
			specs = append(specs, spec)
			count++
		}
	}

	l.Info("Generating reduced knowledge base", esl.Int("commandCount", len(specs)), esl.Int("limit", z.CommandLimit))

	// Get supplemental documents for additional content
	additionalDocs := dc_supplemental.Docs(dc_index.MediaKnowledge)

	// Generate knowledge base content
	c.UI().Progress(MKnowledge.ProgressGenerate)
	knowledge := dc_knowledge.New(dc_index.MediaKnowledge)
	doc := knowledge.GenerateKnowledge(specs, additionalDocs)

	// Determine output path
	outputPath := z.OutputPath
	if outputPath == "" {
		// Default to docs/knowledge/knowledge_reduced.md
		outputPath = filepath.Join("docs", "knowledge", "knowledge_reduced.md")
	}

	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		l.Error("Failed to create directory", esl.Error(err), esl.String("dir", dir))
		return err
	}

	// Write the knowledge file
	l.Info("Writing reduced knowledge base", esl.String("path", outputPath))
	if err := os.WriteFile(outputPath, []byte(doc), 0644); err != nil {
		l.Error("Failed to write knowledge file", esl.Error(err))
		return err
	}

	c.UI().Success(MKnowledge.SuccessGenerate.With("Path", outputPath).With("Commands", len(specs)))

	return nil
}

func (z *Knowledge) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Knowledge{}, func(r rc_recipe.Recipe) {
		m := r.(*Knowledge)
		m.CommandLimit = 5
	})
}