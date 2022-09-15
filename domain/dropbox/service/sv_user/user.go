package sv_user

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_user"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type User interface {
	Features() (feature *mo_user.Feature, err error)
}

func New(ctx dbx_client.Client) User {
	return &userImpl{
		ctx: ctx,
	}
}

type userImpl struct {
	ctx dbx_client.Client
}

func (z userImpl) Features() (feature *mo_user.Feature, err error) {
	l := z.ctx.Log()
	featureTags := []string{
		"paper_as_files",
		"file_locking",
	}
	type FT struct {
		Tag string `json:".tag"`
	}
	type FP struct {
		Features []FT `json:"features"`
	}

	fp := FP{Features: make([]FT, 0)}
	for _, f := range featureTags {
		fp.Features = append(fp.Features, FT{
			Tag: f,
		})
	}
	res := z.ctx.Post("users/features/get_values", api_request.Param(&fp))
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	rj := res.Success().Json()
	feature = &mo_user.Feature{
		PaperAsFiles: false,
		FileLocking:  false,
	}
	if values, ok := rj.FindArray("values"); !ok {
		l.Debug("No feature response found", esl.String("response", rj.RawString()))
		return nil, errors.New("no feature values found in the response")
	} else {
		valMap := make(map[string]bool)
		for _, v := range values {
			if name, ok := v.FindString(".tag"); !ok {
				l.Debug("Skip: tag not found", esl.String("value", v.RawString()))
				continue
			} else {
				if enabled, ok := v.FindBool("enabled"); !ok {
					l.Debug("Skip: enabled tag not found", esl.String("value", v.RawString()))
					continue
				} else {
					valMap[name] = enabled
				}
			}
		}

		valJson, err := json.Marshal(valMap)
		if err != nil {
			l.Debug("Unable to marshal valMap", esl.Error(err))
			return nil, err
		}

		err = json.Unmarshal(valJson, feature)
		if err != nil {
			l.Debug("Unable to unmarshal", esl.Error(err))
			return nil, err
		}
		return feature, nil
	}
}
