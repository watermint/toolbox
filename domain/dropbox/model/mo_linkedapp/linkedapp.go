package mo_linkedapp

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type LinkedApp struct {
	Raw          json.RawMessage
	TeamMemberId string `path:"-" json:"team_member_id"`
	AppId        string `path:"app_id" json:"app_id"`
	AppName      string `path:"app_name" json:"app_name"`
	IsAppFolder  bool   `path:"is_app_folder" json:"is_app_folder"`
	Publisher    string `path:"publisher" json:"publisher"`
	PublisherUrl string `path:"publisher_url" json:"publisher_url"`
	Linked       string `path:"linked" json:"linked"`
}

type MemberLinkedApp struct {
	Raw             json.RawMessage
	TeamMemberId    string `path:"profile.team_member_id" json:"team_member_id"`
	Email           string `path:"profile.email" json:"email"`
	Status          string `path:"profile.status.\\.tag" json:"status"`
	GivenName       string `path:"profile.name.given_name" json:"given_name"`
	Surname         string `path:"profile.name.surname" json:"surname"`
	FamiliarName    string `path:"profile.name.familiar_name" json:"familiar_name"`
	DisplayName     string `path:"profile.name.display_name" json:"display_name"`
	AbbreviatedName string `path:"profile.name.abbreviated_name" json:"abbreviated_name"`
	ExternalId      string `path:"profile.external_id" json:"external_id"`
	AccountId       string `path:"profile.account_id" json:"account_id"`
	AppId           string `path:"linked_app.app_id" json:"app_id"`
	AppName         string `path:"linked_app.app_name" json:"app_name"`
	IsAppFolder     bool   `path:"linked_app.is_app_folder" json:"is_app_folder"`
	Publisher       string `path:"linked_app.publisher" json:"publisher"`
	PublisherUrl    string `path:"linked_app.publisher_url" json:"publisher_url"`
	Linked          string `path:"linked_app.linked" json:"linked"`
}

func NewMemberLinkedApp(member *mo_member.Member, linkedApp *LinkedApp) (mla *MemberLinkedApp) {
	prof := gjson.ParseBytes(member.Raw).Get("profile")
	raws := make(map[string]json.RawMessage)
	if prof.Raw != "" {
		raws["profile"] = json.RawMessage(prof.Raw)
	} else {
		raws["profile"] = nil
	}
	raws["linked_app"] = linkedApp.Raw
	raw := api_parser.CombineRaw(raws)

	mla = &MemberLinkedApp{}
	if err := api_parser.ParseModelRaw(mla, raw); err != nil {
		esl.Default().Error("unable to parse", esl.Error(err))
	}
	return mla
}
