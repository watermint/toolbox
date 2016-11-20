package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/knowledge"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/integration/auth"
	"github.com/watermint/toolbox/service/dupload"
	"os"
)

var (
	AppKey    string = ""
	AppSecret string = ""
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
	Proxy              string
	LocalPaths         []string
	LocalRecursive     bool
	LocalFollowSymlink bool
	DropboxBasePath    string
	CleanupToken       bool
	WorkPath           string
	Concurrency        int
	//BandwidthLimit     int
}

func parseArgs() (*UploadOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	var proxy, workPath string
	var localRecursive, localFollowSymlink, cleanupToken bool
	var concurrency int

	descProxy := "HTTP/HTTPS proxy (hostname:port)"
	f.StringVar(&proxy, "proxy", "", descProxy)

	descRecursive := "Recurse into directories"
	f.BoolVar(&localRecursive, "recursive", true, descRecursive)
	f.BoolVar(&localRecursive, "r", true, descRecursive)

	descSymlink := "Follow symlinks"
	f.BoolVar(&localFollowSymlink, "follow-symlink", false, descSymlink)
	f.BoolVar(&localFollowSymlink, "L", false, descSymlink)

	descCleanup := "Revoke token on exit"
	f.BoolVar(&cleanupToken, "revoke-token", false, descCleanup)

	descWork := "Work directory (default: $HOME/.dupload)"
	f.StringVar(&workPath, "work", "", descWork)

	descConcurrency := "Upload concurrency"
	f.IntVar(&concurrency, "concurrency", 1, descConcurrency)
	f.IntVar(&concurrency, "c", 1, descConcurrency)

	//descBandwidthLimit := "Limit upload bandwidth; KBytes per second (not kbps)"
	//f.IntVar(&bandwidthLimit, "bwlimit", 0, descBandwidthLimit)

	f.SetOutput(os.Stderr)
	f.Parse(os.Args[1:])
	args := f.Args()
	splitPos := len(args) - 1
	if len(args) < 2 {
		usage()
		f.PrintDefaults()
		return nil, errors.New("Missing SRC and/or DEST")
	}
	if concurrency < 1 {
		concurrency = 1
	}

	return &UploadOptions{
		Proxy:              proxy,
		LocalPaths:         args[:splitPos],
		LocalRecursive:     localRecursive,
		LocalFollowSymlink: localFollowSymlink,
		DropboxBasePath:    args[splitPos],
		CleanupToken:       cleanupToken,
		WorkPath:           workPath,
		Concurrency:        concurrency,
		//BandwidthLimit:     bandwidthLimit,
	}, nil
}

func main() {
	opts, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}

	infraOpts := infra.InfraOpts{}
	infra.InfraStartup(infraOpts)
	defer infra.InfraShutdown()

	seelog.Infof("Upload options: %s", util.MarshalObjectToString(opts))

	infra.SetupHttpProxy(opts.Proxy)

	a := auth.DropboxAuthenticator{
		AppKey:    AppKey,
		AppSecret: AppSecret,
	}

	token, err := a.Authorise()
	if err != nil || token == "" {
		seelog.Errorf("Unable to acquire token (error: %s)", err)
		return
	}
	defer auth.RevokeToken(token)

	uc := &dupload.UploadContext{
		LocalRecursive:     opts.LocalRecursive,
		LocalFollowSymlink: opts.LocalFollowSymlink,
		DropboxBasePath:    opts.DropboxBasePath,
		DropboxToken:       token,
	}

	dupload.Upload(opts.LocalPaths, uc, opts.Concurrency)
}
