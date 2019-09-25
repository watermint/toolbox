package app_report_column

import (
	"encoding/json"
	app2 "github.com/watermint/toolbox/legacy/app"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestColumnMarshaller_Row2(t *testing.T) {
	ec := app2.NewExecContextForTest()

	type Name3 struct {
		GivenName string
		Surname   string
	}
	type Zip3 struct {
		Major string
		Minor string
	}
	type Address3 struct {
		Country string
		City    string
		Zip     *Zip3
	}
	type Member3 struct {
		Id      int
		Name    Name3
		Email   string
		Address Address3
		Raw     []byte
		Json    json.RawMessage
	}

	m3 := Member3{
		Id: 123,
		Name: Name3{
			GivenName: "one",
			Surname:   "two-three",
		},
		Email: "123@example.com",
		Address: Address3{
			Country: "Japan",
			City:    "Tokyo",
			Zip: &Zip3{
				Major: "100",
				Minor: "0000",
			},
		},
	}

	ec.Log().Debug("nested - no pointer")
	cz3 := NewRow(m3, ec)
	cols3 := cz3.Header()

	expectedCols3 := []string{
		"Id", "Name.GivenName", "Name.Surname", "Email",
		"Address.Country", "Address.City",
		"Address.Zip.Major", "Address.Zip.Minor",
	}
	expectedVals3 := []string{
		"123",
		"one",
		"two-three",
		"123@example.com",
		"Japan",
		"Tokyo",
		"100",
		"0000",
	}
	vals3 := cz3.ValuesAsString(m3)
	ec.Log().Info("cols3", zap.Strings("cols", cols3))
	ec.Log().Info("vals3", zap.Strings("vals", vals3))

	if !reflect.DeepEqual(expectedCols3, cols3) {
		t.Error("cols3 didn't match")
	}
	if !reflect.DeepEqual(expectedVals3, vals3) {
		t.Error("vals3 didn't match")
	}

	m3s := Member3{
		Id: 123,
		Name: Name3{
			GivenName: "one",
			Surname:   "two-three",
		},
		Email: "123@example.com",
		Address: Address3{
			Country: "Japan",
			City:    "Tokyo",
			Zip:     nil,
		},
	}
	expectedVals3s := []string{
		"123",
		"one",
		"two-three",
		"123@example.com",
		"Japan",
		"Tokyo",
		"",
		"",
	}

	cz3s := NewRow(m3s, ec)
	cols3s := cz3s.Header()
	vals3s := cz3s.ValuesAsString(m3s)
	ec.Log().Info("cols3", zap.Strings("cols", cols3s))
	ec.Log().Info("vals3", zap.Strings("vals", vals3s))

	if !reflect.DeepEqual(expectedCols3, cols3s) {
		t.Error("cols3 didn't match")
	}
	if !reflect.DeepEqual(expectedVals3s, vals3s) {
		t.Error("vals3 didn't match")
	}

	vals3s = cz3s.ValuesAsString(m3s)
	if !reflect.DeepEqual(expectedVals3s, vals3s) {
		t.Error("vals3 didn't match")
	}
}
