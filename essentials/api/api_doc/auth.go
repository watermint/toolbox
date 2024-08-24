package api_doc

import (
	"github.com/watermint/toolbox/domain/deepl/api/deepl_conn_impl"
	"github.com/watermint/toolbox/essentials/api/api_auth_basic"
	"github.com/watermint/toolbox/essentials/api/api_auth_key"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/essentials/api/api_callback"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgApiDoc struct {
	AuthDescAsana       app_msg.Message
	AuthDescDeepl       app_msg.Message
	AuthDescDropbox     app_msg.Message
	AuthDescGithub      app_msg.Message
	AuthDescDropboxSign app_msg.Message
	AuthDescSlack       app_msg.Message
	AuthDescFigma       app_msg.Message
	ServiceUrlAsana     app_msg.Message
	ServiceUrlDropbox   app_msg.Message
	ServiceUrlGithub    app_msg.Message
	ServiceUrlSlack     app_msg.Message
	ServiceUrlFigma     app_msg.Message
}

var (
	MApiDoc = app_msg.Apply(&MsgApiDoc{}).(*MsgApiDoc)
)

type ApiAuthDoc func(cui app_ui.UI)

var (
	ApiDocCuiPreview = map[string]ApiAuthDoc{
		app_definitions.ScopeLabelAsana: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlAsana))
			cui.Break()
		},
		app_definitions.ScopeLabelDeepl: func(cui app_ui.UI) {
			cui.Info(api_auth_key.MConsole.PromptEnterKey)
			cui.AskText(deepl_conn_impl.MDeeplApi.AskApiKey)
		},
		app_definitions.ScopeLabelDropboxIndividual: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		app_definitions.ScopeLabelDropboxTeam: func(cui app_ui.UI) {
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq1.With("Url", cui.Text(MApiDoc.ServiceUrlDropbox)))
			cui.Info(api_auth_oauth.MApiAuth.OauthSeq2)
		},
		app_definitions.ScopeLabelGithub: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlGithub))
			cui.Break()
		},
		app_definitions.ScopeLabelDropboxSign: func(cui app_ui.UI) {
			cui.Info(api_auth_basic.MConsole.PromptEnterUsernameAndPassword)
			cui.AskText(api_auth_key.MConsole.AskKey)
		},
		app_definitions.ScopeLabelSlack: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlSlack))
			cui.Break()
		},
		app_definitions.ScopeLabelFigma: func(cui app_ui.UI) {
			cui.Info(api_callback.MCallback.MsgOpenUrlOnYourBrowser)
			cui.Code(cui.Text(MApiDoc.ServiceUrlFigma))
			cui.Break()
		},
	}

	ApiDocAuthDesc = map[string]app_msg.Message{
		app_definitions.ScopeLabelAsana:             MApiDoc.AuthDescAsana,
		app_definitions.ScopeLabelDeepl:             MApiDoc.AuthDescDeepl,
		app_definitions.ScopeLabelDropboxIndividual: MApiDoc.AuthDescDropbox,
		app_definitions.ScopeLabelDropboxTeam:       MApiDoc.AuthDescDropbox,
		app_definitions.ScopeLabelGithub:            MApiDoc.AuthDescGithub,
		app_definitions.ScopeLabelDropboxSign:       MApiDoc.AuthDescDropboxSign,
		app_definitions.ScopeLabelSlack:             MApiDoc.AuthDescSlack,
		app_definitions.ScopeLabelFigma:             MApiDoc.AuthDescFigma,
	}
)
