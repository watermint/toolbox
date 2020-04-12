# dev release publish 

リリースを公開します 

# 利用方法

このドキュメントは"デスクトップ"フォルダを例として使用します.

## 実行

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe dev release publish -artifact-path /LOCAL/PATH/TO/ARTIFACT
```

macOS, Linux:

```bash
$HOME/Desktop/tbx dev release publish -artifact-path /LOCAL/PATH/TO/ARTIFACT
```

macOS Catalina 10.15以上の場合: macOSは開発者情報を検証します. 現在、`tbx`はそれに対応していません. 実行時の最初に表示されるダイアログではキャンセルします. 続いて、”システム環境設定"のセキュリティーとプライバシーから一般タブを選択します.
次のようなメッセージが表示されています:
> "tbx"は開発元を確認できないため、使用がブロックされました。

"このまま開く"というボタンがあります. リスクを確認の上、開いてください. ２回目の実行ではダイアログに"開く”ボタンがありますので、これを選択します

## オプション

| オプション       | 説明                   | デフォルト             |
|------------------|------------------------|------------------------|
| `-artifact-path` | 成果物へのパス         |                        |
| `-branch`        | Target branch          | master                 |
| `-conn-github`   | Account alias          | default                |
| `-skip-tests`    | Skip end to end tests. | false                  |
| `-test-resource` | テストリソースへのパス | test/dev/resource.json |

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

# ネットワークプロクシの設定

プログラムはシステム設定から自動的にプロクシ設定情報を取得します. しかしながら、それでもエラーが発生する場合には明示的にプロクシを指定することができます. `-proxy` オプションを利用します, `-proxy ホスト名:ポート番号`のように指定してください. なお、現在のところ認証が必要なプロクシには対応していません.

