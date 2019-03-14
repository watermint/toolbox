package app_util

import (
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestLogWrapper_Write(t *testing.T) {
	zl, err := zap.NewDevelopment()
	if err != nil {
		t.Error(err)
		return
	}
	lw := NewLogWrapper(10, zl)
	log.SetOutput(lw)

	log.Println("Hello")
	log.Printf("Hello, World")
	log.Print("Hello, World, Tool, box")
	log.Print("Hello, World, Tool, box, Hello, World, Tool, box, Hello, World, Tool, box, Hello, World, Tool, box, Hello, World, Tool, box")

	lw.Write([]byte("D123456789A123456789"))
	lw.Write([]byte("E123456789B123456789"))
	lw.Write([]byte("F123456789C123456789"))
	lw.Write([]byte(""))
	lw.Flush()

	for i := 0; i < 25; i++ {
		q := make([]byte, i)
		for j := 0; j < i; j++ {
			q[j] = 'A' + byte(j)
		}
		lw.Write(q)
	}

	lw.Flush()
}
