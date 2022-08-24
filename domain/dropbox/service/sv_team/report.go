package sv_team

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/infra/api/api_request"
)

type Report interface {
	Activity(span ReportSpan) (activity mo_team.Activity, err error)
	Devices(span ReportSpan) (devices mo_team.Devices, err error)
	Membership(span ReportSpan) (membership mo_team.Membership, err error)
	Storage(span ReportSpan) (storage mo_team.Storage, err error)
}

func NewReport(ctx dbx_client.Client) Report {
	return &repImpl{
		ctx: ctx,
	}
}

type repImpl struct {
	ctx dbx_client.Client
}

type ReportSpan struct {
	StartDate string `json:"start_date,omitempty" url:"start_date"`
	EndDate   string `json:"end_date,omitempty" url:"end_date"`
}

func NewSpan(startDate, endDate mo_time.TimeOptional) (span ReportSpan) {
	isoDateFormat := "2006-01-02"
	if startDate.Ok() {
		span.StartDate = startDate.Time().Format(isoDateFormat)
	}
	if endDate.Ok() {
		span.EndDate = endDate.Time().Format(isoDateFormat)
	}
	return
}

func (z repImpl) Activity(span ReportSpan) (activity mo_team.Activity, err error) {
	res := z.ctx.Post("team/reports/get_activity", api_request.Param(&span))
	var fail bool
	if err, fail = res.Failure(); fail {
		return
	}

	err = res.Success().Json().Model(&activity)
	return
}

func (z repImpl) Devices(span ReportSpan) (devices mo_team.Devices, err error) {
	res := z.ctx.Post("team/reports/get_devices", api_request.Param(&span))
	var fail bool
	if err, fail = res.Failure(); fail {
		return
	}

	err = res.Success().Json().Model(&devices)
	return
}

func (z repImpl) Membership(span ReportSpan) (membership mo_team.Membership, err error) {
	res := z.ctx.Post("team/reports/get_membership", api_request.Param(&span))
	var fail bool
	if err, fail = res.Failure(); fail {
		return
	}

	err = res.Success().Json().Model(&membership)
	return
}

func (z repImpl) Storage(span ReportSpan) (storage mo_team.Storage, err error) {
	res := z.ctx.Post("team/reports/get_storage", api_request.Param(&span))
	var fail bool
	if err, fail = res.Failure(); fail {
		return
	}

	err = res.Success().Json().Model(&storage)
	return
}
