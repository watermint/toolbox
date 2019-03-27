package mo_member_quota

import "encoding/json"

type Quota struct {
	Raw          json.RawMessage
	TeamMemberId string `path:"user.team_member_id"`
	Quota        int    `path:"quota_gb"`
}

func (z *Quota) IsUnlimited() bool {
	return z.Quota == 0
}
