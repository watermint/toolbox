package uc_team_migration

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/domain/usecase/uc_teamfolder_mirror"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"go.uber.org/zap"
	"io/ioutil"
	"path/filepath"
)

func (z *migrationImpl) Resume(opts ...ResumeOpt) (ctx Context, err error) {
	ro := &resumeOpts{}
	for _, o := range opts {
		o(ro)
	}
	ctxImpl := &contextImpl{}

	{
		b, err := ioutil.ReadFile(filepath.Join(ro.storagePath, "context.json"))
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		err = json.Unmarshal(b, ctxImpl)
		if err != nil {
			z.ctxExec.Log().Error("unable to unmarshal context", zap.Error(err))
			return nil, err
		}
	}

	{
		b, err := ioutil.ReadFile(filepath.Join(ro.storagePath, "namespace_member.json"))
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		j := gjson.ParseBytes(b)
		ctxImpl.MapNamespaceMember = make(map[string]map[string]mo_sharedfolder_member.Member)
		if j.Exists() && j.IsObject() {
			for k, ja := range j.Map() {
				if ja.IsObject() {
					var members map[string]mo_sharedfolder_member.Member
					members = make(map[string]mo_sharedfolder_member.Member)
					for _, je := range ja.Map() {
						member := &mo_sharedfolder_member.Metadata{}
						if err := api_parser.ParseModel(member, je.Get("Raw")); err != nil {
							z.log().Error("Unable to parse", zap.Error(err), zap.String("entry", je.Raw))
							return nil, err
						}
						if u, e := member.User(); e {
							members[u.Email] = u
						}
						if g, e := member.Group(); e {
							members[g.GroupId] = g
						}
						if i, e := member.Invitee(); e {
							members[i.InviteeEmail] = i
						}
					}
					ctxImpl.MapNamespaceMember[k] = members
				}
			}
		}
	}

	{
		tb, err := ioutil.ReadFile(filepath.Join(ro.storagePath, "teamfolder_content.json"))
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		tmc, err := uc_teamfolder_mirror.UnmarshalContext(tb)
		if err != nil {
			z.ctxExec.Log().Error("unable to read stored context", zap.Error(err))
			return nil, err
		}
		ctxImpl.ctxTeamFolder = tmc
	}

	ctxImpl.init(ro.ec)
	z.ctxExec.Log().Info("Context restored", zap.String("path", ro.storagePath))
	return ctxImpl, nil
}
