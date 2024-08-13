package mo_event

import "encoding/json"

type Event struct {
	Raw           json.RawMessage
	Id            string `path:"id" json:"id"`
	Status        string `path:"status" json:"status"`
	Location      string `path:"location" json:"location"`
	StartDate     string `path:"start.date" json:"start_date"`
	EndDate       string `path:"end.date" json:"end_date"`
	StartDateTime string `path:"start.dateTime" json:"start_date_time"`
	EndDateTime   string `path:"end.dateTime" json:"end_date_time"`
	Summary       string `path:"summary" json:"summary"`
	Description   string `path:"description" json:"description"`
}
