package dbx_error

func NewErrorMember(de ErrorInfo) ErrorMember {
	return &errorMemberImpl{
		de: de,
	}
}

type errorMemberImpl struct {
	de ErrorInfo
}

func (z errorMemberImpl) IsNotAMember() bool {
	return z.de.HasPrefix("member_error/not_a_member")
}
