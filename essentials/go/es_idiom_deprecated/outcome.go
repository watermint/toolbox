package es_idiom_deprecated

import "fmt"

// Outcome is the alternative for the `error`. Outcome instance must not be nil.
// And Outcome implementation should implement specific error cases with the prefix `IsXxx`.
// For example: file operation.
//
//	type FileCreateOutcome interface {
//	  Outcome
//	  IsPermissionDenied() bool
//	  IsInvalidPath() bool
//	  IsOperationNotAllowed() bool
//	}
//
// Consumer can handle errors like below.
//
//	f, out := file.Create("/path/to/create")
//	switch {
//	case out.IsOk():
//	    // success
//	case out.IsPermissionDenied(), out.IsOperationNotAllowed():
//	    // handle permission issue
//	case out.IsInvalidPath():
//	    // handle path issue
//	default:
//	    // handle other errors
//	}
type Outcome interface {
	// Stringer Outcome instance returns empty string if an operation succeed, otherwise returns an error string.
	fmt.Stringer

	// IsOk Returns true if an operation succeed.
	IsOk() bool

	// IfOk Perform f if an operation succeed, otherwise does nothing.
	IfOk(f func())

	// IsError Returns true if an operation got an error.
	IsError() bool

	// IfError Perform f if an operation got an error, otherwise does nothing.
	IfError(f func() Outcome) Outcome

	// Cause Return Outcome as an error instance if an operation got an error, otherwise returns nil.
	Cause() error
}

// UnconfirmedOutcome is similar to Outcome, that replaces `error`. But UnconfirmedOutcome does not
// include the operation's actual result. For example, when REST api client interface got a network error,
// that is obviously error. However, if a REST API call result contains an error, the caller need to
// verify response body or status code to determine the operation succeed or not.
type UnconfirmedOutcome interface {
	// HasError Returns true if an operation got at least one error.
	HasError() bool

	// IfError Perform f if an operation got at least one error, otherwise does nothing.
	IfError(f func() UnconfirmedOutcome) UnconfirmedOutcome

	// Cause Returns an error if UnconfirmedOutcome has as an obvious error (such as network error for REST call),
	// otherwise returns nil.
	// Inherited UnconfirmedOutcome interface should define a function like additional information like
	// StatusCode(), LastErr(), etc.
	Cause() error
}
