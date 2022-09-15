package mo_filerequest

import (
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"testing"
)

func TestFileRequest(t *testing.T) {
	j := `{
            "id": "oaCAVmEyrqYnkZX9955Y",
            "url": "https://www.dropbox.com/request/oaCAVmEyrqYnkZX9955Y",
            "title": "Homework submission",
            "created": "2015-10-05T17:00:00Z",
            "is_open": true,
            "file_count": 3,
            "destination": "/File Requests/Homework",
            "deadline": {
                "deadline": "2020-10-12T17:00:00Z",
                "allow_late_uploads": {
                    ".tag": "seven_days"
                }
            }
        }`
	fr := &FileRequest{}
	if err := api_parser.ParseModelString(fr, j); err != nil {
		t.Error(err)
	}
	if fr.Id != "oaCAVmEyrqYnkZX9955Y" {
		t.Error("invalid")
	}
	if fr.Url != "https://www.dropbox.com/request/oaCAVmEyrqYnkZX9955Y" {
		t.Error("invalid")
	}
	if fr.Title != "Homework submission" {
		t.Error("invalid")
	}
	if fr.Created != "2015-10-05T17:00:00Z" {
		t.Error("invalid")
	}
	if !fr.IsOpen {
		t.Error("invalid")
	}
	if fr.FileCount != 3 {
		t.Error("invalid")
	}
	if fr.Destination != "/File Requests/Homework" {
		t.Error("invalid")
	}
	if fr.Deadline != "2020-10-12T17:00:00Z" {
		t.Error("invalid")
	}
	if fr.DeadlineAllowLateUploads != "seven_days" {
		t.Error("invalid")
	}
}
