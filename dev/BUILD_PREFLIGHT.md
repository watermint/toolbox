# Build Preflight Command Documentation

## Overview

The `dev build preflight` command is a critical development tool that ensures the project is in a consistent state before commits or releases. It validates message resources, generates documentation, and maintains code quality.

## What It Does

### 1. Documentation Cleanup
- Removes old generated documentation files from `docs/commands/` and `docs/guides/` directories
- Ensures a clean slate for regeneration

### 2. Documentation Generation
- Runs `dev build doc` internally to generate all documentation
- Creates documentation for all supported languages (English and Japanese by default)
- Generates:
  - README files
  - Command documentation
  - Web documentation
  - Security documentation
  - Knowledge base for AI systems

### 3. Message Resource Validation
- Verifies all message keys used in code are defined in message files
- Identifies missing translations
- Reports any undefined message keys

### 4. Message Resource Optimization
- Identifies unused message keys
- Automatically removes unused keys from message files
- Sorts message keys alphabetically for consistency
- Maintains clean and organized message files

### 5. Recipe and Component Verification
- Validates all recipes (commands) have proper message definitions
- Checks ingredients (reusable components) for proper messages
- Verifies features have appropriate opt-in messages
- Ensures all user-facing text is internationalized

## Command Usage

### Full Preflight Check
```bash
go run tbx.go dev build preflight
```

### Quick Check (English Only)
```bash
go run tbx.go dev build preflight -quick
```

## Implementation Details

### File Location
`recipe/dev/build/preflight.go`

### Key Methods

#### `Exec(c app_control.Control) error`
Main execution method that orchestrates the preflight process:
1. Iterates through supported languages
2. Cleans documentation directories
3. Generates documentation
4. Verifies messages
5. Sorts and cleans message files

#### `sortMessages(c app_control.Control, filename string) error`
- Reads message JSON files
- Identifies unused keys by comparing with touched keys
- Removes unused keys
- Sorts remaining keys alphabetically
- Writes back clean, sorted JSON

#### `deleteOldGeneratedFiles(c app_control.Control, path string) error`
- Removes all files in specified directory
- Used to clean generated documentation before regeneration
- Skips deletion in test mode

## Message Tracking System

The preflight command uses a sophisticated message tracking system:

### Touch Recording
- As recipes and components are processed, used message keys are "touched"
- The `qt_msgusage.Record()` system tracks all accessed keys
- After processing, unused keys are identified and removed

### Message Sources
1. **Recipes**: Command descriptions, arguments, and help text
2. **Ingredients**: Reusable component messages
3. **Message Objects**: Structured message definitions
4. **Features**: Experimental feature descriptions and agreements

## Integration with Build Process

### Before Commits
The preflight command should be run before every commit to ensure:
- Documentation is up-to-date
- Message files are clean
- No missing translations
- Consistent formatting

### In CI/CD Pipeline
Can be integrated into continuous integration:
```yaml
- name: Run preflight checks
  run: go run tbx.go dev build preflight
```

## Error Handling

### Common Errors

1. **Missing Message Keys**
   - Error: Message key not found in resources
   - Solution: Add missing keys to message files

2. **File Permission Issues**
   - Error: Unable to write to docs directory
   - Solution: Check file permissions

3. **Invalid JSON**
   - Error: Unable to unmarshal message file
   - Solution: Fix JSON syntax in message files

## Best Practices

1. **Run Regularly**: Execute preflight before commits and after adding new features
2. **Check Output**: Review removed keys to ensure they're truly unused
3. **Language Support**: Ensure all languages are properly maintained
4. **Quick Mode**: Use quick mode during development for faster iteration
5. **Full Check**: Always run full check before releases

## Workflow Integration

### Development Workflow
1. Make code changes
2. Run `dev build preflight -quick` for rapid feedback
3. Add/update message keys as needed
4. Run full `dev build preflight` before commit
5. Commit changes including updated documentation

### Release Workflow
1. Update version information
2. Run `dev release candidate`
3. Run `dev build preflight` (full check)
4. Review generated documentation
5. Proceed with release

## Technical Notes

### Performance Considerations
- Full preflight can take several minutes due to documentation generation
- Quick mode significantly reduces execution time
- Message sorting is performed in-memory for efficiency

### File System Operations
- Uses atomic file writes to prevent corruption
- Preserves file permissions when updating
- Handles missing directories gracefully

### Testing
- Test mode prevents actual file system modifications
- Outputs to stdout instead of files during tests
- Allows verification of logic without side effects