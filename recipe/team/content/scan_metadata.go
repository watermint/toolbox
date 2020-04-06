package content

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/kvs/kv_kvs"
	"github.com/watermint/toolbox/infra/kvs/kv_storage"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
	"go.uber.org/zap"
)

type ScanNamespaceMetadata struct {
	metadata kv_storage.Storage
	queue    rc_worker.Queue
}

func (z *ScanNamespaceMetadata) Scan(ctl app_control.Control, ctx dbx_context.Context, namespaceName string, namespaceId string) {
	z.queue.Enqueue(&MetadataScannerWorker{
		Control:       ctl,
		Context:       ctx,
		NamespaceName: namespaceName,
		NamespaceId:   namespaceId,
		Metadata:      z.metadata,
	})
}

type MetadataScannerWorker struct {
	Control       app_control.Control
	Context       dbx_context.Context
	NamespaceName string
	NamespaceId   string
	Metadata      kv_storage.Storage
}

func (z *MetadataScannerWorker) Exec() error {
	l := z.Context.Log().With(zap.String("namespaceId", z.NamespaceId), zap.String("namespaceName", z.NamespaceName))
	ui := z.Control.UI()

	l.Debug("Scanning metadata")
	ui.Progress(MScanMetadata.ProgressScanNamespaceMetadata.With("Name", z.NamespaceName))

	m, err := sv_sharedfolder.New(z.Context).Resolve(z.NamespaceId)
	if err != nil {
		l.Debug("Unable to retrieve metadata", zap.Error(err))
		return err
	}

	return z.Metadata.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(z.NamespaceId, m)
	})
}
