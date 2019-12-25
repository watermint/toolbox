package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"reflect"
)

var (
	valueTypes = []Value{
		newValueBool(),
		newValueRcConnBusinessInfo(rc_conn_impl.DefaultPeerName),
		newValueRpModelRowReport(""),
		newValueFdFileRowFeed(""),
	}
)

// Find value of type.
// Returns nil when the value type is not supported
func valueOfType(t reflect.Type, name string) Value {
	for _, vt := range valueTypes {
		if v := vt.Accept(t, name); v != nil {
			return v
		}
	}
	return nil
}

type repositoryImpl struct {
}

func (z repositoryImpl) Conns() map[string]rc_conn.ConnDropboxApi {
	panic("implement me")
}

func (z *repositoryImpl) Feeds() map[string]fd_file.RowFeed {
	panic("implement me")
}

func (z *repositoryImpl) FeedSpecs() map[string]fd_file.Spec {
	panic("implement me")
}

func (z *repositoryImpl) Reports() map[string]rp_model.Report {
	panic("implement me")
}

func (z *repositoryImpl) Init() {
	panic("implement me")
}

func (z *repositoryImpl) Fork(ctl app_control.Control) {
	panic("implement me")
}

func (z *repositoryImpl) Apply() {
	panic("implement me")
}

func (z *repositoryImpl) SpinUp(ctl app_control.Control) error {
	panic("implement me")
}

func (z *repositoryImpl) SpinDown() error {
	panic("implement me")
}

func (z *repositoryImpl) ApplyFlags(f *flag.FlagSet, ui app_ui.UI) {
	panic("implement me")
}

func (z *repositoryImpl) FlagFieldDesc(name string) app_msg.Message {
	panic("implement me")
}

func (z *repositoryImpl) Debug() map[string]interface{} {
	panic("implement me")
}
