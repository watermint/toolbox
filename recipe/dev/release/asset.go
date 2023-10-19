package release

import (
	"crypto/sha256"
	"github.com/watermint/essentials/eformat/ehex"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_commit"
	"github.com/watermint/toolbox/domain/github/model/mo_content"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Asset struct {
	rc_recipe.RemarkSecret
	Peer gh_conn.ConnGithubRepo

	Branch  string
	Owner   string
	Path    string
	Repo    string
	Text    string
	Message string

	Content rp_model.RowReport
	Commit  rp_model.RowReport
}

func (z *Asset) Preset() {
	z.Content.SetModel(&mo_content.Content{})
	z.Commit.SetModel(&mo_commit.Commit{})
}

func (z *Asset) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Commit.Open(); err != nil {
		return err
	}
	if err := z.Content.Open(); err != nil {
		return err
	}

	shaBytes := sha256.Sum256([]byte(z.Text))
	sha := ehex.ToHexString(shaBytes[:])
	svc := sv_content.New(z.Peer.Client(), z.Owner, z.Repo)
	opts := make([]sv_content.ContentOpt, 0)
	opts = append(opts, sv_content.Branch(z.Branch))
	opts = append(opts, sv_content.Sha(sha))

	l.Debug("Commit", esl.String("sha", sha), esl.String("branch", z.Branch),
		esl.String("path", z.Path), esl.String("message", z.Message),
		esl.String("text", z.Text), esl.Any("opts", opts))

	cts, commit, err := svc.Put(z.Path, z.Message, z.Text, opts...)
	if err != nil {
		l.Debug("Unable to commit the change", esl.Error(err))
		return err
	}
	z.Commit.Row(&commit)
	z.Content.Row(&cts)
	return nil
}

func (z *Asset) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Asset{}, func(r rc_recipe.Recipe) {
		m := r.(*Asset)
		m.Owner = "watermint"
		m.Repo = "toolbox"
		m.Branch = "master"
		m.Path = "release/asset.go"
		m.Text = "https://github.com/watermint/toolbox/blob/main/BUILD.md"
	})
}
