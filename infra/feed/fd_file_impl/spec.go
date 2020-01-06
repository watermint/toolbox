package fd_file_impl

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
)

func newSpec(rf *RowFeed) fd_file.Spec {
	s := &Spec{rf: rf}
	base := ut_reflect.Key(app.Pkg, rf.Model())
	s.colDesc = make(map[string]app_msg.Message)
	s.colExample = make(map[string]app_msg.Message)

	for _, col := range rf.fields {
		s.colDesc[col] = app_msg.M(base + "." + col + ".desc")
		s.colExample[col] = app_msg.M(base + "." + col + ".example")
	}
	return s
}

type Spec struct {
	rf         *RowFeed
	base       string
	colDesc    map[string]app_msg.Message
	colExample map[string]app_msg.Message
}

func (z *Spec) Name() string {
	return z.rf.name
}

func (z *Spec) Desc() app_msg.Message {
	return app_msg.M(z.base + ".desc")
}

func (z *Spec) Columns() []string {
	return z.rf.fields
}

func (z *Spec) ColumnDesc(col string) app_msg.Message {
	return z.colDesc[col]
}

func (z *Spec) ColumnExample(col string) app_msg.Message {
	return z.colExample[col]
}
