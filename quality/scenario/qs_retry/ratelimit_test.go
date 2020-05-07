package qs_retry

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestRetryAfterDropbox(t *testing.T) {
	qt_recipe.TestWithReplayDbxContext(t, "qs_retry-ratelimit-team-get_info.json", func(ctx dbx_context.Context) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}

// RateLimit-Reset: second
// It's not an expected behavior of Dropbox API
func TestRateLimitReset(t *testing.T) {
	qt_recipe.TestWithReplayDbxContext(t, "qs_retry-ratelimit-ratelimit-reset.json", func(ctx dbx_context.Context) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}

// RateLimit-Reset: Fix date
// It's not an expected behavior of Dropbox API
func TestRateLimitResetFixDate(t *testing.T) {
	qt_recipe.TestWithReplayDbxContext(t, "qs_retry-ratelimit-ratelimit-reset-fixdate.json", func(ctx dbx_context.Context) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}

// X-RateLimit-Reset: Unix time
// It's not an expected behavior of Dropbox API
func TestXRateLimitResetUnixTime(t *testing.T) {
	qt_recipe.TestWithReplayDbxContext(t, "qs_retry-ratelimit-x-ratelimit-reset-unixtime.json", func(ctx dbx_context.Context) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}
