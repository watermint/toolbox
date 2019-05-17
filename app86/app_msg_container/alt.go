package app_msg_container

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/app86/app_msg"
)

type Alt struct {
}

func (Alt) Exists(key string) bool {
	return false
}

func (Alt) Compile(m app_msg.Message) string {
	params := make(map[string]interface{})
	for _, p := range m.Params() {
		param := p()
		params[param.Key] = param.Value
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
