# ファイアウォールまたはプロキシサーバーの設定

ツールは、システムからプロキシの設定を自動的に検出します. しかし、それが失敗したり、設定ミスの原因になったりすることがあります. このような場合は、`-proxy` オプションを使って `-proxy 192.168.1.1.1:8080` のようにプロキシサーバのホスト名とポート番号を指定してください (プロキシサーバ 192.168.1.1.1 、ポート番号は8080の場合). 

注意：このツールは、Basic認証やNTLMなどの認証を持つプロキシサーバには対応していません.

# パフォーマンスの問題

コマンドが遅く感じたり、停止したりした場合は、オプション `-verbose` を指定して再実行してみてください. そうすることで、より詳細な進捗状況がわかります. しかし、ほとんどの場合、原因は単純にあなたが処理するためのより大きなデータを持っているだけです. そうでなければ、APIサーバーからのレート制限にすでにヒットしていることになります. レート制限の状態を見たい場合は、キャプチャログやデバッグを参照してください. 

このツールは、APIサーバーからの追加制限を回避するために、並行性を自動的に調整します. 現在の並行性を確認したい場合は、以下のようなコマンドを実行してください. これは、エンドポイントごとの現在のウィンドウサイズ（最大同時実行数）を表示します. デバッグメッセージ"WaiterStatus"は、現在の同時実行とウィンドウサイズを報告します. マップ"runners"は、現在APIサーバーからの結果待ちの操作のためのものですマップ "window "は、各エンドポイントのウィンドウサイズのためのものです. マップ "concurrency "は、現在実行中の操作のためのウィンドウサイズのためのものです. 次の例はエンドポイント "https://api.dropboxapi.com/2/file_requests/create" のために、ツールは1より大きい同時実行のそのエンドポイントを呼ぶことを許可しないことを示します. つまり、一つ一つの操作が必要であり、操作を高速化するための簡単な回避策はありません.
```
tbx job log last -quiet | jq 'select(.msg == "WaiterStatus")' 
{
  "level": "DEBUG",
  "time": "2020-11-10T14:55:57.501+0900",
  "name": "z951.z960.z112064",
  "caller": "nw_congestion/congestion.go:310",
  "msg": "WaiterStatus",
  "goroutine": "gr:284877",
  "runners": {
    "gr:1": {
      "key": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create",
      "go_routine_id": "gr:1",
      "running_since": "2020-11-10T14:55:56.124899+09:00"
    }
  },
  "numRunners": 1,
  "waiters": [],
  "numWaiters": 0,
  "window": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/team/token/get_authenticated_admin": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/save_url/check_job_status": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search/continue_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/search_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/users/get_current_account": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_reference/get": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/copy_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/delete_v2": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/get_metadata": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/files/list_folder": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/sharing/list_mountable_folders": 4,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/download": 5,
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://content.dropboxapi.com/2/files/export": 4
  },
  "concurrency": {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-https://api.dropboxapi.com/2/file_requests/create": 1
  }
}
```

# 文字化けした出力

ツールの出力が文字化けしてしまう場合は、Ctrl+Cでツールを停止してください. この問題は、通常、コンソールに表示するフォントがない場合に発生します. そして、言語に対応したフォントに変更してみてください. あるいは、ツールの言語設定を英語に上書きするオプション `-lang en` を試してみてください.

PowerShellでは、(1)タイトルバーを右クリックし、(2)プロパティをクリックし、(3)フォントタブを選択し、(4)フォントを "MSゴシック "のような適切なフォントに変更することで、フォントを変更することができます.

# ログファイル

既定では、ログファイルは、Windows上のパス"%USERPROFILE%\.toolbox\jobs" (例えば、`C:\Users\USERNAME\.toolbox\jobs`)、またはLinuxまたはmacOS上の"$HOME/.toolbox/jobs" (例えば、`/Users/USERNAME/.toolbox/jobs`)の下に格納されています. ログファイルには、(1)OSの種類/バージョン/環境変数などのランタイム情報、(2)ツールへのランタイムオプション（入力データファイルのコピーを含む）、(3)Dropboxなどのサービスのアカウント情報、(4)APIサーバへのリクエスト/レスポンスデータ、(5)ファイル名、メタデータ、ID、URLなどのサービス内のデータなどの情報が含まれています。コマンドに依存します）.

