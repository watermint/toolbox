package dbx_error

func NewErrorMember(de DropboxError) ErrorMember {
	return &errorMemberImpl{
		de: de,
	}
}

type errorMemberImpl struct {
	de DropboxError
}

func (z errorMemberImpl) IsNotAMember() bool {
	return z.de.HasPrefix("member_error/not_a_member")
}
