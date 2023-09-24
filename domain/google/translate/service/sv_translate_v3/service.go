package sv_translate_v3

import (
	"errors"
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/translate/model/to_translate_v3"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"strings"
)

type Translate interface {
	// Translate the text from source language to target language.
	// Both source and target language are ISO 639 codes.
	// If source language is empty, the language is auto-detected.
	Translate(text, source, target string) (result string, err error)
}

func New(projectId string, client goog_client.Client) Translate {
	return &translateImpl{
		projectId: projectId,
		client:    client,
	}
}

type translateImpl struct {
	projectId string
	client    goog_client.Client
}

func (z translateImpl) Translate(text, source, target string) (result string, err error) {
	if len(strings.TrimSpace(text)) == 0 {
		return "", nil
	}
	to := &to_translate_v3.Translate{
		Contents:           []string{text},
		SourceLanguageCode: source,
		TargetLanguageCode: target,
	}
	res := z.client.Post("projects/"+z.projectId+":translateText", api_request.Param(&to))
	if err, fail := res.Failure(); fail {
		return "", err
	}
	translations, ok := res.Success().Json().FindArray("translations")
	if !ok || len(translations) == 0 {
		return "", errors.New("empty response")
	}
	translatedText, ok := translations[0].FindString("translatedText")
	if !ok {
		return "", errors.New("empty response")
	}
	return translatedText, nil
}
