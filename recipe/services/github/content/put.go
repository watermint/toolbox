package content

import (
	"fmt"
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_commit"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"io/ioutil"
	"os"
	"time"
)

type Put struct {
	Owner      string
	Repository string
	Path       string
	Message    string
	Content    mo_path.ExistingFileSystemPath
	Branch     mo_string.OptionalString
	Commit     rp_model.RowReport
	Peer       gh_conn.ConnGithubRepo
}

func (z *Put) Preset() {
	z.Commit.SetModel(&mo_commit.Commit{})
}

func (z *Put) Exec(c app_control.Control) error {
	l := c.Log().With(esl.String("contentPath", z.Content.Path()))
	content, err := ioutil.ReadFile(z.Content.Path())
	if err != nil {
		l.Debug("Unable to read file", esl.Error(err))
		return err
	}
	if err := z.Commit.Open(); err != nil {
		return err
	}

	opts := make([]sv_content.ContentOpt, 0)
	if z.Branch.IsExists() {
		opts = append(opts, sv_content.Branch(z.Branch.Value()))
	}

	svc := sv_content.New(z.Peer.Context(), z.Owner, z.Repository)
	existing, err := svc.Get(z.Path, opts...)
	if err != nil {
		l.Debug("no prior content found", esl.Error(err))
	} else {
		if f, ok := existing.File(); ok {
			opts = append(opts, sv_content.Sha(f.Sha))
			l.Debug("Adding content sha", esl.String("sha", f.Sha))
		}
	}

	cts, commit, err := svc.Put(z.Path, z.Message, string(content), opts...)
	if err != nil {
		l.Debug("unable to create or update", esl.Error(err))
		return err
	}
	l.Debug("content", esl.Any("content", cts))
	l.Debug("commit", esl.Any("commit", commit))
	z.Commit.Row(commit)

	return nil
}

func (z *Put) Test(c app_control.Control) error {
	f, err := ioutil.TempFile("", "content-put")
	if err != nil {
		return err
	}
	path := f.Name()
	defer func() {
		_ = f.Close()
		_ = os.Remove(path)
	}()
	if _, err := fmt.Fprintf(f, "# content put at %s\n", time.Now().Format(time.RFC3339)); err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Put{}, func(r rc_recipe.Recipe) {
		m := r.(*Put)
		m.Owner = "watermint"
		m.Repository = "toolbox_sandobx"
		m.Path = "README.md"
		m.Message = time.Now().Format(time.RFC3339)
		m.Content = mo_path.NewExistingFileSystemPath(path)
	})
}
