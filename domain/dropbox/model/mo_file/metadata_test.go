package mo_file

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestMetadataFile(t *testing.T) {
	fileJson := `{
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
        }`

	entry := Metadata{}
	if err := api_parser.ParseModelString(&entry, fileJson); err != nil {
		t.Error(err)
	}
	if entry.Tag() != "file" {
		t.Error("invalid")
	}
	if entry.Path().Path() != "/Homework/math/Prime_Numbers.txt" {
		t.Error("invalid")
	}
	if entry.PathDisplay() != "/Homework/math/Prime_Numbers.txt" {
		t.Error("invalid")
	}
	if entry.PathLower() != "/homework/math/prime_numbers.txt" {
		t.Error("invalid")
	}
	if f, e := entry.File(); !e ||
		f.Size != 7212 ||
		f.Revision != "a1c10ce0dd78" ||
		f.ContentHash != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Error("invalid")
	}
	if f, e := entry.Folder(); e || f != nil {
		t.Error("invalid")
	}
	if f, e := entry.Deleted(); e || f != nil {
		t.Error("invalid")
	}
}

func TestMetadataFolder(t *testing.T) {
	folderJson := `{
            ".tag": "folder",
            "name": "math",
            "id": "id:a4ayc_80_OEAAAAAAAAAXz",
            "path_lower": "/homework/math",
            "path_display": "/Homework/math",
            "sharing_info": {
                "read_only": false,
                "parent_shared_folder_id": "84528192421",
                "traverse_only": false,
                "no_access": false
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
            ]
        }`

	entry := Metadata{}
	if err := api_parser.ParseModelString(&entry, folderJson); err != nil {
		t.Error(err)
	}
	if entry.Tag() != "folder" {
		t.Error("invalid")
	}
	if entry.Path().Path() != "/Homework/math" {
		t.Error("invalid")
	}
	if entry.PathDisplay() != "/Homework/math" {
		t.Error("invalid")
	}
	if entry.PathLower() != "/homework/math" {
		t.Error("invalid")
	}
	if f, e := entry.Folder(); !e || f == nil {
		t.Error("invalid")
	}
	if f, e := entry.File(); e || f != nil {
		t.Error("invalid")
	}
	if f, e := entry.Deleted(); e || f != nil {
		t.Error("invalid")
	}
}

func TestMetadataDeleted(t *testing.T) {
	deletedJson := `{
      ".tag": "deleted",
      "name": "プロジェクトG",
      "path_lower": "/プロジェクトg",
      "path_display": "/プロジェクトG"
    }`

	entry := Metadata{}
	if err := api_parser.ParseModelString(&entry, deletedJson); err != nil {
		t.Error(err)
	}
	if entry.Tag() != "deleted" {
		t.Error("invalid")
	}
	if entry.Path().Path() != "/プロジェクトG" {
		t.Error("invalid")
	}
	if entry.PathDisplay() != "/プロジェクトG" {
		t.Error("invalid")
	}
	if entry.PathLower() != "/プロジェクトg" {
		t.Error("invalid")
	}
	if f, e := entry.Deleted(); !e ||
		f.Name() != "プロジェクトG" {
		t.Error("invalid")
	}
	if f, e := entry.Folder(); e || f != nil {
		t.Error("invalid")
	}
	if f, e := entry.File(); e || f != nil {
		t.Error("invalid")
	}
}
