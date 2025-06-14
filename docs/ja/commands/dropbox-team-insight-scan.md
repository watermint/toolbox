---
layout: command
title: コマンド `dropbox team insight scan`
lang: ja
---

# dropbox team insight scan

チームデータをスキャンして分析 

このコマンドは、チームフォルダ内のファイル、パーミッション、共有リンクなど、さまざまなチームデータを収集し、データベースに保存します。
収集したデータは、`dropbox team insight report teamfoldermember`のようなコマンドや、SQLite全般をサポートするデータベースツールで分析できます。

スキャンにかかる時間について:

チームのスキャンには長い時間がかかります。特に保存されているファイル数が多い場合、時間はファイル数に直線的に比例します。スキャン速度を上げるには、`-concurrency`オプションを使用して並列処理を行うのがよいでしょう。
ただし、並列性が高すぎるとDropboxサーバーからのエラー率が高くなるので、バランスを考慮する必要があります。いくつかのベンチマークの結果によると、`concurrency`オプションの並列度は12～24が良いようです。
スキャンに要する時間はDropboxサーバーのレスポンスに依存しますが、1000万ファイルあたり20～30時間程度でした（`-concurrency 16`の場合）。

スキャン中、ユーザーはファイルを削除、移動、追加するかもしれません。コマンドは、これらの違いをすべて把握して正確な結果を報告することを目的としているわけではなく、大まかな情報をできるだけ早く提供することを目的としています。

データベースのファイルサイズについて:

このコマンドはチームのファイルを含むすべてのメタデータを取得するため、データベースのサイズはそれらのメタデータのサイズに応じて増加します。ベンチマーク結果によると、チーム内に保存されている1000万ファイルあたりのデータベース・サイズは約10～12GBでした。実行する前に `-database` で指定したパスに十分な空き容量があることを確認してください。

スキャンエラーについて:

スキャンの実行時に Dropbox サーバーがエラーを返す場合があります。コマンドは自動的に何度かスキャンの再実行を試みますが、サーバーの混雑や状態により、一定時間エラーが解決されない場合があります。この場合、コマンドは再実行を停止し、エラーが発生したスキャンタスクをデータベースに記録します。
失敗したスキャンを再実行したい場合は、`dropbox team insight scanretry` コマンドを使用してスキャンを再実行してください。
再実行を繰り返しても問題が解決せず、現在のスキャンのカバレッジだけを分析したい場合は、分析の前に集計タスクを実行する必要があります。集計作業は `dropbox team insight summary` コマンドで行うことができます。

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

| 説明 |
|------|

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
.\tbx.exe dropbox team insight scan -database /LOCAL/PATH/TO/database
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team insight scan -database /LOCAL/PATH/TO/database
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-base-path**
: ファイルパス標準を選択します。これは、特にDropbox for Teams向けのオプションです。Dropboxの個人版を使用している場合は、選択したものが何であってもほとんど問題ありません。Dropbox for Teamsでは、更新されたチームスペースで`home`を選択すると、ユーザー名付きの個人フォルダが選択されます。これは、個人フォルダ内のファイルを参照したりアップロードしたりする場合に便利です。なぜなら、パスにユーザー名付きのフォルダ名を指定する必要がないからです。一方、`root`を指定すると、アクセス権のあるすべてのフォルダにアクセスできます。ただし、個人フォルダにアクセスする場合には、個人フォルダの名前を含むパスを指定する必要があります。. Options: root (Full access to all folders with permissions), home (Access limited to personal home folder). Default: root

**-database**
: データベースへのパス

**-max-retries**
: 最大リトライ回数. Default: 3

**-peer**
: アカウントの別名. Default: default

**-scan-member-folders**
: メンバーフォルダのスキャン. Default: false

**-skip-summarize**
: 集計タスクをスキップします. Default: false

## 共通のオプション:

**-auth-database**
: 認証データベースへのカスタムパス (デフォルト: $HOME/.toolbox/secrets/secrets.db)

**-auto-open**
: 成果物フォルダまたはURLを自動で開く. Default: false

**-bandwidth-kb**
: コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない. Default: 0

**-budget-memory**
: メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます). Options: low, normal. Default: normal

**-budget-storage**
: ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します). Options: low, normal, unlimited. Default: normal

**-concurrency**
: 指定した並列度で並列処理を行います. Default: プロセッサー数

**-debug**
: デバッグモードを有効にする. Default: false

**-experiment**
: 実験的機能を有効化する

**-extra**
: 追加パラメータファイルのパス

**-lang**
: 表示言語. Options: auto, en, ja. Default: auto

**-output**
: 出力書式 (none/text/markdown/json). Options: text, markdown, json, none. Default: text

**-output-filter**
: 出力フィルタ・クエリ（jq構文）。レポートの出力はjq構文を使ってフィルタリングされる。このオプションは、レポートがJSONとして出力される場合にのみ適用される。

**-proxy**
: HTTP/HTTPS プロクシ (hostname:port). プロキシの設定を省略したい場合は`DIRECT`を指定してください

**-quiet**
: エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します. Default: false

**-retain-job-data**
: ジョブデータ保持ポリシー. Options: default, on_error, none. Default: default

**-secure**
: トークンをファイルに保存しません. Default: false

**-skip-logging**
: ローカルストレージへのログ保存をスキップ. Default: false

**-verbose**
: 現在の操作を詳細に表示します.. Default: false

**-workspace**
: ワークスペースへのパス

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: errors

エラーレポート
このコマンドはレポートを3種類の書式で出力します. `errors.csv`, `errors.json`, ならびに `errors.xlsx`.

| 列       | 説明             |
|----------|------------------|
| category | エラーカテゴリー |
| message  | エラーメッセージ |
| tag      | エラータグ       |
| detail   | エラーの詳細     |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `errors_0000.xlsx`, `errors_0001.xlsx`, `errors_0002.xlsx`, ...


