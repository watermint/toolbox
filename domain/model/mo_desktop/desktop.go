package mo_desktop

const (
	TypePersonal = "personal"
	TypeBusiness = "business"
)

type Desktop struct {
	DropboxType      string `json:"-"`
	Path             string `json:"path"`
	Host             int64  `json:"host"`
	IsTeam           bool   `json:"is_team"`
	SubscriptionType string `json:"subscription_type"`
}
