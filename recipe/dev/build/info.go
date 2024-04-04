package build

import (
	"encoding/json"
	"errors"
	"github.com/go-git/go-git/v5"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/security/sc_zap"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/resources"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Info struct {
	rc_recipe.RemarkSecret
	FailFast bool
}

func (z *Info) Preset() {
}

func (z *Info) Exec(c app_control.Control) error {
	l := c.Log()

	productionReady := true

	prjBase, err := es_project.DetectRepositoryRoot()
	if err != nil {
		l.Warn("Unable to detect the repository root", esl.Error(err))
		return err
	}

	prjGit := filepath.Join(prjBase, ".git")
	repo, err := git.PlainOpen(prjGit)
	if err != nil {
		l.Warn("Unable to open the .git", esl.Error(err))
		return err
	}

	hash, err := repo.ResolveRevision("HEAD")
	if err != nil {
		l.Warn("Unable to detect the hash", esl.Error(err))
		return err
	}

	head, err := repo.Head()
	if err != nil {
		l.Warn("Unable to detect the head", esl.Error(err))
		return err
	}

	headName := string(head.Name())
	if !strings.HasPrefix(headName, "refs/heads") {
		l.Warn("Unexpected ref format", esl.String("head", headName))
		return errors.New("unexpected git refs")
	}

	branch := strings.ReplaceAll(headName, "refs/heads/", "")

	xap, found := os.LookupEnv(app_definitions.EnvNameToolboxBuilderKey)
	if !found || len(xap) < 10 {
		l.Warn("Builder key not found or too short. Please set the build key for production release", esl.String("key", app_definitions.EnvNameToolboxBuilderKey), esl.Int("length", len(xap)))
		xap = ""
		productionReady = false
		if z.FailFast {
			return errors.New("builder key not found")
		}
	}

	var zap string
	zap = sc_zap.NewZap(hash.String())
	appKeyData, found := os.LookupEnv(app_definitions.EnvNameToolboxAppKeys)
	if !found {
		l.Warn("App key data not found. Please set the build key for production release", esl.String("key", app_definitions.EnvNameToolboxAppKeys))
		zap = ""
		productionReady = false
		if z.FailFast {
			return errors.New("app key data not found")
		}
	} else {
		if !gjson.Valid(appKeyData) {
			l.Warn("App key data is not look like a JSON data", esl.Int("length", len(appKeyData)))
			return errors.New("invalid app key data format")
		}
		if err := sc_zap.Zap(zap, prjBase, []byte(appKeyData)); err != nil {
			l.Warn("Unable to zap the data", esl.Error(err))
			return err
		}
	}

	licenseSalt, found := os.LookupEnv(app_definitions.EnvNameToolboxLicenseSalt)
	if !found {
		l.Warn("License salt not found", esl.String("key", app_definitions.EnvNameToolboxLicenseSalt))
		licenseSalt = ""
		productionReady = false
		if z.FailFast {
			return errors.New("license salt not found")
		}
	}

	buildTimestamp := time.Now().UTC()
	info := resources.BuildInfo{
		Version:     app_definitions.BuildId,
		Hash:        hash.String(),
		Branch:      branch,
		Timestamp:   buildTimestamp.Format(time.RFC3339),
		Year:        buildTimestamp.Year(),
		Zap:         zap,
		Xap:         xap,
		Production:  productionReady,
		LicenseSalt: licenseSalt,
	}

	infoPath := filepath.Join(prjBase, "resources/build", "info.json")
	l.Info("Build info", esl.Any("branch", branch), esl.Any("hash", info.Hash), esl.String("version", app_definitions.BuildId), esl.Bool("releaseReady", productionReady))
	infoData, err := json.Marshal(info)
	if err != nil {
		l.Debug("Unable to marshal the data", esl.Error(err))
		return err
	}

	if err := os.WriteFile(infoPath, infoData, 0600); err != nil {
		l.Warn("Unable to write the file", esl.Error(err))
		return err
	}
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
