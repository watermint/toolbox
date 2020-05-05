package app_msg_container_impl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"text/template"
)

// Load single language container
func NewSingle(la lang.Lang) (c app_msg_container.Container, err error) {
	return newFromBytes(la, func(name string) (i []byte, e error) {
		return app_resource.Bundle().Messages().Bytes(name)
	})
}

func NewSingleWithMessages(msgs map[string]string) app_msg_container.Container {
	return &sglContainer{
		messages: msgs,
	}
}

func newFromBytes(la lang.Lang, loader func(name string) ([]byte, error)) (c app_msg_container.Container, err error) {
	l := es_log.Default().With(es_log.String("lang", la.String()))

	resName := fmt.Sprintf("messages%s.json", la.Suffix())
	l = l.With(es_log.String("name", resName))
	resData, err := loader(resName)
	if err != nil {
		return nil, err
	}
	resMsgs := make(map[string]string)
	if err = json.Unmarshal(resData, &resMsgs); err != nil {
		l.Error("Unable to unmarshal message resource", es_log.Error(err))
		return nil, err
	}

	return &sglContainer{
		messages: resMsgs,
	}, nil
}

// Single language container.
type sglContainer struct {
	messages map[string]string
}

func (z sglContainer) Text(key string) string {
	if msg, ok := z.messages[key]; !ok {
		return AltText(key)
	} else {
		return msg
	}
}

func (z *sglContainer) Exists(key string) bool {
	_, ok := z.messages[key]
	return ok
}

func (z *sglContainer) Compile(m app_msg.Message) string {
	l := es_log.Default()
	key := m.Key()
	if msg, ok := z.messages[key]; !ok {
		return AltCompile(m)
	} else {
		params := make(map[string]interface{})
		for _, p := range m.Params() {
			for k, v := range p {
				params[k] = v
			}
		}
		t, err := template.New(key).Parse(msg)
		if err != nil {
			l.Warn("Unable to compile message",
				es_log.String("key", key),
				es_log.String("msg", msg),
				es_log.Error(err),
			)
			return AltCompile(m)
		}
		var buf bytes.Buffer
		if err = t.Execute(&buf, params); err != nil {
			l.Warn("Unable to format message",
				es_log.String("key", key),
				es_log.String("msg", msg),
				es_log.Error(err),
			)
			return AltCompile(m)
		}

		return buf.String()
	}
}
