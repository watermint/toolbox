package release

import (
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/service/sv_graphql"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_announcement"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"os"
	"path/filepath"
	"text/template"
)

type Announcement struct {
	rc_recipe.RemarkSecret
	Peer          gh_conn.ConnGithubRepo
	Repository    string
	Owner         string
	CategoryId    string
	Announcements rp_model.RowReport
}

var (
	AnnounceDiscussionQuery = `
query {
	repository(owner: "{{.Owner}}", name: "{{.Repository}}") {
		discussions(first: 11, categoryId: "{{.CategoryId}}", states: [OPEN], orderBy: {field: UPDATED_AT, direction: DESC}) {
			nodes {
				number
				title
				url
				updatedAt
			}
		}
	}
}
`
)

func (z *Announcement) Preset() {
	z.Owner = "watermint"
	z.Repository = "toolbox"
	z.CategoryId = "DIC_kwDOBFqRm84CQesd"
	z.Announcements.SetModel(&dc_announcement.AnnouncementNode{})
}

func (z *Announcement) Exec(c app_control.Control) error {
	if err := z.Announcements.Open(); err != nil {
		return err
	}
	root, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}

	announceFilePath := filepath.Join(root, "resources", "release", "announcements.json")
	l := c.Log().With(esl.String("announceFilePath", announceFilePath))

	l.Debug("Querying announcements")
	query := bytes.Buffer{}
	queryTmpl := template.Must(template.New("query").Parse(AnnounceDiscussionQuery))
	err = queryTmpl.Execute(&query, map[string]string{
		"Owner":      z.Owner,
		"Repository": z.Repository,
		"CategoryId": z.CategoryId,
	})
	l.Debug("Query", esl.String("query", query.String()))

	result, err := sv_graphql.NewQuery(z.Peer.Client()).Query(query.String())
	if err != nil {
		l.Debug("Unable to query", esl.Error(err))
		return err
	}
	var announcements map[string]interface{}
	if err := json.Unmarshal(result.Raw(), &announcements); err != nil {
		l.Debug("Unable to parse", esl.Error(err))
		return err
	}
	announcementsIndent, err := json.MarshalIndent(announcements, "", "  ")
	if err != nil {
		l.Debug("Unable to indent", esl.Error(err))
		return err
	}

	err = result.FindArrayEach("data.repository.discussions.nodes", func(e es_json.Json) error {
		a := &dc_announcement.AnnouncementNode{}
		if err := e.Model(a); err != nil {
			l.Debug("Unable to parse announcement", esl.Error(err))
			return err
		}
		z.Announcements.Row(a)
		return nil
	})
	if err != nil {
		l.Debug("Unable to find array", esl.Error(err))
		return err
	}

	return os.WriteFile(announceFilePath, announcementsIndent, 0644)
}

func (z *Announcement) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Announcement{}, func(r rc_recipe.Recipe) {
		m := r.(*Announcement)
		m.Repository = "toolbox_sandbox"
		m.Owner = "watermint"
		m.CategoryId = "DIC_xxxx"
	})
}
