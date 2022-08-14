---
layout: release
title: Changes of Release 106
lang: en
---

# Changes between `Release 106` to `Release 107`

# Commands added

| Command                             | Title                                                                    |
|-------------------------------------|--------------------------------------------------------------------------|
| dev build compile                   | Create build script                                                      |
| dev build target                    | Generate target build script                                             |
| dev kvs benchmark                   | KVS engine benchmark                                                     |
| dev stage encoding                  | Encoding test command (upload a dummy file with specified encoding name) |
| services google calendar event list | List Google Calendar events                                              |
| util archive unzip                  | Extract the zip archive file                                             |
| util archive zip                    | Compress target files into the zip archive                               |
| util database exec                  | Execute query on SQLite3 database file                                   |
| util database query                 | Query SQLite3 database                                                   |
| util file hash                      | Print file digest                                                        |
| util image exif                     | Print EXIF metadata of image file                                        |
| util monitor client                 | Start device monitor client                                              |
| util net download                   | Download a file                                                          |
| util text case down                 | Print lower case text                                                    |
| util text case up                   | Print upper case text                                                    |
| util text encoding from             | Convert text encoding to UTF-8 text file from specified encoding.        |
| util text encoding to               | Convert text encoding to specified encoding from UTF-8 text file.        |

# Commands deleted

| Command    | Title                                   |
|------------|-----------------------------------------|
| image info | Show EXIF information of the image file |

