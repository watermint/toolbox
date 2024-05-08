package mo_signature

import (
	"github.com/watermint/toolbox/domain/dropboxsign/model/mo_list"
	"github.com/watermint/toolbox/domain/dropboxsign/model/mo_warning"
	"time"
)

func formatTime(t int64) string {
	if t == 0 {
		return ""
	} else {
		return time.Unix(t, 0).Format(time.RFC3339)
	}
}

type Request struct {
	SignatureRequestId    string       `json:"signature_request_id" path:"signature_request_id"`
	RequesterEmailAddress string       `json:"requester_email_address" path:"requester_email_address"`
	Title                 string       `json:"title" path:"title"`
	Subject               string       `json:"subject" path:"subject"`
	Message               string       `json:"message" path:"message"`
	CreatedAt             int64        `json:"created_at" path:"created_at"`
	CreatedAtRFC3339      string       `json:"created_at_rfc3339" path:"-"`
	ExpiresAt             int64        `json:"expires_at" path:"expires_at"`
	ExpiresAtRFC3339      string       `json:"expires_at_rfc3339" path:"-"`
	IsComplete            bool         `json:"is_complete" path:"is_complete"`
	IsDeclined            bool         `json:"is_declined" path:"is_declined"`
	Signatures            []*Signature `json:"signatures" path:"signatures"`
}

func (z Request) FormatTime() Request {
	z.ExpiresAtRFC3339 = formatTime(z.ExpiresAt)
	z.CreatedAtRFC3339 = formatTime(z.CreatedAt)
	return z
}

func (z Request) SignatureList() []*SignatureOfRequest {
	x := z.FormatTime()
	signatures := make([]*SignatureOfRequest, 0)
	for _, s := range z.Signatures {
		t := s.FormatTime()
		signatures = append(signatures, &SignatureOfRequest{
			SignatureRequestId:    z.SignatureRequestId,
			RequesterEmailAddress: z.RequesterEmailAddress,
			Title:                 z.Title,
			Subject:               z.Subject,
			Message:               z.Message,
			CreatedAt:             z.CreatedAt,
			CreatedAtRFC3339:      x.CreatedAtRFC3339,
			ExpiresAt:             z.ExpiresAt,
			ExpiresAtRFC3339:      x.ExpiresAtRFC3339,
			IsComplete:            z.IsComplete,
			IsDeclined:            z.IsDeclined,
			SignatureId:           s.SignatureId,
			SignerEmailAddress:    s.SignerEmailAddress,
			SignerName:            s.SignerName,
			SignerRole:            s.SignerRole,
			Order:                 s.Order,
			StatusCode:            s.StatusCode,
			DeclineReason:         s.DeclineReason,
			SignedAt:              s.SignedAt,
			SignedAtRFC3339:       t.SignedAtRFC3339,
		})
	}
	return signatures
}

type Signature struct {
	SignatureId        string `json:"signature_id" path:"signature_id"`
	SignerEmailAddress string `json:"signer_email_address" path:"signer_email_address"`
	SignerName         string `json:"signer_name" path:"signer_name"`
	SignerRole         string `json:"signer_role" path:"signer_role"`
	Order              int    `json:"order" path:"order"`
	StatusCode         string `json:"status_code" path:"status_code"`
	DeclineReason      string `json:"decline_reason" path:"decline_reason"`
	SignedAt           int64  `json:"signed_at" path:"signed_at"`
	SignedAtRFC3339    string `json:"signed_at_rfc3339" path:"-"`
}

func (z Signature) FormatTime() Signature {
	z.SignedAtRFC3339 = formatTime(z.SignedAt)
	return z
}

type SignatureOfRequest struct {
	SignatureRequestId string `json:"signature_request_id" path:"signature_request_id"`
	SignatureId        string `json:"signature_id" path:"signature_id"`

	RequesterEmailAddress string `json:"requester_email_address" path:"requester_email_address"`
	Title                 string `json:"title" path:"title"`
	Subject               string `json:"subject" path:"subject"`
	Message               string `json:"message" path:"message"`
	CreatedAt             int64  `json:"created_at" path:"created_at"`
	CreatedAtRFC3339      string `json:"created_at_rfc3339" path:"-"`
	ExpiresAt             int64  `json:"expires_at" path:"expires_at"`
	ExpiresAtRFC3339      string `json:"expires_at_rfc3339" path:"-"`
	IsComplete            bool   `json:"is_complete" path:"is_complete"`
	IsDeclined            bool   `json:"is_declined" path:"is_declined"`

	SignerEmailAddress string `json:"signer_email_address" path:"signer_email_address"`
	SignerName         string `json:"signer_name" path:"signer_name"`
	SignerRole         string `json:"signer_role" path:"signer_role"`
	Order              int    `json:"order" path:"order"`
	StatusCode         string `json:"status_code" path:"status_code"`
	DeclineReason      string `json:"decline_reason" path:"decline_reason"`
	SignedAt           int64  `json:"signed_at" path:"signed_at"`
	SignedAtRFC3339    string `json:"signed_at_rfc3339" path:"-"`
}

type RequestList struct {
	SignatureRequests []*Request            `json:"signature_requests" path:"signature_requests"`
	ListInfo          mo_list.ListInfo      `json:"list_info" path:"list_info"`
	Warnings          []*mo_warning.Warning `json:"warnings" path:"warnings"`
}
