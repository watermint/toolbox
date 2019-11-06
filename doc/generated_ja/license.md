# license 

ライセンス情報を表示します

# Usage

This document uses the Desktop folder for command example. 

## Run

Windows:

```powershell
cd $HOME\Desktop
.\tbx.exe license 
```

macOS, Linux:

```bash
$HOME/Desktop/tbx license 
```

## Options

Common options:

| オプション     | 説明                                                                | デフォルト     |
|----------------|---------------------------------------------------------------------|----------------|
| `-concurrency` | 指定した並列度で並列処理を行います                                  | プロセッサー数 |
| `-debug`       | デバッグモードを有効にする                                          | false          |
| `-proxy`       | HTTP/HTTPS プロクシ (ホスト名:ポート番号)                           |                |
| `-quiet`       | エラー以外のメッセージを抑制し、出力をJSONLフォーマットに変更します | false          |
| `-secure`      | トークンをファイルに保存しません                                    | false          |
| `-workspace`   | ワークスペースへのパス                                              |                |

## Network configuration: Proxy

The executable automatically detects your proxy configuration from the environment.
However, if you got an error or you want to specify explicitly, please add -proxy option, like -proxy hostname:port.
Currently, the executable doesn't support proxies which require authentication.

