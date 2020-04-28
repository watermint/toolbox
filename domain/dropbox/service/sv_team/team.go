package sv_team

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/api/api_request"
	"go.uber.org/zap"
)

var (
	ErrorUnexpectedFormat = errors.New("unexpected response format")
)

type Team interface {
	Info() (info *mo_team.Info, err error)
	Feature() (feature *mo_team.Feature, err error)
}

func New(ctx dbx_context.Context) Team {
	return &teamImpl{
		ctx: ctx,
	}
}

type teamImpl struct {
	ctx dbx_context.Context
}

func (z *teamImpl) Info() (info *mo_team.Info, err error) {
	res := z.ctx.Post("team/get_info")
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	info = &mo_team.Info{}
	err = res.Success().Json().Model(info)
	return
}

func (z *teamImpl) Feature() (feature *mo_team.Feature, err error) {
	featureTags := []string{
		"upload_api_rate_limit",
		"has_team_shared_dropbox",
		"has_team_file_events",
		"has_team_selective_sync",
	}
	type FT struct {
		Tag string `json:".tag"`
	}
	type FP struct {
		Values []FT `json:"features"`
	}

	features := make(map[string]json.RawMessage)

	for _, tag := range featureTags {
		z.ctx.Log().Debug("Feature", zap.String("tag", tag))
		p := FP{Values: []FT{{Tag: tag}}}
		res := z.ctx.Post("team/features/get_values", api_request.Param(p))
		if err, fail := res.Failure(); fail {
			return nil, err
		}
		firstValue, found := res.Success().Json().Find("values.0")
		if !found {
			return nil, ErrorUnexpectedFormat
		}
		valueTag, found := firstValue.FindString("\\.tag")
		if !found {
			return nil, ErrorUnexpectedFormat
		}
		value, found := firstValue.Find(valueTag)
		if !found {
			return nil, ErrorUnexpectedFormat
		}

		features[tag] = value.Raw()
	}

	raw := api_parser.CombineRaw(features)
	feature = &mo_team.Feature{}
	err = es_json.MustParse(raw).Model(feature)
	return
}
