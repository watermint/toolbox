package mo_activity

import "encoding/json"

type Event struct {
	Raw           json.RawMessage
	Timestamp     string `path:"timestamp"`
	EventCategory string `path:"event_category.\\.tag"`
	EventType     string `path:"event_type.\\.tag"`
	EventTypeDesc string `path:"event_type.description"`
}
