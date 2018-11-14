package oper

import (
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/poc/oper/oper_ui"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

type Operation interface {
}

type Executable interface {
	Exec()
}

type Group interface {
	Operations() []Operation
}

type Resource struct {
	Title   string            `json:"title,omitempty"`
	Desc    string            `json:"desc,omitempty"`
	Options map[string]string `json:"options,omitempty"`
}

type Context struct {
	Logger *zap.Logger
	Box    *rice.Box
	Lang   language.Tag
	UI     oper_ui.UI
}

func (c Context) LangBase() string {
	base, _, _ := c.Lang.Raw()
	return base.String()
}
