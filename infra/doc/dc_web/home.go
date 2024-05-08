package dc_web

import (
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/doc/dc_announcement"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/doc/dc_supplemental"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

const (
	TaglineImage = "![watermint toolbox]({{ site.baseurl }}/images/logo.png){: width=\"160\" }"
	TaglineDemo  = "![Demo]({{ site.baseurl }}/images/demo.gif)"
)

func Home() dc_section.Document {
	return &homeDoc{}
}

type homeDoc struct {
	Tagline app_msg.Message
}

func (z homeDoc) DocId() dc_index.DocId {
	return dc_index.DocWebHome
}

func (z homeDoc) DocDesc() app_msg.Message {
	return z.Tagline
}

func (z homeDoc) Sections() []dc_section.Section {
	return []dc_section.Section{
		&homeTagline{},
		&homeUsage{},
		dc_readme.NewRelease(),
		dc_readme.NewLicense(),
		dc_announcement.NewAnnouncements(),
		dc_readme.NewSecuritySection(),
	}
}

type homeTagline struct {
	Header  app_msg.Message
	Tagline app_msg.Message
}

func (z homeTagline) Title() app_msg.Message {
	return z.Header
}

func (z homeTagline) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(TaglineImage))
	ui.Break()
	ui.Info(z.Tagline)
}

type homeUsage struct {
	Header      app_msg.Message
	Desc1       app_msg.Message
	Desc2       app_msg.Message
	Refs        app_msg.Message
	HeadDocName app_msg.Message
}

func (z homeUsage) Title() app_msg.Message {
	return z.Header
}

func (z homeUsage) Body(ui app_ui.UI) {
	lg := ui.Messages().Lang()

	numPublicCommands := 0
	for _, r := range app_catalogue.Current().Recipes() {
		spec := rc_spec.New(r)
		if !spec.IsSecret() {
			numPublicCommands++
		}
	}

	ui.Info(z.Desc1.With("NumCommands", numPublicCommands))
	ui.Break()
	ui.Info(app_msg.Raw(TaglineDemo))
	ui.Break()
	ui.Info(z.Desc2)
	ui.Break()
	ui.Info(z.Refs)
	ui.Break()

	docs := []dc_section.Document{
		NewCommands(ui.Messages()),
		dc_supplemental.NewDropboxBusiness(dc_index.MediaWeb),
	}
	ui.WithTable("references", func(t app_ui.Table) {
		t.Header(z.HeadDocName)

		for _, doc := range docs {
			compiledDoc := app_msg.Apply(doc).(dc_section.Document)
			path := dc_index.DocName(dc_index.MediaWeb, compiledDoc.DocId(), lg, dc_index.RefPath(true))
			name := ui.Text(compiledDoc.DocDesc())
			t.RowRaw("[" + name + "](" + path + ")")
		}
	})
}
