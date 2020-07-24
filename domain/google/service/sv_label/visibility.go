package sv_label

const (
	VisibilityLabelListHide         = "labelHide"
	VisibilityLabelListShow         = "labelShow"
	VisibilityLabelListShowIfUnread = "labelShowIfUnread"
	VisibilityMessageListHide       = "hide"
	VisibilityMessageListShow       = "show"
)

var (
	VisibilityLabelList = []string{
		VisibilityLabelListHide,
		VisibilityLabelListShow,
		VisibilityLabelListShowIfUnread,
	}
	VisibilityMessageList = []string{
		VisibilityMessageListHide,
		VisibilityMessageListShow,
	}
)
