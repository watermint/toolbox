package es_color

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"testing"
)

func TestFprintf(t *testing.T) {
	w := es_stdout.NewDefaultOut(true)
	Colorfln(w, ColorBlue, false, "Hello %d", 123)
	Colorfln(w, ColorCyan, true, "Hello %d", 123)
	Colorfln(w, ColorGreen, false, "Hello %d %s", 123, "world")

	Boldfln(w, "Hello %d", 123)
	Boldfln(w, "Hello %d", 123)
	Boldfln(w, "Hello %d %s", 123, "world")
}
