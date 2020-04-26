package api_request

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/util/ut_io"
	"go.uber.org/zap"
)

const (
	ReqHeaderContentType           = "Content-Type"
	ReqHeaderAccept                = "Accept"
	ReqHeaderContentLength         = "Content-Length"
	ReqHeaderAuthorization         = "Authorization"
	ReqHeaderUserAgent             = "User-Agent"
	ReqHeaderDropboxApiSelectUser  = "Dropbox-API-Select-User"
	ReqHeaderDropboxApiSelectAdmin = "Dropbox-API-Select-Admin"
	ReqHeaderDropboxApiPathRoot    = "Dropbox-API-Path-Root"
	ReqHeaderDropboxApiArg         = "Dropbox-API-Arg"
)

type RequestData struct {
	p interface{}
	h map[string]string
	c ut_io.ReadRewinder
}

// Returns JSON form of param. Returns `null` string if an error occurred.
func (z RequestData) ParamJson() json.RawMessage {
	l := app_root.Log()
	q, err := json.Marshal(z.p)
	if err != nil {
		l.Debug("unable to marshal param", zap.Error(err), zap.Any("p", z.p))
		return json.RawMessage("null")
	} else {
		return q
	}
}

// Returns query string like "?key=value&key2=value2". Returns empty string if an error occurred.
func (z RequestData) ParamQuery() string {
	l := app_root.Log()
	q, err := query.Values(z.p)
	if err != nil {
		l.Debug("unable to make query", zap.Error(err), zap.Any("p", z.p))
		return ""
	} else {
		return "?" + q.Encode()
	}
}

func (z RequestData) Param() interface{} {
	return z.p
}

func (z RequestData) Headers() map[string]string {
	if z.h == nil {
		return map[string]string{}
	}
	return z.h
}
func (z RequestData) Content() ut_io.ReadRewinder {
	return z.c
}

type RequestDatum func(d RequestData) RequestData

func Param(p interface{}) RequestDatum {
	return func(d RequestData) RequestData {
		d.p = p
		return d
	}
}

func Header(name, value string) RequestDatum {
	return func(d RequestData) RequestData {
		h := make(map[string]string)
		for k, v := range d.h {
			h[k] = v
		}
		h[name] = value
		d.h = h
		return d
	}
}

func Content(c ut_io.ReadRewinder) RequestDatum {
	return func(d RequestData) RequestData {
		d.c = c
		return d
	}
}

// Combine datum into data
func Combine(rds []RequestDatum) RequestData {
	rd := RequestData{}
	for _, d := range rds {
		rd = d(rd)
	}
	return rd
}
