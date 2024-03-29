package rc_compatibility

import (
	"os"
	"testing"
)

func TestLoadCompatibilityDefinition(t *testing.T) {
	content, err := os.ReadFile("compatibility_test.json")
	if err != nil {
		t.Error(err)
	}
	cds, err := ParseCompatibilityDefinition(content)
	if err != nil {
		t.Error(err)
	}
	if cds.PathChanges[0].Current.Name != "lemon" {
		t.Error(cds.PathChanges[0].Current.Name)
	}
}

func TestCompatibilityDefinitions_Find(t *testing.T) {
	content, err := os.ReadFile("compatibility_test.json")
	if err != nil {
		t.Error(err)
	}
	cds, err := ParseCompatibilityDefinition(content)
	if err != nil {
		t.Error(err)
	}
	{
		cd, found := cds.FindPathChange([]string{"fruit", "citrus"}, "lemon")
		if !found || cd.Current.Name != "lemon" {
			t.Error(cd)
		}
	}
	{
		cd, found := cds.FindPathChange([]string{"fruit", "citrus"}, "lime")
		if !found || cd.Current.Name != "lime" {
			t.Error(cd)
		}
	}
	{
		cd, found := cds.FindPathChange([]string{"fruit", "citrus"}, "orange")
		if found {
			t.Error(cd)
		}
	}
}

func TestCompatibilityDefinitions_FindAlive(t *testing.T) {
	content, err := os.ReadFile("compatibility_test.json")
	if err != nil {
		t.Error(err)
	}
	cds, err := ParseCompatibilityDefinition(content)
	if err != nil {
		t.Error(err)
	}
	{
		// expired
		cd, found := cds.FindAlivePathChange([]string{"fruit", "citrus"}, "lemon")
		if found {
			t.Error(cd)
		}
	}
	{
		// alive
		cd, found := cds.FindAlivePathChange([]string{"fruit", "citrus"}, "lime")
		if !found || cd.Current.Name != "lime" {
			t.Error(cd)
		}
	}
	{
		// alive
		cd, found := cds.FindAlivePathChange([]string{"fruit", "citrus"}, "tangerine")
		if !found || cd.Current.Name != "tangerine" {
			t.Error(cd)
		}
	}
}

func TestCompatibilityDefinitions_FindPruned(t *testing.T) {
	content, err := os.ReadFile("compatibility_test.json")
	if err != nil {
		t.Error(err)
	}
	cds, err := ParseCompatibilityDefinition(content)
	if err != nil {
		t.Error(err)
	}
	{
		// expired
		cd, found := cds.FindPlannedPrune([]string{"chemical"}, "lead")
		if found && cd.Current.Name != "lead" {
			t.Error(cd)
		}
	}
	{
		// alive
		cd, found := cds.FindPlannedPrune([]string{"weapon"}, "nuclear")
		if !found || cd.Current.Name != "nuclear" {
			t.Error(cd)
		}
	}
}
