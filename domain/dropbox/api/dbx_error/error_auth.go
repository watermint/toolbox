package dbx_error

func NewErrorAuth(de ErrorInfo) ErrorAuth {
	return &errorAuthImpl{
		de: de,
	}
}

type errorAuthImpl struct {
	de ErrorInfo
}

func (z errorAuthImpl) IsInvalidAccessToken() bool {
	return z.de.HasPrefix("invalid_access_token")
}

func (z errorAuthImpl) IsInvalidSelectUser() bool {
	return z.de.HasPrefix("invalid_select_user")
}

func (z errorAuthImpl) IsInvalidSelectAdmin() bool {
	return z.de.HasPrefix("invalid_select_admin")
}

func (z errorAuthImpl) IsUserSuspended() bool {
	return z.de.HasPrefix("user_suspended")
}

func (z errorAuthImpl) IsExpiredAccessToken() bool {
	return z.de.HasPrefix("expired_access_token")
}

func (z errorAuthImpl) IsMissingScope() bool {
	return z.de.HasPrefix("missing_scope")
}

func (z errorAuthImpl) IsRouteAccessDenied() bool {
	return z.de.HasPrefix("route_access_denied")
}
