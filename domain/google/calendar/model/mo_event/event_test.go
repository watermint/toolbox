package mo_event

import (
	"github.com/watermint/toolbox/infra/api/api_parser"
	"testing"
)

const (
	eventSampleJson = `{
   "kind": "calendar#event",
   "etag": "\"xxxxxxxxxxxxxxxx\"",
   "id": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
   "status": "confirmed",
   "htmlLink": "https://www.google.com/calendar/event?eid=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
   "created": "2017-06-10T06:14:54.000Z",
   "updated": "2017-07-11T04:21:06.086Z",
   "summary": "xxxx xx xxxxx",
   "description": "xx xxx xxxxxxxx xxxxxxxxxxx xxx xxxxxxxxxxxxx xxxxxxx xxxxxx xxxx xxxx xxx, xxx xxx xxxxxxxx xxxxxx xxxxxxxx xxx. xxxxx://x.xx/xxxxxxxx",
   "location": "xxxxxxxxxxxxxxxxxx",
   "creator": {
    "email": "xxxxxxxx@xxxxxxxxx.xxx",
    "self": true
   },
   "organizer": {
    "email": "unknownorganizer@calendar.google.com"
   },
   "start": {
    "date": "2017-07-13"
   },
   "end": {
    "date": "2017-07-15"
   },
   "transparency": "transparent",
   "visibility": "private",
   "iCalUID": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
   "sequence": 0,
   "attendees": [
    {
     "email": "xxxxxxxx@xxxxxxxxx.xxx",
     "self": true,
     "responseStatus": "accepted"
    }
   ],
   "guestsCanInviteOthers": false,
   "privateCopy": true,
   "reminders": {
    "useDefault": false
   },
   "eventType": "default"
  }
`
)

func TestEvent(t *testing.T) {
	e := &Event{}
	err := api_parser.ParseModelString(e, eventSampleJson)
	if err != nil {
		t.Error(err)
	}
	if e.Id != "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
		t.Error(e.Id)
	}
	if e.StartDate != "2017-07-13" {
		t.Error(e.StartDate)
	}
	if e.EndDate != "2017-07-15" {
		t.Error(e.EndDate)
	}
}
