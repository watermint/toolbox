package main

import (
	"errors"
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/service/dupload"
	"os"
	"github.com/watermint/toolbox/integration/auth"
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
	proxy := flag.String("proxy", "", "HTTP(S) proxy (hostname:port)")
	localPath := flag.String("localPath", "", "Path to upload (required)")
	localRecursive := flag.Bool("recursive", true, "Upload child directories")
	localFollowSymlink := flag.Bool("followSymlink", false, "Follow symlink")
	dropboxPath := flag.String("dropboxPath", "", "Base path in Dropbox (required)")

	seelog.Flush()
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Parse()

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

	infra.SetupHttpProxy(opts.Proxy)

	a := auth.DropboxAuthenticator{
		AppKey:    AppKey,
		AppSecret: AppSecret,
	}

	token, err := a.Authorise()
	if err != nil || token == "" {
		seelog.Error("Unable to acquire token")
		return
	}

	uc := &dupload.UploadContext{
		LocalPath:          opts.LocalPath,
		LocalRecursive:     opts.LocalRecursive,
		LocalFollowSymlink: opts.LocalFollowSymlink,
		DropboxBasePath:    opts.DropboxBasePath,
		DropboxToken:       token,
	}

	dupload.Upload(uc)
}
