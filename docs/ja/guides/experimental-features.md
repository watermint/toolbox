---
layout: page
title: 実験的機能
lang: ja
---

# 実験的機能

実験的な機能スイッチは、テストや早期アクセス機能にアクセスするためのものです. これらの機能は `-experiment` オプションで有効にすることができます. 複数の機能を指定する場合は、カンマで結合された機能を選択してください. (例: `-experiment feature1,feature2`).

| 名前                                     | 説明                                                                                                                                                                                  |
|------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| batch_balance                            | 大きいバッチから順に実行する                                                                                                                                                          |
| batch_non_durable                        | 非耐久性バッチフレームワークの使用                                                                                                                                                    |
| batch_random                             | ランダムなバッチIDの順番でバッチを実行します.                                                                                                                                         |
| batch_sequential                         | 同じバッチIDで順次バッチを実行します.                                                                                                                                                 |
| congestion_window_aggressive             | 積極的な初期混雑ウィンドウサイズの適用                                                                                                                                                |
| congestion_window_no_limit               | 輻輳ウィンドウでの同時実行を制限しない.                                                                                                                                               |
| dbx_auth_course_grained_scope            | コマンドで定義されたスコープではなく、すべてのDropboxの認証スコープを要求します。これは、コマンドで定義された認可範囲ではプログラムが正常に動作しない場合の回避策として使用されます。 |
| dbx_auth_redirect                        | Dropboxへの認証処理にリダイレクト処理を使用する。                                                                                                                                     |
| dbx_client_conditioner_error100          | サーバーエラーをシミュレートします. リクエストの100%がサーバーエラーで失敗します.                                                                                                     |
| dbx_client_conditioner_error20           | サーバーエラーをシミュレートします. リクエストの20%がサーバーエラーで失敗します                                                                                                       |
| dbx_client_conditioner_error40           | サーバーエラーをシミュレートします. リクエストの40%がサーバーエラーで失敗します                                                                                                       |
| dbx_client_conditioner_narrow100         | レートリミットエラーをシミュレートします. 100%のリクエストはレート制限で失敗します.                                                                                                   |
| dbx_client_conditioner_narrow20          | レートリミットエラーをシミュレートします. 20%のリクエストはレート制限で失敗します                                                                                                     |
| dbx_client_conditioner_narrow40          | レートリミットエラーをシミュレートします. 40%のリクエストはレート制限で失敗します                                                                                                     |
| dbx_disable_auto_path_root               | 自動パスルートを無効にする。無効にすると、ユーザーのホーム名前空間がルート名前空間と異なる場合、API呼び出しのデフォルトとしてユーザー ホーム名前空間が使用されます。                  |
| dbx_download_block                       | ダウンロードファイルをブロック単位で分割（並行処理性の向上）                                                                                                                          |
| file_sync_disable_reduce_create_folder   | ファイルシステムを同期する際に reduce create_folder を無効にします. これでフォルダの同期中に空のフォルダが作成されます.                                                               |
| legacy_local_to_dbx_connector            | 古いローカルとDropboxの同期コネクタを使用                                                                                                                                             |
| use_no_cache_dbxfs                       | ノンキャッシュのDropboxファイルシステムの使用                                                                                                                                         |
| kvs_badger                               | BadgerをKVSエンジンに使用                                                                                                                                                             |
| kvs_badger_turnstile                     | Badger+TurnstileをKVSエンジンとして使用                                                                                                                                               |
| kvs_bitcask                              | KVSエンジンとしてBitcaskを使用                                                                                                                                                        |
| kvs_bitcask_turnstile                    | Bitcask+turnstileのキー・バリュー・ストアとして使う                                                                                                                                   |
| kvs_sqlite                               | KVSエンジンとしてSqlite3を使用                                                                                                                                                        |
| kvs_sqlite_turnstile                     | SQLite+turnstileをキー・バリュー・ストアとして使う                                                                                                                                    |
| profile_cpu                              | CPUプロファイラの有効化                                                                                                                                                               |
| profile_memory                           | メモリプロファイラの有効化                                                                                                                                                            |
| report_all_columns                       | データ構造として定義されているすべての列を表示します.                                                                                                                                 |
| suppress_progress                        | 進捗表示の抑制                                                                                                                                                                        |
| validate_network_connection_on_bootstrap | 起動時にネットワーク接続を検証する                                                                                                                                                    |


