package content

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
)

type MemberScannerWorker struct {
	Control             app_control.Control
	Context             api_context.DropboxApiContext
	Member              *mo_member.Member
	TeamOwnedNamespaces map[string]bool // namespace Id -> true
	Scanner             ScanNamespace
}

func (z *MemberScannerWorker) Exec() error {
	l := z.Context.Log().With(zap.String("member", z.Member.Email))
	ui := z.Control.UI()

	l.Debug("Scanning member")
	ui.Progress(MScanMetadata.ProgressScanMember.With("Email", z.Member.Email))

	folders, err := sv_sharedfolder.New(z.Context).List()
	if err != nil {
		return err
	}

	for _, f := range folders {
		if z.TeamOwnedNamespaces[f.SharedFolderId] {
			l.Debug("Skip team owned folder", zap.Any("folder", f))
			continue
		}
		z.Scanner.Scan(z.Control, z.Context, f.Name, f.SharedFolderId)
	}
	return nil
}