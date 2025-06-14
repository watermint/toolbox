package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgErrorHandlingGuide struct {
	DocDesc                app_msg.Message
	Title                  app_msg.Message
	CommonErrors           app_msg.Message
	NetworkErrors          app_msg.Message
	AuthenticationErrors   app_msg.Message
	FileSystemErrors       app_msg.Message
	RateLimitErrors       app_msg.Message
	APIErrors             app_msg.Message
	DebugTechniques       app_msg.Message
}

var (
	MErrorHandlingGuide = app_msg.Apply(&MsgErrorHandlingGuide{}).(*MsgErrorHandlingGuide)
)

type ErrorHandlingGuide struct {
}

func (z *ErrorHandlingGuide) DocId() dc_index.DocId {
	return dc_index.DocSupplementalTroubleshooting
}

func (z *ErrorHandlingGuide) DocDesc() app_msg.Message {
	return MErrorHandlingGuide.DocDesc
}

func (z *ErrorHandlingGuide) Sections() []dc_section.Section {
	return []dc_section.Section{
		&CommonErrorsSection{},
		&NetworkErrorsSection{},
		&AuthenticationErrorsSection{},
		&FileSystemErrorsSection{},
		&RateLimitErrorsSection{},
		&APIErrorsSection{},
		&DebugTechniquesSection{},
	}
}

// CommonErrorsSection covers the most frequent errors
type CommonErrorsSection struct{}

func (z *CommonErrorsSection) Title() app_msg.Message {
	return MErrorHandlingGuide.CommonErrors
}

func (z *CommonErrorsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Common Errors and Solutions

This section covers the most frequently encountered errors and their solutions:

1. Command Not Found:
   - Ensure you're using the correct command syntax
   - Use 'help' or '--help' to see available commands
   - Check for typos in command names

2. Invalid Arguments:
   - Verify required arguments are provided
   - Check argument formats (paths, emails, etc.)
   - Use quotes for arguments containing spaces

3. Configuration Issues:
   - Check if required configuration files exist
   - Verify configuration file permissions
   - Use default configurations when in doubt

4. Output Directory Issues:
   - Ensure output directories exist and are writable
   - Check disk space availability
   - Avoid paths with special characters

5. General Troubleshooting Steps:
   - Run with -debug flag for detailed logging
   - Check the command documentation
   - Verify your environment meets requirements
   - Try with minimal arguments first
`))
}

// NetworkErrorsSection covers network-related issues
type NetworkErrorsSection struct{}

func (z *NetworkErrorsSection) Title() app_msg.Message {
	return MErrorHandlingGuide.NetworkErrors
}

func (z *NetworkErrorsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Network Errors

Network connectivity issues and solutions:

1. Connection Timeout:
   - Check internet connectivity
   - Verify DNS resolution for dropbox.com and dropboxapi.com
   - Try again after network issues are resolved
   - Consider using -bandwidth-kb flag to limit transfer speed

2. SSL/TLS Errors:
   - Update your system certificates
   - Check if corporate firewalls are interfering
   - Verify system time is accurate

3. Proxy Issues:
   - Configure system proxy settings if needed
   - Check proxy authentication requirements
   - Test direct connection when possible

4. DNS Resolution Failures:
   - Try using different DNS servers
   - Check /etc/hosts file for conflicts
   - Verify network configuration

5. Intermittent Connection Issues:
   - Implement retry logic with delays
   - Use smaller batch sizes for large operations
   - Monitor network stability during transfers

6. Corporate Network Restrictions:
   - Work with IT team to whitelist required domains
   - Request access to necessary ports (443, 80)
   - Consider using mobile hotspot for testing
`))
}

// AuthenticationErrorsSection covers auth-specific errors
type AuthenticationErrorsSection struct{}

func (z *AuthenticationErrorsSection) Title() app_msg.Message {
	return MErrorHandlingGuide.AuthenticationErrors
}

