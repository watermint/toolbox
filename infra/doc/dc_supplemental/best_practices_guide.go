package dc_supplemental

import (
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type MsgBestPracticesGuide struct {
	DocDesc              app_msg.Message
	Title                app_msg.Message
	GeneralPractices     app_msg.Message
	PerformanceOptimization app_msg.Message
	SecurityPractices    app_msg.Message
	AutomationTips       app_msg.Message
	DataManagement       app_msg.Message
	TeamCollaboration    app_msg.Message
	MaintenanceAndUpdates app_msg.Message
}

var (
	MBestPracticesGuide = app_msg.Apply(&MsgBestPracticesGuide{}).(*MsgBestPracticesGuide)
)

type BestPracticesGuide struct {
}

func (z *BestPracticesGuide) DocId() dc_index.DocId {
	return dc_index.DocSupplementalTroubleshooting
}

func (z *BestPracticesGuide) DocDesc() app_msg.Message {
	return MBestPracticesGuide.DocDesc
}

func (z *BestPracticesGuide) Sections() []dc_section.Section {
	return []dc_section.Section{
		&GeneralPracticesSection{},
		&PerformanceOptimizationSection{},
		&SecurityPracticesSection{},
		&AutomationTipsSection{},
		&DataManagementSection{},
		&TeamCollaborationSection{},
		&MaintenanceAndUpdatesSection{},
	}
}

// GeneralPracticesSection covers general best practices
type GeneralPracticesSection struct{}

func (z *GeneralPracticesSection) Title() app_msg.Message {
	return MBestPracticesGuide.GeneralPractices
}

func (z *GeneralPracticesSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
General Best Practices

Follow these general guidelines for effective toolbox usage:

1. Command Preparation:
   - Always read command documentation before use
   - Test commands with sample data first
   - Use --help flag to understand available options
   - Verify required permissions and prerequisites

2. Data Backup:
   - Create backups before major operations
   - Test restore procedures regularly
   - Use version control for important files
   - Document backup and recovery procedures

3. Error Handling:
   - Enable debug logging for troubleshooting (-debug flag)
   - Keep logs of important operations
   - Implement proper error checking in scripts
   - Have rollback procedures for critical operations

4. Resource Management:
   - Monitor disk space before large operations
   - Use appropriate concurrency settings (-concurrency flag)
   - Manage memory usage with -budget-memory flag
   - Clean up temporary files and logs regularly

5. Documentation:
   - Document custom workflows and procedures
   - Keep track of configuration changes
   - Maintain inventory of automated scripts
   - Document troubleshooting steps for common issues

6. Testing:
   - Test commands in development environment first
   - Use small datasets for initial testing
   - Validate results before processing large batches
   - Implement automated testing for critical workflows
`))
}

// PerformanceOptimizationSection covers performance tips
type PerformanceOptimizationSection struct{}

func (z *PerformanceOptimizationSection) Title() app_msg.Message {
	return MBestPracticesGuide.PerformanceOptimization
}

func (z *PerformanceOptimizationSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Performance Optimization

Optimize toolbox performance with these techniques:

1. Concurrency Management:
   - Adjust -concurrency flag based on system resources
   - Higher concurrency for I/O intensive operations
   - Lower concurrency for CPU intensive operations
   - Monitor system resources during operations

2. Bandwidth Optimization:
   - Use -bandwidth-kb flag to limit network usage
   - Schedule large transfers during off-peak hours
   - Consider network conditions and limitations
   - Monitor transfer speeds and adjust accordingly

3. Memory Management:
   - Use -budget-memory=low for memory-constrained environments
   - Process data in smaller chunks for large datasets
   - Monitor memory usage during operations
   - Clear caches and temporary data regularly

4. Storage Optimization:
   - Use -budget-storage=low to reduce storage usage
   - Clean up logs and temporary files regularly
   - Use appropriate output formats (avoid verbose formats when not needed)
   - Compress data when possible

5. Batch Operations:
   - Group similar operations together
   - Use batch commands when available
   - Minimize API calls with efficient operations
   - Process multiple items in single commands

6. Caching Strategies:
   - Leverage local caching for frequently accessed data
   - Avoid redundant API calls
   - Use incremental operations when possible
   - Cache authentication tokens properly

7. Network Optimization:
   - Use stable, high-speed network connections
   - Avoid wireless connections for large transfers
   - Consider geographic proximity to servers
   - Implement retry logic with exponential backoff
`))
}

