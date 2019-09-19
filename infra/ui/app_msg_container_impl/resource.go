package app_msg_container_impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"text/template"
)

type Resource struct {
	Messages map[string]string
	Prefix   string
}

func (z *Resource) WithPrefix(prefix string) app_msg_container.Container {
	return &Resource{
		Messages: z.Messages,
		Prefix:   prefix,
	}
}

func (z *Resource) Exists(key string) bool {
	_, ok := z.Messages[keyWithPrefix(z.Prefix, key)]
	return ok
}

func (z *Resource) Compile(m app_msg.Message) string {
	key := keyWithPrefix(z.Prefix, m.Key())
	if msg, ok := z.Messages[key]; !ok {
		c := Alt{}
		return c.Compile(m)
	} else {
		params := make(map[string]interface{})
		for _, p := range m.Params() {
			for k, v := range p {
				params[k] = v
			}
		}
		t, err := template.New(key).Parse(msg)
		if err != nil {
			app_root.Log().Warn("Unable to compile message",
				zap.String("key", key),
				zap.String("msg", msg),
				zap.Error(err),
			)
			c := Alt{}
			return c.Compile(m)
		}
		var buf bytes.Buffer
		if err = t.Execute(&buf, params); err != nil {
			app_root.Log().Warn("Unable to format message",
				zap.String("key", key),
				zap.String("msg", msg),
				zap.Error(err),
			)
			c := Alt{}
			return c.Compile(m)
		}

		return buf.String()
	}
}

func NewResource(lang language.Tag, box *rice.Box) (c app_msg_container.Container, err error) {
	l := app_root.Log().With(zap.String("lang", lang.String()))

	resName := "messages.json"
	b, _ := lang.Base()
	if b.String() != "en" {
		resName = fmt.Sprintf("messages_%s.json", b)
	}
	l = l.With(zap.String("name", resName))
	resData, err := box.Bytes(resName)
	if err != nil {
		return nil, err
	}
	resMsgs := make(map[string]string)
	if err = json.Unmarshal(resData, &resMsgs); err != nil {
		l.Error("Unable to unmarshal message resource", zap.Error(err))
		return nil, err
	}

	return &Resource{
		Messages: resMsgs,
	}, nil
}
