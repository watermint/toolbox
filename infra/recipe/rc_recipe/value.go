package rc_recipe

import (
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"reflect"
)

type Value interface {
	// Returns forked instance when the type is acceptable
	// Otherwise returns nil
	Accept(t reflect.Type, v0 interface{}, name string) Value

	// Return value reference of the instance
	Bind() interface{}

	// Initialize instance, and returns the instance to set
	Init() (v interface{})

	// Apply preset value
	ApplyPreset(v0 interface{})

	// Apply internal state (bind'ed value) to the instance.
	Apply() (v interface{})

	// Debug information
	Debug() interface{}

	// Spin up for run
	SpinUp(ctl app_control.Control) error

	// Spin down after run
	SpinDown(ctl app_control.Control) error

	// Value spec
	Spec() (typeName string, typeAttr interface{})
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

type ValueReports interface {
	Value

	// Returns reports that used by this value
	Reports() map[string]rp_model.Report
}

type ValueConn interface {
	Value

	// True when the value is type of connection, and return the instance of the connection
	Conn() (conn api_conn.Connection, valid bool)
}

type ValueMessage interface {
	Value

	// True when the value is type of message, and return the instance eof the conection
	Message() (msg app_msg.Message, valid bool)
}
