# Build Guide

## Prerequisites

### API Keys Setup

Create `resources/toolbox.appkeys` with your API credentials:

```json
{
  "user_full.key": "xxxxxxxxxxxxxx",
  "user_full.secret": "xxxxxxxxxxxxxx", 
  "business_info.key": "xxxxxxxxxxxxxx",
  "business_info.secret": "xxxxxxxxxxxxxx",
  "business_file.key": "xxxxxxxxxxxxxx",
  "business_file.secret": "xxxxxxxxxxxxxx",
  "business_management.key": "xxxxxxxxxxxxxx",
  "business_management.secret": "xxxxxxxxxxxxxx",
  "business_audit.key": "xxxxxxxxxxxxxx",
  "business_audit.secret": "xxxxxxxxxxxxxx"
}
```

Register applications at:
- Dropbox: https://www.dropbox.com/developers/apps
- GitHub: https://github.com/settings/developers
- Other services as needed

## Build Commands

### Essential Commands

```bash
# ALWAYS run before committing
go run . dev build preflight

# Quick check (English only, faster)
go run . dev build preflight -quick

# Update command catalogue after adding recipes
go run . dev build catalogue

# Generate documentation only
go run . dev build doc
```

### The Preflight Command

`dev build preflight` is critical for maintaining code quality. It:

1. **Cleans old documentation** - Removes outdated files
2. **Generates all docs** - Creates command docs, README, guides
3. **Validates messages** - Ensures all UI text is defined
4. **Optimizes resources** - Removes unused message keys
5. **Sorts message files** - Maintains consistent ordering

#### What Preflight Validates
- All recipe messages have required keys (`.desc`, `.readme`, `.arg.*`)
- All ingredients have proper messages
- Features have opt-in messages
- Message templates are valid
- No missing translations

#### Common Issues
- **Missing message key**: Add to `resources/messages/*/messages.json`
- **Invalid JSON**: Check syntax in message files
- **Permission denied**: Check file permissions in docs/

## Building Executables

### Local Build
```bash
# Build for current platform
go build -o tbx

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o tbx-linux
GOOS=darwin GOARCH=amd64 go build -o tbx-mac
GOOS=windows GOARCH=amd64 go build -o tbx.exe
```

### Docker Build
```bash
docker-compose build
docker-compose run build
```

### CI/CD Environment Variables

For GitHub Actions:

1. **TOOLBOX_APPKEYS** - JSON string with API keys
2. **TOOLBOX_BUILDERKEY** - Random string for obfuscation
3. **TOOLBOX_DEPLOY_TOKEN** - Dropbox token for deployment
4. **TOOLBOX_REPLAY_URL** - Shared link for replay data
5. **TOOLBOX_BUILD_TARGET** - Target like `darwin/amd64`

## Testing

### Setup Integration Tests
```bash
# Connect test accounts (one-time setup)
go run . dev ci auth connect
```

### Run Tests
```bash
# All tests (use -p 1 for stateful tests)
go test -p 1 ./...

# Specific package
go test ./recipe/dev/build/...

# With coverage
go test -cover ./...

# Verbose output
go test -v ./...
```

### Test Categories
- **Unit tests** - Fast, isolated, no external deps
- **Integration tests** - Require credentials, network access
- **End-to-end tests** - Full command execution

## Release Process

### 1. Update Release Info
```bash
# Edit version number
vim resources/release/release

# Edit release notes  
vim resources/release/release_notes
```

### 2. Prepare Release
```bash
# Validates and updates resources
go run . dev release candidate
```

### 3. Final Preflight
```bash
# Full validation
go run . dev build preflight
```

### 4. Publish Release
```bash
# Creates GitHub release and uploads assets
go run . dev release publish
```

## Development Workflow

### Before Starting
1. Pull latest changes
2. Run `go mod download`
3. Verify with `dev build preflight -quick`

### While Developing
1. Make changes
2. Add/update messages
3. Run relevant tests
4. Use `-debug` flag for troubleshooting

### Before Committing
1. Run `dev build preflight` (full check)
2. Review changed files
3. Ensure tests pass
4. Commit including generated docs

### Debug Tools
```bash
# Enable debug logging
go run . -debug <command>

# Capture detailed logs
go run . -capture-path logs <command>

# Experiment mode
go run . -experiment <command>
```

## Project Website

### Local Preview
```bash
docker run --rm \
  --volume="$(pwd):/srv/jekyll" \
  -p 4000:4000 \
  jekyll/jekyll:stable \
  jekyll serve \
  --config /srv/jekyll/docs/_config.yml,/srv/jekyll/docs/_config_dev.yml \
  --destination /tmp/staging \
  --source /srv/jekyll/docs \
  --watch
```

Visit http://localhost:4000

## Performance Tips

- Use `-quick` flag during development
- Run full preflight before commits
- Parallel test execution except for stateful tests
- Cache dependencies with `go mod download`

## Troubleshooting

### Build Failures
- Check Go version compatibility
- Verify all dependencies: `go mod tidy`
- Clear cache: `go clean -modcache`

### Test Failures  
- Check for required env variables
- Verify test account credentials
- Look for rate limit errors
- Run with `-v` for details

### Documentation Issues
- Ensure write permissions
- Check disk space
- Validate JSON syntax
- Review preflight output