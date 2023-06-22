package setting

import (
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestUpdateSetting_Split(t *testing.T) {
	{
		u1 := UpdateSetting{
			Path: "/Sales/Report",
		}
		teamFolder, path := u1.Split()
		if teamFolder != "Sales" {
			t.Error(teamFolder)
		}
		if path != "/Report" {
			t.Error(path)
		}
	}

	{
		u1 := UpdateSetting{
			Path: "Sales/Report",
		}
		teamFolder, path := u1.Split()
		if teamFolder != "Sales" {
			t.Error(teamFolder)
		}
		if path != "/Report" {
			t.Error(path)
		}
	}

	{
		u1 := UpdateSetting{
			Path: "/Sales",
		}
		teamFolder, path := u1.Split()
		if teamFolder != "Sales" {
			t.Error(teamFolder)
		}
		if path != "" {
			t.Error(path)
		}
	}

	{
		u1 := UpdateSetting{
			Path: "Sales",
		}
		teamFolder, path := u1.Split()
		if teamFolder != "Sales" {
			t.Error(teamFolder)
		}
		if path != "" {
			t.Error(path)
		}
	}
}

func TestUpdate_Exec(t *testing.T) {
	qtr_endtoend.TestRecipe(t, &Update{})
}
