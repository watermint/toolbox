package dbx_async

//import (
//	"github.com/tidwall/gjson"
//	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
//	"testing"
//)
//
//var (
//	fileJson = `{
//		".tag": "file",
//		"name": "Prime_Numbers.txt",
//		"id": "id:a4ayc_80_OEAAAAAAAAAXw",
//		"client_modified": "2015-05-12T15:50:38Z",
//		"server_modified": "2015-05-12T15:50:38Z",
//		"rev": "a1c10ce0dd78",
//		"size": 7212,
//		"path_lower": "/homework/math/prime_numbers.txt",
//		"path_display": "/Homework/math/Prime_Numbers.txt",
//		"sharing_info": {
//			"read_only": true,
//			"parent_shared_folder_id": "84528192421",
//			"modified_by": "dbid:AAH4f99T0taONIb-OurWxbNQ6ywGRopQngc"
//		},
//		"property_groups": [
//			{
//				"template_id": "ptid:1a5n2i6d3OYEAAAAAAAAAYa",
//				"fields": [
//					{
//						"name": "Security Policy",
//						"value": "Confidential"
//					}
//				]
//			}
//		],
//		"has_explicit_shared_members": false,
//		"content_hash": "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
//	}`
//)
//
//func TestResponseImpl_Json(t *testing.T) {
//	r := &responseImpl{
//		complete:       gjson.Parse("{}"),
//		completeExists: true,
//	}
//	if j, err := r.Json(); err != nil {
//		t.Error(err)
//	} else if j.Raw != "{}" {
//		t.Error(j.Raw)
//	}
//}
//
//func TestResponseImpl_Model(t *testing.T) {
//	r := &responseImpl{
//		complete:       gjson.Parse(fileJson),
//		completeExists: true,
//	}
//	m := &mo_file.Metadata{}
//	if err := r.Model(m); err != nil {
//		t.Error(err)
//	} else if m.Name() != "Prime_Numbers.txt" {
//		t.Error(m)
//	}
//}
//
//func TestResponseImpl_ModelWithPath(t *testing.T) {
//	r := &responseImpl{completeExists: false}
//	if err := r.ModelWithPath(&mo_file.Metadata{}, "file"); err != ErrorNoResult {
//		t.Error(err)
//	}
//
//	fj := `{"entry":` + fileJson + `}`
//	r = &responseImpl{
//		complete:       gjson.Parse(fj),
//		completeExists: true,
//	}
//	m := &mo_file.Metadata{}
//	if err := r.ModelWithPath(m, "entry"); err != nil {
//		t.Error(err)
//	} else if m.Name() != "Prime_Numbers.txt" {
//		t.Error(m)
//	}
//}
//
//
