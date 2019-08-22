package app_ui

import (
	"encoding/json"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/cloudfoundry-attic/jibber_jabber"
	"github.com/watermint/toolbox/legacy/app/app_util"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

var (
	supportedLanguages = []language.Tag{
		language.English,
		language.Japanese,
	}
)

type UIMessage interface {
	// UIMessage with arguments for printf style data
	WithArg(a ...interface{}) UIMessage

	// UIMessage with template data structure
	WithData(d interface{}) UIMessage

	// Text part of this UIMessage
	T() string

	// Tell message
	Tell()

	// Tell error message
	TellError()

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

	AskConfirm() bool
}

var (
	missingResources = make(map[string]bool)
)

func Missing() map[string]bool {
	return missingResources
}

type UIMessageContainer struct {
	resources     *rice.Box
	baseMessages  map[string]UIMessage
	localMessages map[string]UIMessage
	userInterface UI
	logger        *zap.Logger
	isTest        bool
	isDebug       bool
	lang          string
}

func NewUIMessageContainer(bx *rice.Box, ui UI, logger *zap.Logger) *UIMessageContainer {
	return &UIMessageContainer{
		resources:     bx,
		userInterface: ui,
		logger:        logger,
	}
}

func (z *UIMessageContainer) detectLanguage() language.Tag {
	bcp47, err := jibber_jabber.DetectIETF()
	if err != nil {
		z.logger.Debug("unable to detect language", zap.Error(err))
		return language.English
	}

	return z.chooseLanguage(bcp47)
}

func (z *UIMessageContainer) chooseLanguage(bcp47 string) language.Tag {
	if bcp47 == "" {
		return language.English
	}
	tag, err := language.Parse(bcp47)
	if err != nil {
		z.logger.Debug("unable to parse language into tag", zap.String("bcp47", bcp47), zap.Error(err))
		return language.English
	}
	m := language.NewMatcher(supportedLanguages)
	l, _, c := m.Match(tag)
	z.logger.Debug("detect language", zap.Any("lang", l), zap.String("confidence", c.String()))

	return l
}

func (z *UIMessageContainer) loadResource(lang language.Tag) (map[string]UIMessage, error) {
	resName := "messages.json"
	b, _ := lang.Base()
	if b.String() != "en" {
		resName = fmt.Sprintf("messages_%s.json", b)
	}
	z.logger.Debug("Loading message resource", zap.String("name", resName))

	baseAppMsgBytes, err := z.resources.Bytes(resName)
	if err != nil {
		z.logger.Error("unable to load base app msg", zap.String("name", resName), zap.Error(err))
		return nil, err
	} else {
		baseAppMsg := make(map[string]string)
		err = json.Unmarshal(baseAppMsgBytes, &baseAppMsg)
		if err != nil {
			z.logger.Error("unable to unmarshal app msg", zap.String("name", resName), zap.Error(err))
			return nil, err
		} else {
			return NewMessageMap(baseAppMsg, z.userInterface, z.logger), nil
		}
	}
}

func (z *UIMessageContainer) UpdateLang(bcp47 string) {
	l := z.chooseLanguage(bcp47)
	z.logger.Debug("Updating language", zap.String("bcp47", bcp47), zap.Any("chosen", l))

	if l == language.English {
		z.localMessages = nil
	} else {
		z.logger.Debug("Loading additional language resource", zap.Any("lang", l))
		lmc, err := z.loadResource(l)
		if err != nil {
			return
		}
		z.localMessages = lmc
	}
}

func (z *UIMessageContainer) Load() {
	z.baseMessages = make(map[string]UIMessage)
	if z.resources == nil {
		z.logger.Debug("no box found. skip loading messages.")
		z.isTest = true
		return
	}

	// load base message
	bmc, err := z.loadResource(language.English)
	if err != nil {
		return
	}
	z.baseMessages = bmc

	lang := z.detectLanguage()
	if lang != language.English {
		z.logger.Debug("Loading additional language resource", zap.Any("lang", lang))
		lmc, err := z.loadResource(lang)
		if err != nil {
			return
		}
		z.localMessages = lmc
	}
}

func (z *UIMessageContainer) MsgExists(key string) bool {
	if z.baseMessages == nil {
		return false
	}
	if z.localMessages != nil {
		if _, e := z.localMessages[key]; e {
			return true
		}
		// fallback to base messages if the message not found in local messages
	}
	if _, e := z.baseMessages[key]; e {
		return true
	}
	return false
}

func (z *UIMessageContainer) Msg(key string) UIMessage {
	if z.baseMessages == nil {
		return NewAltMessage(key, z.userInterface)
	}
	if z.localMessages != nil {
		if m, e := z.localMessages[key]; e {
			return m
		}
		// fallback to base messages if the message not found in local messages
	}
	if m, e := z.baseMessages[key]; e {
		return m
	}
	if !z.isTest {
		z.logger.Error("resource not found for key", zap.String("key", key))
	}
	missingResources[key] = true

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

func (z *TextMessage) T() string {
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

func (z *TextMessage) AskConfirm() bool {
	return z.userInterface.AskConfirm(z)
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

func (z *AltMessage) T() string {
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

func (z *AltMessage) AskConfirm() bool {
	return z.userInterface.AskConfirm(z)
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

func (z *Message) AskConfirm() bool {
	return z.userInterface.AskConfirm(z)
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

func (z *Message) T() string {
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
