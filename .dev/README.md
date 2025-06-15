# Development Guide

Quick reference for watermint toolbox development.

## Quick Start

```bash
# Clone and setup
git clone https://github.com/watermint/toolbox.git
cd toolbox

# Install dependencies
go mod download

# Setup API keys (see BUILD.md for details)
cp resources/toolbox.appkeys.example resources/toolbox.appkeys
# Edit toolbox.appkeys with your API credentials

# Verify setup
go run . dev build preflight -quick

# Run a command
go run . version
```

## Key Commands

```bash
# Before committing - ALWAYS run this
go run . dev build preflight

# Development commands
go run . dev build doc          # Generate documentation
go run . dev build catalogue    # Update command registry
go run . dev test recipe       # Run tests
go run . dev release candidate # Prepare release
```

## Documentation

- **[ARCHITECTURE.md](ARCHITECTURE.md)** - Project structure and command system
- **[BUILD.md](BUILD.md)** - Building, testing, and releasing  
- **[DOCUMENTATION.md](DOCUMENTATION.md)** - Documentation and message system

## Project Structure

```
toolbox/
├── recipe/          # Toolbox utilities and dev commands
├── citron/          # External service integrations
├── infra/           # Infrastructure and framework
├── domain/          # Domain logic and models
├── essentials/      # Core utilities
├── resources/       # Messages, templates, keys
└── docs/            # Generated documentation
```

## Adding a New Command

1. Create recipe file in appropriate directory
2. Add message resources to `resources/messages/*/messages.json`
3. Run `go run . dev build catalogue`
4. Run `go run . dev build preflight`
5. Test your command

## Testing

```bash
# Run all tests (use -p 1 for stateful tests)
go test -p 1 ./...

# Run specific package tests
go test ./recipe/dev/build/...

# Enable debug output
go run . -debug <command>
```

## Common Tasks

| Task | Command |
|------|---------|
| Add new command | Create recipe → Add messages → Run catalogue → Run preflight |
| Update docs | Edit messages → Run `dev build doc` → Run preflight |
| Debug issue | Use `-debug` flag or `-capture-path logs` |
| Check setup | Run `dev build preflight` |

## Important Notes

- **Always run `dev build preflight` before committing**
- Follow existing code patterns and conventions
- Add messages for ALL user-facing text
- Test with both English and Japanese locales
- See individual guides for detailed information