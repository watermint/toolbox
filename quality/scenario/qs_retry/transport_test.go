package qs_retry

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestTransportErrorRequestCancelled(t *testing.T) {
	qtr_endtoend.TestWithReplayDbxContext(t, "qs_retry-transport-request-cancelled.json", func(ctx dbx_client.Client) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}

func TestTransportErrorTcp1(t *testing.T) {
	qtr_endtoend.TestWithReplayDbxContext(t, "qs_retry-transport-read-tcp1.json", func(ctx dbx_client.Client) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}

func TestTransportErrorTcp2(t *testing.T) {
	qtr_endtoend.TestWithReplayDbxContext(t, "qs_retry-transport-read-tcp2.json", func(ctx dbx_client.Client) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}

func TestTransportErrorTcp3(t *testing.T) {
	qtr_endtoend.TestWithReplayDbxContext(t, "qs_retry-transport-read-tcp3.json", func(ctx dbx_client.Client) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}

func TestTransportErrorTcp4(t *testing.T) {
	qtr_endtoend.TestWithReplayDbxContext(t, "qs_retry-transport-read-tcp4.json", func(ctx dbx_client.Client) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}
