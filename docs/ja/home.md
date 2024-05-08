# watermint toolbox

![watermint toolbox]({{ site.baseurl }}/images/logo.png){: width="160" }

Dropbox、Figma、Google、GitHubなどのウェブサービス用の多目的ユーティリティコマンドラインツールです。

# watermint toolboxでもっとできること

watermint toolboxには、日々の作業を解決するための304コマンドが用意されています. 例えば、あなたがDropbox for teamsの管理者で、グループを管理する必要がある場合。グループコマンドを使って、グループを一括作成したり、グループにメンバーを一括追加することができます.

![Demo]({{ site.baseurl }}/images/demo.gif)

watermint toolboxは、Windows、macOS（Darwin）、Linux上で、追加のライブラリなしで動作します. バイナリをダウンロードして解凍した直後に、コマンドを実行することができます.

詳細はコマンド リファレンスをご参照ください.

| 参照                                                                              |
|-----------------------------------------------------------------------------------|
| [コマンド]({{ site.baseurl }}/ja/commands/toc.html)                               |
| [チーム向けDropboxのコマンド]({{ site.baseurl }}/ja/guides/dropbox-business.html) |

# ビルド済み実行ファイル

コンパイル済みバイナリは [最新のリリース](https://github.com/watermint/toolbox/releases/latest) からダウンロードいただけます. ソースコードからビルドする場合には [BUILD.md](BUILD.md) を参照してください.

## macOS/LinuxでHomebrewを使ってインストールする。

まずHomebrewをインストールします. 手順は [オフィシャルサイト](https://brew.sh/)を参照してください. 次のコマンドを実行してwatermint toolboxをインストールします.
```
brew tap watermint/toolbox
brew install toolbox
```

# ライセンスと免責条項

watermint toolboxはApache License, Version 2.0でライセンスされています。
詳細はファイル LICENSE.mdまたは LICENSE.txt ご参照ください.

以下にご留意ください:
> Unless required by applicable law or agreed to in writing, Licensor provides the Work (and each Contributor provides its Contributions) on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied, including, without limitation, any warranties or conditions of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR PURPOSE.

# お知らせ

* [#813 License change : MIT License to Apache License, Version 2.0](https://github.com/watermint/toolbox/discussions/813)
* [#815 Lifecycle: Availability period for each release](https://github.com/watermint/toolbox/discussions/815)
* [#793 Google commands require re-authentication on Release 130](https://github.com/watermint/toolbox/discussions/793)
* [#799 Path change: Dropbox and Dropbox for teams commands have been  moved to under `dropbox`](https://github.com/watermint/toolbox/discussions/799)
* [#797 Path change: commands under `services` have been moved to a new location](https://github.com/watermint/toolbox/discussions/797)
* [#796 Deprecation: Dropbox Team space commands will be removed](https://github.com/watermint/toolbox/discussions/796)

# セキュリティとプライバシー

## 情報は収集しません 

watermint toolboxは、第三者のサーバーに情報を収集することはありません. 

watermint toolboxは、Dropbox のようなサービスとご自身のアカウントでやりとりするためのものです. 第三者のアカウントは関与していません. コマンドは、PCのストレージにAPIトークン、ログ、ファイル、またはレポートを保存します.

## 機密データ

APIトークンなどの機密データのほとんどは、難読化されてアクセス制限された状態でPCのストレージに保存されています. しかし、それらのデータを秘密にするのはあなたの責任です.
特に、ツールボックスのワークパスの下にある`secrets`フォルダ(デフォルトでは`C:\Users\<ユーザー名>\.toolbox`、または`$HOME/.toolbox`フォルダ以下)は共有しないでください。

## 使用するコマンドの対象となるサービス以外のインターネットアクセス

watermintツールボックスには、重大なバグやセキュリティ上の問題がある特定のリリースを無効にする機能があります。これは、約30日に1度、GitHubにホストされているリポジトリからデータを取得し、リリースのステータスをチェックすることによって行われます。
このアクセスは、あなたの個人データ（Dropbox、Google、ローカルファイル、トークンなど）を収集することはありません。これはリリースステータスをチェックするだけですが、副次的な効果として、データをダウンロードする際にあなたのIPアドレスがGitHubに送信されます。IPアドレスもPIIであることは承知しています。しかし、これは一般のウェブサイトを訪問するのと同じであり、特別な操作ではありません。
watermint toolboxプロジェクトの管理者は、これらのファイルが何回ダウンロードされたかを確認することもできません。

