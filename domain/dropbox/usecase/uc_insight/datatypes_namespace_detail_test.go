package uc_insight

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"reflect"
	"testing"
)

var (
	sampleNamespaceDetail = `{
    "access_inheritance": {
        ".tag": "inherit"
    },
    "access_type": {
        ".tag": "owner"
    },
    "is_inside_team_folder": false,
    "is_team_folder": false,
    "link_metadata": {
        "audience_options": [
            {
                ".tag": "public"
            },
            {
                ".tag": "team"
            },
            {
                ".tag": "members"
            }
        ],
        "current_audience": {
            ".tag": "public"
        },
        "link_permissions": [
            {
                "action": {
                    ".tag": "change_audience"
                },
                "allow": true
            }
        ],
        "password_protected": false,
        "url": ""
    },
    "name": "dir",
    "path_lower": "/dir",
    "permissions": [],
    "policy": {
        "acl_update_policy": {
            ".tag": "owner"
        },
        "member_policy": {
            ".tag": "anyone"
        },
        "resolved_member_policy": {
            ".tag": "team"
        },
        "shared_link_policy": {
            ".tag": "anyone"
        }
    },
    "preview_url": "https://www.dropbox.com/scl/fo/fir9vjelf",
    "shared_folder_id": "84528192421",
    "time_invited": "2016-01-20T00:00:00Z"
}`
)

func TestNewNamespaceDetail(t *testing.T) {
	nd, err := NewNamespaceDetail(es_json.MustParseString(sampleNamespaceDetail))
	if err != nil {
		t.Error(err)
	}

	if nd.NamespaceId != "84528192421" {
		t.Error(nd.NamespaceId)
	}
	if nd.Name != "dir" {
		t.Error(nd.Name)
	}
	if nd.PolicyAclUpdate != "owner" {
		t.Error(nd.PolicyAclUpdate)
	}
	if nd.PolicyMemberPolicy != "anyone" {
		t.Error(nd.PolicyMemberPolicy)
	}
	if nd.PolicyResolvedAcl != "team" {
		t.Error(nd.PolicyResolvedAcl)
	}
	if nd.PolicySharedLink != "anyone" {
		t.Error(nd.PolicySharedLink)
	}
	if nd.AccessInheritance != "inherit" {
		t.Error(nd.AccessInheritance)
	}
	if nd.IsInsideTeamFolder != false {
		t.Error(nd.IsInsideTeamFolder)
	}
	if nd.IsTeamFolder != false {
		t.Error(nd.IsTeamFolder)
	}

	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		db, err := ctl.NewOrmOnMemory()
		if err != nil {
			t.Error(err)
			return
		}
		if err = db.AutoMigrate(&NamespaceDetail{}); err != nil {
			t.Error(err)
			return
		}
		if err = db.Create(nd).Error; err != nil {
			t.Error(err)
			return
		}

		nd2 := &NamespaceDetail{}
		db.First(nd2)
		if !reflect.DeepEqual(nd, nd2) {
			t.Error(nd2)
		}
	})
}
