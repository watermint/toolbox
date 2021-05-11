package dbx_error

import (
	"strings"
)

type ErrorInfo struct {
	ErrorTag          string `path:"error.\\.tag" json:"error,omitempty"`
	ErrorSummary      string `path:"error_summary" json:"error_summary,omitempty"`
	UserMessageLocale string `path:"user_message.locale" json:"user_message_lang,omitempty"`
	UserMessage       string `path:"user_message.text" json:"user_message,omitempty"`
}

type ErrorBadRequest struct {
	Reason string `json:"reason"`
}

func (z ErrorBadRequest) Error() string {
	return z.Reason
}

func (z ErrorInfo) Error() string {
	if z.UserMessage != "" {
		return z.UserMessage
	}
	return z.ErrorSummary
}

func (z ErrorInfo) HasPrefix(prefix string) bool {
	return strings.HasPrefix(z.ErrorSummary, prefix)
}

func NewErrors(err error) DropboxError {
	if err == nil {
		return nil
	}
	if ei, ok := err.(*ErrorInfo); ok {
		return &errorsImpl{ei: *ei}
	} else {
		return &errorsImpl{ei: ErrorInfo{}}
	}
}

type DropboxError interface {
	error

	Auth() ErrorAuth
	Access() ErrorAccess
	Path() ErrorEndpointPath
	Endpoint() ErrorEndpoint
	To() ErrorWrite
	BadPath() ErrorSharePath
	Member() ErrorMember

	// too_many_write_operations
	IsTooManyWriteOperations() bool
	// too_many_files
	IsTooManyFiles() bool
	// folder_name_already_used
	IsFolderNameAlreadyUsed() bool
	// group_name_already_used
	IsGroupNameAlreadyUsed() bool
	// duplicate user
	IsDuplicateUser() bool
	// member not in group
	IsMemberNotInGroup() bool
	// id not found
	IsIdNotFound() bool
	// shared_link_already_exists
	IsSharedLinkAlreadyExists() bool
	// It may not exist or may point to a deleted file.
	IsInvalidRevision() bool

	Summary() string
}

// 401: Bad or expired token. This can happen if the access token is expired or if the access token has been revoked by Dropbox or the user.
type ErrorAuth interface {
	// The access token is invalid.
	IsInvalidAccessToken() bool
	// The user specified in 'Dropbox-API-Select-User' is no longer on the team.
	IsInvalidSelectUser() bool
	// The user specified in 'Dropbox-API-Select-Admin' is not a Dropbox Business team admin.
	IsInvalidSelectAdmin() bool
	// The user has been suspended.
	IsUserSuspended() bool
	// The access token has expired.
	IsExpiredAccessToken() bool
	// The access token does not have the required scope to access the route.
	IsMissingScope() bool
	// The route is not available to public.
	IsRouteAccessDenied() bool
}

// 403: The user or team account doesn't have access to the endpoint or feature.
type ErrorAccess interface {
	// Current account type cannot access the resource.
	IsInvalidAccountType() bool
	// Current account cannot access Paper.
	IsPaperAccessDenied() bool
}

// 409: Endpoint specific error, Path
type ErrorEndpointPath interface {
	// There is nothing at the given path.
	IsNotFound() bool
	// Not a file
	IsNotFile() bool
	// The given path does not satisfy the required path format
	IsMalformedPath() bool
	// Couldn't write to the target path because there was something in the way.
	IsConflict() bool
	// There are too many write operations in user's Dropbox. Please retry this request.
	IsTooManyWriteOperations() bool
}

// 409: Other endpoint specific error
type ErrorEndpoint interface {
	IsRateLimit() bool
}

// 409: WriteError
type ErrorWrite interface {
	IsConflict() bool
}

// 409: SharePathError
type ErrorSharePath interface {
	IsAlreadyShared() bool
}

// 409: ErrorMember
type ErrorMember interface {
	IsNotAMember() bool
}

type errorsImpl struct {
	ei ErrorInfo
}

func (z errorsImpl) IsInvalidRevision() bool {
	return z.ei.HasPrefix("invalid_revision")
}

func (z errorsImpl) IsSharedLinkAlreadyExists() bool {
	return z.ei.HasPrefix("shared_link_already_exists")
}

func (z errorsImpl) Error() string {
	return z.Summary()
}

func (z errorsImpl) IsIdNotFound() bool {
	return z.ei.HasPrefix("id_not_found")
}

func (z errorsImpl) IsMemberNotInGroup() bool {
	return z.ei.HasPrefix("member_not_in_group")
}

func (z errorsImpl) IsDuplicateUser() bool {
	return z.ei.HasPrefix("duplicate_user")
}

func (z errorsImpl) IsGroupNameAlreadyUsed() bool {
	return z.ei.HasPrefix("group_name_already_used")
}

func (z errorsImpl) BadPath() ErrorSharePath {
	return NewSharePath("bad_path", z.ei)
}

func (z errorsImpl) IsFolderNameAlreadyUsed() bool {
	return z.ei.HasPrefix("folder_name_already_used")
}

func (z errorsImpl) IsTooManyFiles() bool {
	return z.ei.HasPrefix("too_many_files")
}

func (z errorsImpl) IsTooManyWriteOperations() bool {
	return z.ei.HasPrefix("too_many_write_operations") || z.Path().IsTooManyWriteOperations()
}

func (z errorsImpl) Summary() string {
	return z.ei.ErrorSummary
}

func (z errorsImpl) Auth() ErrorAuth {
	return NewErrorAuth(z.ei)
}

func (z errorsImpl) Access() ErrorAccess {
	return NewErrorAccess(z.ei)
}

func (z errorsImpl) Path() ErrorEndpointPath {
	return NewErrorPath(z.ei)
}

func (z errorsImpl) Endpoint() ErrorEndpoint {
	return NewErrorEndpoint(z.ei)
}

func (z errorsImpl) To() ErrorWrite {
	return NewErrorWrite("to", z.ei)
}

func (z errorsImpl) Member() ErrorMember {
	return NewErrorMember(z.ei)
}
