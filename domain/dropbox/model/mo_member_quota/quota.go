package mo_member_quota

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/api/api_parser"
)

type Quota struct {
	Raw          json.RawMessage
	TeamMemberId string `path:"user.team_member_id" json:"team_member_id"`
	Quota        int    `path:"quota_gb" json:"quota"`
}

func (z *Quota) IsUnlimited() bool {
	return z.Quota == 0
}

type MemberQuota struct {
	Raw json.RawMessage
	//TeamMemberId    string `path:"member.profile.team_member_id" json:"team_member_id"`
	Email string `path:"member.profile.email" json:"email"`
	//EmailVerified   bool   `path:"member.profile.email_verified" json:"email_verified"`
	//Status          string `path:"member.profile.status.\\.tag" json:"status"`
	//GivenName       string `path:"member.profile.name.given_name" json:"given_name"`
	//Surname         string `path:"member.profile.name.surname" json:"surname"`
	//FamiliarName    string `path:"member.profile.name.familiar_name" json:"familiar_name"`
	//DisplayName     string `path:"member.profile.name.display_name" json:"display_name"`
	//AbbreviatedName string `path:"member.profile.name.abbreviated_name" json:"abbreviated_name"`
	//MemberFolderId  string `path:"member.profile.member_folder_id" json:"member_folder_id"`
	//ExternalId      string `path:"member.profile.external_id" json:"external_id"`
	//AccountId       string `path:"member.profile.account_id" json:"account_id"`
	//PersistentId    string `path:"member.profile.persistent_id" json:"persistent_id"`
	//JoinedOn        string `path:"member.profile.joined_on" json:"joined_on"`
	//Role            string `path:"member.role.\\.tag" json:"role"`
	//Tag             string `path:"member.\\.tag" json:"tag"`
	Quota int `path:"quota.quota_gb" json:"quota"`
}

func NewMemberQuota(member *mo_member.Member, quota *Quota) (mq *MemberQuota) {
	raws := make(map[string]json.RawMessage)
	raws["member"] = member.Raw
	raws["quota"] = quota.Raw
	raw := api_parser.CombineRaw(raws)

	mq = &MemberQuota{}
	if err := api_parser.ParseModelRaw(mq, raw); err != nil {
		es_log.Default().Warn("unexpected data format", es_log.Error(err))
		// return empty
		return mq
	}
	return mq
}
