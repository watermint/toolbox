# Architecture Guide

## Project Structure Overview

The watermint toolbox uses a modular architecture with clear separation between internal utilities and external service integrations.

### Directory Layout

```
toolbox/
├── recipe/          # Internal utilities and dev tools
├── citron/          # External service integrations  
├── infra/           # Infrastructure and framework
├── domain/          # Domain logic and models
├── essentials/      # Core utilities and helpers
├── resources/       # Messages, templates, API keys
└── docs/            # Generated documentation
```

## Recipe vs Citron Architecture

### Recipe Directory
**Purpose**: Internal toolbox functionality, utilities, and development tools

Contains commands that:
- Manage the toolbox itself (configuration, logging)
- Provide general utilities (text processing, UUID generation)
- Support development workflows (build, test, release)
- Work offline without external authentication

Examples:
```bash
tbx dev build preflight      # Development tool
tbx util text case up       # Text utility
tbx config feature list     # Configuration
tbx log cat last           # Logging
```

### Citron Directory  
**Purpose**: External service integrations

Contains commands that:
- Integrate with third-party APIs (Dropbox, GitHub, Slack, etc.)
- Require authentication tokens
- Handle service-specific data formats
- Implement service workflows

Examples:
```bash
tbx dropbox file copy       # Dropbox integration
tbx github release list     # GitHub integration
tbx slack conversation list # Slack integration
```

## Command Implementation

### Recipe Interface
All commands implement the `rc_recipe.Recipe` interface:

```go
type Recipe interface {
    Preset()                           // Initialize defaults
    Exec(c app_control.Control) error  // Execute command
    Test(c app_control.Control) error  // Test command
}
```

### Creating a New Command

1. **Choose the right directory**:
   - `recipe/` for utilities and tools
   - `citron/<service>/` for external integrations

2. **Create the recipe file**:
```go
package mycommand

import (
    "github.com/watermint/toolbox/infra/control/app_control"
    "github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type MyCommand struct {
    rc_recipe.RemarkIrreversible     // Add appropriate remarks
    Path  mo_path.DropboxPath        // Input parameters
    Force mo_bool.Bool                // Flags
}

func (z *MyCommand) Preset() {
    z.Force.SetBool(false)  // Set defaults
}

func (z *MyCommand) Exec(c app_control.Control) error {
    // Implementation
    return nil
}

func (z *MyCommand) Test(c app_control.Control) error {
    return rc_spec.RecipeSpec(t, &MyCommand{})
}
```

3. **Add message resources** (see DOCUMENTATION.md)

4. **Update catalogue**:
```bash
go run . dev build catalogue
```

### Command Naming Convention

Commands follow the directory structure:
- File: `citron/dropbox/file/copy.go`
- Command: `dropbox file copy`
- Message prefix: `dropbox.file.copy`

## Recipe Components

### Remarks
Provide metadata about commands:
- `RemarkIrreversible` - Makes permanent changes
- `RemarkExperimental` - Experimental feature
- `RemarkConsole` - Outputs to console
- `RemarkSecret` - Handles sensitive data

### Value Types
Common input parameter types:
- `mo_string.String` - String input
- `mo_path.DropboxPath` - Dropbox path
- `mo_path.FileSystemPath` - Local path
- `mo_int.Int` - Integer input
- `mo_bool.Bool` - Boolean flag
- `mo_filter.Filter` - Filter expressions

### Connections
For external services:
- `dbx_conn.ConnScopedUser` - Dropbox user
- `dbx_conn.ConnScopedTeam` - Dropbox team admin
- `gh_conn.ConnGithubPublic` - GitHub public
- `sv_slack.ConnSlack` - Slack workspace

### Reports
Output data structures:
```go
type Results struct {
    File   mo_path.DropboxPath
    Size   int64
    Status string
}

func (z *MyCommand) Exec(c app_control.Control) error {
    rp := c.NewReport("results")
    rp.Row(&Results{...})
    return nil
}
```

## Message Variables in Recipes

Define UI messages as struct fields:

```go
type MyCommand struct {
    ProgressScan    app_msg.Message  // Auto-mapped to recipe.mycommand.progress_scan
    ErrorNotFound   app_msg.Message  // Auto-mapped to recipe.mycommand.error_not_found
    SuccessComplete app_msg.Message  // Auto-mapped to recipe.mycommand.success_complete
}

func (z *MyCommand) Exec(c app_control.Control) error {
    ui := c.UI()
    ui.Progress(z.ProgressScan.With("Path", "/some/path"))
    ui.Error(z.ErrorNotFound.With("File", "test.txt"))
    ui.Success(z.SuccessComplete.With("Count", 10))
    return nil
}
```

## Testing Strategy

### Recipe Tests
- Focus on unit tests
- Mock external dependencies
- Run without network access
- Fast execution

Example:
```go
func TestMyCommand_Exec(t *testing.T) {
    rc_spec.RecipeSpec(t, &MyCommand{})
}
```

### Citron Tests
- Include integration tests
- May require test accounts
- Handle rate limits gracefully
- Test error scenarios

Example:
```go
func TestDropboxCommand_Integration(t *testing.T) {
    if !rc_compatibility.IsTestEnabled() {
        t.Skip("Integration test requires credentials")
    }
    // Integration test
}
```

## Best Practices

1. **Single Responsibility** - Each command does one thing well
2. **Clear Naming** - Descriptive names indicating purpose
3. **Error Handling** - Return meaningful errors with context
4. **Progress Feedback** - Show progress for long operations
5. **Validation** - Validate inputs early in Exec()
6. **Documentation** - Comprehensive message resources
7. **Testing** - Both unit and integration tests

## Adding a New Service

To add a new service integration:

1. Create service directory:
```
citron/newservice/
```

2. Organize by domain:
```
citron/newservice/
├── user/
├── project/  
└── admin/
```

3. Implement authentication:
```go
// domain/newservice/conn/conn.go
type ConnNewService interface {
    // Service-specific methods
}
```

4. Follow existing patterns from similar services

## Command Discovery

The toolbox automatically discovers commands:
1. Scans for types implementing `Recipe` interface
2. Registers in command catalogue
3. Generates CLI paths from package structure
4. Creates documentation from messages

This is handled by `dev build catalogue` command.