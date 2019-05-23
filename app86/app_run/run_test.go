package app_run

import (
	"os"
	"testing"
)

func TestCatalogue(t *testing.T) {
	c := catalogue()
	c.Run(os.Args)
}
