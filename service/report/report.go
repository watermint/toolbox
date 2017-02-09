package report

import (
	"flag"
	"os"
	"sync"
)

type ReportRow interface {
}

type ReportHeader struct {
	Headers []string
}

type ReportData struct {
	Data []string
}

type ReportEOF struct {
}

type ReportOpts struct {
	OutputCsv  string
	OutputXlsx string
	wg         sync.WaitGroup
}

func PrepareReportFlags(f *flag.FlagSet) *ReportOpts {
	opts := ReportOpts{}

	descCsv := "Output CSV path"
	f.StringVar(&opts.OutputCsv, "csv", "", descCsv)

	descXlsx := "Output .xlsx path"
	f.StringVar(&opts.OutputXlsx, "xlsx", "", descXlsx)

	return &opts
}

func (r *ReportOpts) Write(rows chan ReportRow) {
	if r.OutputXlsx != "" {
		go WriteXlsx(r.OutputXlsx, "Members", rows, &r.wg)
	} else if r.OutputCsv == "" {
		go WriteCsv(os.Stdout, rows, &r.wg)
	} else {
		go WriteCsvFile(r.OutputCsv, rows, &r.wg)
	}
}

func (r *ReportOpts) Wait() {
	r.wg.Wait()
}
