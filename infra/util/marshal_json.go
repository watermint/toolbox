package util

import (
	"encoding/json"
	"fmt"
)

func MarshalObjectToString(obj interface{}) string {
	j, err := json.Marshal(obj)
	if err != nil {
		return fmt.Sprintf("%v", obj)
	}
	return string(j)
}
