package goog_error

import (
	"encoding/json"
	"testing"
)

const (
	errorSample = `{
  "error": {
    "code": 400,
    "message": "Precondition check failed.",
    "errors": [
      {
        "message": "Precondition check failed.",
        "domain": "global",
        "reason": "failedPrecondition"
      }
    ],
    "status": "FAILED_PRECONDITION"
  }
}
`
)

func TestGoogleError_Error(t *testing.T) {
	e := GoogleError{}
	if err := json.Unmarshal([]byte(errorSample), &e); err != nil {
		t.Error(err)
	}
	if e.Info.Code != 400 || e.Info.Status != "FAILED_PRECONDITION" {
		t.Error(e.Info)
	}
}
