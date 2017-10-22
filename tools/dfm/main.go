package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/file"
	"os"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Move files/folders to destination
{{.Command}} move [OPTION]... SRC DEST
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

func parseMoveArgs(args []string) (mc *file.MoveContext, err error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mc = &file.MoveContext{}
	mc.Infra = infra.PrepareInfraFlags(f)

	descBatchSize := fmt.Sprintf("Move operation batch size (1 < batch_size < %d)", file.MOVE_BATCH_MAX_SIZE)
	f.IntVar(&mc.BatchSize, "batch-size", 100, descBatchSize)

	descPreflight := "Preflight mode (simulation mode)"
	f.BoolVar(&mc.Preflight, "preflight", false, descPreflight)

	descPreflightAnon := "Anonimise file names and folder names on preflight"
	f.BoolVar(&mc.PreflightAnon, "preflight-anon", true, descPreflightAnon)

	descFileByFile := "File by file operation mode"
	f.BoolVar(&mc.FileByFile, "file-by-file", false, descFileByFile)

	f.SetOutput(os.Stderr)
	f.Parse(args)
	remainder := f.Args()
	if len(remainder) != 2 {
		f.PrintDefaults()
		return nil, errors.New("Missing SRC and/or DEST")
	}

	mc.SrcPath = remainder[0]
	mc.DestPath = remainder[1]

	return
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "move":
		mc, err := parseMoveArgs(os.Args[2:])
		if err != nil {
			usage()
			return
		}
		defer mc.Infra.Shutdown()
		err = mc.Infra.Startup()
		if err != nil {
			seelog.Errorf("Unable to start operation: %s", err)
			return
		}
		seelog.Tracef("Options: %s", util.MarshalObjectToString(mc))

		token, err := mc.Infra.LoadOrAuthDropboxFull()
		if err != nil || token == "" {
			seelog.Errorf("Unable to acquire token (error: %s)", err)
			return
		}
		mc.TokenFull = token
		mc.Move()

	default:
		usage()
	}
}
