package mo_user

import "encoding/json"

type User struct {
	Raw    json.RawMessage `json:"-"`
	Id     string          `json:"id" path:"id"`
	TeamId string          `json:"team_id" path:"team_id"`
	Name   string          `json:"name" path:"name"`
}
