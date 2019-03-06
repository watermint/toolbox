package copy_ref

import (
	"encoding/json"
	"github.com/watermint/toolbox/model/dbx_api"
	"testing"
)

func TestModelCopyRef(t *testing.T) {
	res := `{
    "metadata": {
        ".tag": "file",
        "name": "Prime_Numbers.txt",
        "id": "id:a4ayc_80_OEAAAAAAAAAXw",
        "client_modified": "2015-05-12T15:50:38Z",
        "server_modified": "2015-05-12T15:50:38Z",
        "rev": "a1c10ce0dd78",
        "size": 7212,
        "path_lower": "/homework/math/prime_numbers.txt",
        "path_display": "/Homework/math/Prime_Numbers.txt",
        "sharing_info": {
            "read_only": true,
            "parent_shared_folder_id": "84528192421",
            "modified_by": "dbid:AAH4f99T0taONIb-OurWxbNQ6ywGRopQngc"
        },
        "property_groups": [
            {
                "template_id": "ptid:1a5n2i6d3OYEAAAAAAAAAYa",
                "fields": [
                    {
                        "name": "Security Policy",
                        "value": "Confidential"
                    }
                ]
            }
        ],
        "has_explicit_shared_members": false,
        "content_hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
    },
    "copy_reference": "z1X6ATl6aWtzOGq0c3g5Ng",
    "expires": "2045-05-12T15:50:38Z"
}`

	cr := CopyRef{}

	if dbx_api.ParseModelJsonForTest(&cr, json.RawMessage(res)) != nil {
		t.Error("failed to parse")
	}

	if cr.CopyReference != "z1X6ATl6aWtzOGq0c3g5Ng" {
		t.Error("invalid")
	}
	if cr.Expires != "2045-05-12T15:50:38Z" {
		t.Error("invalid")
	}
}
