package sv_spreadsheet

import (
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_spreadsheet"
	"github.com/watermint/toolbox/domain/google/sheets/model/to_spreadsheet"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type Spreadsheet interface {
	Create(title string) (spreadsheet *bo_spreadsheet.Spreadsheet, err error)
	Resolve(id string) (spreadsheet *bo_spreadsheet.Spreadsheet, err error)
}

func New(ctx goog_client.Client) Spreadsheet {
	return &ssImpl{
		ctx: ctx,
	}
}

type ssImpl struct {
	ctx goog_client.Client
}

func (z ssImpl) resolveSpreadsheet(id string, includeGridData bool) (spreadsheet *bo_spreadsheet.Spreadsheet, err error) {
	type Q struct {
		Ranges          string `url:"ranges,omitempty"`
		IncludeGridData *bool  `url:"includeGridData,omitempty"`
	}
	q := &Q{}
	if includeGridData {
		q.IncludeGridData = &includeGridData
	}
	if !to_spreadsheet.IsValidSpreadsheetId(id) {
		l := esl.Default()
		l.Debug("Invalid spreadsheet id format", esl.String("id", id))
		return nil, to_spreadsheet.ErrorInvalidSpreadsheetId
	}

	res := z.ctx.Get("spreadsheets/"+id, api_request.Query(q))
	if err, f := res.Failure(); f {
		return nil, err
	}
	spreadsheet = &bo_spreadsheet.Spreadsheet{}
	err = res.Success().Json().Model(spreadsheet)
	return
}

func (z ssImpl) Resolve(id string) (spreadsheet *bo_spreadsheet.Spreadsheet, err error) {
	return z.resolveSpreadsheet(id, false)
}

func (z ssImpl) Create(title string) (spreadsheet *bo_spreadsheet.Spreadsheet, err error) {
	to := to_spreadsheet.Spreadsheet{
		Properties: &to_spreadsheet.SpreadsheetProperties{
			Title: title,
		},
	}
	res := z.ctx.Post("spreadsheets", api_request.Param(&to))
	if err, f := res.Failure(); f {
		return nil, err
	}
	spreadsheet = &bo_spreadsheet.Spreadsheet{}
	err = res.Success().Json().Model(spreadsheet)
	return
}
