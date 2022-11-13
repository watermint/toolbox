package efs_base

func NewBase(cause error) ErrorBase {
	return &errorBaseImpl{
		ErrorCause: cause,
	}
}

type errorBaseImpl struct {
	ErrorCause error
}

func (z errorBaseImpl) Error() string {
	if z.ErrorCause != nil {
		return z.ErrorCause.Error()
	} else {
		return ""
	}
}

func (z errorBaseImpl) Cause() error {
	return z.ErrorCause
}