// SecurityPracticesSection covers security best practices
type SecurityPracticesSection struct{}

func (z *SecurityPracticesSection) Title() app_msg.Message {
	return MBestPracticesGuide.SecurityPractices
}

func (z *SecurityPracticesSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Security Best Practices

Maintain security while using the toolbox:

1. Authentication Security:
   - Use strong, unique passwords for accounts
   - Enable two-factor authentication when available
   - Regularly review and rotate credentials
   - Use dedicated service accounts for automation

2. Token Management:
   - Protect authentication database files
   - Use appropriate file permissions (600 or 700)
   - Avoid sharing authentication databases
   - Regularly audit configured accounts

3. Data Protection:
   - Encrypt sensitive data at rest and in transit
   - Use secure protocols (HTTPS, SSH) for all communications
   - Implement proper access controls
   - Regular security audits of data access

4. Environment Security:
   - Use secure workstations for operations
   - Keep systems updated with security patches
   - Use anti-virus and anti-malware protection
   - Secure physical access to systems

5. Network Security:
   - Use VPN for remote access
   - Avoid public WiFi for sensitive operations
   - Implement network segmentation where appropriate
   - Monitor network traffic for anomalies

6. Audit and Monitoring:
   - Log all significant operations
   - Monitor account activity regularly
   - Set up alerts for unusual activity
   - Maintain audit trails for compliance

7. Incident Response:
   - Have incident response procedures
   - Know how to revoke access quickly
   - Maintain contact information for security teams
   - Practice incident response scenarios
`))
}

// AutomationTipsSection covers automation best practices
type AutomationTipsSection struct{}

func (z *AutomationTipsSection) Title() app_msg.Message {
	return MBestPracticesGuide.AutomationTips
}

func (z *AutomationTipsSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Automation Best Practices

Best practices for automating toolbox operations:

1. Script Development:
   - Use version control for all scripts
   - Implement proper error handling and logging
   - Add comments and documentation
   - Use configuration files for parameters

2. Scheduling and Execution:
   - Use cron jobs or task schedulers appropriately
   - Implement proper locking to prevent concurrent runs
   - Set up monitoring and alerting for failures
   - Use appropriate user accounts for automation

3. Parameter Management:
   - Use configuration files instead of hardcoded values
   - Implement parameter validation
   - Use environment variables for sensitive data
   - Provide default values where appropriate

4. Error Handling:
   - Implement comprehensive error checking
   - Use appropriate exit codes
   - Log errors with sufficient detail
   - Implement retry logic with backoff

5. Testing:
   - Test scripts in development environment
   - Use test data for validation
   - Implement automated testing where possible
   - Validate results automatically

6. Monitoring:
   - Log all significant operations
   - Monitor script execution times
   - Set up alerts for failures
   - Track resource usage

7. Maintenance:
   - Regular review and updates of scripts
   - Monitor for deprecated features
   - Keep dependencies updated
   - Document maintenance procedures
`))
}

// DataManagementSection covers data management practices
type DataManagementSection struct{}

func (z *DataManagementSection) Title() app_msg.Message {
	return MBestPracticesGuide.DataManagement
}

