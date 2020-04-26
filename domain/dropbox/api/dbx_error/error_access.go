package dbx_error

func NewErrorAccess(de DropboxError) ErrorAccess {
	return &errorAccessImpl{
		de: de,
	}
}

type errorAccessImpl struct {
	de DropboxError
}

func (z errorAccessImpl) IsInvalidAccountType() bool {
	return z.de.HasPrefix("invalid_account_type")
}

func (z errorAccessImpl) IsPaperAccessDenied() bool {
	return z.de.HasPrefix("paper_access_denied")
}
