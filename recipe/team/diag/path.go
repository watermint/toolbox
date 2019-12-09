package diag

import (
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"strings"
)

type PathVO struct {
	PeerFile    app_conn.ConnBusinessFile
	PeerAudit   app_conn.ConnBusinessFile
	MemberEmail string
	Path        string
}

type Path struct {
}

func (z *Path) Hidden() {
}

func (z *Path) Console() {
}

func (z *Path) Requirement() app_vo.ValueObject {
	return &PathVO{}
}

func (z *Path) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*PathVO)
	l := k.Log()
	cf, err := vo.PeerFile.Connect(k.Control())
	if err != nil {
		return err
	}

	member, err := sv_member.New(cf).ResolveByEmail(vo.MemberEmail)
	if err != nil {
		return err
	}
	l = l.With(zap.String("TeamMemberId", member.TeamMemberId))

	cfm := cf.AsMemberId(member.TeamMemberId)

	mounts, err := sv_sharedfolder_mount.New(cfm).List()
	if err != nil {
		return err
	}

	var pathMount *mo_sharedfolder.SharedFolder
	pathRel := ""
	pathLower := strings.ToLower(vo.Path)
	for _, mount := range mounts {
		if mount.PathLower == "" {
			continue
		}
		ll := l.With(zap.String("mountPathLower", mount.PathLower), zap.String("mountNamespaceId", mount.SharedFolderId))
		ll.Debug("Scan")
		if strings.HasPrefix(pathLower, mount.PathLower) {
			pr, err := ut_filepath.Rel(mount.PathLower, pathLower)
			if err != nil {
				ll.Debug("Unable to calc rel path", zap.Error(err))
				continue
			}
			switch pr {
			case ".":
				ll.Debug("Matched to path")
				pathRel = ""
				pathMount = mount
				break
			default:
				if len(pathRel) < len(mount.PathLower) {
					ll.Debug("Select path")
					pathRel = pr
					pathMount = mount
				} else {
					ll.Debug("Skip mount")
				}
			}
		}
	}
	l.Debug("Mount", zap.Any("pathMount", pathMount))

	return nil
}

func (z *Path) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}

func (z *Path) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}
