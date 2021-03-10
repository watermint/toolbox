# 実験的機能

実験的な機能スイッチは、テストや早期アクセス機能にアクセスするためのものです. これらの機能は `-experiment` オプションで有効にすることができます. 複数の機能を指定する場合は、カンマで結合された機能を選択してください. (例: `-experiment feature1,feature2`).

| 名前                                   | 説明                                                                                                                    |
|----------------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| dbx_client_conditioner_narrow20        | レートリミットエラーをシミュレートします. 20%のリクエストはレート制限で失敗します                                       |
| dbx_client_conditioner_narrow40        | レートリミットエラーをシミュレートします. 40%のリクエストはレート制限で失敗します                                       |
| dbx_client_conditioner_narrow100       | レートリミットエラーをシミュレートします. 100%のリクエストはレート制限で失敗します.                                     |
| dbx_client_conditioner_error20         | サーバーエラーをシミュレートします. リクエストの20%がサーバーエラーで失敗します                                         |
| dbx_client_conditioner_error40         | サーバーエラーをシミュレートします. リクエストの40%がサーバーエラーで失敗します                                         |
| dbx_client_conditioner_error100        | サーバーエラーをシミュレートします. リクエストの100%がサーバーエラーで失敗します.                                       |
| batch_balance                          | Execute batch from the largest batch                                                                                    |
| batch_random                           | ランダムなバッチIDの順番でバッチを実行します.                                                                           |
| batch_sequential                       | 同じバッチIDで順次バッチを実行します.                                                                                   |
| congestion_window_no_limit             | 輻輳ウィンドウでの同時実行を制限しない.                                                                                 |
| congestion_window_aggressive           | 積極的な初期混雑ウィンドウサイズの適用                                                                                  |
| file_sync_disable_reduce_create_folder | ファイルシステムを同期する際に reduce create_folder を無効にします. これでフォルダの同期中に空のフォルダが作成されます. |
| legacy_local_to_dbx_connector          | Use legacy local to dropbox sync connector                                                                              |
| use_no_cache_dbxfs                     | Use non-cache dropbox file system                                                                                       |

