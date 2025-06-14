# Recipe vs Citron Architecture

## Overview

The watermint toolbox organizes commands (recipes) into two main directories: `recipe/` and `citron/`. This separation provides a clear architectural boundary between toolbox utilities and external service integrations.

## Directory Purpose

### `recipe/` Directory
**Purpose**: Internal toolbox functionality, utilities, and development tools

This directory contains commands that:
- Manage the toolbox itself (configuration, logging)
- Provide general-purpose utilities (text processing, file operations)
- Support development workflows (build, test, release)
- Don't require external service authentication

### `citron/` Directory
**Purpose**: External service integrations

This directory contains commands that:
- Integrate with third-party services (Dropbox, GitHub, Slack, etc.)
- Require API authentication
- Perform service-specific operations
- Handle service-specific data formats

## Naming Origin

The name "citron" was chosen to represent external integrations, similar to how citrus fruits have distinct segments - each service integration is a self-contained segment with its own structure and functionality.

## Architecture Benefits

### 1. Clear Separation of Concerns
- Internal tools vs external integrations
- Easy to identify dependencies
- Simplified security review

### 2. Organized Codebase
- Service-specific code is isolated
- Common patterns within each directory
- Easier navigation for developers

### 3. Dependency Management
- `recipe/` commands have minimal external dependencies
- `citron/` commands handle service-specific SDKs
- Clear boundary for API client code

### 4. Testing Strategy
- `recipe/` commands can run without external services
- `citron/` commands may require integration testing
- Different test approaches for each category

## Command Examples

### Recipe Commands
```bash
# Development tools
tbx dev build preflight
tbx dev test recipe

# Utilities
tbx util text case up
tbx util uuid v4
tbx util file hash

# Configuration
tbx config feature list
tbx config auth list

# Logging
tbx log job list
tbx log cat last
```

### Citron Commands
```bash
# Dropbox
tbx dropbox file copy
tbx dropbox team member list

# GitHub
tbx github release list
tbx github issue list

# Slack
tbx slack conversation list

# Figma
tbx figma file list

# DeepL
tbx deepl translate text
```

## Implementation Guidelines

### When to Add to `recipe/`
Add new commands to `recipe/` when they:
- Don't require external service authentication
- Provide utility functions
- Support toolbox development
- Work with local resources only

### When to Add to `citron/`
Add new commands to `citron/` when they:
- Integrate with external APIs
- Require service authentication
- Handle service-specific data
- Implement service workflows

## Directory Structure

### Recipe Structure
```
recipe/
├── config/          # Configuration management
├── dev/             # Development tools
│   ├── build/       # Build commands
│   ├── test/        # Test utilities
│   └── release/     # Release management
├── log/             # Logging utilities
├── util/            # General utilities
│   ├── archive/     # Archive operations
│   ├── text/        # Text processing
│   └── ...          # Other utilities
└── license.go       # License display
```

### Citron Structure
```
citron/
├── <service>/       # Service name
│   ├── <domain>/    # Domain within service
│   │   ├── command1.go
│   │   └── command2.go
│   └── auth.go      # Authentication helpers
└── ...
```

## Service Integration Pattern

When adding a new service to `citron/`:

1. **Create service directory**
   ```
   citron/newservice/
   ```

2. **Organize by domain**
   ```
   citron/newservice/user/
   citron/newservice/project/
   citron/newservice/admin/
   ```

3. **Implement commands**
   ```go
   // citron/newservice/user/list.go
   package user
   
   type List struct {
       Peer ns_conn.ConnNewService
   }
   ```

4. **Add connection types**
   Define in `domain/newservice/api/` if needed

5. **Add message resources**
   Follow the naming convention for service commands

## Message Key Convention

### Recipe Commands
```
util.text.case.up.desc
dev.build.preflight.desc
config.auth.list.desc
```

### Citron Commands
```
dropbox.file.copy.desc
github.release.list.desc
slack.conversation.history.desc
```

## Testing Approach

### Recipe Tests
- Focus on unit tests
- Mock external dependencies
- Run without network access
- Fast execution

### Citron Tests
- Include integration tests
- May require test accounts
- Handle rate limits
- Test error scenarios

## Future Considerations

### Potential New Services
When adding new services to `citron/`:
- Follow existing patterns
- Maintain consistent structure
- Document authentication requirements
- Include comprehensive tests

### Recipe Evolution
As the toolbox grows:
- Keep utilities generic
- Avoid service-specific code
- Maintain backward compatibility
- Focus on developer productivity

## Summary

The recipe/citron architecture provides a clean separation between internal toolbox functionality and external service integrations. This structure makes the codebase more maintainable, testable, and understandable for both users and developers.