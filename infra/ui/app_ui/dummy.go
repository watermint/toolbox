package app_ui

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func NewDummy() UI {
	app_root.Log().Warn("Dummy UI generated")
	return &Dummy{}
}

type Dummy struct {
}

func (z *Dummy) Header(m app_msg.Message) {
}

func (z *Dummy) Text(m app_msg.Message) string {
	return ""
}

func (z *Dummy) TextOrEmpty(m app_msg.Message) string {
	return ""
}

func (z *Dummy) Info(m app_msg.Message) {
}

func (z *Dummy) Error(m app_msg.Message) {
}

func (z *Dummy) HeaderK(key string, p ...app_msg.P) {
}

func (z *Dummy) InfoK(key string, p ...app_msg.P) {
}

func (z *Dummy) InfoTable(name string) Table {
	return &DummyTable{}
}

func (z *Dummy) ErrorK(key string, p ...app_msg.P) {
}

func (z *Dummy) Break() {
}

func (z *Dummy) TextK(key string, p ...app_msg.P) string {
	return ""
}

func (z *Dummy) TextOrEmptyK(key string, p ...app_msg.P) string {
	return ""
}

func (z *Dummy) AskCont(key string, p ...app_msg.P) (cont bool, cancel bool) {
	return false, true
}

func (z *Dummy) AskText(key string, p ...app_msg.P) (text string, cancel bool) {
	return "", true
}

func (z *Dummy) AskSecure(key string, p ...app_msg.P) (secure string, cancel bool) {
	return "", true
}

func (z *Dummy) OpenArtifact(path string) {
}

func (z *Dummy) Success(key string, p ...app_msg.P) {
}

func (z *Dummy) Failure(key string, p ...app_msg.P) {
}

func (z *Dummy) IsConsole() bool {
	return false
}

func (z *Dummy) IsWeb() bool {
	return false
}

type DummyTable struct {
}

func (z *DummyTable) Header(h ...app_msg.Message) {
}

func (z *DummyTable) HeaderRaw(h ...string) {
}

func (z *DummyTable) Row(m ...app_msg.Message) {
}

func (z *DummyTable) RowRaw(m ...string) {
}

func (z *DummyTable) Flush() {
}
