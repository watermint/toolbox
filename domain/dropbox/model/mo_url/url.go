package mo_url

import (
	"net/url"
)

type Url interface {
	Scheme() string
	Authority() string
	Path() string
	Query() string
	Fragment() string
	Value() string
}

func NewEmptyUrl() Url {
	return &urlImpl{u: &url.URL{}}
}

func NewUrl(u string) (z Url, err error) {
	r, err := url.Parse(u)
	if err != nil {
		return nil, err
	}
	return &urlImpl{u: r}, nil
}

type urlImpl struct {
	u *url.URL
}

func (z *urlImpl) Value() string {
	return z.u.String()
}

func (z *urlImpl) Scheme() string {
	return z.u.Scheme
}

func (z *urlImpl) Authority() string {
	return z.u.Host
}

func (z *urlImpl) Path() string {
	return z.u.Path
}

func (z *urlImpl) Query() string {
	return z.u.Path
}

func (z *urlImpl) Fragment() string {
	return z.u.Fragment
}
