package sv_file_tag

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"reflect"
	"testing"
)

func TestParseTags(t *testing.T) {
	res1 := `{
    "paths_to_tags": [
        {
            "path": "/Prime_Numbers.txt",
            "tags": [
                {
                    ".tag": "user_generated_tag",
                    "tag_text": "my_tag"
                }
            ]
        }
    ]
}`

	res2 := `{
    "paths_to_tags": [
        {
            "path": "/Prime_Numbers.txt",
            "tags": [
                {
                    ".tag": "user_generated_tag",
                    "tag_text": "prime"
                },
                {
                    ".tag": "user_generated_tag",
                    "tag_text": "numbers"
                }
            ]
        }
    ]
}`

	rj1 := es_json.MustParseString(res1)
	tags1, err := parseTags(rj1)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tags1, []string{"my_tag"}) {
		t.Error(tags1)
	}

	rj2 := es_json.MustParseString(res2)
	tags2, err := parseTags(rj2)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(tags2, []string{"prime", "numbers"}) {
		t.Error(tags2)
	}
}
