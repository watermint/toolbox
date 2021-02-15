package mo_message

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_label"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"net/mail"
)

type Message struct {
	Raw      json.RawMessage
	Id       string `json:"id" path:"id"`
	ThreadId string `json:"thread_id" path:"threadId"`
	Date     string `json:"date" path:"payload.headers.#(name==\"Date\").value"`
	Subject  string `json:"subject" path:"payload.headers.#(name==\"Subject\").value"`
	To       string `json:"to" path:"payload.headers.#(name==\"To\").value"`
	Cc       string `json:"cc" path:"payload.headers.#(name==\"Cc\").value"`
	From     string `json:"from" path:"payload.headers.#(name==\"From\").value"`
	ReplyTo  string `json:"reply_to" path:"payload.headers.#(name==\"Reply-To\").value"`
}

func (z Message) Processed() (p Processed, err error) {
	l := esl.Default()

	// Label Ids
	j, err := es_json.Parse(z.Raw)
	if err != nil {
		l.Debug("Unable to parse message", esl.Error(err))
		return p, err
	}
	if labelIds, found := j.FindArray("labelIds"); found {
		p.LabelIds = make([]string, 0)
		for _, labelId := range labelIds {
			if id, valid := labelId.String(); valid {
				p.LabelIds = append(p.LabelIds, id)
			}
		}
	}

	// Id & orig data
	p.Id = z.Id
	p.ThreadId = z.ThreadId
	p.Original = z.Raw
	p.Subject = z.Subject

	if v, found := j.FindString("historyId"); found {
		p.HistoryId = v
	}

	// Size
	if v, found := j.FindNumber("sizeEstimate"); found {
		p.SizeEstimate = v.Int()
	}

	// Date
	header := mail.Header{
		"Date": []string{z.Date},
	}
	date, err := header.Date()
	if err != nil {
		l.Debug("Unable to parse date", esl.Error(err), esl.String("date", z.Date))
	} else {
		p.Date8601 = date.Format("2006-01-02T15:04:05")
		p.DateUnix = date.Unix()
	}
	if v, found := j.FindNumber("internalDate"); found {
		p.DateInternal = v.Int64() / 1000
	} else if v, found := j.FindString("internalDate"); found {
		n := es_number.New(v)
		p.DateInternal = n.Int64() / 1000
	}

	// Addresses
	if z.To != "" {
		addrs, err := mail.ParseAddressList(z.To)
		if err != nil {
			l.Debug("Unable to parse `to`", esl.Error(err), esl.String("to", z.To))
			return p, err
		}
		p.To = make([]*Address, 0)
		for _, addr := range addrs {
			p.To = append(p.To, &Address{
				Name:    addr.Name,
				Address: addr.Address,
			})
		}
	}

	if z.Cc != "" {
		addrs, err := mail.ParseAddressList(z.Cc)
		if err != nil {
			l.Debug("Unable to parse `cc`", esl.Error(err), esl.String("cc", z.Cc))
			return p, err
		}
		p.Cc = make([]*Address, 0)
		for _, addr := range addrs {
			p.Cc = append(p.Cc, &Address{
				Name:    addr.Name,
				Address: addr.Address,
			})
		}
	}

	if z.From != "" {
		addr, err := mail.ParseAddress(z.From)
		if err != nil {
			l.Debug("unable to parse `from", esl.Error(err), esl.String("from", z.From))
			return p, err
		}
		p.From = &Address{
			Name:    addr.Name,
			Address: addr.Address,
		}
	}

	if z.ReplyTo != "" {
		addr, err := mail.ParseAddress(z.ReplyTo)
		if err != nil {
			l.Debug("unable to parse `reply-to", esl.Error(err), esl.String("reply-to", z.ReplyTo))
			return p, err
		}
		p.ReplyTo = &Address{
			Name:    addr.Name,
			Address: addr.Address,
		}
	}
	return p, nil
}

type Address struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Processed struct {
	Id              string            `json:"id"`
	ThreadId        string            `json:"thread_id"`
	HistoryId       string            `json:"history_id"`
	DateInternal    int64             `json:"date_internal"`
	Date8601        string            `json:"date_8601"`
	DateUnix        int64             `json:"date_unix"`
	Subject         string            `json:"subject"`
	To              []*Address        `json:"to,omitempty"`
	Cc              []*Address        `json:"cc,omitempty"`
	From            *Address          `json:"from,omitempty"`
	ReplyTo         *Address          `json:"reply_to,omitempty"`
	LabelIds        []string          `json:"label_ids"`
	LabelNames      []string          `json:"label_names"`
	LabelTypeUser   []*mo_label.Label `json:"label_type_user"`
	LabelTypeSystem []*mo_label.Label `json:"label_type_system"`
	SizeEstimate    int               `json:"size_estimate"`
	Original        json.RawMessage   `json:"original"`
}
