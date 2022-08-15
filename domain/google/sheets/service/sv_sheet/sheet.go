package sv_sheet

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_sheet"
	"github.com/watermint/toolbox/domain/google/sheets/model/to_cell"
	"github.com/watermint/toolbox/domain/google/sheets/model/to_spreadsheet"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
	"net/url"
)

type Sheet interface {
	Clear(spreadsheetId, sheetRange string) (clearedRange string, err error)
	Export(spreadsheetId, sheetRange string, opts ...RenderOpt) (value to_cell.ValueRange, err error)
	Import(spreadsheetId, sheetRange string, values [][]interface{}, rawInput bool) (updated bo_sheet.ValueUpdate, err error)
	Append(spreadsheetId, sheetRange string, values [][]interface{}, rawInput bool) (appended bo_sheet.ValueAppend, err error)

	Create(spreadsheetId, sheetTitle string, cols, rows int) (sheet bo_sheet.Sheet, err error)
	Delete(spreadsheetId, sheetId string) error
}

var (
	ValueRenderOptionFormatted   = "FORMATTED_VALUE"
	ValueRenderOptionUnformatted = "UNFORMATTED_VALUE"
	ValueRenderOptionFormula     = "FORMULA"
	ValueRenderOptionsFull       = []string{
		ValueRenderOptionFormatted,
		ValueRenderOptionUnformatted,
		ValueRenderOptionFormula,
	}

	ValueRenderOptionAliasFormatted   = "formatted"
	ValueRenderOptionAliasUnformatted = "unformatted"
	ValueRenderOptionAliasFormula     = "formula"
	ValueRenderOptionAliases          = []string{
		ValueRenderOptionAliasFormatted,
		ValueRenderOptionAliasFormatted,
		ValueRenderOptionAliasFormula,
	}

	DateTimeRenderOptionSerialNumber = "SERIAL_NUMBER"
	DateTimeRenderOptionFormatted    = "FORMATTED_STRING"
	DateTimeRenderOptionsFull        = []string{
		DateTimeRenderOptionSerialNumber,
		DateTimeRenderOptionFormatted,
	}

	DateTimeRenderOptionAliasSerialNumber = "serial"
	DateTimeRenderOptionAliasFormatted    = "formatted"
	DateTimeRenderOptionAliases           = []string{
		DateTimeRenderOptionAliasSerialNumber,
		DateTimeRenderOptionAliasFormatted,
	}
)

type RenderOpt func(o RenderOpts) RenderOpts

type RenderOpts struct {
	ValueRenderOption    string `url:"valueRenderOption,omitempty"`
	DateTimeRenderOption string `url:"dateTimeRenderOption,omitempty"`
}

func (z RenderOpts) Apply(opts []RenderOpt) RenderOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

func ValueRenderOption(opt string) RenderOpt {
	return func(o RenderOpts) RenderOpts {
		switch opt {
		case ValueRenderOptionFormatted, ValueRenderOptionAliasFormatted:
			o.ValueRenderOption = ValueRenderOptionFormatted
		case ValueRenderOptionUnformatted, ValueRenderOptionAliasUnformatted:
			o.ValueRenderOption = ValueRenderOptionUnformatted
		case ValueRenderOptionFormula, ValueRenderOptionAliasFormula:
			o.ValueRenderOption = ValueRenderOptionFormula
		default:
			l := esl.Default()
			l.Warn("Undefined value render option", esl.String("opt", opt))
		}
		return o
	}
}
func DateTimeRenderOption(opt string) RenderOpt {
	return func(o RenderOpts) RenderOpts {
		switch opt {
		case DateTimeRenderOptionSerialNumber, DateTimeRenderOptionAliasSerialNumber:
			o.DateTimeRenderOption = DateTimeRenderOptionSerialNumber
		case DateTimeRenderOptionFormatted, DateTimeRenderOptionAliasFormatted:
			o.DateTimeRenderOption = DateTimeRenderOptionFormatted
		default:
			l := esl.Default()
			l.Warn("Undefined date time option", esl.String("opt", opt))
		}
		return o
	}
}

func New(ctx goog_context.Context) Sheet {
	return &shImpl{
		ctx: ctx,
	}
}

func parseSheetFromBatchUpdateAddSheetResponse(d es_json.Json) (sheet bo_sheet.Sheet, err error) {
	err = d.FindModel("replies|0|addSheet", &sheet)
	return
}

type shImpl struct {
	ctx goog_context.Context
}

func (z shImpl) Delete(spreadsheetId, sheetId string) error {
	bu := &to_spreadsheet.BatchUpdate{
		Requests: []to_spreadsheet.BatchUpdateRequest{
			{
				DeleteSheet: &to_spreadsheet.BatchUpdateRequestDeleteSheet{
					SheetId: sheetId,
				},
			},
		},
	}
	res := z.ctx.Post("spreadsheets/"+spreadsheetId+":batchUpdate", api_request.Param(bu))
	if err, f := res.Failure(); f {
		return err
	}
	return nil
}

