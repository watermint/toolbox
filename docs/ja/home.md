# watermint toolbox

![watermint toolbox]({{ site.baseurl }}/images/logo.png){: width="160" }

Dropbox、Figma、Google、GitHubなどのウェブサービス用の多目的ユーティリティコマンドラインツールです。

# watermint toolboxでもっとできること

watermint toolboxには、日々の作業を解決するための305コマンドが用意されています. 例えば、あなたがDropbox for teamsの管理者で、グループを管理する必要がある場合。グループコマンドを使って、グループを一括作成したり、グループにメンバーを一括追加することができます.

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

* [#836 Remove binaries that are more than six months old after release](https://github.com/watermint/toolbox/discussions/836)
* [#835 Google commands deprecation](https://github.com/watermint/toolbox/discussions/835)
* [#813 License change : MIT License to Apache License, Version 2.0](https://github.com/watermint/toolbox/discussions/813)
* [#815 Lifecycle: Availability period for each release](https://github.com/watermint/toolbox/discussions/815)
* [#799 Path change: Dropbox and Dropbox for teams commands have been  moved to under `dropbox`](https://github.com/watermint/toolbox/discussions/799)
* [#797 Path change: commands under `services` have been moved to a new location](https://github.com/watermint/toolbox/discussions/797)
* [#796 Deprecation: Dropbox Team space commands will be removed](https://github.com/watermint/toolbox/discussions/796)

# セキュリティとプライバシー

watermint toolboxは、クラウドサービスAPIの使用を簡素化するように設計されています。あなたの意図に反する方法でデータを使用することはありません。

watermint toolboxは、指定されたコマンドの意図に反して、リンクされたクラウドサービスAPIを介して取得したデータを別のサーバーに保存しません。

例えば、watermint toolboxを使ってクラウドサービスからデータを取得した場合、そのデータは自分のPCにしか保存されません。さらに、ファイルまたはデータをクラウドサービスにアップロードするコマンドの場合、それらはお客様のアカウントによって指定された場所にのみ保存されます。

## データ保護

watermint toolboxを使ってクラウドサービスのAPIからデータを取得すると、レポートデータやログデータとしてPCに保存されます。クラウドサービスAPIの認証トークンなど、より機密性の高い情報もPCに保存されます。

あなたのPCに保存されているこれらのデータを安全に保つことは、あなたの責任です。

認証トークンのような重要な情報は難読化されているため、その内容を簡単に読み取ることはできません。しかし、この難読化はセキュリティを強化するためではなく、意図しない操作ミスを防ぐためのものです。悪意のある第三者があなたのトークン情報を別のPCにコピーした場合、あなたが意図しないクラウドサービスにアクセスされる可能性があります。

## 利用

前述の通り、watermint toolboxはPCまたはクラウドアカウントにデータを保存するように設計されています。あなたが意図した操作以外のプロセスには、以下に説明するように、リリースのライフサイクル管理のためのデータ検索が含まれます。

watermintツールボックスには、重大なバグやセキュリティ上の問題がある特定のリリースを無効にする機能があります。これは、GitHubにホストされているリポジトリから約30日ごとにデータを取得し、リリースのステータスを評価することで達成されます。このアクセスによって個人情報（クラウドアカウント情報、ローカルファイル、トークンなど）が収集されることはありません。これは単にリリース状況をチェックするだけだが、副次的な効果として、データをダウンロードする際にあなたのIPアドレスがGitHubに送信されます。

このアクセス情報（日時、IPアドレス）は、今後、各リリースの利用状況を推定するために使用することがありますので、あらかじめご了承ください。

## 共有

watermint toolboxプロジェクトは現在、IPアドレスを含むデータを管理・取得していない。これは、プロジェクトをホストするGitHub社のみがアクセス可能な情報です。ただし、プロジェクトは将来的にこの情報を公開する可能性があり、プロジェクトの運営上必要と判断される場合には、匿名化されたリリースごとの使用状況をプロジェクト・メンバーに開示することがあります。

このような変更は、変更が有効になる少なくとも30日前までに、告知ページおよびこのセキュリティ＆プライバシーポリシーページで発表されます。

