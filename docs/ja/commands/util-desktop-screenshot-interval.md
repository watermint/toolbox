---
layout: command
title: コマンド
lang: ja
---

# util desktop screenshot interval

Take screenshots at regular intervals 

# インストール

[最新リリース](https://github.com/watermint/toolbox/releases/latest)からコンパイル済みのバイナリをダウンロードしてください. Windowsをお使いの方は、`tbx-xx.x.xxx-win.zip`のようなzipファイルをダウンロードしてください. その後、アーカイブを解凍し、デスクトップ フォルダに `tbx.exe` を配置します.
watermint toolboxは、システムで許可されていれば、システム内のどのパスからでも実行できます. しかし、説明書のサンプルでは、デスクトップ フォルダを使用しています. デスクトップ フォルダ以外にバイナリを配置した場合は、パスを読み替えてください.

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:
```
cd $HOME\Desktop
.\tbx.exe util desktop screenshot interval -path /LOCAL/PATH/TO/SCREENSHOT/DIR -interval INTERVAL_SECONDS
```

macOS, Linux:
```
$HOME/Desktop/tbx util desktop screenshot interval -path /LOCAL/PATH/TO/SCREENSHOT/DIR -interval INTERVAL_SECONDS
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション:

| オプション           | 説明                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | デフォルト                       |
|----------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------|
| `-count`             | Number of screenshots to take. If the value is less than 1, the screenshot is taken until the process is killed.                                                                                                                                                                                                                                                                                                                                                                                                                                       | -1                               |
| `-display-id`        | Display ID to take screenshot. To get the display ID, run `util desktop display list` command.                                                                                                                                                                                                                                                                                                                                                                                                                                                         | 0                                |
| `-interval`          | Interval seconds between screenshots.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | 10                               |
| `-name-pattern`      | Name pattern of screenshot file. You can use the following placeholders:`<no value>` .. date (yyyy-MM-dd), `<no value>` .. date in UTC (yyyy-MM-dd), `<no value>` .. display height, `<no value>` .. display ID, `<no value>` .. display width, `<no value>` .. display horizontal offset, `<no value>` .. display vertical offset, `<no value>` .. 5 digit sequence number, `<no value>` .. time (HH-mm-ss), `<no value>` .. time in UTC (HH-mm-ss), `<no value>` .. timestamp (yyyyMMdd-HHmmss), `<no value>` .. timestamp in UTC (yyyyMMdd-HHmmss). | {% raw %}{{.{% endraw %}Sequence}}_{% raw %}{{.{% endraw %}Timestamp}}.png |
| `-path`              | Path to the folder to save screenshots.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |                                  |
| `-skip-if-no-change` | Skip taking screenshot if the screen is not changed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | false                            |

## 共通のオプション:

| オプション         | 説明                                                                                               | デフォルト     |
|--------------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auth-database`   | 認証データベースへのカスタムパス (デフォルト: $HOME/.toolbox/secrets/secrets.db)                   |                |
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
| `-retain-job-data` | ジョブデータ保持ポリシー                                                                           | default        |
| `-secure`          | トークンをファイルに保存しません                                                                   | false          |
| `-skip-logging`    | ローカルストレージへのログ保存をスキップ                                                           | false          |
| `-verbose`         | 現在の操作を詳細に表示します.                                                                      | false          |
| `-workspace`       | ワークスペースへのパス                                                                             |                |

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.


