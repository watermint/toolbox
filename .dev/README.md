# Development Documentation

Welcome to the watermint toolbox development documentation. This directory contains comprehensive guides for developers working on the toolbox project.

## Documentation Index

### Core Development Guides

1. **[BUILD.md](BUILD.md)** - Building and compilation instructions
   - Prerequisites and environment setup
   - Docker build process
   - CI/CD configuration
   - Release process

2. **[DOCUMENTATION_SYSTEM.md](DOCUMENTATION_SYSTEM.md)** - How documentation is generated
   - Documentation generation commands
   - Message resource system
   - Internationalization support
   - Adding new documentation

3. **[COMMAND_STRUCTURE.md](COMMAND_STRUCTURE.md)** - Recipe/command architecture
   - Creating new commands
   - Recipe components and patterns
   - Testing strategies
   - Best practices

4. **[RECIPE_CITRON_ARCHITECTURE.md](RECIPE_CITRON_ARCHITECTURE.md)** - Recipe vs Citron directories
   - Architectural separation
   - When to use each directory
   - Service integration patterns
   - Directory organization

5. **[BUILD_PREFLIGHT.md](BUILD_PREFLIGHT.md)** - Preflight command details
   - What preflight does
   - Message validation system
   - Integration with development workflow
   - Troubleshooting

6. **[MESSAGE_RESOURCES.md](MESSAGE_RESOURCES.md)** - Internationalization system
   - Message file structure
   - Key naming conventions
   - Adding translations
   - Template variables

7. **[DOCUMENTATION_FLOW.md](DOCUMENTATION_FLOW.md)** - Documentation generation flow
   - Step-by-step process
   - File generation details
   - Error handling
   - Performance considerations

## Quick Start for Developers

### Setting Up Development Environment
1. Clone the repository
2. Set up application keys (see [BUILD.md](BUILD.md))
3. Run `go mod download` to fetch dependencies
4. Run `go run tbx.go dev build preflight` to verify setup

### Common Development Tasks

#### Adding a New Command
1. Create recipe file in appropriate directory
2. Add message resources
3. Run `dev build catalogue`
4. Run `dev build preflight`
5. Test your command

#### Updating Documentation
1. Modify message resources or documentation templates
2. Run `dev build doc` to regenerate
3. Run `dev build preflight` to validate
4. Review generated files in `docs/`

#### Before Committing
Always run:
```bash
go run tbx.go dev build preflight
```

### Development Commands

Key development commands you'll use:

- `dev build preflight` - Validate and prepare for commit
- `dev build doc` - Generate documentation
- `dev build catalogue` - Update command catalogue
- `dev test recipe` - Run recipe tests
- `dev release candidate` - Prepare release

## Architecture Overview

### Project Structure
```
toolbox/
├── recipe/          # Utility and dev commands
├── citron/          # Service-specific commands
├── infra/           # Infrastructure code
├── domain/          # Domain-specific logic
├── essentials/      # Core utilities
├── resources/       # Messages and templates
├── docs/            # Generated documentation
└── dev/             # Development documentation
```

### Key Concepts
- **Recipes**: Command implementations
- **Ingredients**: Reusable components
- **Messages**: Internationalized text resources
- **Connections**: External service integrations
- **Reports**: Output formatting

## Contributing

### Code Style
- Follow Go conventions
- Use meaningful variable names
- Add comprehensive error handling
- Include tests for new functionality

### Documentation Style
- Keep documentation concise and clear
- Include examples where helpful
- Update relevant docs when changing functionality
- Maintain both English and Japanese messages

## Testing

### Running Tests
```bash
# Run all tests
go test -p 1 ./...

# Run specific package tests
go test ./recipe/dev/build/...

# Run with coverage
go test -cover ./...
```

### Test Categories
- Unit tests: Fast, isolated tests
- Integration tests: Tests with external dependencies
- End-to-end tests: Full command execution tests

## Debugging

### Enable Debug Logging
```bash
tbx -debug <command>
```

### Capture Detailed Logs
```bash
tbx -capture-path logs <command>
```

### Experiment Mode
```bash
tbx -experiment <command>
```

## Resources

### Internal Documentation
- Message files: `resources/messages/`
- Templates: `resources/templates/`
- Build scripts: `resources/build/`

### External Resources
- [Project Website](https://toolbox.watermint.org)
- [GitHub Repository](https://github.com/watermint/toolbox)
- [Release Notes](../docs/releases/)

## Getting Help

If you need help:
1. Check the documentation in this directory
2. Look at existing code for examples
3. Run commands with `-help` flag
4. Check the test files for usage examples

## Maintenance

### Regular Tasks
- Update dependencies: `go mod tidy`
- Run security checks: `go mod audit`
- Update documentation: `dev build preflight`
- Clean build artifacts: `go clean`

### Release Checklist
1. Update version in `resources/release/release`
2. Update release notes
3. Run `dev release candidate`
4. Run `dev build preflight`
5. Create and push release

---

Remember: Always run `dev build preflight` before committing changes!