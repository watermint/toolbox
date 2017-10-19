package tree

import (
	"github.com/watermint/toolbox/infra"
	"flag"
	"os"
	"errors"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra/util"
)

type MockupOpts struct {
	Infra       *infra.InfraOpts
	ReportPath  string
	IncludePath string
	ExcludePath string
}

func ParseMockupOptions(args []string) (*MockupOpts, error) {
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	opts := &MockupOpts{}
	opts.Infra = infra.PrepareInfraFlags(f)

	descReportPath := "Path for the report"
	f.StringVar(&opts.ReportPath, "in", "", descReportPath)

	f.SetOutput(os.Stderr)
	f.Parse(args)

	return opts, nil
}

func ExecMockup(args []string) error {
	opts, err := ParseMockupOptions(args)
	if err != nil {
		return err
	}
	if opts.ReportPath == "" {
		return errors.New("please specify report file path")
	}

	defer opts.Infra.Shutdown()
	err = opts.Infra.Startup()
	if err != nil {
		seelog.Errorf("Unable to start operation: %s", err)
		return err
	}
	seelog.Tracef("options: %s", util.MarshalObjectToString(opts))

}
