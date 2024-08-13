package to_translate_v3

type Translate struct {
	Contents           []string `json:"contents"`
	MimeType           string   `json:"mimeType,omitempty"`
	SourceLanguageCode string   `json:"sourceLanguageCode,omitempty"`
	TargetLanguageCode string   `json:"targetLanguageCode"`
	Model              string   `json:"model,omitempty"`
}

type TranslationResponse struct {
	Translations []*Translation `json:"translations"`
}

type Translation struct {
	TranslatedText       string `json:"translatedText" path:"translatedText"`
	Model                string `json:"model" path:"model"`
	DetectedLanguageCode string `json:"detectedLanguageCode" path:"detectedLanguageCode"`
}
