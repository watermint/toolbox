package es_resource

import (
	"errors"
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/essentials/http/es_filesystem"
	"net/http"
)

const (
	BundleResource = "resources"
	BundleWeb      = "web"
)

type Bundle interface {
	Templates() Resource
	Messages() Resource
	Web() Resource
	Keys() Resource
	Images() Resource
	Data() Resource
}

type Resource interface {
	Bytes(key string) (bin []byte, err error)
	HttpFileSystem() http.FileSystem
}

func New(tpl, msg, web, key, img, dat *rice.Box) Bundle {
	return &bundleImpl{
		tpl: NewResource(tpl),
		msg: NewSecureResource(msg),
		web: NewResource(web),
		key: NewSecureResource(key),
		img: NewResource(img),
		dat: NewSecureResource(dat),
	}
}

func EmptyBundle() Bundle {
	return &bundleImpl{
		tpl: EmptyResource(),
		msg: EmptyResource(),
		web: EmptyResource(),
		key: EmptyResource(),
		img: EmptyResource(),
		dat: EmptyResource(),
	}
}

type bundleImpl struct {
	tpl Resource
	msg Resource
	web Resource
	key Resource
	img Resource
	dat Resource
}

func (z bundleImpl) Templates() Resource {
	return z.tpl
}

func (z bundleImpl) Messages() Resource {
	return z.msg
}

func (z bundleImpl) Web() Resource {
	return z.web
}

func (z bundleImpl) Keys() Resource {
	return z.key
}

func (z bundleImpl) Images() Resource {
	return z.img
}

func (z bundleImpl) Data() Resource {
	return z.dat
}

func NewResource(b *rice.Box) Resource {
	return &resBox{
		b: b,
	}
}

// rice.Box wrapper
type resBox struct {
	b *rice.Box
}

func (z resBox) Bytes(key string) (bin []byte, err error) {
	return z.b.Bytes(key)
}

func (z resBox) HttpFileSystem() http.FileSystem {
	return z.b.HTTPBox()
}

func NewSecureResource(b *rice.Box) Resource {
	return &resSecureBox{
		b: b,
	}
}

// rice.Box wrapper, but do not return http.FileSystem
type resSecureBox struct {
	b *rice.Box
}

func (z resSecureBox) Bytes(key string) (bin []byte, err error) {
	return z.b.Bytes(key)
}

func (z resSecureBox) HttpFileSystem() http.FileSystem {
	return es_filesystem.Empty{}
}

var (
	ErrorAlwaysFail = errors.New("always fail")
)

func EmptyResource() Resource {
	return &resEmpty{}
}

// empty resource
type resEmpty struct {
}

func (z resEmpty) Bytes(key string) (bin []byte, err error) {
	return nil, ErrorAlwaysFail
}

func (z resEmpty) HttpFileSystem() http.FileSystem {
	return es_filesystem.Empty{}
}
