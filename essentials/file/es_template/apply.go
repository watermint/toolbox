package es_template

import "github.com/watermint/toolbox/essentials/file/es_filesystem"

type HandlerApplyTagAdd func(path es_filesystem.Path)

type ApplyOpts struct {
	HandlerTagAdd HandlerApplyTagAdd
}

type Apply interface {
	// Apply template to the path
	Apply(path es_filesystem.Path, template Root) (err error)
}

func NewApply(fs es_filesystem.FileSystem, opts ApplyOpts) Apply {
	return &applyImpl{
		fs: fs,
		ao: opts,
	}
}

type applyImpl struct {
	fs es_filesystem.FileSystem
	ao ApplyOpts
}

func (z applyImpl) Apply(path es_filesystem.Path, template Root) (err error) {
	//TODO implement me
	panic("implement me")
}
