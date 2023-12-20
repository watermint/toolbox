package es_resource

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func MergedMessageResource(bundles []Bundle) Resource {
	dataFiles := make(map[string][]byte)
	for _, lang := range es_lang.Supported {
		dataFiles[fmt.Sprintf("%s/messages.json", lang.CodeString())] = MergedMessages(bundles, lang)
	}
	return NewEmbedResource(dataFiles)
}

func MergedMessages(bundles []Bundle, lang es_lang.Lang) []byte {
	l := esl.Default()
	all := make(map[string]string)
	for i := len(bundles) - 1; i >= 0; i-- {
		bundle := bundles[i]
		messageFile := fmt.Sprintf("%s/messages.json", lang.CodeString())
		messageBody, err := bundle.Messages().Bytes(messageFile)
		if err != nil {
			l.Debug("No message file", esl.Error(err), esl.String("file", messageFile))
			continue
		}
		messages := make(map[string]string)
		if err := json.Unmarshal(messageBody, &messages); err != nil {
			l.Error("Unable to parse message file", esl.Error(err), esl.String("file", messageFile))
			panic(err)
		}
		for k, v := range messages {
			all[k] = v
		}
	}
	body, err := json.Marshal(all)
	if err != nil {
		l.Error("Unable to marshal messages", esl.Error(err))
		panic(err)
	}
	return body
}
