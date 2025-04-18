package ig_homebrew

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_commit"
	"github.com/watermint/toolbox/domain/github/service/sv_content"
	"github.com/watermint/toolbox/essentials/file/es_filehash"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"path/filepath"
	"text/template"
	"time"
)

var (
	ErrorNotAFile = errors.New("not a file")
)

type Formula struct {
	AssetPathLinuxArm     mo_path.ExistingFileSystemPath
	AssetPathLinuxIntel   mo_path.ExistingFileSystemPath
	AssetPathMacArm       mo_path.ExistingFileSystemPath
	DownloadUrlLinuxArm   string
	DownloadUrlLinuxIntel string
	DownloadUrlMacArm     string
	Message               string
	Branch                string
	Owner                 string
	Repository            string
	FormulaName           string
	Peer                  gh_conn.ConnGithubRepo
	Commit                rp_model.RowReport
}

func (z *Formula) Preset() {
	z.Commit.SetModel(&mo_commit.Commit{})
}

func (z *Formula) makeFormula(c app_control.Control) (formula string, err error) {
	l := c.Log()
	h := es_filehash.NewHash(l)
	sha256LinuxIntel, err := h.SHA256(z.AssetPathLinuxIntel.Path())
	if err != nil {
		l.Debug("Unable to calculate SHA sum of the asset", esl.Error(err))
		return "", err
	}
	sha256LinuxArm, err := h.SHA256(z.AssetPathLinuxArm.Path())
	if err != nil {
		l.Debug("Unable to calculate SHA sum of the asset", esl.Error(err))
		return "", err
	}
	sha256MacArm, err := h.SHA256(z.AssetPathMacArm.Path())
	if err != nil {
		l.Debug("Unable to calculate SHA sum of the asset", esl.Error(err))
		return "", err
	}

	resBundle := app_resource.Bundle()
	formulaSrc, err := resBundle.Templates().Bytes("homebrew-toolbox.rb.tmpl")
	if err != nil {
		l.Debug("Unable to find a template resource", esl.Error(err))
		return "", err
	}

	formulaTmpl, err := template.New("formula").Parse(string(formulaSrc))
	if err != nil {
		l.Debug("Unable to parse", esl.Error(err))
		return "", err
	}

	var buf bytes.Buffer
	err = formulaTmpl.Execute(&buf, map[string]string{
		"Version":               app_definitions.BuildId,
		"DownloadUrlLinuxArm":   z.DownloadUrlLinuxArm,
		"DownloadUrlLinuxIntel": z.DownloadUrlLinuxIntel,
		"DownloadUrlMacArm":     z.DownloadUrlMacArm,
		"Sha256LinuxArm":        sha256LinuxArm,
		"Sha256LinuxIntel":      sha256LinuxIntel,
		"Sha256MacArm":          sha256MacArm,
	})
	if err != nil {
		l.Debug("Unable to compile template", esl.Error(err))
		return "", err
	}
	return buf.String(), nil
}

func (z *Formula) getCurrentSha(c app_control.Control) (sha string, err error) {
	l := c.Log()
	svc := sv_content.New(z.Peer.Client(), z.Owner, z.Repository)
	opts := make([]sv_content.ContentOpt, 0)
	opts = append(opts, sv_content.Ref(z.Branch))

	cts, err := svc.Get(z.FormulaName, opts...)
	if err != nil {
		l.Debug("Unable to retrieve content metadata", esl.Error(err))
		return "", err
	}
	if f, ok := cts.File(); ok {
		l.Debug("Content metadata", esl.Any("file", f))
		return f.Sha, nil
	}
	l.Debug("not a file", esl.Any("cts", cts))
	return "", ErrorNotAFile
}

func (z *Formula) updateFormula(c app_control.Control, formula, sha string) error {
	l := c.Log()
	svc := sv_content.New(z.Peer.Client(), z.Owner, z.Repository)
	opts := make([]sv_content.ContentOpt, 0)
	opts = append(opts, sv_content.Branch(z.Branch))
	opts = append(opts, sv_content.Sha(sha))

	cts, commit, err := svc.Put(z.FormulaName, z.Message, formula, opts...)
	if err != nil {
		l.Debug("Unable to commit the change", esl.Error(err))
		return err
	}
	l.Debug("contents metadata", esl.Any("contents", cts))
	l.Debug("commit metadata", esl.Any("commit", commit))

	z.Commit.Row(commit)

	return nil
}

func (z *Formula) Exec(c app_control.Control) error {
	if err := z.Commit.Open(); err != nil {
		return err
	}

	formula, err := z.makeFormula(c)
	if err != nil {
		return err
	}

	sha, err := z.getCurrentSha(c)
	if err != nil {
		return err
	}

	return z.updateFormula(c, formula, sha)
}

func (z *Formula) Test(c app_control.Control) error {
	root, err := es_project.DetectRepositoryRoot()
	if err != nil {
		return err
	}

	return rc_exec.Exec(c, &Formula{}, func(r rc_recipe.Recipe) {
		m := r.(*Formula)
		m.AssetPathLinuxArm = mo_path.NewExistingFileSystemPath(filepath.Join(root, "SECURITY_AND_PRIVACY.md"))
		m.AssetPathLinuxIntel = mo_path.NewExistingFileSystemPath(filepath.Join(root, "README.md"))
		m.AssetPathMacArm = mo_path.NewExistingFileSystemPath(filepath.Join(root, "LICENSE.md"))
		m.DownloadUrlLinuxArm = "https://raw.githubusercontent.com/watermint/toolbox/master/SECURITY_AND_PRIVACY.md"
		m.DownloadUrlLinuxIntel = "https://raw.githubusercontent.com/watermint/toolbox/master/README.md"
		m.DownloadUrlMacArm = "https://raw.githubusercontent.com/watermint/toolbox/master/LICENSE.md"
		m.Message = "Release:" + time.Now().Format(time.RFC3339)
		m.FormulaName = "toolbox.rb"
		m.Branch = "current"
		m.Owner = "watermint"
		m.Repository = "toolbox_sandbox"
	})
}
