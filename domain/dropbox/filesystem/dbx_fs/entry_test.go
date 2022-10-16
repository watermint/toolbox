package dbx_fs

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestNewEntry(t *testing.T) {
	rawJson := `{
      ".tag": "file",
      "name": "メモ.txt",
      "path_lower": "/レポート/メモ.txt",
      "path_display": "/レポート/メモ.txt",
      "id": "id:xxxxxxxxxxxxxxxxxxxxxx",
      "client_modified": "2018-12-27T23:03:29Z",
      "server_modified": "2018-12-27T23:03:29Z",
      "rev": "xxxxxxxxxxx",
      "size": 22,
      "is_downloadable": true,
      "content_hash": "6654c51cd45a62c5eb5110fb683dd21b8d93e35ee9090fdcc775df6a1b274c45"
    }`

	dbxEntry := &mo_file.Metadata{}
	if err := api_parser.ParseModelString(dbxEntry, rawJson); err != nil {
		t.Error(err)
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		fsEntry := NewEntry(dbxEntry)

		// round trip test
		fsEntryData := fsEntry.AsData()

		fsEntryDataJson, jsErr := json.Marshal(fsEntryData)
		if jsErr != nil {
			t.Error(jsErr)
			return
		}
		fsEntryData2 := &es_filesystem.EntryData{}
		if jsErr := json.Unmarshal(fsEntryDataJson, fsEntryData2); jsErr != nil {
			t.Error(jsErr)
			return
		}

		dbxFs := NewFileSystem(dbx_client_impl.NewMock("mock", ctl))

		dbxEntry1, fsErr := dbxFs.Entry(fsEntryData)
		if fsErr != nil {
			t.Error(fsErr)
			return
		}

		dbxEntry2, fsErr := dbxFs.Entry(*fsEntryData2)
		if fsErr != nil {
			t.Error(fsErr)
			return
		}

		if dbxEntry1.Name() != dbxEntry.Name() {
			t.Error(dbxEntry1)
		}
		if dbxEntry2.Name() != dbxEntry.Name() {
			t.Error(dbxEntry1)
		}
	})
}
