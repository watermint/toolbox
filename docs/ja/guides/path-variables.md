---
layout: page
title: パス変数
lang: ja
---

# パス変数

パス変数は、実行時に置き換えられる定義済みの変数です. 例えば、`{% raw %}{{.{% endraw %}DropboxPersonal}}/Pictures`のように変数でパスを指定すると、そのパスは個人用のDropbox フォルダへの実際のパスに置き換えられます. しかし、このツールはその存在や精度を保証するものではありません.

| パス変数                       | 説明                                                                                                                |
|--------------------------------|---------------------------------------------------------------------------------------------------------------------|
| {% raw %}{{.{% endraw %}DropboxPersonal}}           | Dropbox 個人アカウントのルート フォルダへのパス.                                                                    |
| {% raw %}{{.{% endraw %}DropboxBusiness}}           | Dropbox for teamsアカウントのルートフォルダへのパス。                                                               |
| {% raw %}{{.{% endraw %}DropboxBusinessOrPersonal}} | Dropbox for teamsアカウントのルートフォルダへのパス、または見つからない場合はPersonal Dropboxアカウントへのパス。   |
| {% raw %}{{.{% endraw %}DropboxPersonalOrBusiness}} | Dropbox Personal アカウントのルート フォルダへのパス、または見つからない場合は Business Dropbox アカウントへのパス. |
| {% raw %}{{.{% endraw %}Home}}                      | 現在のユーザーのホームフォルダ.                                                                                     |
| {% raw %}{{.{% endraw %}Username}}                  | 現在のユーザーの名前.                                                                                               |
| {% raw %}{{.{% endraw %}Hostname}}                  | 現在のコンピュータのホスト名.                                                                                       |
| {% raw %}{{.{% endraw %}ExecPath}}                  | このプログラムへのパス.                                                                                             |
| {% raw %}{{.{% endraw %}Rand8}}                     | 0で始まるランダムな8桁の数字.                                                                                       |
| {% raw %}{{.{% endraw %}Year}}                      | ローカル日時の年 (フォーマット 'yyyy', 例: 2021)                                                                    |
| {% raw %}{{.{% endraw %}Month}}                     | ローカル日時の月 (フォーマット 'mm', 例: 01)                                                                        |
| {% raw %}{{.{% endraw %}Day}}                       | ローカル日時の日 (フォーマット'dd', 例: 05)                                                                         |
| {% raw %}{{.{% endraw %}Date}}                      | yyyyy-mm-dd形式の現在のローカル日付.                                                                                |
| {% raw %}{{.{% endraw %}Time}}                      | 現在の現地時間をHH-MM-SSの形式で表示します.                                                                         |
| {% raw %}{{.{% endraw %}DateUTC}}                   | 現在のUTCの日付をyyyy-mm-dd形式で表示.                                                                              |
| {% raw %}{{.{% endraw %}TimeUTC}}                   | 現在のUTC時間をHH-MM-SS形式で表示します.                                                                            |


