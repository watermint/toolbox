---
layout: page
title: 隠しカラム表示、出力フィルタリング、JSON処理を含む高度なレポート機能のガイド
lang: ja
---

# レポートガイド

# Advanced Reporting Features

Many commands in watermint toolbox generate reports with tabular data. There are three powerful features that can extend and modify command output for advanced reporting needs:

## 1. Show All Columns with -experimental report_all_columns

By default, commands display only the most commonly used columns to keep output readable. However, many commands have additional hidden columns that can be revealed using the experimental flag.

### Usage
```bash
go run . [command] -experimental report_all_columns
```

### Example: Dropbox Team Member List

**Standard output (limited columns):**
```bash
go run . dropbox team member list
```
Shows: Email, Status, Role, etc.

**Extended output (all columns):**
```bash
go run . dropbox team member list -experimental report_all_columns
```
Shows: Email, Status, Role, Team Member ID, Account ID, External ID, Profile, Date Joined, Groups, etc.

This reveals internal IDs and metadata that are often needed for automation and integration scenarios.

## 2. Filter Output with -output-filter

The output filter option allows you to run queries on command results using a SQL-like syntax. This requires `-output json` to be specified.

### Usage
```bash
go run . [command] -output json -output-filter "query expression"
```

### Query Syntax
- `select`: Choose columns to display
- `where`: Filter rows based on conditions
- `order by`: Sort results
- `limit`: Limit number of results

### Examples

**Filter active team members:**
```bash
go run . dropbox team member list -output json -output-filter "select email, status where status = 'active'"
```

**Find members by domain:**
```bash
go run . dropbox team member list -output json -output-filter "select email, team_member_id where email like '%@company.com'"
```

**Get top 10 largest files:**
```bash
go run . dropbox file list -output json -output-filter "select name, size order by size desc limit 10"
```

## 3. Post-Process with util json query

The `util json query` command provides jq-like functionality for processing JSON output from other commands. It's particularly useful for complex data transformations.

### Usage
```bash
go run . [command] -output json | go run . util json query -query "jq expression"
```

### Common Patterns

**Extract specific fields:**
```bash
go run . dropbox team member list -output json | go run . util json query -query ".[] | {email, team_member_id}"
```

**Group and count:**
```bash
go run . dropbox team member list -output json | go run . util json query -query "group_by(.status) | map({status: .[0].status, count: length})"
```

**Filter and transform:**
```bash
go run . dropbox file list -output json | go run . util json query -query ".[] | select(.size > 1000000) | .name"
```

## Complete Example: Team Member ID and Email Report

A common administrative task is to generate a report of all team member IDs with email addresses. Here's how to accomplish this using the three features:

### Step 1: Get all columns including team_member_id
```bash
go run . dropbox team member list -experimental report_all_columns -output json
```

### Step 2: Filter to get only team_member_id and email
```bash
go run . dropbox team member list -experimental report_all_columns -output json -output-filter "select team_member_id, email"
```

### Step 3: Alternative using util json query
```bash
go run . dropbox team member list -experimental report_all_columns -output json | go run . util json query -query ".[] | {team_member_id, email}"
```

### Step 4: Export to CSV for spreadsheet use
```bash
go run . dropbox team member list -experimental report_all_columns -output csv -output-filter "select team_member_id, email" > team_members.csv
```

## Best Practices

### 1. Discover Available Columns
Always start with `-experimental report_all_columns` to see what data is available:
```bash
go run . [command] -experimental report_all_columns | head -5
```

### 2. Use Output Filter for Simple Queries
For straightforward filtering and column selection, `-output-filter` is more efficient:
```bash
# Good for simple cases
go run . dropbox team member list -output json -output-filter "select email where status = 'active'"
```

### 3. Use util json query for Complex Processing
For complex transformations, aggregations, or when you need jq's full power:
```bash
# Good for complex cases
go run . dropbox team member list -output json | go run . util json query -query "group_by(.role) | map({role: .[0].role, members: map(.email)})"
```

### 4. Combine with Standard Tools
These features integrate well with standard Unix tools:
```bash
# Count active members
go run . dropbox team member list -output json -output-filter "select email where status = 'active'" | jq length

# Sort and save
go run . dropbox team member list -output csv -output-filter "select email, team_member_id order by email" > sorted_members.csv
```

## Performance Considerations

- `-experimental report_all_columns` may return more data and take slightly longer
- `-output-filter` is processed server-side and is generally faster for large datasets
- `util json query` processes data client-side and is better for complex transformations
- For very large datasets, consider using `-output-filter` to reduce data transfer

## Troubleshooting

### Common Issues

**Column not found in output-filter:**
```bash
# First check available columns
go run . [command] -experimental report_all_columns | head -1
```

**JSON parsing errors:**
```bash
# Ensure -output json is specified
go run . [command] -output json -output-filter "..."
```

**Complex jq expressions:**
```bash
# Test expressions incrementally
go run . [command] -output json | go run . util json query -query ".[0]"  # First record
go run . [command] -output json | go run . util json query -query ".[] | keys"  # Available fields
```

These advanced reporting features provide powerful ways to extract, filter, and transform data from watermint toolbox commands, enabling sophisticated reporting and automation workflows.


