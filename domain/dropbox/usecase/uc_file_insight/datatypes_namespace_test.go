package uc_file_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	sampleNamespace = `{
            "name": "Franz Ferdinand",
            "namespace_id": "123456789",
            "namespace_type": {
                ".tag": "team_member_folder"
            },
            "team_member_id": "dbmid:1234567"
        }`
)

func TestNewNamespaceFromJson(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		ns, err := NewNamespaceFromJson(es_json.MustParseString(sampleNamespace))
		if err != nil {
			t.Error(err)
			return
		}
		if ns.Name != "Franz Ferdinand" {
			t.Error(ns)
		}
		if ns.NamespaceId != "123456789" {
			t.Error(ns)
		}
		if ns.NamespaceType != "team_member_folder" {
			t.Error(ns)
		}
		if ns.TeamMemberId != "dbmid:1234567" {
			t.Error(ns)
		}

		db, err := ctl.NewOrmOnMemory()
		if err := db.AutoMigrate(&Namespace{}); err != nil {
			t.Error(err)
			return
		}
		db.Create(ns)

		ns1 := &Namespace{}
		db.First(ns1)

		if !reflect.DeepEqual(ns, ns1) {
			t.Error(ns1)
		}
	})
}
