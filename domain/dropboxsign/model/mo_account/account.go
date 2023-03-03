package mo_account

import "encoding/json"

type Account struct {
	Raw          json.RawMessage `json:"-"`
	AccountId    string          `path:"account_id" json:"account_id"`
	EmailAddress string          `path:"email_address" json:"email_address"`
	Locale       string          `path:"locale" json:"locale"`
}
