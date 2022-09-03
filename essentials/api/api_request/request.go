package api_request

import (
	"encoding/json"
	"github.com/google/go-querystring/query"
	"github.com/watermint/toolbox/essentials/io/es_rewinder"
	"github.com/watermint/toolbox/essentials/log/esl"
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
	d []RequestDatum
	p interface{}
	q interface{}
	h map[string]string
	c es_rewinder.ReadRewinder
}

func (z RequestData) WithDatum(datum RequestDatum) RequestData {
	all := make([]RequestDatum, len(z.d)+1)
	copy(all[:], z.d[:])
	all[len(z.d)] = datum
	return Combine(all)
}

func (z RequestData) Data() []RequestDatum {
	return z.d
}

// Convert into JSON form of param. Returns `null` string if an error occurred.
// Returns empty string when the parameter is null.
func (z RequestData) ParamJson() json.RawMessage {
	l := esl.Default()
	if z.p == nil {
		l.Debug("Parameter is null")
		return json.RawMessage{}
	}
	q, err := json.Marshal(z.p)
	if err != nil {
		l.Debug("unable to marshal param", esl.Error(err), esl.Any("p", z.p))
		return json.RawMessage("null")
	} else {
		return q
	}
}

// Returns query string like "?key=value&key2=value2". Returns empty string if an error occurred.
func (z RequestData) ParamQuery() string {
	l := esl.Default()
	if z.q == nil {
		return ""
	}
	q, err := query.Values(z.q)
	if err != nil {
		l.Debug("unable to make query", esl.Error(err), esl.Any("q", z.q))
		return ""
	} else {
		return "?" + q.Encode()
	}
}

// Returns raw param data
func (z RequestData) Param() interface{} {
	return z.p
}

// Returns raw query data
func (z RequestData) Query() interface{} {
	return z.q
}

func (z RequestData) Headers() map[string]string {
	if z.h == nil {
		return map[string]string{}
	}
	return z.h
}
func (z RequestData) Content() es_rewinder.ReadRewinder {
	return z.c
}

type RequestDatum func(d RequestData) RequestData

func Query(q interface{}) RequestDatum {
	return func(d RequestData) RequestData {
		d.q = q
		return d
	}
}

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

func Content(c es_rewinder.ReadRewinder) RequestDatum {
	return func(d RequestData) RequestData {
		d.c = c
		return d
	}
}

// Combine datum into data
func Combine(rds []RequestDatum) RequestData {
	rd := RequestData{
		d: rds,
	}
	for _, d := range rds {
		rd = d(rd)
	}
	return rd
}
