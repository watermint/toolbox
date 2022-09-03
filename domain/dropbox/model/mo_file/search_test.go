package mo_file

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestMatch_Concrete(t *testing.T) {
	j := `{
            "metadata": {
                ".tag": "metadata",
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
                    "is_downloadable": true,
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
                    "content_hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
                    "file_lock_info": {
                        "is_lockholder": true,
                        "lockholder_name": "Imaginary User",
                        "created": "2015-05-12T15:50:38Z"
                    }
                }
            }
        }`
	m := &Match{}
	if err := api_parser.ParseModelString(m, j); err != nil {
		t.Error(err)
	}
	if m.Concrete().Name != m.EntryName || m.EntryName != "Prime_Numbers.txt" {
		t.Error(m.EntryName)
	}
}
