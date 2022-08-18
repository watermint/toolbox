package dc_contributor

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgDeveloper struct {
	DeveloperDesc                  app_msg.Message
	RecipeValueTitle               app_msg.Message
	RecipeValueTypeImpl            app_msg.Message
	RecipeValueTypeConn            app_msg.Message
	RecipeValueTypeConns           app_msg.Message
	RecipeValueTypeCustomValueText app_msg.Message
	RecipeValueTypeErrorHandler    app_msg.Message
	RecipeValueTypeFeed            app_msg.Message
	RecipeValueTypeGridDataInput   app_msg.Message
	RecipeValueTypeGridDataOutput  app_msg.Message
	RecipeValueTypeJsonInput       app_msg.Message
	RecipeValueTypeMessage         app_msg.Message
	RecipeValueTypeMessages        app_msg.Message
	RecipeValueTypeReport          app_msg.Message
	RecipeValueTypeReports         app_msg.Message
	RecipeValueTypeTextInput       app_msg.Message
	RecipeValueConnValueTypes      app_msg.Message
	RecipeValueConnScopeLabel      app_msg.Message
	RecipeValueConnServiceName     app_msg.Message
}

var (
	MDeveloper = app_msg.Apply(&MsgDeveloper{}).(*MsgDeveloper)
)

func Docs(media dc_index.MediaType) []dc_section.Document {
	return []dc_section.Document{
		&RecipeValues{},
	}
}
