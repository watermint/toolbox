package mo_usage

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

type Usage struct {
	Raw                               json.RawMessage
	Used                              uint64 `path:"used" json:"used"`
	Allocation                        string `path:"allocation.\\.tag" json:"allocation"`
	Allocated                         uint64 `path:"allocation.allocated" json:"allocated"`
	TeamUsed                          uint64 `path:"allocation.used" json:"team_used"`
	TeamUserWithinTeamSpaceAllocated  uint64 `path:"allocation.user_within_team_space_allocated" json:"team_user_within_team_space_allocated"`
	TeamUserWithinTeamSpaceLimitType  string `path:"allocation.user_within_team_space_limit_type.\\.tag" json:"team_user_within_team_space_limit_type"`
	TeamUserWithinTeamSpaceUsedCached uint64 `path:"allocation.user_within_team_space_used_cached" json:"team_user_within_team_space_used_cached"`
}

type MemberUsage struct {
	Raw        json.RawMessage
	Email      string  `path:"member.profile.email" json:"email"`
	UsedGB     float64 `path:"usage.used" json:"used_gb"`
	UsedBytes  uint64  `path:"usage.used" json:"used_bytes"`
	Allocation string  `path:"usage.allocation.\\.tag" json:"allocation"`
	Allocated  uint64  `path:"usage.allocation.allocated" json:"allocated"`
	//TeamUsed                          uint64  `path:"usage.allocation.used" json:"team_used"`
	//TeamUserWithinTeamSpaceAllocated  uint64  `path:"usage.allocation.user_within_team_space_allocated" json:"team_user_within_team_space_allocated"`
	//TeamUserWithinTeamSpaceLimitType  string  `path:"usage.allocation.user_within_team_space_limit_type.\\.tag" json:"team_user_within_team_space_limit_type"`
	//TeamUserWithinTeamSpaceUsedCached uint64  `path:"usage.allocation.user_within_team_space_used_cached" json:"team_user_within_team_space_used_cached"`
}

func NewMemberUsage(member *mo_member.Member, usage *Usage) (mu *MemberUsage) {
	raws := make(map[string]json.RawMessage)
	raws["member"] = member.Raw
	raws["usage"] = usage.Raw
	raw := api_parser.CombineRaw(raws)

	mu = &MemberUsage{}
	if err := api_parser.ParseModelRaw(mu, raw); err != nil {
		app_root.Log().Warn("unexpected data format", zap.Error(err))
		// return empty
		return mu
	}
	mu.UsedGB = float64(mu.UsedBytes) / 1_073_741_824.0
	return mu
}
