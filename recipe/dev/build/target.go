package build

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/resources"
	"os"
	"path/filepath"
	"strings"
)

type Target struct {
	rc_recipe.RemarkSecret
	TargetName string
	BuildPath  string
	DistPath   string
	DeployPath string
}

func (z *Target) Preset() {
}

func (z *Target) needUpdateInfo(l esl.Logger, prjBase string) bool {
	infoPath := filepath.Join(prjBase, "resources/build", "info.json")
	infoData, err := os.ReadFile(infoPath)
	if os.IsNotExist(err) {
		l.Debug("The info.json does not found", esl.Error(err))
		return true
	}
	info := resources.BuildInfo{}
	err = json.Unmarshal(infoData, &info)
	if err != nil {
		l.Debug("Unable to parse info.json", esl.Error(err))
		return true
	}

	if info.Version != app.BuildId {
		l.Debug("Need to update build info")
		return true
	}

	return false
}

func (z *Target) osArch(target string) (targetOs, targetArch string) {
	switch target {
	case "win":
		return "windows", "amd64"
	//case "win-arm":
	//	return "windows", "arm64"
	case "linux":
		return "linux", "amd64"
	case "linux-arm":
		return "linux", "arm64"
	case "darwin":
		return "darwin", "amd64"
	case "darwin-arm":
		return "darwin", "arm64"
	default:
		panic("unsupported os/arch combination [" + target + "]")
	}
}

func (z *Target) Exec(c app_control.Control) error {
	l := c.Log()
	prjBase, err := es_project.DetectRepositoryRoot()
	if err != nil {
		l.Debug("Unable to detect the repository root", esl.Error(err))
		return err
	}

	if z.needUpdateInfo(l, prjBase) && !c.Feature().IsTest() {
		l.Info("Updating info.json")
		if err = rc_exec.Exec(c, &Info{}, rc_recipe.NoCustomValues); err != nil {
			l.Debug("Unable to create info.json", esl.Error(err))
			return err
		}
	}

	targetBuildPath := filepath.Join(z.BuildPath, z.TargetName)
	if err := os.MkdirAll(targetBuildPath, 0755); err != nil {
		l.Debug("Unable to create target build path", esl.Error(err))
		return err
	}
	if err := os.MkdirAll(z.DistPath, 0755); err != nil {
		l.Debug("Unable to create dist path", esl.Error(err))
		return err
	}

	compileBuildFile := filepath.Join(targetBuildPath, "build-compile.sh")

	targetOs, targetArch := z.osArch(z.TargetName)
	targetPath := ""
	if targetOs == "windows" {
		targetPath = filepath.Join(targetBuildPath, "tbx.exe")
	} else {
		targetPath = filepath.Join(targetBuildPath, "tbx")
	}
	err = rc_exec.Exec(c, &Compile{}, func(r rc_recipe.Recipe) {
		m := r.(*Compile)
		m.Os = targetOs
		m.Arch = targetArch
		m.Package = "github.com/watermint/toolbox"
		m.Path = targetPath
		m.Out = compileBuildFile
	})
	if err != nil {
		l.Debug("Unable to create compile build file", esl.Error(err))
		return err
	}

	targetBuildFile := filepath.Join(z.BuildPath, "build-target.sh")
	{
		buildPkg := func() string {
			return strings.Join([]string{
				"go run tbx.go dev build package",
				"-build-path \"" + targetPath + "\"",
				"-dist-path \"" + z.DistPath + "\"",
				"-deploy-path \"" + z.DeployPath + "\"",
				"-platform " + z.TargetName,
			}, " ")
		}
		lines := []string{
			"echo Build target: " + z.TargetName,
			"echo Build path:   " + z.BuildPath,
			"echo Dist path:    " + z.DistPath,
			"echo Build file:   " + targetBuildFile,
			"bash " + compileBuildFile,
			buildPkg(),
		}
		return os.WriteFile(targetBuildFile, []byte(strings.Join(lines, "\n")), 0755)
	}
}

func (z *Target) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("compile", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()
	err = rc_exec.Exec(c, &Target{}, func(r rc_recipe.Recipe) {
		m := r.(*Target)
		m.TargetName = "win"
		m.DistPath = p
		m.BuildPath = p
		m.DeployPath = p
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Target{}, func(r rc_recipe.Recipe) {
		m := r.(*Target)
		m.TargetName = "darwin"
		m.DistPath = p
		m.BuildPath = p
		m.DeployPath = p
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Target{}, func(r rc_recipe.Recipe) {
		m := r.(*Target)
		m.TargetName = "darwin-arm"
		m.DistPath = p
		m.BuildPath = p
		m.DeployPath = p
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Target{}, func(r rc_recipe.Recipe) {
		m := r.(*Target)
		m.TargetName = "linux"
		m.DistPath = p
		m.BuildPath = p
		m.DeployPath = p
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Target{}, func(r rc_recipe.Recipe) {
		m := r.(*Target)
		m.TargetName = "linux-arm"
		m.DistPath = p
		m.BuildPath = p
		m.DeployPath = p
	})
	if err != nil {
		return err
	}

	return nil
}