これらのログには、パスワード、クレデンシャル、または API トークンが含まれていません. しかし、APIトークンは、Windows上のパス"%USERPROFILE%\.toolbox\secrets" (例えば、`C:\ Users\\USERNAME\.toolbox\secrets`)や、LinuxやmacOS上のパス"$HOME/.toolbox/secrets" (例えば、`/Users/USERNAME/.toolbox/secrets`)の下に格納されています. これらの秘密のフォルダファイルは難読化されていますが、Dropboxのサポートなどのサービスプロバイダのサポートを含む誰にも共有しないようにしてください.

## ログ書式

`jobs` フォルダの下には、いくつかのフォルダとファイルが保存されています. まず、"yyyyMMdd-HHmmSS.xxx"という形式の名前（内部的にはJob Idと呼ばれています）で、実行するたびにジョブフォルダが作成されます. 最初の "yyyyMMdd-HHmmSS "は、コマンド開始のローカル日時です. 2番目の部分".xxx"は、同時実行との競合を避けるために、シーケンシャルまたはランダムな3文字のIDです.

ジョブフォルダの下にはサブフォルダがあり、(1) `logs`: リクエスト/レスポンスデータやパラメータ、デバッグ情報を含む実行時のログ、(2) `reports`: 生成されたレポートを管理するためのレポートフォルダ、(3) `kvs` : KVSフォルダは実行時のデータベースフォルダです. 

トラブルシューティングでは、`logs`以下のファイルは実行時に何が起こったかを理解するために必要不可欠です. このツールは、いくつかの種類のログを生成します. これらのログは、JSON Lines形式です. 注：JSON Linesは、データを行区切り文字で区切るフォーマットです. 仕様の詳細は [JSON Lines](https://jsonlines.org/) をご覧ください.

一部のログはgzip形式で圧縮されています. ログが圧縮されている場合は、ファイルの接尾辞が '.gz' になります. さらに、captureログやtoolboxログなどのログは、一定の大きさに分割されています. ログを解析したい場合は、`job log` コマンドの利用を検討してください. 例えば、`job log last -quiet` は最新のジョブのtoolboxログを解凍したうえで、連結して出力します.

## デバッグログ

このツールは、すべてのデバッグ情報を"toolbox"という接頭辞を持つデバッグログに記録します. すべてのレコードには、操作時のソースコードファイル名と行が記載されています. 怪しいエラーを見つけたら、ソースコードを見てデバッグしましょう. トラブルシューティングの中には、パフォーマンスチューニングやメモリ切れなどの統計解析を必要とするものがあります. `grep` や [jq](https://stedolan.github.io/jq/) のようなツールを使って作業するのが良いでしょう. 

時系列でヒープサイズのデータを見たい場合は、以下のようなコマンドを実行してください. そうすると、時間＋ヒープサイズがCSV形式で表示されます.
```
tbx job log last -quiet | jq -r 'select(.msg == "Heap stats") | [.time, .HeapInuse] | @csv'
"2020-11-10T14:55:45.725+0900",18604032
"2020-11-10T14:55:50.725+0900",15130624
"2020-11-10T14:55:55.725+0900",17408000
"2020-11-10T14:56:00.725+0900",17014784
"2020-11-10T14:56:05.726+0900",19193856
"2020-11-10T14:56:10.725+0900",19136512
"2020-11-10T14:56:15.726+0900",16637952
"2020-11-10T14:56:20.725+0900",16678912
"2020-11-10T14:56:25.727+0900",16678912
"2020-11-10T14:56:30.730+0900",16678912
"2020-11-10T14:56:35.726+0900",16678912
```
## APIトランザクションのログ

トールはAPIリクエストとレスポンスを、接頭辞"capture"を持つキャプチャログに記録しますこのキャプチャログにはOAuthのリクエストとレスポンスは含まれていません. さらに、APIトークンの文字列は `<secret>` に置き換えられます.

