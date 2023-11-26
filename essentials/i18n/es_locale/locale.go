package es_locale

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
	"github.com/watermint/toolbox/essentials/strings/es_case"
	"strings"
)

const (
	TagEnglish  = "en"
	TagJapanese = "ja"
)

var (
	// Default fallback locale
	Default  = mustParse(TagEnglish)
	English  = mustParse(TagEnglish)
	Japanese = mustParse(TagJapanese)
)

// Locale represents a specific geopolitical region.
type Locale interface {
	// Stringer Language tag
	fmt.Stringer

	// Language ISO 639 code (2-3 letter code, always lower case)
	Language() string

	// LanguageTwoLetter ISO 639-1 Two-letter code (always lower case).
	// Returns empty if the language is not defined in ISO 639-1.
	LanguageTwoLetter() string

	// LanguageExtended Extended language subtag
	LanguageExtended() string

	// Extension provide a mechanism for extending language tags for use in various applications.
	Extension() string

	// Script subtags are used to indicate the script or writing system
	//   variations that distinguish the written forms of a language or its
	//   dialects. ISO 15924.
	//   First letter always upper case. Following letters are lower case like Xxxx (upper camel case).
	Script() string

	// Variant subtags are used to indicate additional, well-recognized
	//   variations that define a language or its dialects that are not
	//   covered by other available subtags.
	Variant() string

	// Region subtags are used to indicate linguistic variations associated
	//   with or appropriate to a specific country, territory, or region.
	//   Typically, a region subtag is used to indicate variations such as
	//   regional dialects or usage, or region-specific spelling conventions.
	//   ISO 3166 country code or UN M49 region code.
	//   (always upper case).
	Region() string

	Data() LocaleData
}

type LocaleData struct {
	// BCP 47 language tag
	Tag           string `json:"tag"`
	Lang          string `json:"lang"`
	LangExtended  string `json:"lang_extended"`
	Extension     string `json:"extension"`
	Grandfathered string `json:"grandfathered"`
	Region        string `json:"region"`
	Script        string `json:"script"`
	Variant       string `json:"variant"`
	Private1      string `json:"private1"`
	Private2      string `json:"private2"`
	CodePage      string `json:"code_page"`
}

func mustParse(langTag string) Locale {
	if lc, out := Parse(langTag); out.IsError() {
		panic(out)
	} else {
		return lc
	}
}

func Parse(langTag string) (local Locale, outcome eoutcome.ParseOutcome) {
	lowerCaseLangTag := strings.ToLower(langTag)
	if lowerCaseLangTag == "c" || lowerCaseLangTag == "posix" ||
		strings.HasPrefix(lowerCaseLangTag, "c.") ||
		strings.HasPrefix(lowerCaseLangTag, "posix.") {
		return mustParse(TagEnglish), eoutcome.NewParseSuccess()
	}

	// accept tag like "ja_JP" as "ja-JP" (BCP 47 compliant)
	if strings.Index(langTag, "_") >= 0 {
		langTag = strings.ReplaceAll(langTag, "_", "-")
	}

	matches := bcp47Lang.FindStringSubmatch(langTag)
	if len(matches) < 1 {
		return nil, eoutcome.NewParseInvalidFormat("the given lang-tag does not comply lang tag format")
	}
	language := strings.ToLower(matches[1])

	detail, match := bcp47Re.MatchSubExp(langTag)
	if !match {
		return nil, eoutcome.NewParseInvalidFormat("the given lang-tag does not comply lang tag format")
	}

	data := LocaleData{
		Tag:           langTag,
		Lang:          language,
		Extension:     detail["extension"],
		LangExtended:  detail["extlang"],
		Grandfathered: detail["grandfathered"],
		Region:        strings.ToUpper(detail["region"]),
		Script:        es_case.TokenToCamelCase(detail["script"]),
		Variant:       detail["variant"],
		Private1:      detail["privateUse"],
		Private2:      detail["privateUse2"],
		CodePage:      detail["codepage"],
	}

	return &localeImpl{data: data}, eoutcome.NewParseSuccess()
}
