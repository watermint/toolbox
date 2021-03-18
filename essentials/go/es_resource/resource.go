package es_resource

import (
	"embed"
	"errors"
	"github.com/watermint/toolbox/essentials/http/es_filesystem"
	"io/fs"
	"net/http"
	"path/filepath"
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

func New(tpl, msg, web, key, img, dat Resource) Bundle {
	return &bundleImpl{
		tpl: tpl,
		msg: msg,
		web: web,
		key: key,
		img: img,
		dat: dat,
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

func NewResource(prefix string, fs embed.FS) Resource {
	return &resBox{
		prefix: prefix,
		fs:     fs,
	}
}

// go embed wrapper
type resBox struct {
	prefix string
	fs     embed.FS
}

func (z resBox) Bytes(key string) (bin []byte, err error) {
	return z.fs.ReadFile(filepath.Join(z.prefix, key))
}

func (z resBox) HttpFileSystem() http.FileSystem {
	f, err := fs.Sub(z.fs, z.prefix)
	if err != nil {
		panic(err)
	}
	return http.FS(f)
}

func NewSecureResource(prefix string, fs embed.FS) Resource {
	return &resSecureBox{
		prefix: prefix,
		fs:     fs,
	}
}

// embed.FS wrapper, but do not return http.FileSystem
type resSecureBox struct {
	prefix string
	fs     embed.FS
}

func (z resSecureBox) Bytes(key string) (bin []byte, err error) {
	return z.fs.ReadFile(filepath.Join(z.prefix, key))
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
