package gui

import "encoding/base64"

type Command struct {
	Command string `uri:"command" binding:"required"`
}

func (z Command) EncodeCommandUrl() string {
	return base64.URLEncoding.EncodeToString([]byte(z.Command))
}

func (z Command) DecodeCommandName() (name string, err error) {
	nameRaw, err := base64.URLEncoding.DecodeString(z.Command)
	if err != nil {
		return "", err
	}
	return string(nameRaw), nil
}
