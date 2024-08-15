package api_doc

import (
	"github.com/watermint/toolbox/domain/deepl/api/deepl_conn_impl"
	"github.com/watermint/toolbox/essentials/api/api_auth_basic"
	"github.com/watermint/toolbox/essentials/api/api_auth_key"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/essentials/api/api_callback"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgApiDoc struct {
	AuthDescAsana       app_msg.Message
	AuthDescDeepl       app_msg.Message
	AuthDescDropbox     app_msg.Message
	AuthDescGithub      app_msg.Message
	AuthDescGoogle      app_msg.Message
	AuthDescDropboxSign app_msg.Message
	AuthDescSlack       app_msg.Message
	AuthDescFigma       app_msg.Message
	ServiceUrlAsana     app_msg.Message
	ServiceUrlDropbox   app_msg.Message
	ServiceUrlGithub    app_msg.Message
	ServiceUrlGoogle    app_msg.Message
	ServiceUrlSlack     app_msg.Message
	ServiceUrlFigma     app_msg.Message
}

var (
	MApiDoc = app_msg.Apply(&MsgApiDoc{}).(*MsgApiDoc)
)

type ApiAuthDoc func(cui app_ui.UI)

var (
	ApiDocCuiPreview = map[string]ApiAuthDoc{
		api_conn.ScopeLabelAsana: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlAsana))
			cui.Break()
		},
		api_conn.ScopeLabelDeepl: func(cui app_ui.UI) {
			cui.Info(api_auth_key.MConsole.PromptEnterKey)
			cui.AskText(deepl_conn_impl.MDeeplApi.AskApiKey)
		},
		api_conn.ScopeLabelDropbox: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		api_conn.ScopeLabelDropboxBusiness: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		api_conn.ScopeLabelGithub: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGithub))
			cui.Break()
		},
		api_conn.ScopeLabelGoogleCalendar: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ScopeLabelGoogleMail: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ScopeLabelGoogleSheets: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ScopeLabelGoogleTranslate: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ScopeLabelDropboxSign: func(cui app_ui.UI) {
			cui.Info(api_auth_basic.MConsole.PromptEnterUsernameAndPassword)
			cui.AskText(api_auth_key.MConsole.AskKey)
		},
		api_conn.ScopeLabelSlack: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlSlack))
			cui.Break()
		},
		api_conn.ScopeLabelFigma: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlFigma))
			cui.Break()
		},
	}

	ApiDocAuthDesc = map[string]app_msg.Message{
		api_conn.ScopeLabelAsana:           MApiDoc.AuthDescAsana,
		api_conn.ScopeLabelDeepl:           MApiDoc.AuthDescDeepl,
		api_conn.ScopeLabelDropbox:         MApiDoc.AuthDescDropbox,
		api_conn.ScopeLabelDropboxBusiness: MApiDoc.AuthDescDropbox,
		api_conn.ScopeLabelGithub:          MApiDoc.AuthDescGithub,
		api_conn.ScopeLabelGoogleCalendar:  MApiDoc.AuthDescGoogle,
		api_conn.ScopeLabelGoogleMail:      MApiDoc.AuthDescGoogle,
		api_conn.ScopeLabelGoogleSheets:    MApiDoc.AuthDescGoogle,
		api_conn.ScopeLabelGoogleTranslate: MApiDoc.AuthDescGoogle,
		api_conn.ScopeLabelDropboxSign:     MApiDoc.AuthDescDropboxSign,
		api_conn.ScopeLabelSlack:           MApiDoc.AuthDescSlack,
		api_conn.ScopeLabelFigma:           MApiDoc.AuthDescFigma,
	}
)
