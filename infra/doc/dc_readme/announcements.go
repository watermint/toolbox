package dc_readme

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"os"
	"path/filepath"
)

type AnnouncementNode struct {
	Number    int    `path:"number" json:"number"`
	Title     string `path:"title" json:"title"`
	Url       string `path:"url" json:"url"`
	UpdatedAt string `path:"updated_at" json:"updatedAt"`
}

func NewAnnouncements() dc_section.Section {
	return &Announcements{}
}

type Announcements struct {
	SectionTitle    app_msg.Message
	SectionHeader   app_msg.Message
	NoAnnouncements app_msg.Message
	AnnounceItem    app_msg.Message
}

func (z Announcements) Title() app_msg.Message {
	return z.SectionTitle
}

func (z Announcements) Body(ui app_ui.UI) {
	_ = ui.Text(z.NoAnnouncements)

	root, err := es_project.DetectRepositoryRoot()
	if err != nil {
		panic("Unable to detect repository root: " + err.Error())
	}
	path := filepath.Join(root, "resources", "release", "announcements.json")
	ad, err := os.ReadFile(path)
	if err != nil {
		panic("Unable to load announcements.json: " + err.Error())
	}
	if len(ad) == 0 {
		ui.Info(z.NoAnnouncements)
		return
	}

	aj, err := es_json.Parse(ad)
	if err != nil {
		panic("Unable to parse announcements.json: " + err.Error())
	}

	err = aj.FindArrayEach("data.repository.discussions.nodes", func(e es_json.Json) error {
		n := &AnnouncementNode{}
		if err := e.Model(n); err != nil {
			return err
		}
		ui.Info(z.AnnounceItem.With("Number", n.Number).With("Title", n.Title).With("Url", n.Url))
		return nil
	})
	if err != nil {
		panic("Unable to retrieve announcement data: " + err.Error())
	}
}
