package mo_sendas

import "encoding/json"

type SendAs struct {
	Raw                json.RawMessage
	SendAsEmail        string `json:"send_as_email" path:"sendAsEmail"`
	DisplayName        string `json:"display_name" path:"displayName"`
	ReplyToAddress     string `json:"reply_to_address" path:"replyToAddress"`
	Signature          string `json:"signature" path:"signature"`
	IsPrimary          bool   `json:"is_primary" path:"isPrimary"`
	IsDefault          bool   `json:"is_default" path:"isDefault"`
	TreatAsAlias       string `json:"treat_as_alias" path:"treatAsAlias"`
	VerificationStatus string `json:"verification_status" path:"verificationStatus"`
}
