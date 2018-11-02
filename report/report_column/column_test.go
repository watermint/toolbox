package report_column

import (
	"encoding/json"
	"go.uber.org/zap"
	"reflect"
	"testing"
)

func TestColumnMarshaller_Row2(t *testing.T) {
	log, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
	}

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

	log.Debug("nested - no pointer")
	cz := ColumnZ{}

	cols3 := cz.Header(m3)

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
	vals3 := cz.Values(cols3, m3)
	log.Info("cols3", zap.Strings("cols", cols3))
	log.Info("vals3", zap.Strings("vals", vals3))

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

	cols3s := cz.Header(m3s)
	vals3s := cz.Values(cols3s, m3s)
	log.Info("cols3", zap.Strings("cols", cols3s))
	log.Info("vals3", zap.Strings("vals", vals3s))

	if !reflect.DeepEqual(expectedCols3, cols3s) {
		t.Error("cols3 didn't match")
	}
	if !reflect.DeepEqual(expectedVals3s, vals3s) {
		t.Error("vals3 didn't match")
	}

}