func (z shImpl) Create(spreadsheetId, sheetTitle string, cols, rows int) (sheet bo_sheet.Sheet, err error) {
	bu := &to_spreadsheet.BatchUpdate{
		Requests: []to_spreadsheet.BatchUpdateRequest{
			{
				AddSheet: &to_spreadsheet.BatchUpdateRequestAddSheet{
					Properties: to_spreadsheet.BatchUpdateRequestAddSheetProperties{
						Title: sheetTitle,
						GridProperties: &to_spreadsheet.BatchUpdateRequestGridProperties{
							RowCount:    &rows,
							ColumnCount: &cols,
						},
					},
				},
			},
		},
	}
	res := z.ctx.Post("spreadsheets/"+spreadsheetId+":batchUpdate", api_request.Param(bu))
	if err, f := res.Failure(); f {
		return bo_sheet.Sheet{}, err
	}
	return parseSheetFromBatchUpdateAddSheetResponse(res.Success().Json())
}

func (z shImpl) Append(spreadsheetId, sheetRange string, values [][]interface{}, rawInput bool) (appended bo_sheet.ValueAppend, err error) {
	encodedRange := url.QueryEscape(sheetRange)
	tvr := to_cell.ValueRange{
		Range:          sheetRange,
		MajorDimension: "ROWS",
		Values:         values,
	}
	type VIO struct {
		ValueInputOption string `url:"valueInputOption,omitempty"`
	}
	var q VIO
	if rawInput {
		q.ValueInputOption = "RAW"
	} else {
		q.ValueInputOption = "USER_ENTERED"
	}

	content, err := json.Marshal(tvr)
	if err != nil {
		return bo_sheet.ValueAppend{}, err
	}
	res := z.ctx.Post("spreadsheets/"+spreadsheetId+"/values/"+encodedRange+":append",
		api_request.Query(&q),
		api_request.Content(es_rewinder.NewReadRewinderOnMemory(content)),
	)
	if err, f := res.Failure(); f {
		return bo_sheet.ValueAppend{}, err
	}
	err = res.Success().Json().Model(&appended)
	return
}

func (z shImpl) Import(spreadsheetId, sheetRange string, values [][]interface{}, rawInput bool) (updated bo_sheet.ValueUpdate, err error) {
	encodedRange := url.QueryEscape(sheetRange)
	tvr := to_cell.ValueRange{
		Range:          sheetRange,
		MajorDimension: "ROWS",
		Values:         values,
	}
	type VIO struct {
		ValueInputOption string `url:"valueInputOption,omitempty"`
	}
	var q VIO
	if rawInput {
		q.ValueInputOption = "RAW"
	} else {
		q.ValueInputOption = "USER_ENTERED"
	}

	content, err := json.Marshal(tvr)
	if err != nil {
		return bo_sheet.ValueUpdate{}, err
	}
	res := z.ctx.Put("spreadsheets/"+spreadsheetId+"/values/"+encodedRange,
		api_request.Query(&q),
		api_request.Content(es_rewinder.NewReadRewinderOnMemory(content)),
	)
	if err, f := res.Failure(); f {
		return bo_sheet.ValueUpdate{}, err
	}
	err = res.Success().Json().Model(&updated)
	return
}

func (z shImpl) Export(spreadsheetId, sheetRange string, opts ...RenderOpt) (value to_cell.ValueRange, err error) {
	ro := RenderOpts{}.Apply(opts)
	encodedRange := url.QueryEscape(sheetRange)
	if !to_spreadsheet.IsValidSpreadsheetId(spreadsheetId) {
		return value, to_spreadsheet.ErrorInvalidSpreadsheetId
	}
	res := z.ctx.Get("spreadsheets/"+spreadsheetId+"/values/"+encodedRange, api_request.Query(&ro))
	if err, f := res.Failure(); f {
		return value, err
	}
	jb := res.Success().BodyString()
	err = json.Unmarshal([]byte(jb), &value)
	return
}

func (z shImpl) Clear(spreadsheetId, sheetRange string) (clearedRange string, err error) {
	encodedRange := url.QueryEscape(sheetRange)
	if !to_spreadsheet.IsValidSpreadsheetId(spreadsheetId) {
		return "", to_spreadsheet.ErrorInvalidSpreadsheetId
	}
	res := z.ctx.Post("spreadsheets/" + spreadsheetId + "/values/" + encodedRange + ":clear")
	if err, f := res.Failure(); f {
		return "", err
	}
	type ClearResponse struct {
		SpreadsheetId string `path:"spreadsheetId"`
		ClearedRange  string `path:"clearedRange"`
	}
	cr := &ClearResponse{}
	err = res.Success().Json().Model(cr)
	return cr.ClearedRange, err
}
