package mo_image

import "encoding/json"

type Exif struct {
	Raw              json.RawMessage
	DateTimeOriginal string `json:"date_time_original" path:"DateTimeOriginal"`
	DateTime         string `json:"date_time" path:"DateTime"`
	Make             string `json:"make" path:"Make"`
	Model            string `json:"model" path:"Model"`
}
