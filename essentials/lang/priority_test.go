package lang

import "testing"

func TestPriority(t *testing.T) {
	{
		pri := Priority(Japanese)
		if pri[0].CodeString() != "ja" || pri[1].CodeString() != "en" {
			t.Error(pri)
		}
	}
	{
		pri := Priority(English)
		if pri[0].CodeString() != "en" || len(pri) != 1 {
			t.Error(pri)
		}
	}
}
