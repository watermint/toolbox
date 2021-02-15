package sv_spreadsheet

import (
	"errors"
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/sheets/model/bo_spreadsheet"
	"github.com/watermint/toolbox/domain/google/sheets/model/to_spreadsheet"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
	"regexp"
)

var (
	// https://developers.google.com/sheets/api/guides/concepts#spreadsheet_id
	SpreadsheetIdPattern = regexp.MustCompile(`^([a-zA-Z0-9-_]+)$`)

	ErrorInvalidSpreadsheetId = errors.New("invalid spreadsheet id")
)

func isValidSpreadsheetId(id string) bool {
	return SpreadsheetIdPattern.MatchString(id)
}

type Spreadsheet interface {
	Create(title string) (spreadsheet *bo_spreadsheet.Spreadsheet, err error)
	Resolve(id string) (spreadsheet *bo_spreadsheet.Spreadsheet, err error)
}

func New(ctx goog_context.Context) Spreadsheet {
	return &ssImpl{
		ctx: ctx,
	}
}

type ssImpl struct {
	ctx goog_context.Context
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
	if !isValidSpreadsheetId(id) {
		l := esl.Default()
		l.Debug("Invalid spreadsheet id format", esl.String("id", id))
		return nil, ErrorInvalidSpreadsheetId
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
