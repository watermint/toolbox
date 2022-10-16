package efs_util

import "github.com/watermint/toolbox/essentials/file/efs_base"

type FilePutTextOpts struct {
	Encoding      string
	LineSeparator string
}
type FilePutTextOpt func(o FilePutTextOpts) FilePutTextOpts

type FilePutJsonOpts struct {
	Indent int
}
type FilePutJsonOpt func(o FilePutJsonOpts) FilePutJsonOpts

type FileGetTextOpts struct {
	Encoding      string
	LineSeparator string
}
type FileGetTextOpt func(o FileGetTextOpts) FileGetTextOpts

type FileUtilOps interface {
	FilePutBin(path efs_base.Path, data []byte) efs_base.FsError
	FilePutText(path efs_base.Path, data string, opts ...FilePutTextOpt) efs_base.FsError
	FilePutJson(path efs_base.Path, data interface{}, opts ...FilePutJsonOpt) efs_base.FsError
	FilePutTextLines(path efs_base.Path, adder func(l TextLineAdder)) efs_base.FsError
	FileGetBin(path efs_base.Path) (data []byte, fsError efs_base.FsError)
	FileGetText(path efs_base.Path, opts ...FileGetTextOpts) (data string, fsError efs_base.FsError)
	FileGetTextLine(path efs_base.Path, handler func(line string) bool, opts FileGetTextOpt) efs_base.FsError
	FileGetJson(path efs_base.Path, model interface{}) efs_base.FsError
}

type TextLineAdder interface {
	AddLine(line string) efs_base.FsError
	AddJson(data interface{}) efs_base.FsError
}
