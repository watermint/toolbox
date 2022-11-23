package efs_base

const (
	FsErrorReasonNotAllowed FsErrorReason = iota
	FsErrorReasonTimeout
	FsErrorReasonConflict
	FsErrorReasonPermission
)

type FsErrorReason int

func NewFsError(reason FsErrorReason, cause error) FsError {
	return &fsError{
		ErrorBase: NewBase(cause),
		reason:    reason,
	}
}

type fsError struct {
	ErrorBase
	reason FsErrorReason
}

func (z fsError) IsPathTooLong() bool {
	//TODO implement me
	panic("implement me")
}

func (z fsError) IsPathInvalidName() bool {
	//TODO implement me
	panic("implement me")
}

func (z fsError) IsNotAllowed() bool {
	return z.reason == FsErrorReasonNotAllowed
}

func (z fsError) IsTimeout() bool {
	return z.reason == FsErrorReasonTimeout
}

func (z fsError) IsConflict() bool {
	return z.reason == FsErrorReasonConflict
}

func (z fsError) IsPermission() bool {
	return z.reason == FsErrorReasonPermission
}
