package sv_team

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"go.uber.org/zap"
)

type Team interface {
	Info() (info *mo_team.Info, err error)
	Feature() (feature *mo_team.Feature, err error)
}

func New(ctx api_context.Context) Team {
	return &teamImpl{
		ctx: ctx,
	}
}

type teamImpl struct {
	ctx api_context.Context
}

func (z *teamImpl) Info() (info *mo_team.Info, err error) {
	info = &mo_team.Info{}
	res, err := z.ctx.Rpc("team/get_info").Call()
	if err != nil {
		return nil, err
	}
	if err = res.Model(info); err != nil {
		return nil, err
	}
	return info, nil
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
		res, err := z.ctx.Rpc("team/features/get_values").Param(p).Call()
		if err != nil {
			return nil, err
		}
		j, err := res.Json()
		if err != nil {
			return nil, err
		}
		values := j.Get("values")
		if !values.IsArray() {
			return nil, err
		}
		first := values.Array()[0]
		valueTag := first.Get("\\.tag")
		if !valueTag.Exists() {
			return nil, err
		}
		value := first.Get(valueTag.String())
		if !value.Exists() {
			return nil, err
		}
		features[tag] = json.RawMessage(value.Raw)
	}

	raw := api_parser.CombineRaw(features)
	feature = &mo_team.Feature{}
	if err = api_parser.ParseModelRaw(feature, raw); err != nil {
		return nil, err
	}
	return feature, nil
}
