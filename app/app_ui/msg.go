package app_ui

import (
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app/app_util"
	"go.uber.org/zap"
)

type UIMessage interface {
	// UIMessage with arguments for printf style data
	WithArg(a ...interface{}) UIMessage

	// UIMessage with template data structure
	WithData(d interface{}) UIMessage

	// Text part of this UIMessage
	Text() string

	// Tell message
	Tell()

	// Tell error message
	TellError()

	// Tell done
	TellDone()

	// Tell success
	TellSuccess()

	// Tell failure
	TellFailure()

	// Ask retry with retry message. Returns true when
	// the user/client agreed retry
	AskRetry() bool

	// Ask a text. UI ask text as required option but,
	// a user/client can enter empty string.
	AskText() string
}

type UIMessageContainer struct {
	resources     *rice.Box
	baseMessages  map[string]UIMessage
	userInterface UI
	logger        *zap.Logger
	isTest        bool
}

func NewUIMessageContainer(bx *rice.Box, ui UI, logger *zap.Logger) *UIMessageContainer {
	return &UIMessageContainer{
		resources:     bx,
		userInterface: ui,
		logger:        logger,
	}
}

func (z *UIMessageContainer) Load() {
	z.baseMessages = make(map[string]UIMessage)
	if z.resources == nil {
		z.logger.Debug("no box found. skip loading messages.")
		z.isTest = true
		return
	}

	baseAppMsgBytes, err := z.resources.Bytes("messages.json")
	if err != nil {
		z.logger.Error("unable to load base app msg `messages.json`", zap.Error(err))
	} else {
		baseAppMsg := make(map[string]string)
		err = json.Unmarshal(baseAppMsgBytes, &baseAppMsg)
		if err != nil {
			z.logger.Error("unable to unmarshal app msg `messages.json`", zap.Error(err))
		} else {
			z.baseMessages = NewMessageMap(baseAppMsg, z.userInterface, z.logger)
		}
	}
}

func (z *UIMessageContainer) Msg(key string) UIMessage {
	if z.baseMessages == nil {
		return NewAltMessage(key, z.userInterface)
	}
	if m, e := z.baseMessages[key]; e {
		return m
	}
	if !z.isTest {
		z.logger.Error("resource not found for key", zap.String("key", key))
	}
	return NewAltMessage(key, z.userInterface)
}

func NewAltMessage(key string, ui UI) UIMessage {
	return &AltMessage{
		key:           key,
		userInterface: ui,
	}
}

func NewTextMessage(text string, ui UI, logger *zap.Logger) UIMessage {
	return &TextMessage{
		text:          text,
		userInterface: ui,
		logger:        logger,
	}
}

type TextMessage struct {
	text          string
	args          []interface{}
	tmplData      interface{}
	logger        *zap.Logger
	userInterface UI
}

func (z *TextMessage) WithArg(a ...interface{}) UIMessage {
	return &TextMessage{
		text:          z.text,
		args:          a,
		logger:        z.logger,
		userInterface: z.userInterface,
	}
}

func (z *TextMessage) WithData(d interface{}) UIMessage {
	return &TextMessage{
		text:          z.text,
		tmplData:      d,
		logger:        z.logger,
		userInterface: z.userInterface,
	}
}

func (z *TextMessage) Text() string {
	if z.tmplData != nil {
		t, err := app_util.CompileTemplate(z.text, z.tmplData)
		if err != nil {
			z.logger.Error(
				"Unable to compile template",
				zap.String("tmpl", z.text),
				zap.Any("data", z.tmplData),
				zap.Error(err),
			)
			return z.text
		}
		return t
	} else if z.args != nil && len(z.args) > 0 {
		return fmt.Sprintf(z.text, z.args...)
	} else {
		return z.text
	}
}

func (z *TextMessage) Tell() {
	z.userInterface.Tell(z)
}

func (z *TextMessage) TellError() {
	z.userInterface.TellError(z)
}

func (z *TextMessage) TellDone() {
	z.userInterface.TellDone(z)
}

func (z *TextMessage) TellSuccess() {
	z.userInterface.TellSuccess(z)
}

