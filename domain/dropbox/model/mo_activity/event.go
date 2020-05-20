package mo_activity

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"time"
)

type Event struct {
	Raw           json.RawMessage
	Timestamp     string `path:"timestamp" json:"timestamp"`
	EventCategory string `path:"event_category.\\.tag" json:"event_category"`
	EventType     string `path:"event_type.\\.tag" json:"event_type"`
	EventTypeDesc string `path:"event_type.description" json:"event_type_desc"`
}

func (z *Event) Compatible() *Compatible {
	l := esl.Default()
	c := &Compatible{}
	if err := api_parser.ParseModelRaw(c, z.Raw); err != nil {
		l.Debug("Unable to parse model", esl.Error(err))
		return nil
	}

	g := gjson.ParseBytes(z.Raw)
	dts := g.Get("timestamp").String()
	if ts, err := time.Parse("2006-01-02T15:04:05Z", dts); err != nil {
		l.Debug("Unable to parse time", esl.Error(err))
		return nil
	} else {
		c.Timestamp = ts.Local().Format("2006-01-02 15:04:05")
	}

	c.Participants = g.Get("participants").Raw
	c.Context = g.Get("context").Raw
	c.Assets = g.Get("assets").Raw
	c.OtherInfo = g.Get("details").Raw

	actorPaths := []string{"user", "admin"}
	for _, ap := range actorPaths {
		a := g.Get("actor." + ap)
		if a.Exists() {
			c.Member = a.Get("display_name").String()
			c.MemberEmail = a.Get("email").String()
		}
	}

	return c
}

// Dropbox Business activities report compatible model
type Compatible struct {
	Raw                   json.RawMessage
	Timestamp             string `json:"timestamp"`
	Member                string `json:"member"`
	MemberEmail           string `json:"member_email"`
	EventType             string `path:"event_type.description" json:"event_type"`
	Category              string `path:"event_category.\\.tag" json:"category"`
	AccessMethod          string `path:"origin.access_method.end_user.\\.tag" json:"access_method"`
	IpAddress             string `path:"origin.geo_location.ip_address" json:"ip_address"`
	Country               string `path:"origin.geo_location.country" json:"country"`
	City                  string `path:"origin.geo_location.city" json:"city"`
	InvolveNonTeamMembers bool   `path:"involve_non_team_member" json:"involve_non_team_members"`
	Participants          string `json:"participants"`
	Context               string `json:"context"`
	Assets                string `json:"assets"`
	OtherInfo             string `json:"other_info"`
}
