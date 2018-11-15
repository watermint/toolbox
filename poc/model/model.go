package model

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"reflect"
)

/* PoC for new data structure to reduce redundancy.
 *
 */

type Raw struct {
	Result *gjson.Result   `json:"-"`
	Json   json.RawMessage `json:"-"`
}

func (z *Raw) Get(path string) (string, bool) {
	if z.Result == nil {
		if !gjson.ValidBytes(z.Json) {
			return "", false
		}

		r := gjson.ParseBytes(z.Json)
		z.Result = &r
	}
	r := z.Result.Get(path)
	return r.String(), r.Exists()
}

func (z *Raw) GetTag(path string) (string, bool) {
	dotTag := "\\.tag"
	if path == "" {
		return z.Get(dotTag)
	} else {
		return z.Get(path + "." + dotTag)
	}
}

func (z *Raw) Tag() (string, bool) {
	return z.GetTag("")
}

type Profile struct {
	Raw
}

func (z *Profile) Email() (string, bool) {
	return z.Get("email")
}
func (z *Profile) TeamMemberId() (string, bool) {
	return z.Get("team_member_id")
}
func (z *Profile) Status() (string, bool) {
	return z.GetTag("status")
}

func ColumnHeader(v interface{}) []string {
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	vt = vv.Type()
	fmt.Println(vt.Name())
	fmt.Println(vt.NumMethod())
	//fmt.Println(vt.NumField())

	nm := vt.NumMethod()
	for i := 0; i < nm; i++ {
		m := vt.Method(i)

		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}
	return nil
}
