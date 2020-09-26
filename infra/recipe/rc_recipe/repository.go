package rc_recipe

import (
	"flag"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type Repository interface {
	// Returns feeds that requested by the recipe
	Feeds() map[string]fd_file.RowFeed

	// Returns feed specifications that requested by the recipe
	FeedSpecs() map[string]fd_file.Spec

	// Returns reports that will created by the recipe
	Reports() map[string]rp_model.Report

	// Returns reports that will created by the recipe
	ReportSpecs() map[string]rp_model.Spec

	// Messages used by the recipe
	Messages() []app_msg.Message

	// List of fields
	FieldNames() []string

	// Text representation of the field value
	FieldValueText(name string) string

	// Field value
	FieldValue(name string) Value

	// Returns connections that requested by the recipe
	Conns() map[string]api_conn.Connection

	// Apply values in the repository to the
	Apply() Recipe

	// Apply custom values to the repository
	ApplyCustom()

	// Serialize
	Capture(ctl app_control.Control) (v interface{}, err error)

	// Deserialize
	Restore(j es_json.Json, ctl app_control.Control) error

	// Prepare values for run recipe
	SpinUp(ctl app_control.Control) (Recipe, error)

	// Spin down value
	SpinDown(ctl app_control.Control) error

	// Apply flag set
	ApplyFlags(f *flag.FlagSet, ui app_ui.UI)

	// Description of the field
	FieldDesc(name string) app_msg.Message

	// Custom description for default value
	FieldCustomDefault(name string) app_msg.MessageOptional

	// Serialize values for debug
	Debug() map[string]interface{}
}
