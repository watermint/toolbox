package cmd_device

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_device"
	"github.com/watermint/toolbox/model/dbx_device/flat_device"
)

type CmdTeamDeviceList struct {
	*cmd.SimpleCommandlet
	report        app_report.Factory
	optDeviceType string
}

func (CmdTeamDeviceList) Name() string {
	return "list"
}

func (CmdTeamDeviceList) Desc() string {
	return "cmd.team.device.list.desc"
}

func (CmdTeamDeviceList) Usage() func(usage cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamDeviceList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descDeviceType := z.ExecContext.Msg("cmd.team.device.list.flag.device").T()
	f.StringVar(&z.optDeviceType, "device", "", descDeviceType)
}

func (z *CmdTeamDeviceList) Exec(args []string) {
	l := dbx_device.DeviceList{
		OnError: z.DefaultErrorHandler,
		OnDesktop: func(desktop *flat_device.DesktopClient) bool {
			z.report.Report(desktop)
			return true
		},
		OnMobile: func(mobile *flat_device.MobileClient) bool {
			z.report.Report(mobile)
			return true
		},
		OnWeb: func(web *flat_device.WebSession) bool {
			z.report.Report(web)
			return true
		},
	}

	switch z.optDeviceType {
	case "all":
		l.IncludeWebSessions = true
		l.IncludeMobileClients = true
		l.IncludeDesktopClients = true

	case "mobile":
		l.IncludeWebSessions = false
		l.IncludeMobileClients = true
		l.IncludeDesktopClients = false

	case "web":
		l.IncludeWebSessions = true
		l.IncludeMobileClients = false
		l.IncludeDesktopClients = false

	case "desktop":
		l.IncludeWebSessions = false
		l.IncludeMobileClients = false
		l.IncludeDesktopClients = true

	default:
		z.ExecContext.Msg("cmd.team.device.list.err.device_type").TellError()
		return
	}

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	l.List(apiFile)
}
