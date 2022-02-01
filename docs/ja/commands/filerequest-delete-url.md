---
layout: command
title: コマンド
lang: ja
---

# filerequest delete url

ファイルリクエストのURLを指定して削除 (非可逆な操作です)

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
* Dropbox (個人アカウント): https://help.dropbox.com/installs-integrations/third-party/third-party-apps

## 認可スコープ

| 説明                                             |
|--------------------------------------------------|
| Dropbox: Dropboxのファイルリクエストを参照       |
| Dropbox: Dropboxのファイルリクエストの表示と管理 |

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します. リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2022 Takayuki Okazaki
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
.\tbx.exe filerequest delete url 
```

macOS, Linux:
```
$HOME/Desktop/tbx filerequest delete url 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション | 説明                                | デフォルト |
|------------|-------------------------------------|------------|
| `-force`   | ファイリクエストを強制的に削除する. | false      |
| `-peer`    | アカウントの別名                    | default    |
| `-url`     | ファイルリクエストのURL             |            |

## 共通のオプション:

| オプション         | 説明                                                                                               | デフォルト     |
|--------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`       | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`    | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`   | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                        | normal         |
| `-budget-storage`  | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                            | normal         |
| `-concurrency`     | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`           | デバッグモードを有効にする                                                                         | false          |
| `-experiment`      | 実験的機能を有効化する                                                                             |                |
| `-extra`           | 追加パラメータファイルのパス                                                                       |                |
| `-lang`            | 表示言語                                                                                           | auto           |
| `-output`          | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`           | HTTP/HTTPS プロクシ (hostname:port). プロキシの設定を省略したい場合は`DIRECT`を指定してください    |                |
| `-quiet`           | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-retain-job-data` | Job data retain policy                                                                             | retain         |
| `-secure`          | トークンをファイルに保存しません                                                                   | false          |
| `-verbose`         | 現在の操作を詳細に表示します.                                                                      | false          |
| `-workspace`       | ワークスペースへのパス                                                                             |                |

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: deleted

このレポートはファイルリクエストの一覧を出力します.
このコマンドはレポートを3種類の書式で出力します. `deleted.csv`, `deleted.json`, ならびに `deleted.xlsx`.

| 列                          | 説明                                                                    |
|-----------------------------|-------------------------------------------------------------------------|
| id                          | ファイルリクエストのID                                                  |
| url                         | ファイルリクエストのURL                                                 |
| title                       | ファイルリクエストのタイトル                                            |
| created                     | ファイルリクエストが作成された日時.                                     |
| is_open                     | このファイルリクエストがオープンしているかどうか                        |
| file_count                  | このファイルリクエストが受け取ったファイル数                            |
| destination                 | ファイルをアップロードするDropbox上のパス                               |
| deadline                    | ファイルリクエストの締め切り.                                           |
| deadline_allow_late_uploads | 設定した場合、ファイルリクエストの期限がきてもアップロードを許可します. |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.


