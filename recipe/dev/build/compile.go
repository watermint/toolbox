package build

import (
	"github.com/watermint/toolbox/essentials/text/es_escape"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"strings"
)

type Compile struct {
	rc_recipe.RemarkSecret
	Os      string
	Arch    string
	Path    string
	Package string
	Out     string
}

func (z *Compile) Preset() {
}

func (z *Compile) infoLine() string {
	return "OS=" + es_escape.ReplaceNonAlNum(z.Os, "_") + ", Arch=" + es_escape.ReplaceNonAlNum(z.Arch, "_")
}

func (z *Compile) outLines(lines []string) error {
	info := z.infoLine()
	head := "echo Compiling: " + info + "\n\n"
	foot := "echo Completed: " + info + "\n\n"
	content := head + strings.Join(lines, "\n") + "\n\n" + foot
	return os.WriteFile(z.Out, []byte(content), 0755)
}

func (z *Compile) goBuildCgoEnabled(env map[string]string) string {
	envAll := ""
	for k, v := range env {
		envAll = envAll + k + "=" + v + " "
	}
	return strings.Join([]string{
		envAll,
		"CGO_ENABLED=1",
		"go build",
		"-o " + z.Path,
	}, " ")
}

func (z *Compile) noSupport() string {
	info := z.infoLine()
	return "echo The platform " + info + " is currently not supported"
}

func (z *Compile) compileWindows() []string {
	switch z.Arch {
	case "amd64":
		return []string{
			z.goBuildCgoEnabled(map[string]string{
				"CC":     "x86_64-w64-mingw32-gcc-posix",
				"CXX":    "x86_64-w64-mingw32-g++-posix",
				"GOOS":   "windows",
				"GOARCH": "amd64",
			}),
		}
	default:
		return []string{z.noSupport()}
	}
}

func (z *Compile) compileLinux() []string {
	switch z.Arch {
	case "amd64":
		return []string{
			z.goBuildCgoEnabled(map[string]string{
				"GOOS":   "linux",
				"GOARCH": "amd64",
			}),
		}

	case "arm64":
		return []string{
			z.goBuildCgoEnabled(map[string]string{
				//"CC":     "aarch64-linux-gnu-gcc-6",
				//"CXX":    "aarch64-linux-gnu-g++-6",
				"GOOS":   "linux",
				"GOARCH": "arm64",
			}),
		}

	default:
		return []string{z.noSupport()}
	}
}

func (z *Compile) compileDarwin() []string {
	switch z.Arch {
	case "amd64":
		//  CC=o64-clang CXX=o64-clang++ GOOS=darwin GOARCH=amd64 CGO_ENABLED=1 go build $V $X $TP $MOD "${T[@]}" --ldflags="$LDSTRIP $V $LD" --gcflags="$GC" $R $BM -o "/build/$NAME-darwin-$PLATFORM-amd64$R`extension darwin`" $PACK_RELPATH
		return []string{
			"export MACOSX_DEPLOYMENT_TARGET=10.13",
			z.goBuildCgoEnabled(map[string]string{
				"CC":     "o64-clang",
				"CXX":    "o64-clang++",
				"GOOS":   "darwin",
				"GOARCH": "amd64",
			}),
			"unset MACOSX_DEPLOYMENT_TARGET",
		}

	case "arm64":
		return []string{
			"export MACOSX_DEPLOYMENT_TARGET=10.16", // M1 support from 10.16
			z.goBuildCgoEnabled(map[string]string{
				"CC":     "o64-clang",
				"CXX":    "o64-clang++",
				"GOOS":   "darwin",
				"GOARCH": "amd64",
			}),
			"unset MACOSX_DEPLOYMENT_TARGET",
		}

	default:
		return []string{z.noSupport()}
	}
}

func (z *Compile) Exec(c app_control.Control) error {
	switch z.Os {
	case "windows":
		return z.outLines(z.compileWindows())
	case "linux":
		return z.outLines(z.compileLinux())
	case "darwin":
		return z.outLines(z.compileDarwin())
	}
	return nil
}

func (z *Compile) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("compile", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()
	err = rc_exec.Exec(c, &Compile{}, func(r rc_recipe.Recipe) {
		m := r.(*Compile)
		m.Os = "windows"
		m.Arch = "amd64"
		m.Path = "/tmp/windows-amd64/tbx"
		m.Package = "github.com/watermint/toolbox"
		m.Out = p + "/build-windows-amd64.sh"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Compile{}, func(r rc_recipe.Recipe) {
		m := r.(*Compile)
		m.Os = "windows"
		m.Arch = "arm64"
		m.Path = "/tmp/windows-arm64/tbx"
		m.Package = "github.com/watermint/toolbox"
		m.Out = p + "/build-windows-arm64.sh"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Compile{}, func(r rc_recipe.Recipe) {
		m := r.(*Compile)
		m.Os = "darwin"
		m.Arch = "amd64"
		m.Path = "/tmp/darwin-amd64/tbx"
		m.Package = "github.com/watermint/toolbox"
		m.Out = p + "/build-darwin-amd64.sh"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Compile{}, func(r rc_recipe.Recipe) {
		m := r.(*Compile)
		m.Os = "darwin"
		m.Arch = "arm64"
		m.Path = "/tmp/darwin-arm64/tbx"
		m.Package = "github.com/watermint/toolbox"
		m.Out = p + "/build-darwin-arm64.sh"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Compile{}, func(r rc_recipe.Recipe) {
		m := r.(*Compile)
		m.Os = "linux"
		m.Arch = "amd64"
		m.Path = "/tmp/linux-amd64/tbx"
		m.Package = "github.com/watermint/toolbox"
		m.Out = p + "/build-linux-amd64.sh"
	})
	if err != nil {
		return err
	}
	err = rc_exec.Exec(c, &Compile{}, func(r rc_recipe.Recipe) {
		m := r.(*Compile)
		m.Os = "linux"
		m.Arch = "arm64"
		m.Path = "/tmp/linux-arm64/tbx"
		m.Package = "github.com/watermint/toolbox"
		m.Out = p + "/build-linux-arm64.sh"
	})
	if err != nil {
		return err
	}

	return nil
}
