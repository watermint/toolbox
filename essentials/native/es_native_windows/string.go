//go:build windows
// +build windows

package es_native_windows

import (
	"fmt"
	"syscall"
	"unsafe"
)

type BufferString interface {
	fmt.Stringer

	Pointer() uintptr

	BufSize() uintptr
}

func NewBufferString(bufSize int) BufferString {
	return &winStr{
		buf: make([]uint16, bufSize),
	}
}

type winStr struct {
	buf []uint16
}

func (z winStr) String() string {
	return syscall.UTF16ToString(z.buf)
}

func (z winStr) Pointer() uintptr {
	return uintptr(unsafe.Pointer(&z.buf[0]))
}

func (z winStr) BufSize() uintptr {
	return uintptr(len(z.buf))
}
