package es_case

import (
	"github.com/watermint/toolbox/essentials/strings/es_tokenizer"
	"strings"
	"unicode"
)

// TokenToCamelCase changes token to camel case. "token"/"TOKEN" to "Token".
// This function expect the token consist of only alpha-numeric characters.
// The function will not trim spaces or tokenize.
func TokenToCamelCase(token string) string {
	switch len(token) {
	case 0:
		return ""
	case 1:
		return strings.ToUpper(token)
	default:
		runes := []rune(token)
		runes[0] = unicode.ToUpper(runes[0])
		for i := 1; i < len(runes); i++ {
			runes[i] = unicode.ToLower(runes[i])
		}
		return string(runes)
	}
}

// ToUpperCamelCase change to upper camel case like from "powered by go-lang" to "PoweredByGoLang".
func ToUpperCamelCase(s string) string {
	input := es_tokenizer.AlphaNumCaseTokenizer().Tokens(s)
	if len(input) < 1 {
		return ""
	}
	output := make([]string, len(input))

	for i := 0; i < len(output); i++ {
		output[i] = TokenToCamelCase(input[i])
	}
	return strings.Join(output, "")
}

// ToLowerCamelCase change to lower camel case like from "Powered BY GO-lang" to "poweredByGoLang".
func ToLowerCamelCase(s string) string {
	input := es_tokenizer.AlphaNumCaseTokenizer().Tokens(s)
	if len(input) < 1 {
		return ""
	}
	output := make([]string, len(input))
	output[0] = strings.ToLower(input[0])
	if len(input) < 2 {
		return output[0]
	}
	for i := 1; i < len(output); i++ {
		output[i] = TokenToCamelCase(input[i])
	}
	return strings.Join(output, "")
}
