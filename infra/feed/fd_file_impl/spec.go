package fd_file_impl

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_doc"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/infra/util/ut_reflect"
)

func newSpec(rf *RowFeed) fd_file.Spec {
	s := &Spec{rf: rf}
	s.base = ut_reflect.Key(app.Pkg, rf.Model())
	s.colDesc = make(map[string]app_msg.Message)
	s.colExample = make(map[string]app_msg.Message)

	for _, col := range rf.fields {
		s.colDesc[col] = app_msg.M(s.base + "." + col + ".desc")
		s.colExample[col] = app_msg.M(s.base + "." + col + ".example")
	}
	return s
}

type Spec struct {
	rf         *RowFeed
	base       string
	colDesc    map[string]app_msg.Message
	colExample map[string]app_msg.Message
}

func (z *Spec) Doc(ui app_ui.UI) *rc_doc.Feed {
	cols := make([]*rc_doc.FeedColumn, 0)
	for _, col := range z.Columns() {
		cols = append(cols, &rc_doc.FeedColumn{
			Name:    col,
			Desc:    ui.TextOrEmpty(z.ColumnDesc(col)),
			Example: ui.TextOrEmpty(z.ColumnExample(col)),
		})
	}

	return &rc_doc.Feed{
		Name:    z.Name(),
		Desc:    ui.TextOrEmpty(z.Desc()),
		Columns: cols,
	}
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
