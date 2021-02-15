package sv_sheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/sheets/model/to_spreadsheet"
	"net/url"
)

type Sheet interface {
	Clear(spreadsheetId, sheetRange string) (clearedRange string, err error)
}

func New(ctx goog_context.Context) Sheet {
	return &shImpl{
		ctx: ctx,
	}
}

type shImpl struct {
	ctx goog_context.Context
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
