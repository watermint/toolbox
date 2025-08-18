package es_goroutine

import (
	"runtime"
	"strings"
)

func GetGoRoutineName() string {
	buf := make([]byte, 128)
	runtime.Stack(buf, false)
	name := string(buf)
	name = strings.TrimPrefix(name, "goroutine ")
	name = name[:strings.Index(name, " ")]
	return "gr:" + name
}
