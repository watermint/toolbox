# watermint toolbox

![watermint toolbox]({{ site.baseurl }}/images/logo.png){: width="160" }

Dropbox、Figma、Google、GitHubなどのウェブサービス用の多目的ユーティリティコマンドラインツールです。

# watermint toolboxでもっとできること

watermint toolboxには、日々の作業を解決するための295コマンドが用意されています. 例えば、あなたがDropbox for teamsの管理者で、グループを管理する必要がある場合。グループコマンドを使って、グループを一括作成したり、グループにメンバーを一括追加することができます.

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

watermint toolboxはMITライセンスのもと配布されています.
詳細はファイル LICENSE.mdまたは LICENSE.txt ご参照ください.

以下にご留意ください:
> ソフトウェアは「現状のまま」で、明示であるか暗黙であるかを問わず、何らの保証もなく提供されます。ここでいう保証とは、商品性、特定の目的への適合性、および権利非侵害についての保証も含みますが、それに限定されるものではありません。

# セキュリティとプライバシー

## 情報は収集しません 

watermint toolboxは、第三者のサーバーに情報を収集することはありません. 

watermint toolboxは、Dropbox のようなサービスとご自身のアカウントでやりとりするためのものです. 第三者のアカウントは関与していません. コマンドは、PCのストレージにAPIトークン、ログ、ファイル、またはレポートを保存します.

## 機密データ

APIトークンなどの機密データのほとんどは、難読化されてアクセス制限された状態でPCのストレージに保存されています. しかし、それらのデータを秘密にするのはあなたの責任です.
特に、ツールボックスのワークパスの下にある`secrets`フォルダ(デフォルトでは`C:\Users\<ユーザー名>\.toolbox`、または`$HOME/.toolbox`フォルダ以下)は共有しないでください。

