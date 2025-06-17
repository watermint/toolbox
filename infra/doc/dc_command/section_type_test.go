package dc_command

import (
	"testing"
)

func TestSectionType_Priority(t *testing.T) {
	testCases := []struct {
		name     string
		section  SectionType
		expected int
	}{
		{"Header", SectionTypeHeader, 1},
		{"CommandSecurity", SectionTypeCommandSecurity, 2},
		{"CommandAuth", SectionTypeCommandAuth, 3},
		{"Install", SectionTypeInstall, 4},
		{"Usage", SectionTypeUsage, 5},
		{"Feed", SectionTypeFeed, 6},
		{"Report", SectionTypeReport, 7},
		{"GridDataInput", SectionTypeGridDataInput, 8},
		{"GridDataOutput", SectionTypeGridDataOutput, 9},
		{"TextInput", SectionTypeTextInput, 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if priority := tc.section.Priority(); priority != tc.expected {
				t.Errorf("Expected priority %d, got %d", tc.expected, priority)
			}
		})
	}
}