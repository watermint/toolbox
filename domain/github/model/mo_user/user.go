package mo_user

import "encoding/json"

type User struct {
	Raw   json.RawMessage
	Id    string `path:"id" json:"id"`
	Login string `path:"login" json:"login"`
	Name  string `path:"name" json:"name"`
	Url   string `path:"url" json:"url"`
}
