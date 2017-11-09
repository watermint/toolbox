package main

import (
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/tools/dtm/commands"
	"os"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Detach member(s) from the team

{{.Command}} detach -user user@example.com
{{.Command}} detach -csv user-list.csv


List member(s) of the team
{{.Command}} list
{{.Command}} list -csv members.csv
{{.Command}} list -status invited

External Id
{{.Command}} extid -list -user user@example.com
{{.Command}} extid -list -all-users
{{.Command}} extid -assign-pseudo-id -user user@example.com
{{.Command}} extid -assign-pseudo-id -all-users

Enforce shared link expiration

Expire shared links at +7 days if expiration not set
{{.Command}} link-expire -days 7

`

	data := struct {
		AppName    string
		AppVersion string
		AppHash    string
		Command    string
	}{
		AppName:    knowledge.AppName,
		AppVersion: knowledge.AppVersion,
		AppHash:    knowledge.AppHash,
		Command:    os.Args[0],
	}
	infra.ShowUsage(tmpl, data)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	commandArgs := os.Args[2:]

	switch os.Args[1] {
	case "detach":
		err := commands.Detach(commandArgs)
		if err != nil {
			seelog.Error(err)
		}

	case "list":
		err := commands.List(commandArgs)
		if err != nil {
			seelog.Error(err)
		}

	case "space":
		err := commands.Space(commandArgs)
		if err != nil {
			seelog.Error(err)
		}

	case "extid":
		err := commands.ExtId(commandArgs)
		if err != nil {
			seelog.Error(err)
		}

	case "link-expire":
		err := commands.LinkExpire(commandArgs)
		if err != nil {
			seelog.Error(err)
		}

	default:
		usage()
	}
}
