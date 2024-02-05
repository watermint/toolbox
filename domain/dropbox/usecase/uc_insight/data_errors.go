package uc_insight

import "github.com/watermint/toolbox/domain/dropbox/api/dbx_error"

type ApiErrorRecord interface {
	// ToParam Convert to retry parameter
	ToParam() interface{}
}

type ApiError struct {
	Error    string `json:"error"`
	ErrorTag string `json:"error_tag"`
}

func ApiErrorFromError(err error) ApiError {
	dbxErr := dbx_error.NewErrors(err)
	if dbxErr != nil {
		return ApiError{
			Error:    dbxErr.Error(),
			ErrorTag: dbxErr.Summary(),
		}
	}
	return ApiError{
		Error:    err.Error(),
		ErrorTag: "",
	}
}

type ApiErrorReport struct {
	Category string `json:"category"`
	Message  string `json:"message"`
	Tag      string `json:"tag"`
	Detail   string `json:"detail"`
}
