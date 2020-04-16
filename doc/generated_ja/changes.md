# `リリース 65` から `リリース 66` までの変更点

# 追加されたコマンド

| コマンド                               | タイトル                                   |
|----------------------------------------|--------------------------------------------|
| dev catalogue                          | Generate catalogue                         |
| services github release asset download | Download assets                            |
| services github release asset upload   | Upload assets file into the GitHub Release |
| team filerequest clone                 | Clone file requests by given data          |



# 削除されたコマンド

| コマンド                         | タイトル                                   |
|----------------------------------|--------------------------------------------|
| services github release asset up | Upload assets file into the GitHub Release |
| web                              | Launch web console                         |



# コマンド仕様の変更: `services github release asset list`



## 変更されたレポート: assets

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

