package content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
	"go.uber.org/zap"
)

type ScanNamespaceMembership struct {
	membership kv_storage.Storage
	queue      rc_worker.Queue
}

func (z *ScanNamespaceMembership) Scan(ctl app_control.Control, ctx dbx_context.Context, namespaceName string, namespaceId string) {
	z.queue.Enqueue(&NamespaceMemberScannerWorker{
		Control:       ctl,
		Context:       ctx,
		NamespaceName: namespaceName,
		NamespaceId:   namespaceId,
		Membership:    z.membership,
	})
}

type NamespaceMemberScannerWorker struct {
	Control       app_control.Control
	Context       dbx_context.Context
	NamespaceName string
	NamespaceId   string
	Membership    kv_storage.Storage
}

func (z *NamespaceMemberScannerWorker) Exec() error {
	l := z.Context.Log().With(zap.String("namespaceId", z.NamespaceId), zap.String("namespaceName", z.NamespaceName))
	ui := z.Control.UI()

	l.Debug("Scanning membership")
	ui.Progress(MScanMetadata.ProgressScanNamespaceMember.With("Name", z.NamespaceName))

	members, err := sv_sharedfolder_member.NewBySharedFolderId(z.Context, z.NamespaceId).List()
	if err != nil {
		l.Debug("Unable to retrieve membership", zap.Error(err))
		return err
	}

	return z.Membership.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(z.NamespaceId, members)
	})
}
