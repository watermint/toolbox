package efs_util

import (
	"bytes"
	"github.com/watermint/toolbox/essentials/file/efs_base"
	"strings"
)

type fileImpl struct {
	base efs_base.FileSystemBase
}

func (z fileImpl) FilePutBin(path efs_base.Path, data []byte) efs_base.FsError {
	return z.base.FilePut(path, bytes.NewReader(data))
}

func (z fileImpl) FilePutText(path efs_base.Path, data string, opts ...FilePutTextOpt) efs_base.FsError {
	return z.base.FilePut(path, strings.NewReader(data))
}

func (z fileImpl) FilePutJson(path efs_base.Path, data interface{}, opts ...FilePutJsonOpt) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z fileImpl) FilePutTextLines(path efs_base.Path, adder func(l TextLineAdder)) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z fileImpl) FileGetBin(path efs_base.Path) (data []byte, fsError efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}

func (z fileImpl) FileGetText(path efs_base.Path, opts ...FileGetTextOpts) (data string, fsError efs_base.FsError) {
	//TODO implement me
	panic("implement me")
}

func (z fileImpl) FileGetTextLine(path efs_base.Path, handler func(line string) bool, opts FileGetTextOpt) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}

func (z fileImpl) FileGetJson(path efs_base.Path, model interface{}) efs_base.FsError {
	//TODO implement me
	panic("implement me")
}
