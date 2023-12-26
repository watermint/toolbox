package fd_file_impl

import (
	"github.com/watermint/toolbox/essentials/go/es_reflect"
	"github.com/watermint/toolbox/infra/doc/dc_recipe"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func newSpec(rf *RowFeed) fd_file.Spec {
	s := &Spec{rf: rf}
	if rf.Model() == nil {
		panic("Feed model is not defined")
	}
	s.base = es_reflect.Key(rf.Model())
	s.colDesc = make(map[string]app_msg.Message)
	s.colExample = make(map[string]app_msg.Message)

	for _, col := range rf.fields {
		s.colDesc[col] = app_msg.CreateMessage(s.base + "." + col + ".desc")
		s.colExample[col] = app_msg.CreateMessage(s.base + "." + col + ".example")
	}
	return s
}

type Spec struct {
	rf         *RowFeed
	base       string
	colDesc    map[string]app_msg.Message
	colExample map[string]app_msg.Message
}

func (z *Spec) Doc(ui app_ui.UI) *dc_recipe.Feed {
	cols := make([]*dc_recipe.FeedColumn, 0)
	for _, col := range z.Columns() {
		cols = append(cols, &dc_recipe.FeedColumn{
			Name:    col,
			Desc:    ui.TextOrEmpty(z.ColumnDesc(col)),
			Example: ui.TextOrEmpty(z.ColumnExample(col)),
		})
	}

	return &dc_recipe.Feed{
		Name:    z.Name(),
		Desc:    ui.TextOrEmpty(z.Desc()),
		Columns: cols,
	}
}

func (z *Spec) Name() string {
	return z.rf.name
}

func (z *Spec) Desc() app_msg.Message {
	return app_msg.CreateMessage(z.base + ".desc")
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
