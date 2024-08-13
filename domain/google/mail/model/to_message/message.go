package to_message

import (
	"encoding/base64"
	"net/mail"
	"strings"
)

const (
	HeaderTo      = "To"
	HeaderCc      = "Cc"
	HeaderBcc     = "Bcc"
	HeaderFrom    = "From"
	HeaderSubject = "Subject"
)

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MessagePartBody struct {
	AttachmentId string `json:"attachmentId,omitempty"`
	Size         int    `json:"size"`
	Data         string `json:"data"`
}

type MessagePart struct {
	PartId   string          `json:"partId,omitempty"`
	MimeType string          `json:"mimeType,omitempty"`
	Filename string          `json:"filename,omitempty"`
	Headers  []Header        `json:"headers,omitempty"`
	Body     MessagePartBody `json:"body"`
}

func (z MessagePart) WithHeader(name, value string) MessagePart {
	headers := make([]Header, 0)
	headers = append(headers, z.Headers...)
	headers = append(headers, Header{
		Name:  name,
		Value: value,
	})

	z.Headers = headers
	return z
}

func (z MessagePart) WithTo(addrs ...string) MessagePart {
	return z.WithHeader(HeaderTo, strings.Join(addrs, ","))
}

func (z MessagePart) WithCc(addrs ...string) MessagePart {
	return z.WithHeader(HeaderCc, strings.Join(addrs, ","))
}

func (z MessagePart) WithBcc(addrs ...string) MessagePart {
	return z.WithHeader(HeaderBcc, strings.Join(addrs, ","))
}

func (z MessagePart) WithFrom(addr string) MessagePart {
	return z.WithHeader(HeaderFrom, addr)
}

func (z MessagePart) WithSubject(s string) MessagePart {
	return z.WithHeader(HeaderSubject, s)
}

func (z MessagePart) WithBodyText(text string) MessagePart {
	t := base64.URLEncoding.EncodeToString([]byte(text))
	z.Body = MessagePartBody{
		Size: len(t),
		Data: t,
	}
	return z
}

type Message struct {
	Payload MessagePart `json:"payload"`
}

func Address(name, email string) string {
	adr := &mail.Address{
		Name:    name,
		Address: email,
	}
	return adr.String()
}
