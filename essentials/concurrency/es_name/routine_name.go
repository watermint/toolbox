package es_name

import (
	"runtime"
	"strings"
)

func CurrentRoutineName() string {
	buf := make([]byte, 128)
	runtime.Stack(buf, false)
	name := string(buf)
	name = strings.TrimPrefix(name, "goroutine ")
	name = name[:strings.Index(name, " ")]
	return name
}
