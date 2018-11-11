package dbx_group

import (
	"github.com/watermint/toolbox/model/dbx_api"
	"go.uber.org/zap"
)

func GroupMembers(ctx *dbx_api.Context, log *zap.Logger, handler func(annotation dbx_api.ErrorAnnotation) bool) map[string][]*GroupMember {
	groups := make(map[string][]*GroupMember)

	log.Debug("Expand group")
	gl := GroupList{
		OnError: handler,
		OnEntry: func(group *Group) bool {
			log.Debug("onEntry",
				zap.String("group_id", group.GroupId),
				zap.String("group_name", group.GroupName),
			)

			gml := GroupMemberList{
				OnError: handler,
				OnEntry: func(gm *GroupMember) bool {

					if g, ok := groups[group.GroupId]; ok {
						g = append(g, gm)
						groups[group.GroupId] = g

						log.Debug("onEntry",
							zap.String("group_id", group.GroupId),
							zap.Int("group_members", len(g)),
						)
					} else {
						g = make([]*GroupMember, 1)
						g[0] = gm
						groups[group.GroupId] = g

						log.Debug("onEntry",
							zap.String("group_id", group.GroupId),
							zap.Int("group_members", len(g)),
						)
					}

					return true
				},
			}
			gml.List(ctx, group)
			return true
		},
	}
	if !gl.List(ctx) {
		return nil
	}

	for k, v := range groups {
		log.Debug("Group summary",
			zap.String("group_id", k),
			zap.Int("member_count", len(v)),
		)
	}

	return groups
}
