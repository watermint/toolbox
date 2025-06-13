---
layout: command
title: コマンド `dropbox team activity event`
lang: ja
---

# dropbox team activity event

詳細なチームアクティビティイベントログをフィルタリングオプション付きで取得、セキュリティ監査とコンプライアンス監視に必須 

リリース91以降では、`-start-time`または`-end-time`を`-24h`（24時間）または`-10m`（10分）のようなフォーマットで現在からの相対的な時間として解析します.
もし、1時間ごとにイベントを取得したい場合は、次のように実行します:

```
tbx team activity event -start-time -1h -output json > latest_events.json
```

必要であれば、最新の部分をログ全体に追加します.

```
cat latest_events.json >> all.json
```

より正確には、1時間ごとにすこし重複したイベントを取得します.
```
tbx team activity event -start-time -1h5m -output json > latest_events.json
```

そして、オーバーラップしたイベントを連結し、重複を排除します:
```
cat all.json latest_events.json | sort -u > _all.json && mv _all.json all.json
```

If you prefer CSV format, then use the `jq` command to convert it.
```
cat latest_events.json | jq -r '[.timestamp, .actor[.actor.".tag"].display_name, .actor[.actor.".tag"].email, .event_type.description, .event_category.".tag", .origin.access_method.end_user.".tag", .origin.geo_location.ip_address, .origin.geo_location.country, .origin.geo_location.city, .involve_non_team_member, (.participants | @text), (.context | @text)] | @csv' >> all.csv
```

# セキュリティ

`watermint toolbox`は認証情報をファイルシステム上に保存します. それは次のパスです:

| OS      | パス                                                               |
|---------|--------------------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS   | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux   | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

これらの認証情報ファイルはDropboxサポートを含め誰にも共有しないでください.
不必要になった場合にはこれらのファイルを削除しても問題ありません. 認証情報の削除を確実にしたい場合には、アプリケーションアクセス設定または管理コンソールからアプリケーションへの許可を取り消してください.

方法は次のヘルプセンター記事をご参照ください:
* Dropbox for teams: https://help.dropbox.com/installs-integrations/third-party/business-api#manage

## 認可スコープ

| 説明                                                                                |
|-------------------------------------------------------------------------------------|
| Dropbox for teams：チームのアクティビティログを表示                                 |
| Dropbox for teams：名前、ユーザー数、チーム設定など、チームの基本情報を表示します。 |

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します.
リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

1. 次のURLを開き認証ダイアログを開いてください:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. 'Allow'をクリックします (先にログインしておく必要があります):
3. 認証コードをコピーします:
認証コードを入力してください
```

# インストール

[最新リリース](https://github.com/watermint/toolbox/releases/latest)からコンパイル済みのバイナリをダウンロードしてください. Windowsをお使いの方は、`tbx-xx.x.xxx-win.zip`のようなzipファイルをダウンロードしてください. その後、アーカイブを解凍し、デスクトップ フォルダに `tbx.exe` を配置します.
watermint toolboxは、システムで許可されていれば、システム内のどのパスからでも実行できます. しかし、説明書のサンプルでは、デスクトップ フォルダを使用しています. デスクトップ フォルダ以外にバイナリを配置した場合は、パスを読み替えてください.

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe dropbox team activity event 
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team activity event 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション    | 説明                                                                                  | デフォルト |
|---------------|---------------------------------------------------------------------------------------|------------|
| `-category`   | 一つのイベントカテゴリのみを返すようなフィルター条件. このフィールドはオプションです. |            |
| `-end-time`   | 終了日時 (該当同時刻を含まない).                                                      |            |
| `-peer`       | アカウントの別名                                                                      | default    |
| `-start-time` | 開始日時 (該当時刻を含む)                                                             |            |

## 共通のオプション:

| オプション         | 説明                                                                                                                                                       | デフォルト     |
|--------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------|
| `-auth-database`   | 認証データベースへのカスタムパス (デフォルト: $HOME/.toolbox/secrets/secrets.db)                                                                           |                |
| `-auto-open`       | 成果物フォルダまたはURLを自動で開く                                                                                                                        | false          |
| `-bandwidth-kb`    | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない                                                         | 0              |
| `-budget-memory`   | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                                                                                | normal         |
| `-budget-storage`  | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                                                                                    | normal         |
| `-concurrency`     | 指定した並列度で並列処理を行います                                                                                                                         | プロセッサー数 |
| `-debug`           | デバッグモードを有効にする                                                                                                                                 | false          |
| `-experiment`      | 実験的機能を有効化する                                                                                                                                     |                |
| `-extra`           | 追加パラメータファイルのパス                                                                                                                               |                |
| `-lang`            | 表示言語                                                                                                                                                   | auto           |
| `-output`          | 出力書式 (none/text/markdown/json)                                                                                                                         | text           |
| `-output-filter`   | 出力フィルタ・クエリ（jq構文）。レポートの出力はjq構文を使ってフィルタリングされる。このオプションは、レポートがJSONとして出力される場合にのみ適用される。 |                |
| `-proxy`           | HTTP/HTTPS プロクシ (hostname:port). プロキシの設定を省略したい場合は`DIRECT`を指定してください                                                            |                |
| `-quiet`           | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                                                                        | false          |
| `-retain-job-data` | ジョブデータ保持ポリシー                                                                                                                                   | default        |
| `-secure`          | トークンをファイルに保存しません                                                                                                                           | false          |
| `-skip-logging`    | ローカルストレージへのログ保存をスキップ                                                                                                                   | false          |
| `-verbose`         | 現在の操作を詳細に表示します.                                                                                                                              | false          |
| `-workspace`       | ワークスペースへのパス                                                                                                                                     |                |

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: event

このレポートは、主にDropbox for teamsのアクティビティログと互換性のあるアクティビティログを表示します。
このコマンドはレポートを3種類の書式で出力します. `event.csv`, `event.json`, ならびに `event.xlsx`.

| 列                       | 説明                                                            |
|--------------------------|-----------------------------------------------------------------|
| timestamp                | このアクションが実行されたDropbox側でのタイムスタンプ.          |
| member                   | ユーザーの表示名                                                |
| member_email             | ユーザーのメールアドレス                                        |
| event_type               | 実行されたアクションのタイプ                                    |
| category                 | 監査ログイベントのカテゴリー                                    |
| access_method            | アクションが実行された方法.                                     |
| ip_address               | IPアドレス.                                                     |
| country                  | 国                                                              |
| city                     | 市町村                                                          |
| involve_non_team_members | 1名以上のチーム外のユーザーがこのアクションに関連した場合はTrue |
| participants             | このアクションによって影響を受けたユーザーまたはグループ        |
| context                  | アクターがアクションを実行したユーザーまたはチーム              |
| assets                   | アクションに関連したコンテンツ資産.                             |
| other_info               | このタイプのアクションに適用可能な可変イベントスキーマ.         |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `event_0000.xlsx`, `event_0001.xlsx`, `event_0002.xlsx`, ...


