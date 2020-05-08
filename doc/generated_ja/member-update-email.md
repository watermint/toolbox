# member update email 

メンバーのメールアドレス処理 (非可逆な操作です)

# セキュリティ

`watermint toolbox`は認証情報をファイルシステム上に保存します. それは次のパスです:

| OS       | Path                                                               |
| -------- | ------------------------------------------------------------------ |
| Windows  | `%HOMEPATH%\.toolbox\secrets` (e.g. C:\Users\bob\.toolbox\secrets) |
| macOS    | `$HOME/.toolbox/secrets` (e.g. /Users/bob/.toolbox/secrets)        |
| Linux    | `$HOME/.toolbox/secrets` (e.g. /home/bob/.toolbox/secrets)         |

これらの認証情報ファイルはDropboxサポートを含め誰にも共有しないでください.
不必要になった場合にはこれらのファイルを削除しても問題ありません. 認証情報の削除を確実にしたい場合には、アプリケーションアクセス設定または管理コンソールからアプリケーションへの許可を取り消してください.

方法は次のヘルプセンター記事をご参照ください:
* Dropbox Business: https://help.dropbox.com/teams-admins/admin/app-integrations

このコマンドは次のアクセスタイプを処理に利用します:

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe member update email -file /path/to/data/file.csv
```

macOS, Linux:

```bash
$HOME/Desktop/tbx member update email -file /path/to/data/file.csv
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション

| オプション           | 説明                                                                                                                                                  | デフォルト |
|----------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------|------------|
| `-file`              | データファイル                                                                                                                                        |            |
| `-peer`              | アカウントの別名                                                                                                                                      | default    |
| `-update-unverified` | アカウントのメールアドレスが確認されていない場合でも変更する. アカウントのメールアドレスが確認されていない場合、フォルダへの招待を失う場合があります. | false      |

共通のオプション:

| オプション        | 説明                                                                                               | デフォルト     |
|-------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`      | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`   | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`  | メモリの割り当て目標 (メモリ使用量を減らすために幾つかの機能が制限されます)                        | normal         |
| `-budget-storage` | ストレージの利用目標 (ストレージ利用を減らすためログ、機能を限定します)                            | normal         |
| `-concurrency`    | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`          | デバッグモードを有効にする                                                                         | false          |
| `-output`         | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`          | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

# ファイル書式

## 書式: File

チームメンバーのメールアドレスを変更処理するためのデータファイル. 

| 列         | 説明                 | 値の例                 |
|------------|----------------------|------------------------|
| from_email | 現在のメールアドレス | john@example.com       |
| to_email   | 新しいメールアドレス | john.smith@example.net |

最初の行はヘッダ行です. プログラムはヘッダ行がない場合も認識します.

```csv
from_email,to_email
john@example.com,john.smith@example.net
```

# 認可

最初の実行では、`tbx`はあなたのDropboxアカウントへの認可を要求します. リンクをブラウザにペーストしてください. その後、認可を行います. 認可されると、Dropboxは認証コードを表示します. `tbx`にこの認証コードをペーストしてください.

```

[1mwatermint toolbox xx.x.xxx[0m
[1m==========================[0m

[37m© 2016-2020 Takayuki Okazaki[0m
[37mオープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.[0m

[37m1. 次のURLを開き認証ダイアログを開いてください:

https://www.dropbox.com/oauth2/authorize?client_id=xxxxxxxxxxxxxxx&response_type=code&state=xxxxxxxx

2. 'Allow'をクリックします (先にログインしておく必要があります):
3. 認証コードをコピーします:[0m
[37m認証コードを入力してください[0m

```

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## レポート: operation_log 
このレポートは処理結果を出力します.
レポートファイルは次の3種類のフォーマットで出力されます;
* `operation_log.csv`
* `operation_log.xlsx`
* `operation_log.json`

`-budget-memory low`オプションを指定した場合、レポートはJSON形式のみで生成されます

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます;
`operation_log_0000.xlsx`, `operation_log_0001.xlsx`, `operation_log_0002.xlsx`...   

| 列                    | 説明                                                                                            |
|-----------------------|-------------------------------------------------------------------------------------------------|
| status                | 処理の状態                                                                                      |
| reason                | 失敗またはスキップの理由                                                                        |
| input.from_email      | 現在のメールアドレス                                                                            |
| input.to_email        | 新しいメールアドレス                                                                            |
| result.email          | ユーザーのメールアドレス                                                                        |
| result.email_verified | trueの場合、ユーザーのメールアドレスはユーザーによって所有されていることが確認されています.     |
| result.status         | チームにおけるメンバーのステータス(active/invited/suspended/removed)                            |
| result.given_name     | 名                                                                                              |
| result.surname        | 名字                                                                                            |
| result.display_name   | ユーザーのDropboxアカウントの表示名称                                                           |
| result.joined_on      | メンバーがチームに参加した日時.                                                                 |
| result.role           | ユーザーのチームでの役割 (team_admin, user_management_admin, support_admin, または member_only) |
| result.tag            | 処理のタグ                                                                                      |

