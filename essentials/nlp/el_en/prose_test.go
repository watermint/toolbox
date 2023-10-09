package el_en

import (
	"github.com/watermint/toolbox/essentials/cache/ec_file"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"os"
	"testing"
)

func TestContainerImpl_NewDocument(t *testing.T) {
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
	text := "James had always loved the bustling city of New York. From the tall skyscrapers that seemed to touch the clouds to the busy streets filled with a myriad of people, each with their own story to tell. Every morning, he would walk past Central Park, taking in the fresh air and watching as joggers made their way through the winding paths. The aroma from the nearby coffee shops was intoxicating, beckoning him to take a moment and savor a cup."
	doc, err := dc.NewDocument(text)
	if err != nil {
		t.Error(err)
		return
	}
	sentences := doc.Sentences()
	if len(sentences) != 4 {
		t.Error("invalid sentence count", len(sentences))
		return
	}
}
