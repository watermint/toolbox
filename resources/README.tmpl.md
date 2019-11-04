# watermint toolbox
{{.Badges}}

{{ msg "doc.readme.title" }}

# {{ msg "doc.readme.head.license" }}

{{ msg "doc.readme.body.license" }}

{{ msg "doc.readme.body.license_note" }}

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.

# {{ msg "doc.readme.head.usage" }}

{{ msg "doc.readme.body.usage" }}

```
% ./tbx
watermint toolbox xx.x.xxx
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.


Tools for Dropbox and Dropbox Business

Usage:
./tbx  command

Available commands:
   file          File operation
   group         Group management (Dropbox Business)
   license       Show license information
   member        Team member management (Dropbox Business)
   sharedfolder  Shared folder
   sharedlink    Shared Link of Personal account
   team          Dropbox Business Team
   teamfolder    Team folder management (Dropbox Business)
   web           Launch web console (experimental)
```

## {{ msg "doc.readme.head.commands" }}

{{.Commands}}

## {{ msg "doc.readme.head.authentication" }}

{{ msg "doc.readme.body.authentication" }}

```
watermint toolbox xx.x.xxx
© 2016-2019 Takayuki Okazaki
Licensed under open source licenses. Use the `license` command for more detail.

Testing network connection...
Done

1. Visit the URL for the auth dialog:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. Click 'Allow' (you might have to login first):
3. Copy the authorisation code:
Enter the authorisation code
```

{{ msg "doc.readme.body.authentication_token_file" }}

## {{ msg "doc.readme.head.proxy" }}

{{ msg "doc.readme.body.proxy" }}
