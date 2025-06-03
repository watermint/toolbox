# Command Structure Guide

This guide explains how commands (recipes) are structured in the watermint toolbox project.

## Overview

Commands in watermint toolbox are implemented as "recipes" - self-contained units of functionality that can be executed through the CLI. Each recipe follows a consistent structure and naming convention.

## Recipe Architecture

### 1. Recipe Interface
All recipes implement the `rc_recipe.Recipe` interface, which includes:
- `Preset()`: Initialize default values
- `Exec(c app_control.Control) error`: Execute the recipe
- `Test(c app_control.Control) error`: Test the recipe

### 2. Recipe Location
Recipes are organized in two main directories:

#### `recipe/` Directory
Contains utility and development commands:
```
recipe/
├── config/         # Configuration commands
├── dev/            # Development commands
│   ├── build/      # Build-related commands
│   ├── test/       # Test commands
│   └── release/    # Release commands
├── util/           # Utility commands
├── log/            # Logging commands
└── license.go      # License command
```

#### `citron/` Directory
Contains service-specific commands:
```
citron/
├── asana/          # Asana commands
│   ├── team/       # Team operations
│   └── workspace/  # Workspace operations
├── deepl/          # DeepL translation commands
│   └── translate/  # Translation operations
├── dropbox/        # Dropbox commands
│   ├── file/       # File operations
│   ├── paper/      # Paper operations
│   ├── sign/       # Dropbox Sign operations
│   └── team/       # Team operations
├── figma/          # Figma commands
│   ├── account/    # Account operations
│   ├── file/       # File operations
│   └── project/    # Project operations
├── github/         # GitHub commands
│   ├── content/    # Repository content
│   ├── issue/      # Issue operations
│   └── release/    # Release operations
├── local/          # Local file operations
│   └── file/       # File operations
└── slack/          # Slack commands
    └── conversation/ # Conversation operations
```

### 3. Command Naming Convention
The command path is derived from the package structure:
- Package: `citron/dropbox/file/copy.go`
- Command: `dropbox file copy`
- Package: `recipe/util/text/case/up.go`
- Command: `util text case up`

## Creating a New Command

### Step 1: Create the Recipe File
Create a new Go file in the appropriate directory:
```go
package mypackage

import (
    "github.com/watermint/toolbox/infra/control/app_control"
    "github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

type MyCommand struct {
    rc_recipe.RemarkIrreversible  // Add appropriate remarks
    Peer  dbx_conn.ConnScopedUser  // Connection for Dropbox
    Path  mo_path.DropboxPath      // Input parameter
}

func (z *MyCommand) Preset() {
    // Set default values
}

func (z *MyCommand) Exec(c app_control.Control) error {
    // Implementation
    return nil
}

func (z *MyCommand) Test(c app_control.Control) error {
    // Test implementation
    return nil
}
```

### Step 2: Add Message Resources
Add entries to `resources/messages/en/messages.json`:
```json
{
  "mypackage.mycommand.desc": "Brief description of the command",
  "mypackage.mycommand.readme": "Detailed description for documentation",
  "mypackage.mycommand.arg.path": "Path",
  "mypackage.mycommand.arg.path.desc": "Path to the file or folder"
}
```

### Step 3: Register the Command
Commands are automatically discovered through the catalogue system. Run:
```bash
go run tbx.go dev build catalogue
```

### Step 4: Generate Documentation
```bash
go run tbx.go dev build preflight
```

## Recipe Components

### 1. Remarks
Remarks provide metadata about the recipe:
- `RemarkIrreversible`: Command makes permanent changes
- `RemarkExperimental`: Command is experimental
- `RemarkConsole`: Command outputs to console
- `RemarkSecret`: Command handles sensitive data

### 2. Value Objects
Input parameters are defined as struct fields with appropriate types:
- `mo_string.String`: String input
- `mo_path.DropboxPath`: Dropbox path
- `mo_int.Int`: Integer input
- `mo_bool.Bool`: Boolean flag

### 3. Connections
For commands that interact with external services:
- `dbx_conn.ConnScopedUser`: Dropbox user connection
- `dbx_conn.ConnScopedTeam`: Dropbox team connection
- `gh_conn.ConnGithubPublic`: GitHub connection

## Command Categories

### Recipe Directory Commands
Commands in the `recipe/` directory are primarily for toolbox management and utilities:

#### Development Commands (`recipe/dev`)
- **build**: Build and compilation commands
- **test**: Testing utilities
- **release**: Release management
- **benchmark**: Performance testing
- **doc**: Documentation generation
- **replay**: API replay utilities
- **ci**: Continuous integration tools

#### Utility Commands (`recipe/util`)
- **archive**: Archive operations (zip/unzip)
- **cert**: Certificate generation
- **database**: Database operations
- **datetime**: Date and time utilities
- **encode/decode**: Encoding operations (base32/base64)
- **image**: Image processing
- **json**: JSON manipulation
- **net**: Network operations
- **text**: Text processing and NLP
- **uuid**: UUID generation

#### Other Recipe Commands
- **config**: Configuration management
- **log**: Log viewing and management
- **license**: License information

### Citron Directory Commands
Commands in the `citron/` directory are service-specific integrations:

#### Dropbox Commands (`citron/dropbox`)
- **file**: File and folder operations
- **paper**: Dropbox Paper operations
- **sign**: Dropbox Sign (formerly HelloSign) operations
- **team**: Team administration

#### GitHub Commands (`citron/github`)
- **content**: Repository content management
- **issue**: Issue tracking
- **release**: Release management
- **tag**: Tag operations
- **profile**: User profile operations

#### Other Service Commands
- **asana**: Project management operations
- **deepl**: Translation services
- **figma**: Design file operations
- **slack**: Messaging and conversation operations
- **local**: Local file system operations

## Testing Commands

### Unit Tests
Each recipe should have a corresponding test file:
```go
func TestMyCommand_Exec(t *testing.T) {
    rc_spec.RecipeSpec(t, &MyCommand{})
}
```

### Integration Tests
For commands that interact with external services:
```go
func TestMyCommand_Integration(t *testing.T) {
    if !rc_compatibility.IsTestEnabled() {
        t.Skip()
    }
    // Integration test
}
```

## Best Practices

1. **Single Responsibility**: Each recipe should do one thing well
2. **Clear Naming**: Use descriptive names that indicate the command's purpose
3. **Comprehensive Messages**: Provide clear descriptions and help text
4. **Error Handling**: Return meaningful errors with context
5. **Testing**: Include both unit and integration tests
6. **Documentation**: Ensure all parameters are documented
7. **Validation**: Validate inputs early in the Exec method

## Command Discovery

The toolbox uses an automatic discovery system:
1. All types implementing `Recipe` interface are discovered
2. Commands are registered in the catalogue
3. CLI paths are generated from package structure
4. Documentation is auto-generated from messages

## Advanced Features

### 1. Feed Operations
For commands that process multiple items:
```go
type MyFeedCommand struct {
    Feed   fd_file.RowFeed
    Output rp_model.Report
}
```

### 2. Grid Data Sources
For commands that need data grid functionality:
```go
type MyGridCommand struct {
    Grid dg_model.GridDataSource
}
```

### 3. Custom Reports
For commands that generate reports:
```go
func (z *MyCommand) Exec(c app_control.Control) error {
    report := c.NewReport("results")
    // Add data to report
    return nil
}
```

## Debugging Commands

### Enable Debug Logging
```bash
tbx -debug mycommand
```

### Capture Logs
```bash
tbx -capture-path logs mycommand
```

### Test Mode
```bash
tbx -experiment mycommand
```