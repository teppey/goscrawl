# Goscribble

`goscribble` is a tool for edit and run Go code instantly on console.

`goscribble` はGoプログラムの作成と実行を手軽に行うためのコマンドラインツールです。"Go Playground"のような気軽さをコンソール上で実現することを意図しています。

## 使い方

`goscribble new` を実行すると「Hello World」を出力するGoコードがエディタで開かれます。

エディタを閉じると編集していたGoコードが実行されます。



```
goscribble <command> [arguments]
```

The commands are:

```
clear  remove all Go files under `$GOSCRIBBLE_DIR`
edit   edit and run Go file
list   list under `$GOSCRIBBLE_DIR` directory
new    create and edit Go file, then run it
run    run Go file
show   display Go file
```

```
clear  `$GOSCRIBBLE_DIR`ディレクトリ以下のファイルを削除します
edit   既存の`0.go`ファイルをエディタで開き、その後それを実行します。
list   `$GOSCRIBBLE_DIR`ディレクトリ以下のファイルのパスを表示します。ｍ
new    `$GOSCRIBBLE_DIR`ディレクトリ以下に`0.go`ファイルを作成してエディタで開き、その後それを実行します。
run    `0.go`ファイルを実行します。
show   `0.go`ファイルの内容を表示します。
```


goscribble clear

goscribble edit

goscribble list

goscribble new [template]

goscribble run

goscribble show




## インストール




環境変数：

`GOSCRIBBLE_DIR`
`GOSCRIBBLE_TEMPLATE`
