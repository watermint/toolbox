---
layout: command
title: コマンド `log cat curl`
lang: ja
---

# log cat curl

キャプチャログを `curl` サンプルとしてフォーマットする 

Convert API request logs into equivalent curl commands that can be executed independently. This is extremely useful for debugging API issues, reproducing requests outside of the toolbox, sharing examples with support, or creating test scripts.

# インストール

[最新リリース](https://github.com/watermint/toolbox/releases/latest)からコンパイル済みのバイナリをダウンロードしてください. Windowsをお使いの方は、`tbx-xx.x.xxx-win.zip`のようなzipファイルをダウンロードしてください. その後、アーカイブを解凍し、デスクトップ フォルダに `tbx.exe` を配置します.
watermint toolboxは、システムで許可されていれば、システム内のどのパスからでも実行できます. しかし、説明書のサンプルでは、デスクトップ フォルダを使用しています. デスクトップ フォルダ以外にバイナリを配置した場合は、パスを読み替えてください.

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe log cat curl 
```

macOS, Linux:
```
$HOME/Desktop/tbx log cat curl 
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

**-buffer-size**
: 各リクエストのバッファサイズ. Default: 65536

**-record**
: コマンドラインオプションでキャプチャログファイルの記録を与える

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


