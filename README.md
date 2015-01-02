newsreader
==========

# はじめに
Go 言語の開発に不慣れなので、作法がよく分かっていません。

# 目的
NHK のニュースを取得して、Open JTalk で読み上げるスクリプトです。

# 事前にインストールしておくもの
* Go 言語
* Open JTalk
* aplay

テストは Debian wheezy で行いました。

# 使い方
## 環境設定
    $ mkdir -p newsreader				# 実行ファイルの置き場所
    $ cd newsreader
    $ go build src/reader.go
    $ go build src/input_nhknews.go
    $ mkdir -p source					# 読み上げるテキストの置き場所

## 実行
    $ ./input_nhknews					# 記事をテキストファイル化
    $ ./reader							# 読み上げ
	$ rm source/*						# テキストを削除

# 説明
オプションは -h で確認してください。

## input_nhknews.go
NHK からニュースを取得して、記事ごとにテキストファイルを出力します。デフォルトでは source/ に出力します。長文を Open JTalk に渡すと動作が不安定になるので、記事ごとに分割しています。

## reader.go
指定ディレクトリにあるテキストファイルを読み上げます。デフォルトでは source/ のテキストファイルが対象で、手で書いたテキストでもかまいません。


# さいごに
今回はニュースを取得するようにしたけれど、TL を取得したり、Jenkins のエラーを取得したりすると実用性が上がると思います。


