package release

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_commit"
	"github.com/watermint/toolbox/domain/github/model/mo_content"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/domain/github/service/sv_release"
	"github.com/watermint/toolbox/domain/github/service/sv_release_asset"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Asseturl struct {
	rc_recipe.RemarkSecret
	Peer         gh_conn.ConnGithubRepo
	SourceRepo   string
	SourceOwner  string
	TargetBranch string
	TargetRepo   string
	TargetOwner  string
	Content      rp_model.RowReport
	Commit       rp_model.RowReport
}

func (z *Asseturl) Preset() {
	z.Content.SetModel(&mo_content.Content{})
	z.Commit.SetModel(&mo_commit.Commit{})
}

func (z *Asseturl) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.Commit.Open(); err != nil {
		return err
	}
	if err := z.Content.Open(); err != nil {
		return err
	}

	latest, err := sv_release.New(z.Peer.Client(), z.SourceOwner, z.SourceRepo).Latest()
	if err != nil {
		return err
	}
	l.Info("Latest release", esl.String("release", latest.Name))
	assets, err := sv_release_asset.New(z.Peer.Client(), z.SourceOwner, z.SourceRepo, latest.Id).List()
	if err != nil {
		return err
	}

	svc := sv_content.New(z.Peer.Client(), z.TargetOwner, z.TargetRepo)
	for _, asset := range assets {
		platform := IdentifyPlatform(asset)
		if platform == AssetPlatformUnknown {
			continue
		}

		path := "/latest/" + platform + ".url"
		existing, err := svc.Get(path)
		var sha string
		if err != nil {
			l.Debug("Unable to get the content, ignore", esl.Error(err))
		} else if f, found := existing.File(); found {
			l.Debug("Existing file", esl.String("sha", f.Sha))
			sha = f.Sha
		} else {
			l.Debug("The path is not a file", esl.Any("existing", existing))
		}

		l.Info("Asset", esl.String("path", path), esl.String("asset", asset.Name), esl.String("platform", platform))
		assetUrl := asset.DownloadUrl

		opts := []sv_content.ContentOpt{
			sv_content.Branch(z.TargetBranch),
			sv_content.Sha(sha),
		}
		msg := "Update Release " + latest.Name + " asset URL of " + platform
		cts, commit, err := svc.Put(path, msg, assetUrl, opts...)
		if err != nil {
			l.Error("Unable to commit the change", esl.Error(err))
			return err
		}
		z.Commit.Row(&commit)
		z.Content.Row(&cts)
	}
	return nil
}

func (z *Asseturl) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Asseturl{}, func(r rc_recipe.Recipe) {
		m := r.(*Asseturl)
		m.SourceOwner = "watermint"
		m.SourceRepo = "toolbox"
		m.TargetOwner = "watermint"
		m.TargetRepo = "homebrew-toolbox"
		m.TargetBranch = "master"
	})
}
