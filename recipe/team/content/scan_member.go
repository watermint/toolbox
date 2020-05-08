package content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type MemberScannerWorker struct {
	Control             app_control.Control
	Context             dbx_context.Context
	Member              *mo_member.Member
	TeamOwnedNamespaces map[string]bool // namespace Id -> true
	Scanner             ScanNamespace
}

func (z *MemberScannerWorker) Exec() error {
	l := z.Context.Log().With(esl.String("member", z.Member.Email))
	ui := z.Control.UI()

	l.Debug("Scanning member")
	ui.Progress(MScanMetadata.ProgressScanMember.With("Email", z.Member.Email))

	folders, err := sv_sharedfolder.New(z.Context).List()
	if err != nil {
		return err
	}

	for _, f := range folders {
		if z.TeamOwnedNamespaces[f.SharedFolderId] {
			l.Debug("Skip team owned folder", esl.Any("folder", f))
			continue
		}
		z.Scanner.Scan(z.Control, z.Context, f.Name, f.SharedFolderId)
	}
	return nil
}
