package mo_linkedapp

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/atbx/app_root"
	"github.com/watermint/toolbox/domain/infra/api_parser"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"go.uber.org/zap"
)

type LinkedApp struct {
	Raw          json.RawMessage
	TeamMemberId string `path:"-"`
	AppId        string `path:"app_id"`
	AppName      string `path:"app_name"`
	IsAppFolder  bool   `path:"is_app_folder"`
	Publisher    string `path:"publisher"`
	PublisherUrl string `path:"publisher_url"`
	Linked       string `path:"linked"`
}

type MemberLinkedApp struct {
	Raw             json.RawMessage
	TeamMemberId    string `path:"profile.team_member_id"`
	Email           string `path:"profile.email"`
	Status          string `path:"profile.status.\\.tag"`
	GivenName       string `path:"profile.name.given_name"`
	Surname         string `path:"profile.name.surname"`
	FamiliarName    string `path:"profile.name.familiar_name"`
	DisplayName     string `path:"profile.name.display_name"`
	AbbreviatedName string `path:"profile.name.abbreviated_name"`
	ExternalId      string `path:"profile.external_id"`
	AccountId       string `path:"profile.account_id"`
	AppId           string `path:"linked_app.app_id"`
	AppName         string `path:"linked_app.app_name"`
	IsAppFolder     bool   `path:"linked_app.is_app_folder"`
	Publisher       string `path:"linked_app.publisher"`
	PublisherUrl    string `path:"linked_app.publisher_url"`
	Linked          string `path:"linked_app.linked"`
}

func NewMemberLinkedApp(member *mo_member.Member, linkedApp *LinkedApp) (mla *MemberLinkedApp) {
	prof := gjson.ParseBytes(member.Raw).Get("profile")
	raws := make(map[string]json.RawMessage)
	raws["profile"] = json.RawMessage(prof.Raw)
	raws["linked_app"] = linkedApp.Raw
	raw := api_parser.CombineRaw(raws)

	mla = &MemberLinkedApp{}
	if err := api_parser.ParseModelRaw(mla, raw); err != nil {
		app_root.Log().Error("unable to parse", zap.Error(err))
	}
	return mla
}
