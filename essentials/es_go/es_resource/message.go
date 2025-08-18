package es_resource

import (
	"encoding/json"
	"fmt"
)

func MergedMessageResource(bundles []Bundle, langCodes []string) Resource {
	dataFiles := make(map[string][]byte)
	for _, lang := range langCodes {
		dataFiles[fmt.Sprintf("%s/messages.json", lang)] = MergedMessages(bundles, lang)
	}
	return NewEmbedResource(dataFiles)
}

func MergedMessages(bundles []Bundle, langCode string) []byte {
	all := make(map[string]string)
	for i := len(bundles) - 1; i >= 0; i-- {
		bundle := bundles[i]
		messageFile := fmt.Sprintf("%s/messages.json", langCode)
		messageBody, err := bundle.Messages().Bytes(messageFile)
		if err != nil {
			continue
		}
		messages := make(map[string]string)
		if err := json.Unmarshal(messageBody, &messages); err != nil {
			panic(err)
		}
		for k, v := range messages {
			all[k] = v
		}
	}
	body, err := json.Marshal(all)
	if err != nil {
		panic(err)
	}
	return body
}
