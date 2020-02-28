package ut_log

import (
	"github.com/watermint/toolbox/infra/control/app_log"
	"log"
	"testing"
)

func TestLogWrapper_Write(t *testing.T) {
	zl := app_log.NewConsoleLogger(true, true)
	lw := NewLogWrapper(zl)
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
	for i := 1; i < 25; i++ {
		q := make([]byte, 25)
		for j := 0; j < 25; j++ {
			q[j] = 'A' + byte(j)
		}
		q[25-i] = '\n'
		lw.Write(q)
	}

	lw.Flush()
}
