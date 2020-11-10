package app_msg_container_impl

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Alt struct {
}

func (z Alt) Lang() lang.Lang {
	return lang.Default
}

func (z Alt) Text(key string) string {
	return AltText(key)
}

func (z Alt) Exists(msg app_msg.Message) bool {
	return false
}

func (z Alt) ExistsKey(key string) bool {
	return false
}

func (z Alt) Compile(m app_msg.Message) string {
	return AltCompile(m)
}

func AltText(key string) string {
	return fmt.Sprintf("Key[%s]", key)
}

func AltCompile(m app_msg.Message) string {
	params := make(map[string]interface{})
	for _, p := range m.Params() {
		for k, v := range p {
			params[k] = v
		}
	}

	alt := struct {
		Key    string                 `json:"key"`
		Params map[string]interface{} `json:"params"`
	}{
		Key:    m.Key(),
		Params: params,
	}
	if j, err := json.Marshal(&alt); err != nil {
		return fmt.Sprintf("Key[%s] Param[%v]", m.Key(), params)
	} else {
		return string(j)
	}
}
