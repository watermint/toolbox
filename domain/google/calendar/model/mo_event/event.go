package mo_event

import "encoding/json"

type Event struct {
	Raw           json.RawMessage
	Id            string `path:"id" json:"id"`
	Status        string `path:"status" json:"status"`
	Location      string `path:"location" json:"location"`
	StartDateTime string `path:"start.dateTime" json:"start_date_time"`
	EndDateTime   string `path:"end.dateTime" json:"end_date_time"`
	Summary       string `path:"summary" json:"summary"`
	Description   string `path:"description" json:"description"`
}
