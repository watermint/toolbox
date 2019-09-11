package mo_activity

import "encoding/json"

type Event struct {
	Raw           json.RawMessage
	Timestamp     string `path:"timestamp" json:"timestamp"`
	EventCategory string `path:"event_category.\\.tag" json:"event_category"`
	EventType     string `path:"event_type.\\.tag" json:"event_type"`
	EventTypeDesc string `path:"event_type.description" json:"event_type_desc"`
}
