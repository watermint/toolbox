# Changes between `Release 65` to `Release 66`

# Commands added

| Command                                | Title                                      |
|----------------------------------------|--------------------------------------------|
| dev catalogue                          | Generate catalogue                         |
| services github release asset download | Download assets                            |
| services github release asset upload   | Upload assets file into the GitHub Release |
| team filerequest clone                 | Clone file requests by given data          |



# Commands deleted

| Command                          | Title                                      |
|----------------------------------|--------------------------------------------|
| services github release asset up | Upload assets file into the GitHub Release |
| web                              | Launch web console                         |



# Command spec changed: `services github release asset list`



## Changed report: assets

```
  &rc_doc.Report{
  	Name: "assets",
  	Desc: "GitHub Release assets",
  	Columns: []*rc_doc.ReportColumn{
  		... // 2 identical elements
  		&{Name: "state", Desc: "State of the asset"},
  		&{Name: "download_count", Desc: "Number of downloads"},
+ 		&{Name: "download_url", Desc: "Download URL"},
  	},
  }

```

