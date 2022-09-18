package sv_sharedlink

import "testing"

func TestToDownloadUrl(t *testing.T) {
	link1dl0 := "https://www.dropbox.com/s/a1b2c3d4ef5gh6/example.docx?dl=0"
	link1dl1, err := ToDownloadUrl(link1dl0)
	if err != nil {
		t.Error(err)
	}
	if link1dl1 != "https://www.dropbox.com/s/a1b2c3d4ef5gh6/example.docx?dl=1" {
		t.Error(link1dl1)
	}

	link2dl0 := "https://www.dropbox.com/scl/fi/a1b2c3d4ef5gh6/X-2020-01.xlsx?dl=0&rlkey=a1b2c3d4ef5gh6"
	link2dl1, err := ToDownloadUrl(link2dl0)
	if err != nil {
		t.Error(err)
	}
	if link2dl1 != "https://www.dropbox.com/scl/fi/a1b2c3d4ef5gh6/X-2020-01.xlsx?dl=1&rlkey=a1b2c3d4ef5gh6" {
		t.Error(link2dl1)
	}
}