func (z *DataManagementSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Data Management Best Practices

Effective data management strategies:

1. Data Organization:
   - Use consistent naming conventions
   - Organize data in logical folder structures
   - Implement proper file and folder hierarchy
   - Use metadata and tags effectively

2. Data Validation:
   - Verify data integrity before and after operations
   - Use checksums for critical data
   - Implement data validation rules
   - Test with sample data before processing

3. Backup Strategies:
   - Implement regular automated backups
   - Test backup restoration procedures
   - Use multiple backup locations (3-2-1 rule)
   - Document backup and recovery procedures

4. Data Lifecycle Management:
   - Define data retention policies
   - Implement automated archiving
   - Clean up old and unnecessary data
   - Monitor storage usage trends

5. Data Synchronization:
   - Use incremental sync when possible
   - Verify sync operations regularly
   - Handle conflicts appropriately
   - Monitor sync performance and errors

6. Data Quality:
   - Implement data quality checks
   - Clean and normalize data regularly
   - Remove duplicates and inconsistencies
   - Validate data formats and standards

7. Compliance and Governance:
   - Follow data governance policies
   - Ensure compliance with regulations
   - Implement proper access controls
   - Maintain audit trails for data operations
`))
}

// TeamCollaborationSection covers team collaboration practices
type TeamCollaborationSection struct{}

func (z *TeamCollaborationSection) Title() app_msg.Message {
	return MBestPracticesGuide.TeamCollaboration
}

func (z *TeamCollaborationSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Team Collaboration Best Practices

Effective team collaboration with toolbox:

1. Account Management:
   - Use dedicated service accounts for shared operations
   - Implement proper access controls and permissions
   - Regular review of account access and privileges
   - Document account usage and responsibilities

2. Configuration Management:
   - Use centralized configuration management
   - Version control for shared configurations
   - Implement configuration validation
   - Document configuration changes

3. Workflow Coordination:
   - Define clear workflows and procedures
   - Implement proper change management
   - Use communication channels for coordination
   - Schedule operations to avoid conflicts

4. Knowledge Sharing:
   - Document procedures and best practices
   - Conduct regular training sessions
   - Share troubleshooting experiences
   - Maintain knowledge base and FAQ

5. Quality Assurance:
   - Implement peer review processes
   - Use testing and validation procedures
   - Define quality standards and metrics
   - Regular audits of processes and results

6. Communication:
   - Establish clear communication protocols
   - Use collaboration tools effectively
   - Document decisions and changes
   - Regular team meetings and updates

7. Incident Management:
   - Define incident response procedures
   - Establish escalation paths
   - Maintain contact information
   - Conduct post-incident reviews
`))
}

// MaintenanceAndUpdatesSection covers maintenance practices
type MaintenanceAndUpdatesSection struct{}

func (z *MaintenanceAndUpdatesSection) Title() app_msg.Message {
	return MBestPracticesGuide.MaintenanceAndUpdates
}

func (z *MaintenanceAndUpdatesSection) Body(ui app_ui.UI) {
	ui.Info(app_msg.Raw(`
Maintenance and Updates

Keep your toolbox installation and workflows updated:

1. Software Updates:
   - Regularly check for toolbox updates
   - Test updates in development environment first
   - Keep dependencies updated
   - Monitor for security updates

2. Configuration Maintenance:
   - Regular review of configurations
   - Update deprecated settings
   - Optimize performance settings
   - Clean up unused configurations

3. Data Maintenance:
   - Regular cleanup of logs and temporary files
   - Archive old data appropriately
   - Optimize storage usage
   - Validate data integrity periodically

4. Documentation Updates:
   - Keep documentation current with changes
   - Update procedures and workflows
   - Review and update troubleshooting guides
   - Maintain version history

5. Performance Monitoring:
   - Monitor system performance regularly
   - Track operation times and success rates
   - Identify and address bottlenecks
   - Optimize based on usage patterns

6. Security Maintenance:
   - Regular security audits
   - Update security configurations
   - Review access permissions
   - Monitor for security vulnerabilities

7. Backup and Recovery:
   - Test backup and recovery procedures
   - Update recovery documentation
   - Verify backup integrity
   - Practice disaster recovery scenarios

8. Training and Skills:
   - Keep team skills current
   - Provide training on new features
   - Share knowledge and best practices
   - Encourage continuous learning
`))
}