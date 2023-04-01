package mo_user

import "encoding/json"

type User struct {
	Raw    json.RawMessage
	Id     string `json:"id" path:"id""`
	Handle string `json:"handle" path:"handle"`
	ImgUrl string `json:"img_url" path:"img_url"`
	Email  string `json:"email" path:"email"`
}
