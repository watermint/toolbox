package cmd_license

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/watermint/toolbox/legacy/app/app_report"
	"github.com/watermint/toolbox/legacy/app/app_ui"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
	"os"
)

type CmdLicense struct {
	*cmd2.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdLicense) Name() string {
	return "license"
}

func (z *CmdLicense) Desc() string {
	return "cmd.license.desc"
}

func (z *CmdLicense) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdLicense) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdLicense) Exec(args []string) {
	lic, err := z.ExecContext.ResourceBytes("licenses.json")
	if err != nil {
		z.ExecContext.Msg("cmd.license.err.no_resources").TellError()
		return
	}
	licenses := make(map[string][]string)
	if err = json.Unmarshal(lic, &licenses); err != nil {
		z.Log().Error("Invalid License file format", zap.Error(err))
		z.ExecContext.Msg("cmd.license.err.no_resources").TellError()
		return
	}

	toolboxPkg := "github.com/watermint/toolbox"
	if toolboxLic, e := licenses[toolboxPkg]; !e {
		z.Log().Error("`toolbox` license not found")
		z.ExecContext.Msg("cmd.license.err.no_resources").TellError()
		return
	} else {
		z.showLicense(toolboxPkg, toolboxLic)
	}

	for pkg, lines := range licenses {
		if pkg == toolboxPkg {
			continue
		}
		z.showLicense(pkg, lines)
	}
}

func (z *CmdLicense) showLicense(pkg string, lines []string) {
	fmt.Println()
	app_ui.ColorPrint(os.Stdout, pkg, app_ui.ColorGreen)
	fmt.Println()
	fmt.Println()
	for _, line := range lines {
		fmt.Print("  ")
		app_ui.ColorPrint(os.Stdout, line, app_ui.ColorWhite)
		fmt.Println()
	}
}
