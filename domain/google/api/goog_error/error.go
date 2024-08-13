package goog_error

// Error types
// https://developers.google.com/drive/api/v3/handle-errors

type ErrorReason struct {
	Message string `json:"message"`
	Domain  string `json:"domain"`
	Reason  string `json:"reason"`
}

type ErrorInfo struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Status  string        `json:"status"`
	Errors  []ErrorReason `json:"errors"`
}

type GoogleError struct {
	Info ErrorInfo `json:"error"`
}

func (z GoogleError) Error() string {
	return z.Info.Message
}
