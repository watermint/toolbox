package main

import (
	"errors"
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/integration/auth"
	"github.com/watermint/toolbox/service/dupload"
	"os"
)

var (
	AppKey    string = ""
	AppSecret string = ""
)

type UploadOptions struct {
	Proxy              string
	LocalPath          string
	LocalRecursive     bool
	LocalFollowSymlink bool
	DropboxBasePath    string
}

func parseArgs() (*UploadOptions, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	proxy := f.String("proxy", "", "HTTP(S) proxy (hostname:port)")
	localPath := f.String("localPath", "", "Path to upload (required)")
	localRecursive := f.Bool("recursive", true, "Upload child directories")
	localFollowSymlink := f.Bool("followSymlink", false, "Follow symlink")
	dropboxPath := f.String("dropboxPath", "", "Base path in Dropbox (required)")

	seelog.Flush()
	f.SetOutput(os.Stdout)
	f.Parse(os.Args[1:])
	for 0 < f.NArg() {
		f.Parse(f.Args()[1:])
	}

	if *localPath == "" || *dropboxPath == "" {
		seelog.Error("Missing required option: `-localPath` and/or `-dropboxPath`")

		flag.Usage()

		return nil, errors.New("Missing required option")
	}

	return &UploadOptions{
		Proxy:              *proxy,
		LocalPath:          *localPath,
		LocalRecursive:     *localRecursive,
		LocalFollowSymlink: *localFollowSymlink,
		DropboxBasePath:    *dropboxPath,
	}, nil
}

func main() {
	infra.InfraStartup()
	defer infra.InfraShutdown()

	opts, err := parseArgs()
	if err != nil {
		return
	}
	seelog.Infof("Upload from [%s](Local) to [%s](Dropbox)", opts.LocalPath, opts.DropboxBasePath)

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
		LocalPath:          opts.LocalPath,
		LocalRecursive:     opts.LocalRecursive,
		LocalFollowSymlink: opts.LocalFollowSymlink,
		DropboxBasePath:    opts.DropboxBasePath,
		DropboxToken:       token,
	}

	dupload.Upload(uc)
}
