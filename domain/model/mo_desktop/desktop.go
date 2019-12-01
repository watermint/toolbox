package mo_desktop

const (
	TypePersonal = "personal"
	TypeBusiness = "business"
)

type Desktop struct {
	Path             string `path:"path" json:"path"`
	Host             int64  `path:"host" json:"host"`
	IsTeam           bool   `path:"is_team" json:"is_team"`
	SubscriptionType string `path:"subscription_type" json:"subscription_type"`
}
