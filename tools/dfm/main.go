package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/compare"
	"github.com/watermint/toolbox/service/file"
	"github.com/watermint/toolbox/service/report"
	"github.com/watermint/toolbox/service/upload"
	"os"
	"time"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Usage:
  {{.Command}} COMMAND

Commands:
  move       Move files/folders to destination inside Dropbox
  restore    Restore files under path
  compare    Compare local files and Dropbox files
  upload     Upload local files into Dropbox


Command examples:

Move files/folders to destination
{{.Command}} move [OPTION]... SRC DEST

Restore files under path
{{.Command}} restore [OPTION]... PATH

Compare local files and Dropbox files
{{.Command}} compare                   LOCALPATH [DROPBOXPATH]
{{.Command}} compare    -xlsx XLSXPATH LOCALPATH [DROPBOXPATH]
{{.Command}} compare -csv-dir CSVDIR   LOCALPATH [DROPBOXPATH]

Upload files
{{.Command}} upload [OPTION]... SRC [SRC]... DEST
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
	f.IntVar(&mc.BatchSize, "batch-size", 1000, descBatchSize)

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

func parseMoveMockArgs(args []string) (mmc *file.MoveMockContext, err error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	mmc = &file.MoveMockContext{}
	mmc.Infra = infra.PrepareInfraFlags(f)

	f.SetOutput(os.Stderr)
	f.Parse(args)
	remainder := f.Args()
	if len(remainder) != 2 {
		f.PrintDefaults()
		return nil, errors.New("Missing [SQLITE3 DBFILE] or [DEST FOLDER]")
	}
	mmc.DbFile = remainder[0]
	mmc.DestPath = remainder[1]

	return
}

func parseUploadArgs(args []string) (*upload.UploadContext, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	uo := upload.UploadContext{}
	uo.Infra = infra.PrepareInfraFlags(f)

	descRecursive := "Recurse into directories"
	f.BoolVar(&uo.LocalRecursive, "recursive", true, descRecursive)
	f.BoolVar(&uo.LocalRecursive, "r", true, descRecursive)

	descSymlink := "Follow symlinks"
	f.BoolVar(&uo.LocalFollowSymlink, "follow-symlink", false, descSymlink)
	f.BoolVar(&uo.LocalFollowSymlink, "L", false, descSymlink)

	descConcurrency := "Upload concurrency"
	f.IntVar(&uo.Concurrency, "concurrency", 1, descConcurrency)
	f.IntVar(&uo.Concurrency, "c", 1, descConcurrency)

	descBandwidthLimit := "Limit upload bandwidth; KBytes per second (not kbps)"
	f.IntVar(&uo.BandwidthLimit, "bwlimit", 0, descBandwidthLimit)

	descDeleteAfterUpload := "Delete file after upload completed"
	f.BoolVar(&uo.DeleteAfterUpload, "migrate", false, descDeleteAfterUpload)

	f.SetOutput(os.Stderr)
	f.Parse(args)
	remainder := f.Args()
	splitPos := len(remainder) - 1
	if len(remainder) < 2 {
		usage()
		f.PrintDefaults()
		return nil, errors.New("Missing SRC and/or DEST")
	}
	if uo.Concurrency < 1 {
		uo.Concurrency = 1
	}
	if uo.BandwidthLimit < 0 {
		uo.BandwidthLimit = 0
	} else {
		uo.BandwidthLimit = uo.BandwidthLimit * 1024
	}

	uo.LocalPaths = remainder[:splitPos]
	uo.DropboxBasePath = remainder[splitPos]

	return &uo, nil
}

func parseCmpArgs(arg []string) (*compare.CompareOpts, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := compare.CompareOpts{}
	opts.InfraOpts = infra.PrepareInfraFlags(f)
	opts.ReportOpts = report.PrepareMultiReportFlags(f)

	f.SetOutput(os.Stderr)
	f.Parse(arg)
	remainder := f.Args()
	if len(remainder) < 1 {
		usage()
		f.PrintDefaults()
		return nil, errors.New("Missing LOCALPATH and/or DROPBOXPATH")
	}

	opts.LocalBasePath = remainder[0]
	if len(remainder) == 1 {
		opts.DropboxBasePath = ""
	} else {
		opts.DropboxBasePath = remainder[1]
	}

	return &opts, nil
}

