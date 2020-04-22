package app_msg_container_impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"go.uber.org/zap"
	"text/template"
)

type Resource struct {
	Messages map[string]string
}

func (z *Resource) Text(key string) string {
	if msg, ok := z.Messages[key]; !ok {
		c := Alt{}
		return c.Text(key)
	} else {
		return msg
	}
}

func (z *Resource) Exists(key string) bool {
	_, ok := z.Messages[key]
	return ok
}

func (z *Resource) Compile(m app_msg.Message) string {
	key := m.Key()
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

func newFromBytes(la lang.Lang, loader func(name string) ([]byte, error)) (c app_msg_container.Container, err error) {
	l := app_root.Log().With(zap.String("lang", la.String()))

	resName := fmt.Sprintf("messages%s.json", la.Suffix())
	l = l.With(zap.String("name", resName))
	resData, err := loader(resName)
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

func New(la lang.Lang, ctl app_control.Control) (c app_msg_container.Container, err error) {
	return newFromBytes(la, func(name string) (i []byte, e error) {
		return ctl.Resource(name)
	})
}

func NewResource(la lang.Lang, box *rice.Box) (c app_msg_container.Container, err error) {
	return newFromBytes(la, func(name string) (i []byte, e error) {
		return box.Bytes(name)
	})
}
