---
layout: page
title: パス変数
---

# パス変数

パス変数は、実行時に置き換えられる定義済みの変数です. 例えば、`{% raw %}{{{% endraw %}.DropboxPersonal{% raw %}}}{% endraw %}/Pictures`のように変数でパスを指定すると、そのパスは個人用のDropbox フォルダへの実際のパスに置き換えられます. しかし、このツールはその存在や精度を保証するものではありません.

| パス変数                       | 説明                                                                                                                |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------|
| {% raw %}{{{% endraw %}.DropboxPersonal{% raw %}}}{% endraw %}           | Dropbox 個人アカウントのルート フォルダへのパス.                                                                    |
| {% raw %}{{{% endraw %}.DropboxBusiness{% raw %}}}{% endraw %}           | Dropbox Business アカウントのルート フォルダへのパス.                                                               |
| {% raw %}{{{% endraw %}.DropboxBusinessOrPersonal{% raw %}}}{% endraw %} | Dropbox Business アカウントのルート フォルダへのパス、または見つからない場合は Personal Dropbox アカウントへのパス. |
| {% raw %}{{{% endraw %}.DropboxPersonalOrBusiness{% raw %}}}{% endraw %} | Dropbox Personal アカウントのルート フォルダへのパス、または見つからない場合は Business Dropbox アカウントへのパス. |
| {% raw %}{{{% endraw %}.Home{% raw %}}}{% endraw %}                      | 現在のユーザーのホームフォルダ.                                                                                     |
| {% raw %}{{{% endraw %}.Username{% raw %}}}{% endraw %}                  | 現在のユーザーの名前.                                                                                               |
| {% raw %}{{{% endraw %}.Hostname{% raw %}}}{% endraw %}                  | 現在のコンピュータのホスト名.                                                                                       |
| {% raw %}{{{% endraw %}.ExecPath{% raw %}}}{% endraw %}                  | このプログラムへのパス.                                                                                             |
| {% raw %}{{{% endraw %}.Rand8{% raw %}}}{% endraw %}                     | 0で始まるランダムな8桁の数字.                                                                                       |
| {% raw %}{{{% endraw %}.Year{% raw %}}}{% endraw %}                      | ローカル日時の年 (フォーマット 'yyyy', 例: 2021)                                                                    |
| {% raw %}{{{% endraw %}.Month{% raw %}}}{% endraw %}                     | ローカル日時の月 (フォーマット 'mm', 例: 01)                                                                        |
| {% raw %}{{{% endraw %}.Day{% raw %}}}{% endraw %}                       | ローカル日時の日 (フォーマット'dd', 例: 05)                                                                         |
| {% raw %}{{{% endraw %}.Date{% raw %}}}{% endraw %}                      | yyyyy-mm-dd形式の現在のローカル日付.                                                                                |
| {% raw %}{{{% endraw %}.Time{% raw %}}}{% endraw %}                      | 現在の現地時間をHH-MM-SSの形式で表示します.                                                                         |
| {% raw %}{{{% endraw %}.DateUTC{% raw %}}}{% endraw %}                   | 現在のUTCの日付をyyyy-mm-dd形式で表示.                                                                              |
| {% raw %}{{{% endraw %}.TimeUTC{% raw %}}}{% endraw %}                   | 現在のUTC時間をHH-MM-SS形式で表示します.                                                                            |


