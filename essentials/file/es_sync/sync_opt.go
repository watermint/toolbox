package es_sync

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
)

const (
	SkipSame   SkipReason = "same"
	SkipExists SkipReason = "exists"
	SkipFilter SkipReason = "filter"
)

type Opts struct {
	syncDelete                  bool
	syncOverwrite               bool
	syncDontCompareTime         bool
	syncDontCompareContent      bool
	listenerCopySuccess         ListenerCopySuccess
	listenerCopyFailure         ListenerCopyFailure
	listenerDeleteSuccess       ListenerDeleteSuccess
	listenerDeleteFailure       ListenerDeleteFailure
	listenerCreateFolderSuccess ListenerCreateFolderSuccess
	listenerCreateFolderFailure ListenerCreateFolderFailure
	listenerSkip                ListenerSkip
	entryNameFilter             mo_filter.Filter
}

func (z Opts) SyncOverwrite() bool {
	return z.syncOverwrite
}

func (z Opts) SyncDelete() bool {
	return z.syncDelete
}

func (z Opts) SyncDontCompareTime() bool {
	return z.syncDontCompareTime
}

func (z Opts) SyncDontCompareContent() bool {
	return z.syncDontCompareContent
}

func (z Opts) OnCreateFolderSuccess(target es_filesystem.Path) {
	if z.listenerCreateFolderSuccess != nil {
		z.listenerCreateFolderSuccess(target)
	}
}

func (z Opts) OnCreateFolderFailure(target es_filesystem.Path, err es_filesystem.FileSystemError) {
	if z.listenerCreateFolderFailure != nil {
		z.listenerCreateFolderFailure(target, err)
	}
}

func (z Opts) OnCopySuccess(source es_filesystem.Entry, target es_filesystem.Entry) {
	if z.listenerCopySuccess != nil {
		z.listenerCopySuccess(source, target)
	}
}

func (z Opts) OnCopyFailure(source es_filesystem.Path, err es_filesystem.FileSystemError) {
	if z.listenerCopyFailure != nil {
		z.listenerCopyFailure(source, err)
	}
}

func (z Opts) OnDeleteSuccess(target es_filesystem.Path) {
	if z.listenerDeleteSuccess != nil {
		z.listenerDeleteSuccess(target)
	}
}

func (z Opts) OnDeleteFailure(target es_filesystem.Path, err es_filesystem.FileSystemError) {
	if z.listenerDeleteFailure != nil {
		z.listenerDeleteFailure(target, err)
	}
}

func (z Opts) OnSkip(reason SkipReason, source es_filesystem.Entry, target es_filesystem.Path) {
	if z.listenerSkip != nil {
		z.listenerSkip(reason, source, target)
	}
}

func (z Opts) Apply(opt []Opt) Opts {
	switch len(opt) {
	case 0:
		return z
	case 1:
		return opt[0](z)
	default:
		return opt[0](z).Apply(opt[1:])
	}
}

type Opt func(o Opts) Opts

type SkipReason string

type ListenerCopySuccess func(source es_filesystem.Entry, target es_filesystem.Entry)
type ListenerCopyFailure func(source es_filesystem.Path, fsErr es_filesystem.FileSystemError)
type ListenerDeleteSuccess func(target es_filesystem.Path)
type ListenerDeleteFailure func(target es_filesystem.Path, fsErr es_filesystem.FileSystemError)
type ListenerSkip func(reason SkipReason, source es_filesystem.Entry, target es_filesystem.Path)
type ListenerCreateFolderSuccess func(target es_filesystem.Path)
type ListenerCreateFolderFailure func(target es_filesystem.Path, fsErr es_filesystem.FileSystemError)

func SyncDelete(enabled bool) Opt {
	return func(o Opts) Opts {
		o.syncDelete = enabled
		return o
	}
}

func SyncOverwrite(enabled bool) Opt {
	return func(o Opts) Opts {
		o.syncOverwrite = enabled
		return o
	}
}

func SyncDontCompareTime(enabled bool) Opt {
	return func(o Opts) Opts {
		o.syncDontCompareTime = enabled
		return o
	}
}

func SyncDontCompareContent(enabled bool) Opt {
	return func(o Opts) Opts {
		o.syncDontCompareContent = enabled
		return o
	}
}

func OnCopySuccess(l ListenerCopySuccess) Opt {
	return func(o Opts) Opts {
		o.listenerCopySuccess = l
		return o
	}
}

func OnCopyFailure(l ListenerCopyFailure) Opt {
	return func(o Opts) Opts {
		o.listenerCopyFailure = l
		return o
	}
}

func OnDeleteSuccess(l ListenerDeleteSuccess) Opt {
	return func(o Opts) Opts {
		o.listenerDeleteSuccess = l
		return o
	}
}

func OnDeleteFailure(l ListenerDeleteFailure) Opt {
	return func(o Opts) Opts {
		o.listenerDeleteFailure = l
		return o
	}
}

func OnCreateFolderSuccess(l ListenerCreateFolderSuccess) Opt {
	return func(o Opts) Opts {
		o.listenerCreateFolderSuccess = l
		return o
	}
}

func OnCreateFolderFailure(l ListenerCreateFolderFailure) Opt {
	return func(o Opts) Opts {
		o.listenerCreateFolderFailure = l
		return o
	}
}

func OnSkip(l ListenerSkip) Opt {
	return func(o Opts) Opts {
		o.listenerSkip = l
		return o
	}
}

func WithNameFilter(filter mo_filter.Filter) Opt {
	return func(o Opts) Opts {
		o.entryNameFilter = filter
		return o
	}
}
