package api_doc

import (
	"github.com/watermint/toolbox/domain/deepl/api/deepl_conn_impl"
	"github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn_impl"
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
		api_conn.ServiceTagAsana: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlAsana))
			cui.Break()
		},
		api_conn.ServiceTagDeepl: func(cui app_ui.UI) {
			cui.Info(api_auth_key.MConsole.PromptEnterKey)
			cui.AskText(deepl_conn_impl.MDeeplApi.AskApiKey)
		},
		api_conn.ServiceTagDropbox: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		api_conn.ServiceTagDropboxBusiness: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		api_conn.ServiceTagGithub: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGithub))
			cui.Break()
		},
		api_conn.ServiceTagGoogleCalendar: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceTagGoogleMail: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceTagGoogleSheets: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceTagGoogleTranslate: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceTagDropboxSign: func(cui app_ui.UI) {
			cui.Info(api_auth_basic.MConsole.PromptEnterUsernameAndPassword)
			cui.AskText(hs_conn_impl.MHelloSign.AskApiKey)
		},
		api_conn.ServiceTagSlack: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlSlack))
			cui.Break()
		},
		api_conn.ServiceTagFigma: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlFigma))
			cui.Break()
		},
	}

	ApiDocAuthDesc = map[string]app_msg.Message{
		api_conn.ServiceTagAsana:           MApiDoc.AuthDescAsana,
		api_conn.ServiceTagDeepl:           MApiDoc.AuthDescDeepl,
		api_conn.ServiceTagDropbox:         MApiDoc.AuthDescDropbox,
		api_conn.ServiceTagDropboxBusiness: MApiDoc.AuthDescDropbox,
		api_conn.ServiceTagGithub:          MApiDoc.AuthDescGithub,
		api_conn.ServiceTagGoogleCalendar:  MApiDoc.AuthDescGoogle,
		api_conn.ServiceTagGoogleMail:      MApiDoc.AuthDescGoogle,
		api_conn.ServiceTagGoogleSheets:    MApiDoc.AuthDescGoogle,
		api_conn.ServiceTagGoogleTranslate: MApiDoc.AuthDescGoogle,
		api_conn.ServiceTagDropboxSign:     MApiDoc.AuthDescDropboxSign,
		api_conn.ServiceTagSlack:           MApiDoc.AuthDescSlack,
		api_conn.ServiceTagFigma:           MApiDoc.AuthDescFigma,
	}
)
