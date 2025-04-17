package es_locale

import (
	"fmt"
	"strings"

	"github.com/watermint/toolbox/essentials/go/es_errors"
	"github.com/watermint/toolbox/essentials/strings/es_case"
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
	// stringer returns the language tag
	fmt.Stringer

	// language returns the ISO 639 code (2-3 letter code, always lower case)
	Language() string

	// languageTwoLetter returns the ISO 639-1 two-letter code (always lower case).
	// returns empty if the language is not defined in ISO 639-1.
	LanguageTwoLetter() string

	// languageExtended returns the extended language subtag
	LanguageExtended() string

	// extension provides a mechanism for extending language tags for use in various applications.
	Extension() string

	// script subtags are used to indicate the script or writing system
	//   variations that distinguish the written forms of a language or its
	//   dialects. ISO 15924.
	//   first letter always upper case. following letters are lower case like Xxxx (upper camel case).
	Script() string

	// variant subtags are used to indicate additional, well-recognized
	//   variations that define a language or its dialects that are not
	//   covered by other available subtags.
	Variant() string

	// region subtags are used to indicate linguistic variations associated
	//   with or appropriate to a specific country, territory, or region.
	//   typically, a region subtag is used to indicate variations such as
	//   regional dialects or usage, or region-specific spelling conventions.
	//   ISO 3166 country code or UN M49 region code.
	//   (always upper case).
	Region() string

	Data() LocaleData
}

type LocaleData struct {
	// bcp 47 language tag
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
	if lc, err := Parse(langTag); err != nil {
		panic(err)
	} else {
		return lc
	}
}

func Parse(langTag string) (local Locale, err error) {
	lowerCaseLangTag := strings.ToLower(langTag)
	if lowerCaseLangTag == "c" || lowerCaseLangTag == "posix" ||
		strings.HasPrefix(lowerCaseLangTag, "c.") ||
		strings.HasPrefix(lowerCaseLangTag, "posix.") {
		return mustParse(TagEnglish), nil
	}

	// accept tag like "ja_JP" as "ja-JP" (BCP 47 compliant)
    if strings.Contains(langTag, "_") {
		langTag = strings.ReplaceAll(langTag, "_", "-")
	}

	matches := bcp47Lang.FindStringSubmatch(langTag)
	if len(matches) < 1 {
		return nil, es_errors.NewInvalidFormatError("the given lang-tag does not comply lang tag format")
	}
	language := strings.ToLower(matches[1])

	detail, match := bcp47Re.MatchSubExp(langTag)
	if !match {
		return nil, es_errors.NewInvalidFormatError("the given lang-tag does not comply lang tag format")
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

	return &localeImpl{data: data}, nil
}