func (z *AuthenticationErrorsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Authentication Errors

Authentication-related errors and solutions:

1. Token Expired/Invalid:
   - Run the command again to trigger re-authentication
   - Use 'config auth delete' to remove old credentials
   - Ensure system time is accurate

2. Permission Denied:
   - Verify account has necessary permissions
   - Check if account is suspended or restricted
   - For business accounts, ensure proper team member access

3. OAuth Flow Failures:
   - Try different browser or incognito mode
   - Clear browser cookies for dropbox.com
   - Check if ad blockers are interfering

4. Browser Not Opening:
   - Use -auto-open=false and copy URL manually
   - Check if running in headless environment
   - Verify default browser configuration

5. Multiple Account Conflicts:
   - Use 'config auth list' to see configured accounts
   - Remove conflicting accounts with 'config auth delete'
   - Specify account explicitly in commands

6. Database Access Issues:
   - Check permissions on secrets database file
   - Verify database directory is writable
   - Use -auth-database flag for custom location
`))
}

// FileSystemErrorsSection covers file system issues
type FileSystemErrorsSection struct{}

func (z *FileSystemErrorsSection) Title() app_msg.Message {
	return MErrorHandlingGuide.FileSystemErrors
}

func (z *FileSystemErrorsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
File System Errors

Local file system related errors and solutions:

1. Permission Denied:
   - Check file and directory permissions
   - Ensure user has read/write access as needed
   - Use sudo cautiously and only when necessary

2. Disk Space Issues:
   - Check available disk space with df -h
   - Clean up temporary files and logs
   - Use -budget-storage=low flag to reduce storage usage

3. Path Not Found:
   - Verify file and directory paths exist
   - Use absolute paths when relative paths fail
   - Check for typos in path names

4. File Lock Issues:
   - Close applications that might have files open
   - Check for running processes using files
   - Wait and retry if files are temporarily locked

5. Character Encoding Issues:
   - Ensure file names use valid character encoding
   - Avoid special characters in file names
   - Use UTF-8 encoding for text files

6. Symlink and Junction Issues:
   - Verify symlinks point to valid targets
   - Check permissions on symlink targets
   - Consider using direct paths instead of symlinks

7. Long Path Names:
   - Keep path lengths reasonable (< 260 characters on Windows)
   - Use shorter directory and file names
   - Move operations closer to root directory
`))
}

// RateLimitErrorsSection covers API rate limiting
type RateLimitErrorsSection struct{}

func (z *RateLimitErrorsSection) Title() app_msg.Message {
	return MErrorHandlingGuide.RateLimitErrors
}

func (z *RateLimitErrorsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Rate Limit Errors

API rate limiting errors and solutions:

1. Too Many Requests (429 Error):
   - Wait before retrying (toolbox handles this automatically)
   - Reduce concurrency with -concurrency flag
   - Use smaller batch sizes for bulk operations

2. Daily/Hourly Limits:
   - Spread operations across multiple days
   - Monitor API usage patterns
   - Consider using multiple accounts for large operations

3. Burst Limit Exceeded:
   - Add delays between operations
   - Use batch operations when available
   - Avoid rapid sequential API calls

4. Team Rate Limits:
   - Coordinate with other team members using the API
   - Implement organization-wide rate limiting policies
   - Monitor team-wide API usage

5. Optimization Strategies:
   - Use list operations instead of individual file requests
   - Cache results to avoid repeated API calls
   - Combine multiple operations into single requests where possible

6. Monitoring and Planning:
   - Track API usage patterns
   - Plan large operations during off-peak hours
   - Set up alerts for approaching rate limits
`))
}

// APIErrorsSection covers general API errors
type APIErrorsSection struct{}

func (z *APIErrorsSection) Title() app_msg.Message {
	return MErrorHandlingGuide.APIErrors
}

func (z *APIErrorsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
API Errors

General API errors and solutions:

1. 400 Bad Request:
   - Verify request parameters are correct
   - Check data formats (dates, emails, paths)
   - Ensure required fields are provided

2. 401 Unauthorized:
   - Check authentication credentials
   - Verify token hasn't expired
   - Ensure proper authorization scope

3. 403 Forbidden:
   - Verify account permissions
   - Check if feature is available for account type
   - Ensure API access hasn't been restricted

4. 404 Not Found:
   - Verify file/folder paths exist
   - Check if items have been moved or deleted
   - Ensure proper path formatting

5. 409 Conflict:
   - Handle concurrent modification conflicts
   - Retry with updated information
   - Resolve conflicts manually if needed

6. 500 Internal Server Error:
   - Retry the operation after a delay
   - Check Dropbox status page for service issues
   - Contact support if error persists

7. Service Unavailable (503):
   - Wait and retry (temporary service issues)
   - Check for scheduled maintenance
   - Use exponential backoff for retries
`))
}

// DebugTechniquesSection provides debugging guidance
type DebugTechniquesSection struct{}

func (z *DebugTechniquesSection) Title() app_msg.Message {
	return MErrorHandlingGuide.DebugTechniques
}

func (z *DebugTechniquesSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Debug Techniques

Advanced debugging techniques for troubleshooting:

1. Enable Debug Logging:
   - Use -debug flag for verbose output
   - Check log files for detailed information
   - Look for specific error messages and codes

2. Test with Minimal Parameters:
   - Start with simplest possible command
   - Add parameters one by one to isolate issues
   - Use default values when possible

3. Environment Verification:
   - Check system requirements
   - Verify environment variables
   - Test on different machines if available

4. Network Diagnostics:
   - Use ping/traceroute to test connectivity
   - Check firewall and proxy settings
   - Monitor network traffic during operations

5. API Testing:
   - Use API testing tools to verify endpoints
   - Check API responses manually
   - Verify request formats and parameters

6. Log Analysis:
   - Review application logs systematically
   - Look for patterns in error messages
   - Check timestamps for sequence of events

7. Isolation Testing:
   - Test with different accounts
   - Try operations on different files/folders
   - Use minimal test data sets

8. Community Resources:
   - Search documentation and FAQ
   - Check community forums and discussions
   - Report bugs with detailed reproduction steps
`))
}