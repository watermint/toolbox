package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/service/upload"
	"os"
)

func usage() {
	tmpl := `{{.AppName}} {{.AppVersion}} ({{.AppHash}}):

Usage:
{{.Command}} [OPTION]... SRC [SRC]... DEST
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

type UploadOptions struct {
	Infra              *infra.InfraOpts
	LocalPaths         []string
	LocalRecursive     bool
	LocalFollowSymlink bool
	DropboxBasePath    string
	Concurrency        int
	BandwidthLimit     int
}

func parseArgs() (*UploadOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	uo := UploadOptions{}
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

	f.SetOutput(os.Stderr)
	f.Parse(os.Args[1:])
	args := f.Args()
	splitPos := len(args) - 1
	if len(args) < 2 {
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

	uo.LocalPaths = args[:splitPos]
	uo.DropboxBasePath = args[splitPos]

	return &uo, nil
}

func main() {
	opts, err := parseArgs()
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

	seelog.Tracef("Upload options: %s", util.MarshalObjectToString(opts))

	token, err := opts.Infra.LoadOrAuthDropboxFull()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return
	}

	uc := &upload.UploadContext{
		LocalRecursive:     opts.LocalRecursive,
		LocalFollowSymlink: opts.LocalFollowSymlink,
		DropboxBasePath:    opts.DropboxBasePath,
		DropboxToken:       token,
		BandwidthLimit:     opts.BandwidthLimit,
	}

	upload.Upload(opts.LocalPaths, uc, opts.Concurrency)
}
