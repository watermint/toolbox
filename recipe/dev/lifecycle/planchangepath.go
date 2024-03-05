package lifecycle

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_catalogue_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_compatibility"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Planchangepath struct {
	rc_recipe.RemarkSecret
	CompatibilityFile mo_path.FileSystemPath
	MessageFile       mo_path.FileSystemPath
	CurrentPath       string
	FormerPath        string
	CurrentBase       string
	FormerBase        string
	AnnounceUrl       string
	Date              mo_time.Time
	Compact           bool
}

func (z *Planchangepath) Preset() {
	z.CompatibilityFile = mo_path.NewFileSystemPath("catalogue/catalogue_compatibility.json")
	z.MessageFile = mo_path.NewFileSystemPath("resources/messages/en/messages.json")
	z.CurrentBase = "citron"
	z.FormerBase = "recipe"
}

func (z *Planchangepath) Exec(c app_control.Control) error {
	l := c.Log()

	cds, err := rc_compatibility.LoadOrNewCompatibilityDefinition(z.CompatibilityFile.Path())
	if err != nil {
		l.Error("Unable to load compatibility file", esl.Error(err))
		return err
	}

	rs := app_catalogue.Current().RecipeSpec(z.CurrentPath)
	l.Info("Recipe",
		esl.String("path", rs.CliPath()),
		esl.String("title", c.UI().TextOrEmpty(rs.Title())))

	if d, found := cds.FindPathChange(rs.Path()); found {
		l.Info("Existing pathChange definition found", esl.Any("pathChange", d))
		return nil
	}

	formerPaths := strings.Split(z.FormerPath, " ")
	formerPath := formerPaths[:len(formerPaths)-1]
	formerName := formerPaths[len(formerPaths)-1]

	path, name := rs.Path()
	pcd := rc_compatibility.PathChangeDefinition{
		Announcement:        z.AnnounceUrl,
		PruneAfterBuildDate: z.Date.Iso8601(),
		Current: rc_compatibility.PathPair{
			Path: path,
			Name: name,
		},
		FormerPaths: []rc_compatibility.PathPair{
			{
				Path: formerPath,
				Name: formerName,
			},
		},
	}
	l.Info("New Path", esl.Any("pathChange", pcd))

	cds.PathChanges = append(cds.PathChanges, pcd)
	if err := rc_compatibility.SaveCompatibilityDefinition(z.CompatibilityFile.Path(), cds, z.Compact); err != nil {
		l.Error("Unable to save compatibility file", esl.Error(err))
		return err
	}

	l.Info("Saved", esl.String("path", z.CompatibilityFile.Path()))

	formerMsgPath := z.FormerBase + "." + strings.Join(append(formerPath, formerName), ".")
	currentMsgPath := z.CurrentBase + "." + strings.Join(append(path, name), ".")
	l.Info("Search messages", esl.String("path", formerMsgPath))
	mainMsgFile, err := os.ReadFile(z.MessageFile.Path())
	if err != nil {
		l.Error("Unable to read main message file", esl.Error(err))
		return err
	}
	mainMessages := make(map[string]string)
	if err := json.Unmarshal(mainMsgFile, &mainMessages); err != nil {
		l.Error("Unable to unmarshal main message file", esl.Error(err))
		return err
	}

	newMessages := make(map[string]string)

	for mk, mm := range mainMessages {
		currentKey := strings.Replace(mk, formerMsgPath, currentMsgPath, 1)
		_, found := mainMessages[currentKey]
		if found {
			l.Debug("Key already exists", esl.String("key", currentKey))
			continue
		}

		if strings.HasPrefix(mk, formerMsgPath) {
			l.Info("Found message", esl.String("formerKey", mk), esl.String("newKey", currentKey), esl.String("message", mm))
			newMessages[currentKey] = mm
		}
	}

	if len(newMessages) > 0 {
		for k, v := range newMessages {
			mainMessages[k] = v
		}
		l.Info("Update messages", esl.Int("newMessages", len(newMessages)))
		newMessageBody, err := json.MarshalIndent(mainMessages, "", "  ")
		if err != nil {
			l.Error("Unable to marshal new messages", esl.Error(err))
			return err
		}
		if err := os.WriteFile(z.MessageFile.Path(), newMessageBody, 0644); err != nil {
			l.Error("Unable to write new messages", esl.Error(err))
			return err
		}
	}

	return nil
}

func (z *Planchangepath) Test(c app_control.Control) error {
	d, err := qt_file.MakeTestFolder("prune", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(d)
	}()

	app_catalogue.SetCurrent(rc_catalogue_impl.NewCatalogue(
		[]rc_recipe.Recipe{
			&Planchangepath{},
		},
		[]rc_recipe.Recipe{},
		[]interface{}{},
		[]app_feature.OptIn{},
	))

	msgSnapshot, err := os.ReadFile("resources/messages/en/messages.json")
	if err != nil {
		return err
	}

	pathCompat := filepath.Join(d, "test_compatibility.json")
	pathMessages := filepath.Join(d, "test_messages.json")

	if err := os.WriteFile(pathMessages, msgSnapshot, 0644); err != nil {
		return err
	}

	err = rc_exec.Exec(c, &Planchangepath{}, func(r rc_recipe.Recipe) {
		m := r.(*Planchangepath)
		m.CurrentPath = "dev lifecycle planchangepath"
		m.FormerPath = "dev lifecycle planpathchange"
		m.CompatibilityFile = mo_path.NewFileSystemPath(pathCompat)
		m.MessageFile = mo_path.NewFileSystemPath(pathMessages)
		m.Date = mo_time.New(time.Date(2123, 12, 24, 10, 30, 0, 0, time.UTC))
		m.AnnounceUrl = "https://github.com/watermint/toolbox/issues/781"
		m.Compact = true
	})
	if err != nil {
		return err
	}

	content, err := os.ReadFile(pathCompat)
	if err != nil {
		return err
	}
	expected := `{"path_change":[{"announcement":"https://github.com/watermint/toolbox/issues/781","prune_after_build_date":"2123-12-24T10:30:00Z","current":{"path":["dev","lifecycle"],"name":"planchangepath"},"former_paths":[{"path":["dev","lifecycle"],"name":"planpathchange"}]}],"prune":[]}`
	if string(content) != expected {
		c.Log().Warn("Unexpected content", esl.String("expected", expected), esl.String("actual", string(content)))
		return errors.New("unexpected content")
	}
	return nil
}
