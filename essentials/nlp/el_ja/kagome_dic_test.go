package el_ja

import (
	tokenizer2 "github.com/ikawaha/kagome/v2/tokenizer"
	"github.com/watermint/toolbox/essentials/cache/ec_file"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"os"
	"reflect"
	"testing"
)

func TestDictionaryContainerImpl_Load(t *testing.T) {
	if qt_endtoend.IsSkipEndToEndTest() {
		t.Skip("skip download test")
		return
	}

	cacheRoot, err := os.MkdirTemp("", "cache-dict")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(cacheRoot)
	}()
	l := esl.Default()
	dc := NewContainer(ec_file.New(cacheRoot, l), l)
	d, err := dc.Load("ipa")
	if err != nil {
		t.Error(err)
		return
	}
	tokenizer, err := tokenizer2.New(d)
	if err != nil {
		t.Error(err)
		return
	}
	segments := tokenizer.Wakati("すもももももももものうち")
	if !reflect.DeepEqual(segments, []string{"すもも", "も", "もも", "も", "もも", "の", "うち"}) {
		t.Error("invalid wakati gaki result", segments)
		return
	}
}
