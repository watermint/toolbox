package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"reflect"
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

	// Returns connections that requested by the recipe
	Conns() map[string]rc_conn.ConnDropboxApi

	// Apply values in the repository to the
	Apply() rc_recipe.Recipe

	// Prepare values for run recipe
	SpinUp(ctl app_control.Control) (rc_recipe.Recipe, error)

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

type Value interface {
	// Returns forked instance when the type is acceptable
	// Otherwise returns nil
	Accept(t reflect.Type, r rc_recipe.Recipe, name string) Value

	// Return value reference of the instance
	Bind() interface{}

	// Initialize instance, and returns the instance to set
	Init() (v interface{})

	// Apply internal state (bind'ed value) to the instance
	Apply() (v interface{})

	// Debug information
	Debug() interface{}

	// Spin up for run
	SpinUp(ctl app_control.Control) error

	// Spin down after run
	SpinDown(ctl app_control.Control) error
}

type ValueCustomValueText interface {
	Value

	ValueText() string
}

type ValueFeed interface {
	Value

	// True when the value is type of feed, and returns the instance of the feed
	Feed() (feed fd_file.RowFeed, valid bool)
}

type ValueReport interface {
	Value

	// True when the value is type of report, and returns the instance of the report
	Report() (report rp_model.Report, valid bool)
}

type ValueConn interface {
	Value

	// True when the value is type of connection, and return the instance of the connection
	Conn() (conn rc_conn.ConnDropboxApi, valid bool)
}

type ValueMessage interface {
	Value

	// True when the value is type of message, and return the instance eof the conection
	Message() (msg app_msg.Message, valid bool)
}
