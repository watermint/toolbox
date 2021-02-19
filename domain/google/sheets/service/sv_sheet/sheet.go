package sv_sheet

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/sheets/model/to_cell"
	"github.com/watermint/toolbox/domain/google/sheets/model/to_spreadsheet"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
	"net/url"
)

type Sheet interface {
	Clear(spreadsheetId, sheetRange string) (clearedRange string, err error)
	Export(spreadsheetId, sheetRange string, opts ...RenderOpt) (value to_cell.ValueRange, err error)
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

type shImpl struct {
	ctx goog_context.Context
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
