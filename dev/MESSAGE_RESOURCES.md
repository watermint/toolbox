# Message Resources System

## Overview

The watermint toolbox uses a comprehensive message resource system for internationalization (i18n) and maintaining consistent user-facing text across the application. All user-visible text is stored in JSON files and referenced by keys throughout the codebase.

## Message File Structure

### Location
Message files are located in `resources/messages/`:
```
resources/
└── messages/
    ├── en/
    │   └── messages.json    # English messages
    └── ja/
        └── messages.json    # Japanese messages
```

### JSON Format
Messages are stored as key-value pairs:
```json
{
  "app.name": "watermint toolbox",
  "app.desc": "The multi-purpose utility command-line tool for web services",
  "dropbox.file.copy.desc": "Copy files or folders",
  "dropbox.file.copy.arg.from": "Source path",
  "dropbox.file.copy.arg.to": "Destination path"
}
```

## Message Key Conventions

### General Pattern
Keys follow a hierarchical naming convention:
```
<component>.<subcomponent>.<type>
```

### Recipe (Command) Messages
For commands, the pattern is:
```
<package>.<command>.<message_type>
```

Examples:
- `dropbox.file.copy.desc` - Command description
- `dropbox.file.copy.readme` - Detailed readme text
- `dropbox.file.copy.arg.path` - Argument name
- `dropbox.file.copy.arg.path.desc` - Argument description

### Common Message Types

#### Command Messages
- `.desc` - Brief description (shown in command list)
- `.readme` - Detailed description (shown in documentation)
- `.arg.<name>` - Argument display name
- `.arg.<name>.desc` - Argument description
- `.arg.<name>.value.<value>` - Predefined argument values

#### UI Messages
- `.progress` - Progress messages
- `.error` - Error messages
- `.success` - Success messages
- `.confirm` - Confirmation prompts

#### Feature Messages
- `.feature.<name>.desc` - Feature description
- `.feature.<name>.disclaimer` - Feature disclaimer
- `.feature.<name>.agreement` - User agreement text

## Automatic Key Mapping

### Recipe Mapping
The system automatically maps recipe types to message keys:
- Type: `recipe/dropbox/file/Copy`
- Auto-mapped key prefix: `dropbox.file.copy`

### Value Object Mapping
For value objects (command parameters):
- Type: `CopyOpts`
- Field: `Path`
- Auto-mapped key: `copy_opts.path`

## Adding New Messages

### Step 1: Identify Required Keys
When adding a new command or feature, identify all user-facing text:
- Command description
- Parameter descriptions
- Error messages
- Success confirmations
- Progress indicators

### Step 2: Add to English Messages
Edit `resources/messages/en/messages.json`:
```json
{
  "myfeature.command.desc": "Brief description of the command",
  "myfeature.command.readme": "Detailed description for documentation",
  "myfeature.command.arg.input": "Input file",
  "myfeature.command.arg.input.desc": "Path to the input file",
  "myfeature.command.progress": "Processing {{.current}} of {{.total}}...",
  "myfeature.command.success": "Successfully processed {{.count}} items"
}
```

### Step 3: Add Translations
Add corresponding entries to `resources/messages/ja/messages.json`:
```json
{
  "myfeature.command.desc": "コマンドの簡潔な説明",
  "myfeature.command.readme": "ドキュメント用の詳細な説明",
  "myfeature.command.arg.input": "入力ファイル",
  "myfeature.command.arg.input.desc": "入力ファイルへのパス",
  "myfeature.command.progress": "{{.total}}件中{{.current}}件を処理中...",
  "myfeature.command.success": "{{.count}}件のアイテムを正常に処理しました"
}
```

### Step 4: Validate Messages
Run preflight to ensure all messages are properly defined:
```bash
go run tbx.go dev build preflight
```

## Using Messages in Code

### Basic Usage
```go
// Get message through UI
message := c.UI().Text(messages.MMyMessage)

// With template variables
progress := c.UI().Text(messages.MProgress.With("current", 10).With("total", 100))
```

### Message Objects
Define message objects for complex messages:
```go
type MsgCommandProgress struct {
    Current int
    Total   int
}

func (m MsgCommandProgress) Key() string {
    return "mycommand.progress"
}
```

## Template Variables

Messages support Go template syntax for dynamic content:
```json
{
  "file.process.progress": "Processing {{.filename}} ({{.percent}}% complete)"
}
```

Use in code:
```go
msg := messages.MProcessProgress.With("filename", "data.csv").With("percent", 75)
```

## Message Validation

### Automatic Validation
The preflight command automatically:
- Checks for missing keys
- Identifies unused keys
- Validates template syntax
- Ensures consistency across languages

### Manual Testing
Test messages appear correctly:
```go
func TestMyMessages(t *testing.T) {
    mc := app_msg_container.NewSingle("en")
    text := mc.Text(messages.MMyMessage)
    assert.NotEmpty(t, text)
    assert.NotContains(t, text, "{{")  // No unresolved templates
}
```

## Best Practices

### 1. Consistent Naming
- Use lowercase with dots as separators
- Follow the hierarchical structure
- Be descriptive but concise

### 2. Message Content
- Keep messages concise and clear
- Use active voice
- Include necessary context
- Avoid technical jargon

### 3. Templates
- Use meaningful variable names
- Provide all required variables
- Test with various input lengths

### 4. Translations
- Maintain the same meaning across languages
- Consider cultural context
- Test with native speakers when possible

### 5. Error Messages
- Be specific about what went wrong
- Suggest how to fix the issue
- Include relevant details (paths, values)

## Troubleshooting

### Missing Key Errors
```
Error: message key not found: mycommand.desc
```
Solution: Add the key to message files

### Template Errors
```
Error: template: mycommand.progress:1:2: executing "mycommand.progress" at <.count>: can't evaluate field count
```
Solution: Ensure all template variables are provided

### Unused Keys
Preflight removes unused keys automatically. To preserve a key:
1. Ensure it's actually used in code
2. Or add a reference in test code

## Advanced Topics

### Dynamic Key Generation
Sometimes keys need to be generated dynamically:
```go
key := fmt.Sprintf("status.%s.desc", statusCode)
msg := mc.Text(app_msg.CreateMessage(key))
```

### Fallback Messages
For optional messages with fallbacks:
```go
if mc.HasKey("specific.message") {
    return mc.Text(messages.MSpecific)
}
return mc.Text(messages.MGeneric)
```

### Message Inheritance
Some messages can inherit from more general ones:
- Specific: `dropbox.file.copy.error.permission`
- General: `file.error.permission`
- Fallback: `error.permission`

## Maintenance

### Regular Tasks
1. Run preflight to clean up unused keys
2. Review messages for clarity
3. Update translations when English changes
4. Test template rendering

### Adding Language Support
To add a new language:
1. Create new directory `resources/messages/<lang>/`
2. Copy English messages.json
3. Translate all entries
4. Update `es_lang.Supported` to include new language
5. Test thoroughly

Remember: The message system is crucial for user experience. Take time to write clear, helpful messages that guide users effectively.