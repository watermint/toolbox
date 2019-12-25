package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
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

	// Returns connections that requested by the recipe
	Conns() map[string]rc_conn.ConnDropboxApi

	// Initialize given object & this value repository
	Init()

	// Fork this value repository
	Fork(ctl app_control.Control)

	// Apply values in the repository to the
	Apply()

	// Prepare values for run recipe
	SpinUp(ctl app_control.Control) error

	// Spin down value
	SpinDown() error

	// Apply flag set
	ApplyFlags(f *flag.FlagSet, ui app_ui.UI)

	// Description of the field
	FlagFieldDesc(name string) app_msg.Message

	// Serialize values for debug
	Debug() map[string]interface{}
}

type Value interface {
	// Returns forked instance when the type is acceptable
	// Otherwise returns nil
	Accept(t reflect.Type, name string) Value

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

	// True when the value is type of feed, and returns instance of the feed
	IsFeed() (feed fd_file.RowFeed, valid bool)

	// True when the value is type of report, and returns instance of the report
	IsReport() (report rp_model.Report, valid bool)

	// True when the value is type of connection, and return instance of the connection
	IsConn() (conn rc_conn.ConnDropboxApi, valid bool)
}
