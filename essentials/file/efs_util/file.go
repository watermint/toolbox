package efs_util

import (
	"github.com/watermint/toolbox/essentials/file/efs_base"
	"golang.org/x/text/encoding"
)

type textOpts struct {
	Encoding      encoding.Encoding
	LineSeparator string
}

func TextEncoding(encoding encoding.Encoding) TextOpt {
	return func(o textOpts) textOpts {
		o.Encoding = encoding
		return o
	}
}

type TextOpt func(o textOpts) textOpts

type PutJsonOpts struct {
	Indent int
}
type FilePutJsonOpt func(o PutJsonOpts) PutJsonOpts

type FileUtilOps interface {
	FilePutBin(path efs_base.Path, data []byte) efs_base.FsError
	FilePutText(path efs_base.Path, data string, opts ...TextOpt) efs_base.FsError
	FilePutJson(path efs_base.Path, data interface{}, opts ...FilePutJsonOpt) efs_base.FsError
	FilePutTextLines(path efs_base.Path, adder func(l TextLineAdder)) efs_base.FsError
	FileGetBin(path efs_base.Path) (data []byte, fsError efs_base.FsError)
	FileGetText(path efs_base.Path, opts ...TextOpt) (data string, fsError efs_base.FsError)
	FileGetTextLine(path efs_base.Path, handler func(line string) bool, opts TextOpt) efs_base.FsError
	FileGetJson(path efs_base.Path, model interface{}) efs_base.FsError
}

type TextLineAdder interface {
	AddLine(line string) efs_base.FsError
	AddJson(data interface{}) efs_base.FsError
}
