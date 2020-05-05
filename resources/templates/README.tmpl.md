# watermint toolbox
{{.Badges}}
{{.Logo}}

{{ msg "doc.readme.title" }}

# {{ msg "doc.readme.head.license" }}

{{ msg "doc.readme.body.license" }}

{{ msg "doc.readme.body.license_note" }}

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

{{ if .Release}}
# {{ msg "doc.readme.head.release" }}

{{ msg "doc.readme.body.release" }}

{{ end }}
# {{ msg "doc.readme.head.usage" }}

{{ msg "doc.readme.body.usage" }}

```
% ./tbx
{{.BodyUsage}}
```

## {{ msg "doc.readme.head.commands" }}

{{.Commands}}
