package app_msg_container_impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_msg"
	"github.com/watermint/toolbox/app86/app_msg_container"
	"github.com/watermint/toolbox/app86/app_root"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"text/template"
)

type Resource struct {
	Messages map[string]string
}

func (z *Resource) Exists(key string) bool {
	_, ok := z.Messages[key]
	return ok
}

func (z *Resource) Compile(m app_msg.Message) string {
	if msg, ok := z.Messages[m.Key()]; !ok {
		c := Alt{}
		return c.Compile(m)
	} else {
		params := make(map[string]interface{})
		for _, p := range m.Params() {
			param := p()
			params[param.Key] = param.Value
		}
		t, err := template.New(m.Key()).Parse(msg)
		if err != nil {
			app_root.Root.Log().Warn("Unable to compile message",
				zap.String("key", m.Key()),
				zap.String("msg", msg),
				zap.Error(err),
			)
			c := Alt{}
			return c.Compile(m)
		}
		var buf bytes.Buffer
		if err = t.Execute(&buf, params); err != nil {
			app_root.Root.Log().Warn("Unable to format message",
				zap.String("key", m.Key()),
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
	if lang != language.English {
		b, _ := lang.Base()
		resName = fmt.Sprintf("messages_%s.json", b)
	}
	l = l.With(zap.String("name", resName))
	resData, err := box.Bytes(resName)
	if err != nil {
		l.Error("Unable to load resource data", zap.Error(err))
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
