# Documentation System Guide

This guide explains how the watermint toolbox documentation system works and how to maintain it.

## Overview

The watermint toolbox uses an automated documentation generation system that creates multiple types of documentation from source code and message resources. The primary command for documentation generation is `dev build doc`.

## Key Components

### 1. Documentation Generation Commands

#### `dev build preflight`
- **Purpose**: Validates and prepares the project for documentation generation
- **Location**: `recipe/dev/build/preflight.go`
- **Key Functions**:
  - Cleans up old generated documentation files
  - Generates README and command documentation for all supported languages
  - Verifies message resources
  - Removes unused message keys
  - Sorts message files alphabetically

#### `dev build doc`
- **Purpose**: Main documentation generation command
- **Location**: `recipe/dev/build/doc.go`
- **Key Functions**:
  - Generates README files
  - Generates SECURITY_AND_PRIVACY documents
  - Generates command manuals for all recipes
  - Generates web documentation
  - Generates supplemental documentation
  - Generates contributor documentation
  - Generates knowledge base documentation for LLM training

### 2. Documentation Types

#### Command Documentation
- **Generated to**: `docs/commands/` and `docs/ja/commands/`
- **Content**: Individual markdown files for each command
- **Format**: Uses Jekyll-compatible front matter for web rendering
- **Includes**: 
  - Command description
  - Usage examples
  - Available options
  - Authentication requirements

#### README Files
- **Main README**: Root `README.md`
- **Japanese README**: `README_ja.md`
- **Generated from**: Templates and message resources
- **Includes**: Project overview, installation instructions, command list

#### Web Documentation
- **Location**: `docs/` directory
- **Format**: Jekyll-compatible markdown
- **Includes**: Home page, guides, command references

#### Knowledge Base
- **Purpose**: Training data for LLM systems
- **Location**: `docs/knowledge/`
- **Content**: Consolidated documentation optimized for AI training

### 3. Message Resources

The documentation system heavily relies on message resources for internationalization:

- **Location**: `resources/messages/`
- **Languages**: English (`en/messages.json`) and Japanese (`ja/messages.json`)
- **Key Format**: Follows a naming convention based on recipe paths
  - Example: `file.copy.desc` for the description of the file copy command

### 4. Documentation Generation Process

1. **Clean Up**: Remove old generated files from previous runs
2. **Message Verification**: Check all message keys are defined
3. **Generate Documents**: Create documentation for each type
4. **Sort Messages**: Alphabetically sort message files and remove unused keys
5. **Validation**: Ensure all required documentation is generated

## How to Generate Documentation

### Full Documentation Build
```bash
go run tbx.go dev build doc
```

### Quick Preflight Check (English only)
```bash
go run tbx.go dev build preflight -quick
```

### Before Committing Changes
Always run preflight to ensure documentation is up to date:
```bash
go run tbx.go dev build preflight
```

## Adding New Documentation

### 1. For New Commands
When adding a new recipe/command:
1. Add appropriate message keys in `resources/messages/en/messages.json`
2. Add Japanese translations in `resources/messages/ja/messages.json`
3. Run `dev build preflight` to generate documentation

### 2. For New Guides
1. Create content in the appropriate `dc_supplemental` or `dc_contributor` modules
2. Add message resources for the content
3. Run documentation generation

### 3. Message Key Convention
Follow the existing naming convention:
- `<package>.<command>.desc`: Command description
- `<package>.<command>.arg.<name>`: Argument description
- `<package>.<command>.arg.<name>.desc`: Detailed argument description

## Documentation Infrastructure

### Key Packages
- `infra/doc/dc_command`: Command documentation generation
- `infra/doc/dc_readme`: README generation
- `infra/doc/dc_section`: Section and layout management
- `infra/doc/dc_web`: Web documentation generation
- `infra/doc/dc_knowledge`: Knowledge base generation

### Media Types
- `MediaRepository`: Repository-level documentation (README, etc.)
- `MediaWeb`: Web documentation (Jekyll site)
- `MediaKnowledge`: Knowledge base for AI training

## Troubleshooting

### Missing Message Keys
If you see errors about missing message keys:
1. Check the key name in the error message
2. Add the key to both English and Japanese message files
3. Run preflight again

### Documentation Not Generated
1. Ensure you're not in test mode (`c.Feature().IsTest()`)
2. Check file permissions in the docs directory
3. Look for errors in the log output

### Unused Keys Warning
The preflight command automatically removes unused keys. This helps keep the message files clean.

## Best Practices

1. **Always run preflight before committing** to ensure documentation is synchronized
2. **Add messages for all user-facing text** to support internationalization
3. **Follow the naming convention** for message keys to ensure automatic mapping works
4. **Test documentation generation** after adding new commands or features
5. **Review generated documentation** to ensure it reads well and is accurate