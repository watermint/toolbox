# file sync up

Dropboxと上り方向で同期します (非可逆な操作です)

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

| ラベル    | 説明                     |
|-----------|--------------------------|
| user_full | Dropbox へのフルアクセス |

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します. リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.
```

watermint toolbox xx.x.xxx
==========================

© 2016-2020 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

1. 次のURLを開き認証ダイアログを開いてください:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. 'Allow'をクリックします (先にログインしておく必要があります):
3. 認証コードをコピーします:
認証コードを入力してください
```

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.
## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe file sync up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT
```

macOS, Linux:
```
$HOME/Desktop/tbx file sync up -dropbox-path /DROPBOX/PATH/TO/UPLOAD -local-path /LOCAL/PATH/OF/CONTENT
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション             | 説明                                                | デフォルト |
|------------------------|-----------------------------------------------------|------------|
| `-chunk-size-kb`       | アップロードチャンク容量(Kバイト)                   | 65536      |
| `-delete`              | Delete Dropbox file if a file removed locally       | false      |
| `-dropbox-path`        | 転送先のDropboxパス                                 |            |
| `-local-path`          | ローカルファイルのパス                              |            |
| `-name-disable-ignore` | Filter by name. Filter system file or ignore files. |            |
| `-name-name`           | Filter by name. Filter by exact match to the name.  |            |
| `-name-name-prefix`    | Filter by name. Filter by name match to the prefix. |            |
| `-name-name-suffix`    | Filter by name. Filter by name match to the suffix. |            |
| `-peer`                | アカウントの別名                                    | default    |
| `-skip-existing`       | Skip existing files. Do not overwrite               | false      |

## 共通のオプション:

| オプション        | 説明                                                                                               | デフォルト     |
|-------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`      | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`   | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`  | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                        | normal         |
| `-budget-storage` | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                            | normal         |
| `-concurrency`    | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`          | デバッグモードを有効にする                                                                         | false          |
| `-experiment`     | 実験的機能を有効化する                                                                             |                |
| `-lang`           | 表示言語                                                                                           | auto           |
| `-output`         | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`          | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | パスのパターン                              | 例                                                     |
|---------|---------------------------------------------|--------------------------------------------------------|
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` | C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /Users/bob/.toolbox/jobs/20190909-115959.597/reports   |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports`      | /home/bob/.toolbox/jobs/20190909-115959.597/reports    |

## レポート: deleted

Path
このコマンドはレポートを3種類の書式で出力します. `deleted.csv`, `deleted.json`, ならびに `deleted.xlsx`.

| 列         | 説明 |
|------------|------|
| entry_path | Path |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `deleted_0000.xlsx`, `deleted_0001.xlsx`, `deleted_0002.xlsx`, ...
## レポート: skipped

このレポートは処理結果を出力します.
このコマンドはレポートを3種類の書式で出力します. `skipped.csv`, `skipped.json`, ならびに `skipped.xlsx`.

| 列               | 説明                     |
|------------------|--------------------------|
| status           | 処理の状態               |
| reason           | 失敗またはスキップの理由 |
| input.entry_path | Path                     |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`, ...
## レポート: summary

このレポートはアップロード結果の概要を出力します.
このコマンドはレポートを3種類の書式で出力します. `summary.csv`, `summary.json`, ならびに `summary.xlsx`.

| 列                 | 説明                                             |
|--------------------|--------------------------------------------------|
| upload_start       | アップロード開始日時                             |
| upload_end         | アップロード終了日時                             |
| num_bytes          | 合計アップロードサイズ (バイト)                  |
| num_files_error    | 失敗またはエラーが発生したファイル数.            |
| num_files_upload   | アップロード済みまたはアップロード対象ファイル数 |
| num_files_skip     | スキップ対象またはスキップ予定のファイル数       |
| num_folder_created | Number of created folders.                       |
| num_delete         | Number of deleted entry.                         |
| num_api_call       | この処理によって消費される見積アップロードAPI数  |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `summary_0000.xlsx`, `summary_0001.xlsx`, `summary_0002.xlsx`, ...
## レポート: uploaded

このレポートは処理結果を出力します.
このコマンドはレポートを3種類の書式で出力します. `uploaded.csv`, `uploaded.json`, ならびに `uploaded.xlsx`.

| 列                     | 説明                                                         |
|------------------------|--------------------------------------------------------------|
| status                 | 処理の状態                                                   |
| reason                 | 失敗またはスキップの理由                                     |
| input.path             | Path                                                         |
| result.name            | 名称                                                         |
| result.path_display    | パス (表示目的で大文字小文字を区別する).                     |
| result.client_modified | ファイルの場合、更新日時はクライアントPC上でのタイムスタンプ |
| result.server_modified | Dropbox上で最後に更新された日時                              |
| result.size            | ファイルサイズ(バイト単位)                                   |
| result.content_hash    | ファイルコンテンツのハッシュ                                 |

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます; `uploaded_0000.xlsx`, `uploaded_0001.xlsx`, `uploaded_0002.xlsx`, ...

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

