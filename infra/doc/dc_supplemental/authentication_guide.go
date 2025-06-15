package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

const (
	AuthenticationGuideDescKey = "auth.guide.desc"
	AuthenticationGuideTitle   = "auth.guide.title"
	AuthGuideOverview         = "auth.guide.overview"
	AuthGuideDropboxAuth      = "auth.guide.dropbox"
	AuthGuideTokenManagement  = "auth.guide.token"
	AuthGuideTroubleshooting  = "auth.guide.troubleshooting"
	AuthGuideSecurityTips     = "auth.guide.security"
)

type MsgAuthenticationGuide struct {
	DocDesc            app_msg.Message
	Title              app_msg.Message
	Overview           app_msg.Message
	DropboxAuth        app_msg.Message
	TokenManagement    app_msg.Message
	Troubleshooting    app_msg.Message
	SecurityTips       app_msg.Message
}

var (
	MAuthenticationGuide = app_msg.Apply(&MsgAuthenticationGuide{}).(*MsgAuthenticationGuide)
)

type AuthenticationGuide struct {
}

func (z *AuthenticationGuide) DocId() dc_index.DocId {
	return dc_index.DocSupplementalTroubleshooting
}

func (z *AuthenticationGuide) DocDesc() app_msg.Message {
	return MAuthenticationGuide.DocDesc
}

func (z *AuthenticationGuide) Sections() []dc_section.Section {
	return []dc_section.Section{
		&AuthOverviewSection{},
		&DropboxAuthSection{},
		&TokenManagementSection{},
		&AuthTroubleshootingSection{},
		&SecurityTipsSection{},
	}
}

// AuthOverviewSection provides an overview of authentication in toolbox
type AuthOverviewSection struct{}

func (z *AuthOverviewSection) Title() app_msg.Message {
	return MAuthenticationGuide.Overview
}

func (z *AuthOverviewSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Authentication Overview

The watermint toolbox requires proper authentication to access Dropbox services. The toolbox supports multiple authentication methods and securely manages tokens for seamless operation.

Key authentication concepts:
- OAuth 2.0 flow for secure authorization
- Token-based authentication for API access
- Secure token storage in local database
- Automatic token refresh when possible
- Support for multiple account configurations
`))
}

// DropboxAuthSection explains Dropbox-specific authentication
type DropboxAuthSection struct{}

func (z *DropboxAuthSection) Title() app_msg.Message {
	return MAuthenticationGuide.DropboxAuth
}

func (z *DropboxAuthSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Dropbox Authentication

The toolbox uses OAuth 2.0 to authenticate with Dropbox. This process involves:

1. Initial Authorization:
   - Run any command that requires Dropbox access
   - The toolbox will open a browser window to Dropbox authorization page
   - Sign in to your Dropbox account and grant permissions
   - The toolbox receives an authorization code and exchanges it for access tokens

2. Supported Account Types:
   - Personal Dropbox accounts
   - Dropbox Business accounts (with team member access)
   - Dropbox Business admin accounts (with full team access)

3. Required Permissions:
   - File access permissions (read/write as needed)
   - Team information access (for business accounts)
   - User information access for account identification

4. Authentication Flow:
   - Commands automatically detect when authentication is needed
   - Browser-based OAuth flow ensures secure credential handling
   - No passwords or API keys need to be manually entered
`))
}

// TokenManagementSection explains token management
type TokenManagementSection struct{}

func (z *TokenManagementSection) Title() app_msg.Message {
	return MAuthenticationGuide.TokenManagement
}

func (z *TokenManagementSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Token Management

The toolbox securely manages authentication tokens:

1. Token Storage:
   - Tokens are stored in encrypted local database
   - Default location: $HOME/.toolbox/secrets/secrets.db
   - Custom database path can be specified with -auth-database flag

2. Token Lifecycle:
   - Access tokens are automatically refreshed when possible
   - Expired tokens trigger re-authentication flow
   - Tokens are associated with specific account configurations

3. Multiple Accounts:
   - Support for multiple Dropbox accounts
   - Each account maintains separate token storage
   - Account selection via command-line flags or configuration

4. Token Security:
   - Tokens are encrypted at rest
   - No tokens are logged or exposed in command output
   - Secure deletion when accounts are removed

5. Managing Tokens:
   - Use 'config auth list' to view configured accounts
   - Use 'config auth delete' to remove account configurations
   - Re-authentication is automatic when tokens are invalid
`))
}

// AuthTroubleshootingSection provides troubleshooting guidance
type AuthTroubleshootingSection struct{}

func (z *AuthTroubleshootingSection) Title() app_msg.Message {
	return MAuthenticationGuide.Troubleshooting
}

func (z *AuthTroubleshootingSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Authentication Troubleshooting

Common authentication issues and solutions:

1. Browser Not Opening:
   - Check if you're running in a headless environment
   - Use -auto-open=false flag to disable automatic browser opening
   - Copy the authorization URL manually if needed

2. Permission Denied Errors:
   - Verify you have necessary permissions on your Dropbox account
   - For business accounts, ensure you have appropriate team member access
   - Re-authenticate if permissions have changed

3. Token Expired Errors:
   - Run the command again to trigger automatic re-authentication
   - Check if your account has been suspended or permissions revoked
   - Clear old tokens with 'config auth delete' if needed

4. Database Access Issues:
   - Ensure the secrets database directory is writable
   - Check file permissions on the database file
   - Use -auth-database flag to specify alternative location

5. Network Connectivity:
   - Verify internet connection for OAuth flow
   - Check if corporate firewalls are blocking access
   - Ensure access to dropbox.com and dropboxapi.com domains

6. Multiple Account Conflicts:
   - Use 'config auth list' to see all configured accounts
   - Remove conflicting accounts with 'config auth delete'
   - Specify explicit account in command if needed
`))
}

// SecurityTipsSection provides security best practices
type SecurityTipsSection struct{}

func (z *SecurityTipsSection) Title() app_msg.Message {
	return MAuthenticationGuide.SecurityTips
}

func (z *SecurityTipsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Security Best Practices

Follow these security practices when using authentication:

1. Token Protection:
   - Never share your secrets database file
   - Use appropriate file permissions on the database
   - Regularly review and clean up unused account configurations

2. Account Access:
   - Use dedicated service accounts for automation
   - Regularly review OAuth app authorizations in your Dropbox account
   - Revoke access for unused or suspicious applications

3. Environment Security:
   - Use secure workstations for authentication
   - Avoid authenticating on shared or public computers
   - Clear browser history after authentication if using public computers

4. Network Security:
   - Use secure networks for authentication
   - Avoid public WiFi for initial authentication
   - Consider using VPN for additional security

5. Monitoring:
   - Regularly review Dropbox account activity logs
   - Monitor for unexpected API usage
   - Set up alerts for unusual account activity

6. Backup and Recovery:
   - Keep secure backups of important data
   - Have a recovery plan if authentication is compromised
   - Know how to revoke and re-establish authentication
`))
}