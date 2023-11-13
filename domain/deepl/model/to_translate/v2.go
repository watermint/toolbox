package to_translate

type V2TranslateRequest struct {
	Text               []string `json:"text"`
	SourceLang         string   `json:"source_lang,omitempty"`
	TargetLang         string   `json:"target_lang"`
	Context            string   `json:"context,omitempty"`
	SplitSentences     string   `json:"split_sentences,omitempty"`
	PreserveFormatting *bool    `json:"preserve_formatting,omitempty"`
	Formality          string   `json:"formality,omitempty"`
	GlossaryId         string   `json:"glossary_id,omitempty"`
	TagHandling        string   `json:"tag_handling,omitempty"`
	OutlineDetection   *bool    `json:"outline_detection,omitempty"`
}

type V2TranslateResponses struct {
	Translations []V2TranslateResponse `json:"translations"`
}

type V2TranslateResponse struct {
	Text                   string `json:"text"`
	DetectedSourceLanguage string `json:"detected_source_language,omitempty"`
}
