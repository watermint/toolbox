package dc_command

import (
	"github.com/watermint/toolbox/infra/api/api_conn"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sort"
)

func NewSecurity(spec rc_recipe.Spec) dc_section.Section {
	return &Security{spec: spec}
}

type Security struct {
	spec                   rc_recipe.Spec
	Header                 app_msg.Message
	CredentialLocation     app_msg.Message
	TableHeaderOS          app_msg.Message
	TableHeaderPath        app_msg.Message
	CredentialRemarks      app_msg.Message
	HowToRemoveIt          app_msg.Message
	HowToHelpCenter        app_msg.Message
	Scopes                 app_msg.Message
	TableHeaderScopesLabel app_msg.Message
	TableHeaderScopesDesc  app_msg.Message
}

func (z Security) Title() app_msg.Message {
	return z.Header
}

func (z Security) Body(ui app_ui.UI) {
	ui.Info(z.CredentialLocation)
	ui.WithTable("Locations", func(t app_ui.Table) {
		t.Header(z.TableHeaderOS, z.TableHeaderPath)
		t.RowRaw("Windows", "`%HOMEPATH%\\.toolbox\\secrets` (e.g. C:\\Users\\bob\\.toolbox\\secrets)")
		t.RowRaw("macOS", "`$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)")
		t.RowRaw("Linux", "`$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)")
	})
	ui.Info(z.CredentialRemarks)
	ui.Info(z.HowToRemoveIt)

	ui.Break()
	ui.Info(z.HowToHelpCenter)
	for _, svc := range z.spec.Services() {
		ui.Info(app_msg.ObjMessage(&z, "help_center."+svc))
	}
	ui.Break()

	scopes := make(map[string]bool)
	for _, vn := range z.spec.ValueNames() {
		if vc, ok := z.spec.Value(vn).(rc_recipe.ValueConn); ok {
			if conn, ok := vc.Conn(); ok {
				if cs, ok := conn.(api_conn.ScopedConnection); ok {
					for _, scope := range cs.Scopes() {
						scopes[cs.ServiceName()+"."+scope] = true
					}
				} else {
					scopes[conn.ServiceName()+"."+conn.ScopeLabel()] = true
				}
			}
		}
	}
	scopeMessages := make([]app_msg.Message, 0)
	for scope := range scopes {
		scopeMessages = append(scopeMessages, app_msg.ObjMessage(&z, "scope."+scope))
	}
	sort.Slice(scopeMessages, func(i, j int) bool {
		return scopeMessages[i].Key() < scopeMessages[j].Key()
	})

	ui.SubHeader(z.Scopes)
	ui.WithTable("Scopes", func(t app_ui.Table) {
		t.Header(z.TableHeaderScopesDesc)
		for _, scm := range scopeMessages {
			t.Row(scm)
		}
	})
}
