package report

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/cihub/seelog"
	"io"
	"os"
	"sync"
)

func WriteCsvRow(f io.Writer, row ReportRow) error {
	w := csv.NewWriter(f)
	defer w.Flush()

	switch r := row.(type) {
	case nil:
		return nil

	case *ReportHeader:
		return w.Write(r.Headers)

	case *ReportData:
		rowStr := make([]string, 0)
		for _, a := range r.Data {
			rowStr = append(rowStr, fmt.Sprintf("%v", a))
		}
		return w.Write(rowStr)

	default:
		seelog.Warnf("Unexpected row")
		return errors.New("Unexpected row detected")
	}
	return nil
}

func WriteCsv(f io.Writer, report chan ReportRow, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	for r := range report {
		if r == nil {
			break
		}
		err := WriteCsvRow(f, r)
		if err != nil {
			seelog.Warnf("Unable to write row : error[%s]", err)
			return err
		}
	}
	return nil
}

func WriteCsvFile(file string, report chan ReportRow, wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	f, err := os.Create(file)
	if err != nil {
		seelog.Errorf("Unable to write file[%s] erorr[%s]", file, err)
		return err
	}
	defer f.Close()
	return WriteCsv(f, report, wg)
}
