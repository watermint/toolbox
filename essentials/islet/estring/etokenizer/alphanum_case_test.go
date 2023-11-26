package etokenizer

import (
	"reflect"
	"testing"
)

func TestAlNumCaseTokenizer_Tokens(t *testing.T) {
	alt := AlphaNumCaseTokenizer()
	verify := func(s string, expected []string) {
		x := alt.Tokens(s)
		if !reflect.DeepEqual(x, expected) {
			t.Error(s, x)
		}
	}

	verify("", []string{})
	verify("abc", []string{"abc"})
	verify("abc def ghi", []string{"abc", "def", "ghi"})
	verify("=abc=def=ghi=", []string{"abc", "def", "ghi"})
	verify("ａｂｃ=def=ghi=", []string{"ａｂｃ", "def", "ghi"})
	verify("Powered by GoLang version1.17.", []string{"Powered", "by", "Go", "Lang", "version1", "17"})
	verify(" func TestAlNumCaseTokenizer_Tokens(t *testing.T) {}", []string{"func", "Test", "Al", "Num", "Case", "Tokenizer", "Tokens", "t", "testing", "T"})
}
