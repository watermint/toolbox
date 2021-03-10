package qrcode

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestWifi_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Wifi{})
}
