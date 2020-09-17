package mo_filter

import "github.com/watermint/toolbox/infra/ui/app_msg"

type MsgFilter struct {
	DescFilterName       app_msg.Message
	DescFilterNamePrefix app_msg.Message
	DescFilterNameSuffix app_msg.Message
	DescFilterEmail      app_msg.Message
}

var (
	MFilter = app_msg.Apply(&MsgFilter{}).(*MsgFilter)
)
