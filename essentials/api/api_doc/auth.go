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
		api_conn.ServiceAsana: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlAsana))
			cui.Break()
		},
		api_conn.ServiceDeepl: func(cui app_ui.UI) {
			cui.Info(api_auth_key.MConsole.PromptEnterKey)
			cui.AskText(deepl_conn_impl.MDeeplApi.AskApiKey)
		},
		api_conn.ServiceDropbox: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		api_conn.ServiceDropboxBusiness: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		api_conn.ServiceGithub: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGithub))
			cui.Break()
		},
		api_conn.ServiceGoogleCalendar: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceGoogleMail: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceGoogleSheets: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceGoogleTranslate: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGoogle))
			cui.Break()
		},
		api_conn.ServiceDropboxSign: func(cui app_ui.UI) {
			cui.Info(api_auth_basic.MConsole.PromptEnterUsernameAndPassword)
			cui.AskText(hs_conn_impl.MHelloSign.AskApiKey)
		},
		api_conn.ServiceSlack: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlSlack))
			cui.Break()
		},
		api_conn.ServiceFigma: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlFigma))
			cui.Break()
		},
	}

	ApiDocAuthDesc = map[string]app_msg.Message{
		api_conn.ServiceAsana:           MApiDoc.AuthDescAsana,
		api_conn.ServiceDeepl:           MApiDoc.AuthDescDeepl,
		api_conn.ServiceDropbox:         MApiDoc.AuthDescDropbox,
		api_conn.ServiceDropboxBusiness: MApiDoc.AuthDescDropbox,
		api_conn.ServiceGithub:          MApiDoc.AuthDescGithub,
		api_conn.ServiceGoogleCalendar:  MApiDoc.AuthDescGoogle,
		api_conn.ServiceGoogleMail:      MApiDoc.AuthDescGoogle,
		api_conn.ServiceGoogleSheets:    MApiDoc.AuthDescGoogle,
		api_conn.ServiceGoogleTranslate: MApiDoc.AuthDescGoogle,
		api_conn.ServiceDropboxSign:     MApiDoc.AuthDescDropboxSign,
		api_conn.ServiceSlack:           MApiDoc.AuthDescSlack,
		api_conn.ServiceFigma:           MApiDoc.AuthDescFigma,
	}
)
