package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	sampleGroup = `{
            "group_id": "g:e2db7665347abcd600000000001a2b3c",
            "group_management_type": {
                ".tag": "user_managed"
            },
            "group_name": "Test group",
            "member_count": 10
        }`
)

func TestNewGroupFromJson(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		g, err := NewGroupFromJson(es_json.MustParseString(sampleGroup))
		if err != nil {
			t.Error(err)
		}
		if g.GroupId != "g:1234567890" {
			t.Error(g.GroupId)
		}
		if g.GroupName != "Finance" {
			t.Error(g.GroupName)
		}
		if g.GroupManagementType != "user_managed" {
			t.Error(g.GroupManagementType)
		}

		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		if err := db.AutoMigrate(&Group{}); err != nil {
			t.Error(err)
			return
		}
		if err := db.Create(g).Error; err != nil {
			return
		}

		g1 := &Group{}
		db.First(&g1)
		if !reflect.DeepEqual(g, g1) {
			t.Error(g1)
		}
	})
}
