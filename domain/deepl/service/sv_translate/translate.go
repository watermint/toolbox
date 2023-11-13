package sv_translate

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/deepl/api/deepl_client"
	"github.com/watermint/toolbox/domain/deepl/model/to_translate"
	"github.com/watermint/toolbox/essentials/api/api_request"
)

type Translate interface {
	Translate(text string, sourceLang string, targetLang string) (result to_translate.V2TranslateResponse, err error)
}

func New(client deepl_client.Client) Translate {
	return &translateImpl{
		client: client,
	}
}

type translateImpl struct {
	client deepl_client.Client
}

func (z translateImpl) Translate(text string, sourceLang string, targetLang string) (result to_translate.V2TranslateResponse, err error) {
	req := to_translate.V2TranslateRequest{
		Text:       []string{text},
		SourceLang: sourceLang,
		TargetLang: targetLang,
	}
	res := z.client.Post("translate", api_request.Param(req))
	if err, fail := res.Failure(); fail {
		return result, err
	}
	translations := to_translate.V2TranslateResponses{}
	err = json.Unmarshal(res.Success().Body(), &translations)
	return translations.Translations[0], err
}
