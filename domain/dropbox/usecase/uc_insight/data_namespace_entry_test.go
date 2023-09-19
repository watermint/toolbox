package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

const (
	sampleNamespaceEntry = `{
            ".tag": "file",
            "client_modified": "2015-05-12T15:50:38Z",
            "content_hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
            "file_lock_info": {
                "created": "2015-05-12T15:50:38Z",
                "is_lockholder": true,
                "lockholder_name": "Imaginary User"
            },
            "has_explicit_shared_members": false,
            "id": "id:a4ayc_80_OEAAAAAAAAAXw",
            "is_downloadable": true,
            "name": "Prime_Numbers.txt",
            "path_display": "/Homework/math/Prime_Numbers.txt",
            "path_lower": "/homework/math/prime_numbers.txt",
            "property_groups": [
                {
                    "fields": [
                        {
                            "name": "Security Policy",
                            "value": "Confidential"
                        }
                    ],
                    "template_id": "ptid:1a5n2i6d3OYEAAAAAAAAAYa"
                }
            ],
            "rev": "a1c10ce0dd78",
            "server_modified": "2015-05-12T15:50:38Z",
            "sharing_info": {
                "modified_by": "dbid:AAH4f99T0taONIb-OurWxbNQ6ywGRopQngc",
                "parent_shared_folder_id": "84528192421",
                "read_only": true
            },
            "size": 7212
        }`
)

func TestNewNamespaceEntryFromJson(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ne, err := NewNamespaceEntry("ns:1234", "/Homework/math", es_json.MustParseString(sampleNamespaceEntry))
		if err != nil {
			t.Error(err)
		}
		if ne.PathDisplay != "/Homework/math/Prime_Numbers.txt" {
			t.Error(ne.PathDisplay)
		}
		if ne.PathLower != "/homework/math/prime_numbers.txt" {
			t.Error(ne.PathLower)
		}
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		if err := db.AutoMigrate(&NamespaceEntry{}); err != nil {
			t.Error(err)
			return
		}
		if err := db.Create(ne).Error; err != nil {
			t.Error(err)
			return
		}

		ne2 := &NamespaceEntry{}
		db.First(&ne2)
		if !reflect.DeepEqual(ne, ne2) {
			t.Error(ne2)
		}
	})
}