func parseRestoreArgs(args []string) (rc *file.RestoreContext, err error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	rc = &file.RestoreContext{}
	rc.Infra = infra.PrepareInfraFlags(f)

	descPreflight := "Preflight mode (simulation mode)"
	f.BoolVar(&rc.Preflight, "preflight", false, descPreflight)

	descFilterTimeAfter := "Filter: time after (inclusive)"
	var timeAfter string
	f.StringVar(&timeAfter, "after", "", descFilterTimeAfter)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	if timeAfter != "" {
		rc.FilterTimeAfter, err = parseTimestampOpt(timeAfter)
		if err != nil {
			fmt.Errorf("unable to parse time for `-after`: %s", timeAfter)
			return nil, err
		}
	} else {
		rc.FilterTimeAfter = nil
	}

	remainder := f.Args()
	if len(remainder) < 1 {
		f.PrintDefaults()
		return nil, errors.New("missing [path]")
	}
	rc.BasePath = remainder[0]
	return
}

var (
	TIMESTAMP_OPT_ACCEPTABLE_FORMATS = []string{
		"2006-01-02",
		"2006/01/02",
		"2006-01-02T15:04:05",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05Z0700",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2006-01-02 15:04",
		"2006/01/02 15:04",
	}
)

func parseTimestampOpt(t string) (*time.Time, error) {
	for _, f := range TIMESTAMP_OPT_ACCEPTABLE_FORMATS {
		t, err := time.ParseInLocation(f, t, time.Local)
		u := t.UTC()
		if err == nil {
			return &u, nil
		}
	}
	return nil, errors.New("unable to parse date/time")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	commandArgs := os.Args[2:]

	switch os.Args[1] {
	case "move":
		mc, err := parseMoveArgs(commandArgs)
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

	case "move-mockup":
		// hidden command. undocumented in usage() or README; because it is only for testing purpose.
		//
		// {{.Command}} move-mockup [SQLITE3 DBFILE] [DEST FOLDER]
		// This command creates dummy files and folders by using existing preflight data.
		//
		// [SQLITE3 DBFILE]
		// Preflight data might have actual file/folder names. It depends on PreflightAnon option.
		// But this command uses only for fileId/folderId to create dummy file tree.
		//
		// [DEST FOLDER]
		// Dest folder is the location of destination path of local Dropbox folder.
		// If you want to simulate nested shared folder permission, please specify the team folder.

		mmc, err := parseMoveMockArgs(commandArgs)
		if err != nil {
			usage()
			return
		}
		defer mmc.Infra.Shutdown()
		err = mmc.Infra.Startup()
		if err != nil {
			seelog.Errorf("Unable to start operation: %s", err)
			return
		}
		seelog.Tracef("Options: %s", util.MarshalObjectToString(mmc))

		mmc.MockUp()

	case "restore":
		rc, err := parseRestoreArgs(commandArgs)
		if err != nil {
			usage()
			return
		}
		defer rc.Infra.Shutdown()
		err = rc.Infra.Startup()
		if err != nil {
			seelog.Errorf("Unable to start operation: %s", err)
			return
		}
		seelog.Tracef("Options: %s", util.MarshalObjectToString(rc))

		token, err := rc.Infra.LoadOrAuthDropboxFull()
		if err != nil || token == "" {
			seelog.Errorf("Unable to acquire token (error: %s)", err)
			return
		}
		rc.TokenFull = token
		rc.Restore()

	case "compare":
		opts, err := parseCmpArgs(commandArgs)
		if err != nil {
			usage()
			fmt.Fprintln(os.Stderr, "Error: ", err)
			return
		}
		err = opts.ReportOpts.ValidateMultiReportOpts()
		if err != nil {
			usage()
			fmt.Fprintln(os.Stderr, "Error: ", err)
			return
		}
		defer opts.InfraOpts.Shutdown()
		err = opts.InfraOpts.Startup()
		if err != nil {
			seelog.Errorf("Unable to start operation: %s", err)
			return
		}
		seelog.Tracef("Options: %s", util.MarshalObjectToString(opts))

		token, err := opts.InfraOpts.LoadOrAuthDropboxFull()
		if err != nil || token == "" {
			seelog.Errorf("Unable to acquire token (error: %s)", err)
			return
		}
		opts.DropboxToken = token

		compare.Compare(opts)

	case "upload":
		opts, err := parseUploadArgs(commandArgs)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: ", err)
			return
		}

		defer opts.Infra.Shutdown()
		err = opts.Infra.Startup()
		if err != nil {
			seelog.Errorf("Unable to start operation: %s", err)
			return
		}
		seelog.Tracef("Options: %s", util.MarshalObjectToString(opts))

		token, err := opts.Infra.LoadOrAuthDropboxFull()
		if err != nil || token == "" {
			seelog.Errorf("Unable to acquire token (error: %s)", err)
			return
		}
		opts.DropboxToken = token
		opts.Upload()

	default:
		usage()
	}
}
