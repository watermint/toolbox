package es_case

import (
	"github.com/watermint/toolbox/essentials/strings/es_tokenizer"
	"strings"
)

func ToLowerDelimited(s, delimiter string) string {
	input := es_tokenizer.AlphaNumCaseTokenizer().Tokens(s)
	if len(input) < 1 {
		return ""
	}
	output := make([]string, len(input))

	for i := 0; i < len(output); i++ {
		output[i] = strings.ToLower(input[i])
	}
	return strings.Join(output, delimiter)
}

func ToUpperDelimited(s, delimiter string) string {
	input := es_tokenizer.AlphaNumCaseTokenizer().Tokens(s)
	if len(input) < 1 {
		return ""
	}
	output := make([]string, len(input))

	for i := 0; i < len(output); i++ {
		output[i] = strings.ToUpper(input[i])
	}
	return strings.Join(output, delimiter)
}
