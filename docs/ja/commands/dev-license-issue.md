---
layout: command
title: コマンド `dev license issue`
lang: ja
---

# dev license issue

ライセンスの発行 

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
* GitHub: https://developer.github.com/apps/managing-oauth-apps/deleting-an-oauth-app/

## 認可スコープ

| 説明                                                                                                                                                                                                                                                                                                                                               |
|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| GitHub: プライベートリポジトリを含む、リポジトリへのフルアクセスを許可それには、コードへの読み書き可能なアクセス、コミットステータス、リポジトリや組織のプロジェクト、招待状、共同作業者、チームメンバーの追加、デプロイメントステータス、リポジトリや組織のWebhookなどが含まれます. また、ユーザーのプロジェクトを管理する機能も付与されています. |

# 認可

最初の実行では、`tbx`はあなたのGitHubアカウントへの認可を要求します.
Enterキーを押すと、ブラウザが起動します。その後、サービスが認証を行い、tbxがその結果を受け取ります。認証成功のメッセージが表示されたら、ブラウザのウィンドウを閉じてもかまいません。
```

watermint toolbox xx.x.xxx
==========================

© 2016-2025 Takayuki Okazaki
オープンソースライセンスのもと配布されています. 詳細は`license`コマンドでご覧ください.

認可URLを開きます:
https://github.com/login/oauth/authorize?client_id=xxxxxxxxxxxxxxxxxxxx&redirect_uri=http%3A%2F%2Flocalhost%3A7800%2Fconnect%2Fauth&response_type=code&scope=repo&state=xxxxxxxx

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
.\tbx.exe dev license issue -licensee-email LICENSEE_EMAIL -licensee-name LICENSEE_NAME
```

macOS, Linux:
```
$HOME/Desktop/tbx dev license issue -licensee-email LICENSEE_EMAIL -licensee-name LICENSEE_NAME
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-app-name**
: アプリケーション名. Default: watermint toolbox

**-branch**
: ライセンス・リポジトリ・ブランチ. Default: main

**-expiration**
: ライセンスの有効期限

**-licensee-email**
: ライセンシーのEメール

**-licensee-name**
: ライセンシー名

**-lifecycle-available-after**
: ビルド時間からこの期間後に利用可能なライフサイクル（秒）. Default: 94608000

**-lifecycle-warning-after**
: ビルド時刻からこの期間経過後のライフサイクル警告（秒）. Default: 31536000

**-owner**
: ライセンス・リポジトリの所有者. Default: watermint

**-peer**
: アカウントの別名. Default: default

**-recipe-allowed-prefix**
: レシピの接頭辞

**-recipes-allowed**
: コンマで区切られたレシピのリスト

**-repository**
: ライセンス・リポジトリ. Default: toolbox-supplement

**-scope**
: ライセンス範囲

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


