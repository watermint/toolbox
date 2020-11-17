package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgTroubleshooting struct {
	Title app_msg.Message

	NetworkTitle      app_msg.Message
	NetworkOverview   app_msg.Message
	NetworkProxyNotes app_msg.Message

	GarbledTitle      app_msg.Message
	GarbledOverview   app_msg.Message
	GarbledPowershell app_msg.Message

	SlowTitle     app_msg.Message
	SlowOverview  app_msg.Message
	SlowRateLimit app_msg.Message

	LogTitle                    app_msg.Message
	LogOverview                 app_msg.Message
	LogSecurityDisclaimer       app_msg.Message
	LogFormat                   app_msg.Message
	LogFormatJobFolder          app_msg.Message
	LogFormatJobFolderStructure app_msg.Message
	LogFormatLogsFolder         app_msg.Message
	LogFormatLogFormat          app_msg.Message
	LogDebug                    app_msg.Message
	LogDebugOverview            app_msg.Message
	LogDebugExample             app_msg.Message
	LogCapture                  app_msg.Message
	LogCaptureOverview          app_msg.Message
}

var (
	MTroubleshooting = app_msg.Apply(&MsgTroubleshooting{}).(*MsgTroubleshooting)
)

type Troubleshooting struct {
}

func (z Troubleshooting) DocId() dc_index.DocId {
	return dc_index.DocSupplementalTroubleshooting
}

func (z Troubleshooting) Sections() []dc_section.Section {
	return []dc_section.Section{
		&TroubleshootingNetwork{},
		&TroubleshootingSlow{},
		&TroubleshootingGarbled{},
		&TroubleshootingLog{},
	}
}

type TroubleshootingNetwork struct {
}

func (z TroubleshootingNetwork) Title() app_msg.Message {
	return MTroubleshooting.NetworkTitle
}

func (z TroubleshootingNetwork) Body(ui app_ui.UI) {
	ui.Info(MTroubleshooting.NetworkOverview)
	ui.Break()
	ui.Info(MTroubleshooting.NetworkProxyNotes)
}

type TroubleshootingGarbled struct {
}

func (z TroubleshootingGarbled) Title() app_msg.Message {
	return MTroubleshooting.GarbledTitle
}

func (z TroubleshootingGarbled) Body(ui app_ui.UI) {
	ui.Info(MTroubleshooting.GarbledOverview)
	ui.Break()
	ui.Info(MTroubleshooting.GarbledPowershell)
}

type TroubleshootingSlow struct {
}

func (z TroubleshootingSlow) Title() app_msg.Message {
	return MTroubleshooting.SlowTitle
}

func (z TroubleshootingSlow) Body(ui app_ui.UI) {
	ui.Info(MTroubleshooting.SlowOverview)
	ui.Break()
	ui.Info(MTroubleshooting.SlowRateLimit)
	ui.Code(`tbx job log last -quiet | jq 'select(.msg == "WaiterStatus")' 
{
  "level": "DEBUG",
  "time": "2020-11-10T14:55:57.501+0900",
  "name": "z951.z960.z112064",
  "caller": "nw_congestion/congestion.go:310",
  "msg": "WaiterStatus",
  "goroutine": "gr:284877",
  "runners": {
    "gr:1": {
      "key": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create",
      "go_routine_id": "gr:1",
      "running_since": "2020-11-10T14:55:56.124899+09:00"
    }
  },
  "numRunners": 1,
  "waiters": [],
  "numWaiters": 0,
  "window": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url/check_job_status": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search/continue_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/users/get_current_account": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_reference/get": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/delete_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/get_metadata": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/sharing/list_mountable_folders": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/download": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/export": 4
  },
  "concurrency": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 1
  }
}`)
}

type TroubleshootingLog struct {
}

func (z TroubleshootingLog) Title() app_msg.Message {
	return MTroubleshooting.LogTitle
}

func (z TroubleshootingLog) Body(ui app_ui.UI) {
	ui.Info(MTroubleshooting.LogOverview)
	ui.Break()
	ui.Info(MTroubleshooting.LogSecurityDisclaimer)
	ui.Break()

	ui.SubHeader(MTroubleshooting.LogFormat)
	ui.Info(MTroubleshooting.LogFormatJobFolder)
	ui.Break()
	ui.Info(MTroubleshooting.LogFormatJobFolderStructure)
	ui.Break()
	ui.Info(MTroubleshooting.LogFormatLogsFolder)
	ui.Break()
	ui.Info(MTroubleshooting.LogFormatLogFormat)
	ui.Break()

	ui.SubHeader(MTroubleshooting.LogDebug)
	ui.Info(MTroubleshooting.LogDebugOverview)
	ui.Break()
	ui.Info(MTroubleshooting.LogDebugExample)
	ui.Code(`tbx job log last -quiet | jq -r 'select(.msg == "Heap stats") | [.time, .HeapInuse] | @csv'
"2020-11-10T14:55:45.725+0900",18604032
"2020-11-10T14:55:50.725+0900",15130624
"2020-11-10T14:55:55.725+0900",17408000
"2020-11-10T14:56:00.725+0900",17014784
"2020-11-10T14:56:05.726+0900",19193856
"2020-11-10T14:56:10.725+0900",19136512
"2020-11-10T14:56:15.726+0900",16637952
"2020-11-10T14:56:20.725+0900",16678912
"2020-11-10T14:56:25.727+0900",16678912
"2020-11-10T14:56:30.730+0900",16678912
"2020-11-10T14:56:35.726+0900",16678912
`)

	ui.SubHeader(MTroubleshooting.LogCapture)
	ui.Info(MTroubleshooting.LogCaptureOverview)
}
