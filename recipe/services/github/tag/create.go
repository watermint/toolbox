package tag

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_tag"
	"github.com/watermint/toolbox/domain/github/service/sv_tag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"time"
)

type Create struct {
	Owner      string
	Repository string
	Tag        string
	Message    string
	Sha1       string
	Created    rp_model.TransactionReport
	Peer       gh_conn.ConnGithubRepo
}

type CreateTag struct {
	Owner      string `json:"owner"`
	Repository string `json:"repository"`
	Tag        string `json:"tag"`
	Message    string `json:"message"`
	Sha1       string `json:"sha_1"`
}

func (z *Create) Preset() {
	z.Created.SetModel(&CreateTag{}, &mo_tag.Tag{})
}

func (z *Create) Exec(c app_control.Control) error {
	if err := z.Created.Open(); err != nil {
		return err
	}
	ct := &CreateTag{
		Owner:      z.Owner,
		Repository: z.Repository,
		Tag:        z.Tag,
		Message:    z.Message,
		Sha1:       z.Sha1,
	}
	_, err := sv_tag.New(z.Peer.Context(), z.Owner, z.Repository).List()
	if err != nil {
		return err
	}

	tag, err := sv_tag.New(z.Peer.Context(), z.Owner, z.Repository).Create(z.Tag, z.Message, z.Sha1)
	if err != nil {
		z.Created.Failure(err, ct)
		return err
	}
	z.Created.Success(ct, tag)
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	sha := "4a6b7fea537d8912b1c0ce1422f270dab1e90d82"

	return rc_exec.Exec(c, &Create{}, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.Owner = "watermint"
		m.Repository = "toolbox_sandobx"
		m.Sha1 = sha
		m.Tag = time.Now().Format("20060103150405")
		m.Message = time.Now().Format("20060103150405") + ": test tag"
	})
}
