package app_recipe_poc

import (
	"reflect"
	"testing"
)

func TestPoc(t *testing.T) {
	vo := TeamFolderListVO{}
	vot := reflect.TypeOf(vo)
	println(vot.PkgPath())
}