# Command spec changed: `config disable`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "config disable",
- 	CliArgs:         "",
+ 	CliArgs:         "-key FEATURE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `config enable`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "config enable",
- 	CliArgs:         "",
+ 	CliArgs:         "-key FEATURE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev benchmark local`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev benchmark local",
- 	CliArgs:         "-path /LOCAL/PATH/TO/PROCESS",
+ 	CliArgs:         `-num-files NUM -path /LOCAL/PATH/TO/PROCESS -size-max-kb NUM -size-min-kb NUM"`,
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev benchmark upload`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev benchmark upload",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/PROCESS",
+ 	CliArgs:         "-num-files NUM -path /DROPBOX/PATH/TO/PROCESS -size-max-kb NUM -size-min-kb NUM",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev benchmark uploadlink`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev benchmark uploadlink",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/UPLOAD",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/UPLOAD -size-kb NUM",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev build package`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev build package",
  	CliArgs: strings.Join({
  		"-build-path /LOCAL/PATH/",
- 		"OF/build -deploy-path /DROPBOX/PATH/TO/deploy -dest-path /LOCAL/",
- 		"PATH/TO/save_package",
+ 		"TO/build -dist-path /LOCAL/PATH/TO/dist -platform PLATFORM_TYPE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "BuildPath", Desc: "Full path to the binary", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "DeployPath", Desc: "Deploy destination folder path (remote)", TypeName: "essentials.model.mo_string.opt_string"},
  		&{
- 			Name:    "DestPath",
+ 			Name:    "DistPath",
  			Desc:    "Package destination folder path (local)",
  			Default: "",
  			... // 2 identical fields
  		},
  		&{Name: "Platform", Desc: "Platform name like win/linux/mac", TypeName: "string"},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

# Command spec changed: `dev ci artifact up`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "dev ci artifact up",
  	CliArgs: strings.Join({
  		"-dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF",
  		"/ARTIFACT",
+ 		" -timeout NUM",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev replay approve`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev replay approve",
- 	CliArgs:         "",
+ 	CliArgs:         "-id JOB_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev replay recipe`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev replay recipe",
- 	CliArgs:         "",
+ 	CliArgs:         "-id JOB_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev stage griddata`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev stage griddata",
- 	CliArgs:         "",
+ 	CliArgs:         "-in /LOCAL/PATH/TO/INPUT.csv",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev test echo`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test echo",
- 	CliArgs:         "",
+ 	CliArgs:         "-text VALUE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev test panic`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "dev test panic",
- 	CliArgs:         "",
+ 	CliArgs:         "-panic-type VALUE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `dev test setup teamsharedlink`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "dev test setup teamsharedlink",
- 	CliArgs:         "",
+ 	CliArgs:         "-group GROUP_NAME -num-links-per-member NUM -query QUERY",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `file paper append`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper append",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/append",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/APPEND",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file paper create`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper create",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/create",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/CREATE",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file paper overwrite`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper overwrite",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/overwrite",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/OVERWRIT",
+ 		"E",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file paper prepend`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file paper prepend",
  	CliArgs: strings.Join({
  		"-",
- 		"path /DROPBOX/PATH/TO/prepend",
+ 		"content /LOCAL/PATH/TO/INPUT.txt -path /DROPBOX/PATH/TO/PREPEND",
  		".paper",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file revision download`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "file revision download",
  	CliArgs: strings.Join({
  		"-local-path /LOCAL/PATH/TO/DOWNLOAD",
+ 		" -revision REVISION",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file revision restore`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file revision restore",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/RESTORE -revision REVISION",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file search content`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file search content",
- 	CliArgs:         "",
+ 	CliArgs:         "-query QUERY",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file search name`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file search name",
- 	CliArgs:         "",
+ 	CliArgs:         "-query QUERY",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `file share info`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "file share info",
- 	CliArgs:         "",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/GET_INFO",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{
  			Name:     "Path",
  			Desc:     "File",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.dropbox.model.mo_path.dropbox_path_impl",
  			TypeAttr: nil,
  		},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

# Command spec changed: `filerequest create`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "(Irreversible operation)",
  	Path:    "filerequest create",
  	CliArgs: strings.Join({
  		"-path /DROPBOX/PATH/OF/FILE",
+ 		"_",
  		"REQUEST",
+ 		" -title TITLE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `filerequest delete url`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "filerequest delete url",
- 	CliArgs:         "",
+ 	CliArgs:         "-url URL",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Force", Desc: "Force delete the file request.", Default: "false", TypeName: "bool", ...},
  		&{Name: "Peer", Desc: "Account alias", Default: "default", TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_individual", ...},
  		&{
  			Name:     "Url",
  			Desc:     "URL of the file request.",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "domain.dropbox.model.mo_url.url_impl",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

# Command spec changed: `group add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "group add",
- 	CliArgs:         "",
+ 	CliArgs:         "-name GROUP_NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `group member add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "group member add",
- 	CliArgs:         "",
+ 	CliArgs:         "-group-name GROUP_NAME -member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `group member delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "group member delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-group-name GROUP_NAME -member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `group rename`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "group rename",
- 	CliArgs:         "",
+ 	CliArgs:         "-current-name CURRENT_NAME -new-name NEW_NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `job log jobid`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "job log jobid",
- 	CliArgs:         "",
+ 	CliArgs:         "-id JOB_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `member file lock all release`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "member file lock all release",
  	CliArgs: strings.Join({
  		"-",
+ 		"member-email VALUE -",
  		"path /DROPBOX/PATH/TO/RELEASE/LOCK",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `member file lock list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member file lock list",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST",
+ 	CliArgs:         "-member-email EMAIL -path /DROPBOX/PATH/TO/LIST_LOCK",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `member file lock release`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "member file lock release",
  	CliArgs: strings.Join({
  		"-",
+ 		"member-email VALUE -",
  		"path /DROPBOX/PATH/TO/RELEASE/LOCK",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `member file permdelete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "member file permdelete",
  	CliArgs: strings.Join({
  		"-",
+ 		"member-email EMAIL -",
  		"path /DROPBOX/PATH/TO/",
- 		"PERM_",
  		"DELETE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `member suspend`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member suspend",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `member unsuspend`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "member unsuspend",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github content get`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services github content get",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPOSITORY -path PATH",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github content put`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services github content put",
- 	CliArgs:         "-content /LOCAL/PATH/TO/content",
+ 	CliArgs:         " -owner OWNER -repository REPO -path PATH -content /LOCAL/PATH/TO/content -message MSG",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github issue list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "services github issue list",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github release asset download`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "services github release asset download",
- 	CliArgs:         "-path /LOCAL/PATH/TO/DOWNLOAD",
+ 	CliArgs:         "-owner OWNER -repository REPO -path /LOCAL/PATH/TO/DOWNLOAD -release RELEASE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github release asset list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "services github release asset list",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO -release RELEASE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github release asset upload`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental, and Irreversible operation)",
  	Path:            "services github release asset upload",
- 	CliArgs:         "-asset /LOCAL/PATH/TO/assets",
+ 	CliArgs:         "-owner OWNER -repository REPO -release RELEASE -asset /LOCAL/PATH/TO/assets",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github release draft`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "(Experimental, and Irreversible operation)",
  	Path:    "services github release draft",
  	CliArgs: strings.Join({
  		"-",
- 		"body-file /LOCAL/PATH/TO/body.txt",
+ 		"owner OWNER -repository REPO -body-file /LOCAL/PATH/TO/BODY.txt ",
+ 		"-branch BRANCH -name NAME -tag TAG",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github release list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental)",
  	Path:            "services github release list",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services github tag create`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental, and Irreversible operation)",
  	Path:            "services github tag create",
- 	CliArgs:         "",
+ 	CliArgs:         "-owner OWNER -repository REPO -sha1 SHA -tag TAG",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail filter delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail filter delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-id ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail label add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail label add",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail label delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail label delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail label rename`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail label rename",
- 	CliArgs:         "",
+ 	CliArgs:         "-current-name CURRENT_NAME -new-name NEW_NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail message label add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail message label add",
- 	CliArgs:         "",
+ 	CliArgs:         "-label LABEL -message-id MSG_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail message label delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail message label delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-label LABEL -message-id MSG_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail message send`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "services google mail message send",
- 	CliArgs:         "",
+ 	CliArgs:         "-body /LOCAL/PATH/TO/INPUT.txt -subject SUBJECT -to TO",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail sendas add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail sendas add",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google mail sendas delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google mail sendas delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google sheets sheet append`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet append",
- 	CliArgs:         "",
+ 	CliArgs:         "-data /LOCAL/PATH/TO/INPUT.csv -id GOOGLE_SHEET_ID -range RANGE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google sheets sheet clear`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet clear",
- 	CliArgs:         "",
+ 	CliArgs:         "-id GOOGLE_SPREADSHEET_ID -range RANGE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google sheets sheet export`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet export",
- 	CliArgs:         "",
+ 	CliArgs:         "-id GOOGLE_SPREADSHEET_ID -range RANGE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google sheets sheet import`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet import",
- 	CliArgs:         "",
+ 	CliArgs:         "-data /LOCAL/PATH/TO/INPUT.csv -id VALUE -range VALUE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google sheets sheet list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets sheet list",
- 	CliArgs:         "",
+ 	CliArgs:         "-id GOOGLE_SPREADSHEET_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `services google sheets spreadsheet create`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "services google sheets spreadsheet create",
- 	CliArgs:         "",
+ 	CliArgs:         "-title TITLE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `sharedfolder leave`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedfolder leave",
- 	CliArgs:         "",
+ 	CliArgs:         "-shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `sharedfolder member add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "sharedfolder member add",
  	CliArgs: strings.Join({
  		"-",
- 		"path /SHARED_FOLDER",
+ 		"email EMAIL -path /DROPBOX",
  		"/PATH/TO/ADD",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `sharedfolder member delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "sharedfolder member delete",
  	CliArgs: strings.Join({
  		"-",
- 		"path /SHARED_FOLDER",
+ 		"email EMAIL -path /DROPBOX",
  		"/PATH/TO/DELETE",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `sharedfolder mount add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedfolder mount add",
- 	CliArgs:         "",
+ 	CliArgs:         "-shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `sharedfolder mount delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "sharedfolder mount delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: true,
  	... // 15 identical fields
  }
```

# Command spec changed: `team activity daily event`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team activity daily event",
- 	CliArgs:         "",
+ 	CliArgs:         "-start-date DATE",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team admin group role add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin group role add",
- 	CliArgs:         "",
+ 	CliArgs:         "-group GROUP_NAME -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team admin group role delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin group role delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-exception-group GROUP_NAME -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team admin role add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin role add",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team admin role clear`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin role clear",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team admin role delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team admin role delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-email EMAIL -role-id ROLE_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team runas sharedfolder list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder list",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team runas sharedfolder mount add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount add",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL -shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team runas sharedfolder mount delete`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount delete",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL -shared-folder-id SHARED_FOLDER_ID",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team runas sharedfolder mount list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount list",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team runas sharedfolder mount mountable`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "team runas sharedfolder mount mountable",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `team sharedlink delete member`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "team sharedlink delete member",
- 	CliArgs:         "",
+ 	CliArgs:         "-member-email EMAIL",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `teamfolder add`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Irreversible operation)",
  	Path:            "teamfolder add",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `teamfolder archive`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 16 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "Name", Desc: "Team folder name", TypeName: "string"},
  		&{
  			... // 2 identical fields
  			Default:  "default",
  			TypeName: "domain.dropbox.api.dbx_conn_impl.conn_scoped_team",
  			TypeAttr: []any{
- 				string("team_data.team_space"),
+ 				string("team_data.content.read"),
+ 				string("team_data.content.write"),
  			},
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
  	... // 2 identical fields
  }
```

# Command spec changed: `teamfolder file lock all release`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "teamfolder file lock all release",
  	CliArgs: strings.Join({
  		"-path /DROPBOX/PATH/TO/RELEASE",
- 		"/LOCK",
+ 		" -team-folder NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `teamfolder file lock list`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "teamfolder file lock list",
- 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST",
+ 	CliArgs:         "-path /DROPBOX/PATH/TO/LIST -team-folder NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `teamfolder file lock release`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "teamfolder file lock release",
  	CliArgs: strings.Join({
  		"-path /DROPBOX/PATH/TO/RELEASE",
- 		"/LOCK",
+ 		" -team-folder NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `teamfolder replication`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "(Experimental, and Irreversible operation)",
  	Path:            "teamfolder replication",
- 	CliArgs:         "",
+ 	CliArgs:         "-name NAME",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `util decode base32`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util decode base32",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "No padding", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "Text",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "Text to decode"}},
  	JsonInput:      {},
  }
```

# Command spec changed: `util decode base64`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util decode base64",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "No padding", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "Text",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "Text to decode"}},
  	JsonInput:      {},
  }
```

# Command spec changed: `util encode base32`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util encode base32",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "No padding", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "Text",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "Text to encode"}},
  	JsonInput:      {},
  }
```

# Command spec changed: `util encode base64`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks:         "",
  	Path:            "util encode base64",
- 	CliArgs:         "",
+ 	CliArgs:         "-text /LOCAL/PATH/TO/INPUT.txt",
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		&{Name: "NoPadding", Desc: "No padding", Default: "false", TypeName: "bool", ...},
  		&{
  			Name:     "Text",
  			Desc:     "Text",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "Text to decode"}},
  	JsonInput:      {},
  }
```

# Command spec changed: `util qrcode create`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util qrcode create",
  	CliArgs: strings.Join({
  		"-out /LOCAL/PATH/TO/",
- 		"create_qrcode.png",
+ 		"OUT.png -text /LOCAL/PATH/TO/INPUT.txt",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 8 identical fields
  	Reports: nil,
  	Feeds:   nil,
  	Values: []*dc_recipe.Value{
  		... // 2 identical elements
  		&{Name: "Out", Desc: "Output path with file name", TypeName: "essentials.model.mo_path.file_system_path_impl", TypeAttr: map[string]any{"shouldExist": bool(false)}},
  		&{Name: "Size", Desc: "Image resolution (pixel)", Default: "256", TypeName: "essentials.model.mo_int.range_int", ...},
  		&{
  			Name:     "Text",
  			Desc:     "Text data",
  			Default:  "",
- 			TypeName: "string",
+ 			TypeName: "Text",
  			TypeAttr: nil,
  		},
  	},
  	GridDataInput:  {},
  	GridDataOutput: {},
- 	TextInput:      []*dc_recipe.DocTextInput{},
+ 	TextInput:      []*dc_recipe.DocTextInput{&{Name: "Text", Desc: "Text"}},
  	JsonInput:      {},
  }
```

# Command spec changed: `util qrcode wifi`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util qrcode wifi",
  	CliArgs: strings.Join({
  		"-out /LOCAL/PATH/TO/",
- 		"create_qrcode.png",
+ 		"OUT.png -ssid SSID",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `util xlsx create`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util xlsx create",
  	CliArgs: strings.Join({
  		"-file /LOCAL/PATH/TO/",
- 		"create.xlsx",
+ 		"CREATE.xlsx -sheet SHEET_NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `util xlsx sheet export`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util xlsx sheet export",
  	CliArgs: strings.Join({
  		"-file /LOCAL/PATH/TO/",
- 		"export.xlsx",
+ 		"EXPORT.xlsx -sheet SHEET_NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```

# Command spec changed: `util xlsx sheet import`

## Command configuration changed

```
  &dc_recipe.Recipe{
  	... // 3 identical fields
  	Remarks: "",
  	Path:    "util xlsx sheet import",
  	CliArgs: strings.Join({
  		"-",
- 		"file /LOCAL/PATH/TO/import.xlsx",
+ 		"data /LOCAL/PATH/TO/INPUT.csv -file /LOCAL/PATH/TO/TARGET.xlsx -",
+ 		"sheet SHEET_NAME",
  	}, ""),
  	CliNote:         "",
  	ConnUsePersonal: false,
  	... // 15 identical fields
  }
```
