package to_spreadsheet

import "testing"

func TestIsValidSpreadsheetId(t *testing.T) {
	if x := IsValidSpreadsheetId("1Kokm29I8he7mn7p9iE66veGwj8qyEpTtixv1ulbl7Tw"); !x {
		t.Error(x)
	}
	if x := IsValidSpreadsheetId("    "); x {
		t.Error(x)
	}
	if x := IsValidSpreadsheetId("&redirect_uri=http%3A%2F%2Fexample.com"); x {
		t.Error(x)
	}
}
