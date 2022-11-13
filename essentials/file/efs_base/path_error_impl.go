package efs_base

const (
	PathErrorTooLong PathErrorReason = iota
	PathErrorInvalidName
)

type PathErrorReason int

func NewPathError(reason PathErrorReason, cause error) PathError {
	return &pathErrorImpl{
		ErrorBase: NewBase(cause),
		reason:    reason,
	}
}

type pathErrorImpl struct {
	ErrorBase
	reason PathErrorReason
}

func (z pathErrorImpl) IsPathTooLong() bool {
	return z.reason == PathErrorTooLong
}

func (z pathErrorImpl) IsPathInvalidName() bool {
	return z.reason == PathErrorInvalidName
}