func (z *TextMessage) TellFailure() {
	z.userInterface.TellFailure(z)
}

func (z *TextMessage) AskRetry() bool {
	return z.userInterface.AskRetry(z)
}

func (z *TextMessage) AskText() string {
	return z.userInterface.AskText(z)
}

type AltMessage struct {
	key           string
	args          []interface{}
	tmplData      interface{}
	userInterface UI
}

func (z *AltMessage) WithArg(a ...interface{}) UIMessage {
	return &AltMessage{
		key:           z.key,
		args:          a,
		userInterface: z.userInterface,
	}
}

func (z *AltMessage) WithData(d interface{}) UIMessage {
	return &AltMessage{
		key:           z.key,
		tmplData:      d,
		userInterface: z.userInterface,
	}
}

func (z *AltMessage) Text() string {
	if z.tmplData != nil {
		d, err := json.Marshal(z.tmplData)
		if err == nil {
			return fmt.Sprintf("key:%s data:%s", z.key, string(d))
		}
		return fmt.Sprintf("key:%s data:%v", z.key, z.tmplData)
	} else if z.args != nil && len(z.args) > 0 {
		a, err := json.Marshal(z.args)
		if err == nil {
			return fmt.Sprintf("key:%s args:%s", z.key, string(a))
		}
		return fmt.Sprintf("key:%s args:%v", z.key, a)
	} else {
		return fmt.Sprintf("key:%s", z.key)
	}
}

func (z *AltMessage) Tell() {
	z.userInterface.Tell(z)
}

func (z *AltMessage) TellError() {
	z.userInterface.TellError(z)
}

func (z *AltMessage) TellDone() {
	z.userInterface.TellDone(z)
}

func (z *AltMessage) TellSuccess() {
	z.userInterface.TellSuccess(z)
}

func (z *AltMessage) TellFailure() {
	z.userInterface.TellFailure(z)
}

func (z *AltMessage) AskRetry() bool {
	return z.userInterface.AskRetry(z)
}

func (z *AltMessage) AskText() string {
	return z.userInterface.AskText(z)
}

type Message struct {
	message       string
	args          []interface{}
	tmplData      interface{}
	logger        *zap.Logger
	userInterface UI
}

func (z *Message) Tell() {
	z.userInterface.Tell(z)
}

func (z *Message) TellError() {
	z.userInterface.TellError(z)
}

func (z *Message) TellDone() {
	z.userInterface.TellDone(z)
}

func (z *Message) TellSuccess() {
	z.userInterface.TellSuccess(z)
}

func (z *Message) TellFailure() {
	z.userInterface.TellFailure(z)
}

func (z *Message) AskRetry() bool {
	return z.userInterface.AskRetry(z)
}

func (z *Message) AskText() string {
	return z.userInterface.AskText(z)
}

func (z *Message) WithData(d interface{}) UIMessage {
	return &Message{
		message:       z.message,
		args:          z.args,
		tmplData:      d,
		logger:        z.logger,
		userInterface: z.userInterface,
	}
}

func (z *Message) WithArg(a ...interface{}) UIMessage {
	return &Message{
		message:       z.message,
		args:          a,
		tmplData:      z.tmplData,
		logger:        z.logger,
		userInterface: z.userInterface,
	}
}

func (z *Message) Text() string {
	if z.tmplData != nil {
		t, err := app_util.CompileTemplate(z.message, z.tmplData)
		if err != nil {
			z.logger.Error(
				"Unable to compile template",
				zap.String("tmpl", z.message),
				zap.Any("data", z.tmplData),
				zap.Error(err),
			)
			return z.message
		}
		return t
	} else if z.args != nil && len(z.args) > 0 {
		return fmt.Sprintf(z.message, z.args...)
	} else {
		return z.message
	}
}

func NewMessageMap(res map[string]string, ui UI, log *zap.Logger) map[string]UIMessage {
	mm := make(map[string]UIMessage)
	for k, m := range res {
		mm[k] = &Message{
			message:       m,
			userInterface: ui,
			logger:        log,
		}
	}
	return mm
}
