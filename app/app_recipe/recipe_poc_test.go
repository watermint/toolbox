package app_recipe

import (
	"reflect"
	"testing"
)

func TestPoc(t *testing.T) {
	vo := TeamFolderListVO{}
	vot := reflect.TypeOf(vo)
	println(vot.PkgPath())
}
