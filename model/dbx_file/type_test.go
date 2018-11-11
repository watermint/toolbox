package dbx_file

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"go.uber.org/zap"
	"testing"
)

const (
	listFolderResponse = `{
    "entries": [
        {
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
        {
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
        },
    	{
      	".tag": "deleted",
      	"name": "Get Started with Dropbox.pdf",
      	"path_lower": "/get started with dropbox.pdf",
    	  "path_display": "/Get Started with Dropbox.pdf"
	    },
    ],
    "cursor": "ZtkX9_EHj3x7PMkVuFIhwKYXEpwpLwyxp9vMKomUhllil9q7eWiAu",
    "has_more": false
}`
)

func TestEntryParser_Parse(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
	}

	ep := EntryParser{
		Logger: log,
		OnError: func(annotation dbx_api.ErrorAnnotation) bool {
			t.Error(annotation.UserMessage())
			return false
		},
		OnFile: func(file *File) bool {
			log.Info("onFile", zap.Any("file", file))
			return true
		},
		OnFolder: func(folder *Folder) bool {
			log.Info("onFolder", zap.Any("folder", folder))
			return true
		},
		OnDelete: func(deleted *Deleted) bool {
			log.Info("onDeleted", zap.Any("deleted", deleted))
			return true
		},
	}

	entries := gjson.Get(listFolderResponse, "entries")
	if !entries.IsArray() {
		t.Error("invalid test data")
	}
	for _, e := range entries.Array() {
		ep.Parse(e)
	}
}
