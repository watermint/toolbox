package rc_value

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgValFdFileRowFeed struct {
	HeadFeed       app_msg.Message
	FeedDesc       app_msg.Message
	FeedSample     app_msg.Message
	HeadColName    app_msg.Message
	HeadColDesc    app_msg.Message
	HeadColExample app_msg.Message
}

type MsgRepository struct {
	ErrorMissingRequiredOption       app_msg.Message
	ErrorInvalidValue                app_msg.Message
	ErrorMoPathFsPathNotFound        app_msg.Message
	ErrorMoStringSelectInvalidOption app_msg.Message
	ProgressDoneValueInitialization  app_msg.Message
}

var (
	MValFdFileRowFeed = app_msg.Apply(&MsgValFdFileRowFeed{}).(*MsgValFdFileRowFeed)
	MRepository       = app_msg.Apply(&MsgRepository{}).(*MsgRepository)
)
