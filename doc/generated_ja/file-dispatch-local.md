# file dispatch local 

ローカルファイルを整理します 

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe file dispatch local -file /PATH/TO/DATA_FILE.csv
```

macOS, Linux:

```bash
$HOME/Desktop/tbx file dispatch local -file /PATH/TO/DATA_FILE.csv
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション

| オプション | 説明                   | デフォルト |
|------------|------------------------|------------|
| `-file`    | データファイルへのパス |            |
| `-preview` | プレビューモード       | false      |

共通のオプション:

| オプション      | 説明                                                                                               | デフォルト     |
|-----------------|----------------------------------------------------------------------------------------------------|----------------|
| `-auto-open`    | 成果物フォルダまたはURLを自動で開く                                                                | false          |
| `-bandwidth-kb` | コンテンツをアップロードまたはダウンロードする際の帯域幅制限(Kバイト毎秒). 0の場合、制限を行わない | 0              |
| `-concurrency`  | 指定した並列度で並列処理を行います                                                                 | プロセッサー数 |
| `-debug`        | デバッグモードを有効にする                                                                         | false          |
| `-low-memory`   | 省メモリモード                                                                                     | false          |
| `-output`       | 出力書式 (none/text/markdown/json)                                                                 | text           |
| `-proxy`        | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                                                          |                |
| `-quiet`        | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します                                | false          |
| `-secure`       | トークンをファイルに保存しません                                                                   | false          |
| `-workspace`    | ワークスペースへのパス                                                                             |                |

# ファイル書式

## 書式: File

Data file for dispatch rules. 

| 列                  | 説明                                  | 値の例                                        |
|---------------------|---------------------------------------|-----------------------------------------------|
| suffix              | ファイル名の拡張子                    | .pdf                                          |
| source_path         | 元のパス                              | <no value>/ダウンロード                       |
| source_file_pattern | 転送元ファイル名のパターン (正規表現) | toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2})      |
| dest_path_pattern   | 転送先パスのパターン                  | <no value>/ドキュメント/<no value>-<no value> |
| dest_file_pattern   | 転送先ファイル名のパターン            | TBX_<no value>-<no value>-<no value>          |

最初の行はヘッダ行です. プログラムはヘッダ行がない場合も認識します.

```csv
suffix,source_path,source_file_pattern,dest_path_pattern,dest_file_pattern
.pdf,<no value>/ダウンロード,toolbox-([0-9]{4})-([0-9]{2})-([0-9]{2}),<no value>/ドキュメント/<no value>-<no value>,TBX_<no value>-<no value>-<no value>
```

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.
