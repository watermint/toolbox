---
layout: command
title: コマンド `dropbox team sharedlink update expiry`
lang: ja
---

# dropbox team sharedlink update expiry

チーム内の公開されている共有リンクについて有効期限を更新します (非可逆な操作です)

注：リリース87以降、このコマンドは、アップデートする共有リンクを選択するためのファイルを受け取ります. チーム内のすべての共有リンクの有効期限を更新したい場合は、`dropbox team sharedlink list`の組み合わせをご検討ください。例えば、[jq](https://stedolan.github.io/jq/)というコマンドに慣れていれば、以下のように同等の操作を行うことができます（すべての公開リンクに対して28日以内に強制失効させる）.

```
tbx team sharedlink list -output json -visibility public | jq '.sharedlink.url' | tbx team sharedlink update expiry -file - -at +720h
```
リリース92以降、このコマンドは引数 `-days` を受け取りません. 相対的な日時を設定したい場合は、`+720h`（720時間＝30日）のように`-at +HOURh`を使用してください.

コマンド `dropbox team sharedlink update` は共有リンクに値を設定するためのコマンドです。コマンド `dropbox team sharedlink cap` は共有リンクに上限値を設定するためのコマンドです。例えば、`dropbox team sharedlink update expiry`で有効期限を2021-05-06に設定した場合。このコマンドは、既存のリンクが2021-05-04のように短い有効期限を持っている場合でも、有効期限を2021-05-06に更新します.

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
| Dropbox for teams：チームメンバーの確認                                             |
| Dropbox for teams：Dropboxの共有設定と共同作業者の表示と管理                        |
| Dropbox for teams：チームやメンバーのフォルダ構造を表示                             |
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
.\tbx.exe dropbox team sharedlink update expiry -file /PATH/TO/DATA_FILE.csv -at +720h
```

macOS, Linux:
```
$HOME/Desktop/tbx dropbox team sharedlink update expiry -file /PATH/TO/DATA_FILE.csv -at +720h
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-at**
: 新しい有効期限の日時

**-file**
: データファイルへのパス

**-peer**
: アカウントの別名. Default: default

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

# ファイル書式

## 書式: File

対象となる共有リンク

| 列  | 説明            | 例                                       |
|-----|-----------------|------------------------------------------|
| url | 共有リンクのURL | https://www.dropbox.com/scl/fo/fir9vjelf |

最初の行はヘッダ行です. プログラムは、ヘッダのないファイルを受け入れます.
```
url
https://www.dropbox.com/scl/fo/fir9vjelf
```

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: operation_log

このレポートは処理結果を出力します.
このコマンドはレポートを3種類の書式で出力します. `operation_log.csv`, `operation_log.json`, ならびに `operation_log.xlsx`.

| 列                | 説明                                   |
|-------------------|----------------------------------------|
| status            | 処理の状態                             |
| reason            | 失敗またはスキップの理由               |
| input.url         | 共有リンクのURL                        |
| result.tag        | エントリーの種別 (file, または folder) |
| result.url        | 共有リンクのURL.                       |
| result.name       | リンク先ファイル名称                   |
| result.expires    | 有効期限 (設定されている場合)          |
| result.path_lower | パス (すべて小文字に変換).             |
| result.visibility | 共有リンクの開示範囲                   |
| result.email      | ユーザーのメールアドレス               |
| result.surname    | リンク所有者の名字                     |
| result.given_name | リンク所有者の名                       |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`, ...


