package monitor

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestClient_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Client{})
}
