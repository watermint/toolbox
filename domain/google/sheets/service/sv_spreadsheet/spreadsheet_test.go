package sv_spreadsheet

import "testing"

func TestIsValidSpreadsheetId(t *testing.T) {
	if x := isValidSpreadsheetId("1Kokm29I8he7mn7p9iE66veGwj8qyEpTtixv1ulbl7Tw"); !x {
		t.Error(x)
	}
	if x := isValidSpreadsheetId("    "); !x {
		t.Error(x)
	}
}
