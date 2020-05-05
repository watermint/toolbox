# team sharedlink update expiry 

チーム内の公開されている共有リンクについて有効期限を更新します (非可逆な操作です)

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
* Dropbox Business File access

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe team sharedlink update expiry -days 28
```

macOS, Linux:

```bash
$HOME/Desktop/tbx team sharedlink update expiry -days 28
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション

| オプション    | 説明                       | デフォルト |
|---------------|----------------------------|------------|
| `-at`         | 新しい有効期限の日時       |            |
| `-days`       | 新しい有効期限までの日時   | 0          |
| `-peer`       | アカウントの別名           | default    |
| `-visibility` | 対象となるリンクの公開範囲 | public     |

共通のオプション:

| オプション        | 説明                                                                                               | デフォルト     |
|-------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`      | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb`   | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-budget-memory`  | Memory budget (limits some feature to reduce memory footprint)                                     | normal         |
| `-budget-storage` | Storage budget (limits logs or some feature to reduce storage usage)                               | normal         |
| `-concurrency`    | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`          | デバッグモードを有効にする                                                                         | false          |
| `-output`         | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`          | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`          | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`         | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`      | ワークスペースへのパス                                                                             |                |

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

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

# 実行結果

作成されたレポートファイルのパスはコマンド実行時の最後に表示されます. もしコマンドライン出力を失ってしまった場合には次のパスを確認してください. [job-id]は実行の日時となります. このなかの最新のjob-idを各委任してください.

| OS      | Path                                                                                                      |
| ------- | --------------------------------------------------------------------------------------------------------- |
| Windows | `%HOMEPATH%\.toolbox\jobs\[job-id]\reports` (e.g. C:\Users\bob\.toolbox\jobs\20190909-115959.597\reports) |
| macOS   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /Users/bob/.toolbox/jobs/20190909-115959.597/reports)        |
| Linux   | `$HOME/.toolbox/jobs/[job-id]/reports` (e.g. /home/bob/.toolbox/jobs/20190909-115959.597/reports)         |

## レポート: skipped 
このレポートはチーム内のチームメンバーがもつ共有リンク一覧を出力します.
レポートファイルは次の3種類のフォーマットで出力されます;
* `skipped.csv`
* `skipped.xlsx`
* `skipped.json`

`-low-memory`オプションを指定した場合には、コマンドはJSONフォーマットのレポートのみを出力します.

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます;
`skipped_0000.xlsx`, `skipped_0001.xlsx`, `skipped_0002.xlsx`...   

| 列         | 説明                                                                 |
|------------|----------------------------------------------------------------------|
| tag        | エントリーの種別 (file, または folder)                               |
| url        | 共有リンクのURL.                                                     |
| name       | リンク先ファイル名称                                                 |
| expires    | 有効期限 (設定されている場合)                                        |
| path_lower | パス (すべて小文字に変換).                                           |
| visibility | 共有リンクの開示範囲                                                 |
| email      | ユーザーのメールアドレス                                             |
| status     | チームにおけるメンバーのステータス(active/invited/suspended/removed) |
| surname    | リンク所有者の名字                                                   |
| given_name | リンク所有者の名                                                     |

## レポート: updated 
このレポートは処理結果を出力します.
レポートファイルは次の3種類のフォーマットで出力されます;
* `updated.csv`
* `updated.xlsx`
* `updated.json`

`-low-memory`オプションを指定した場合には、コマンドはJSONフォーマットのレポートのみを出力します.

レポートが大きなものとなる場合、`.xlsx`フォーマットのファイルは次のようにいくつかに分割されて出力されます;
`updated_0000.xlsx`, `updated_0001.xlsx`, `updated_0002.xlsx`...   

| 列               | 説明                                                                 |
|------------------|----------------------------------------------------------------------|
| status           | 処理の状態                                                           |
| reason           | 失敗またはスキップの理由                                             |
| input.tag        | エントリーの種別 (file, または folder)                               |
| input.url        | 共有リンクのURL.                                                     |
| input.name       | リンク先ファイル名称                                                 |
| input.expires    | 有効期限 (設定されている場合)                                        |
| input.path_lower | パス (すべて小文字に変換).                                           |
| input.visibility | 共有リンクの開示範囲                                                 |
| input.email      | ユーザーのメールアドレス                                             |
| input.status     | チームにおけるメンバーのステータス(active/invited/suspended/removed) |
| input.surname    | リンク所有者の名字                                                   |
| input.given_name | リンク所有者の名                                                     |
| result.expires   | 有効期限 (設定されている場合)                                        |

