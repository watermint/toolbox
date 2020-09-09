package es_sync

import "github.com/watermint/toolbox/essentials/file/es_filesystem"

const (
	SkipSame   SkipReason = "same"
	SkipExists SkipReason = "exists"
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

func (z Opts) OnCopySuccess(source es_filesystem.Entry, target es_filesystem.Path) {
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

func (z Opts) OnSkip(reason SkipReason, source, target es_filesystem.Entry) {
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

type ListenerCopySuccess func(source es_filesystem.Entry, targetEntry es_filesystem.Path)
type ListenerCopyFailure func(source es_filesystem.Path, err es_filesystem.FileSystemError)
type ListenerDeleteSuccess func(target es_filesystem.Path)
type ListenerDeleteFailure func(target es_filesystem.Path, err es_filesystem.FileSystemError)
type ListenerSkip func(reason SkipReason, sourceEntry, targetEntry es_filesystem.Entry)
type ListenerCreateFolderSuccess func(target es_filesystem.Path)
type ListenerCreateFolderFailure func(target es_filesystem.Path, err es_filesystem.FileSystemError)
