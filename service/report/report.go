package report

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/tealeg/xlsx"
	"os"
	"path/filepath"
	"sync"
)

type ReportRow interface {
}

type ReportHeader struct {
	Headers []string
}

type ReportData struct {
	Data []interface{}
}

type SimpleReportOpts struct {
	OutputCsv  string
	OutputXlsx string
	wg         sync.WaitGroup
}

func PrepareSimpleReportFlags(f *flag.FlagSet) *SimpleReportOpts {
	opts := SimpleReportOpts{}

	descCsv := "Output CSV path"
	f.StringVar(&opts.OutputCsv, "csv", "", descCsv)

	descXlsx := "Output .xlsx path"
	f.StringVar(&opts.OutputXlsx, "xlsx", "", descXlsx)

	return &opts
}

func (r *SimpleReportOpts) Write(rows chan ReportRow) {
	if r.OutputXlsx != "" {
		go WriteXlsx(r.OutputXlsx, "Report", rows, &r.wg)
	} else if r.OutputCsv == "" {
		go WriteCsv(os.Stdout, rows, &r.wg)
	} else {
		go WriteCsvFile(r.OutputCsv, rows, &r.wg)
	}
}

func (r *SimpleReportOpts) Wait() {
	r.wg.Wait()
}

type MultiReportOpts struct {
	OutputCSVDir    string
	OutputExcelFile string
	wg              sync.WaitGroup

	xlsxFile *xlsx.File
}

func PrepareMultiReportFlags(f *flag.FlagSet) *MultiReportOpts {
	opts := MultiReportOpts{}

	descCsv := "Output CSV directory"
	f.StringVar(&opts.OutputCSVDir, "csv-dir", "", descCsv)

	descXlsx := "Output .xlsx file path"
	f.StringVar(&opts.OutputExcelFile, "xlsx", "", descXlsx)

	return &opts
}

func (r *MultiReportOpts) BeginMultiReport() error {
	err := r.beginMultiReportXlsx()
	if err != nil {
		seelog.Warnf("Unable to begin report session for xlsx : error[%s]", err)
		return err
	}
	err = r.beginMultiReportCsv()
	if err != nil {
		seelog.Warnf("Unable to begin report session for csv : error[%s]", err)
		return err
	}
	return nil
}

func (r *MultiReportOpts) beginMultiReportXlsx() error {
	if r.OutputExcelFile == "" {
		return nil
	}
	r.xlsxFile = xlsx.NewFile()
	return nil
}

func (r *MultiReportOpts) beginMultiReportCsv() error {
	if r.OutputCSVDir == "" {
		return nil
	}

	info, err := os.Lstat(r.OutputCSVDir)
	if err != nil && os.IsNotExist(err) {
		seelog.Warnf("Unable to acquire path info : path[%s] error[%s]", r.OutputCSVDir, err)
		return err
	}
	if os.IsNotExist(err) {
		err = os.MkdirAll(r.OutputCSVDir, 0755)
		if err != nil {
			seelog.Warnf("Unable to create folder : path[%s] error[%s]", r.OutputCSVDir, err)
			return err
		}
	}
	if info.IsDir() {
		return nil
	} else {
		return errors.New(fmt.Sprintf("Path[%s] is not a folder", r.OutputCSVDir))
	}

	return nil
}

func (r *MultiReportOpts) Write(name string, rows chan ReportRow) error {
	err := r.writeXlsx(name, rows)
	if err != nil {
		seelog.Warnf("Unable to write report : name[%s] error[%s]", name, err)
		return err
	}
	err = r.writeCsv(name, rows)
	if err != nil {
		seelog.Warnf("Unable to write report : name[%s] error[%s]", name, err)
		return err
	}

	return nil
}

func (r *MultiReportOpts) writeXlsx(name string, rows chan ReportRow) error {
	r.wg.Add(1)
	defer r.wg.Done()

	sheet, err := r.xlsxFile.AddSheet(name)
	if err != nil {
		seelog.Warnf("Unable to add sheet : name[%s] error[%s]", name, err)
		return err
	}

	for r := range rows {
		if r == nil {
			break
		}
		err = WriteXlsxRow(sheet, r)
		if err != nil {
			seelog.Warnf("Unable to write row : error[%s]", err)
			return err
		}
	}
	return nil
}

func (r *MultiReportOpts) writeCsv(name string, rows chan ReportRow) error {
	csvPath := filepath.Join(r.OutputCSVDir, fmt.Sprintf("%s.csv", name))

	return WriteCsvFile(csvPath, rows, &r.wg)
}

func (r *MultiReportOpts) FlushSingleReport() error {
	r.wg.Wait()

	return nil
}

func (r *MultiReportOpts) EndMultiReport() error {
	if r.OutputExcelFile != "" && r.xlsxFile != nil {
		err := r.xlsxFile.Save(r.OutputExcelFile)
		if err != nil {
			seelog.Warnf("Unable to save xlsx file : path[%s] error[%s]", r.OutputExcelFile, err)
			return err
		}
	}
	return nil
}
