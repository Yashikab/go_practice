# Create a module

- [参考](https://go.dev/doc/tutorial/)

## 概要

2つのモジュールを作る。

- 1つは他のライブラリやアプリケーションでインポートされるもの
- もう一つはアプリケーションで最初に呼び出されるもの

## 習得したいこと

Create a module -- Write a small module with functions you can call from another module.
Call your code from another module -- Import and use your new module.
Return and handle an error -- Add simple error handling.
Return a random greeting -- Handle data in slices (Go's dynamically-sized arrays).
Return greetings for multiple people -- Store key/value pairs in a map.
Add a test -- Use Go's built-in unit testing features to test your code.
Compile and install the application -- Compile and install your code locally.

## Start a module that others can use

- モジュール内では、一つ以上の関数群を収集する。
- goのコードはパッケージへ集約される
- パッケージはモジュールへと集約される
- モジュールが特定の依存関係を必要とする場合は、goのバージョンと必要なモジュールのセットを含ませる必要がある
- モジュールを追加したり改修した場合は新しいバージョンのモジュールとして公開する
  - モジュールを呼ぶ人は更新されたモジュールを読んだりテストしたりする

### 特徴

- 大文字で始まる関数は同じパッケージ内では呼ばれない
  - exported nameとしてGoでは知られている
- `:=`を使うと宣言と初期化を一気にできる。

## Call your code from another module

[参考](https://go.dev/doc/tutorial/call-module-code)

- `hello`ディレクトリを作り、greetingを呼び出す
- 依存関係のトラックをできるようにする

```shell
go mod init github.com/Yashikab/go_practice/tutorials/create_a_go_module/hello
```

- `hello.go`をつくる
- Hello関数を呼び出すコードを `hello.go`に書き、関数の戻り値を出力する
- helloモジュールをローカルのgreetingsモジュールが使えるように編集する
  - productionで使用する場合は、greetingsモジュールをリポジトリから公開する(Go toolsがDLするために見つけられる場所で)
  - 今はまだ公開していないので、ローカルから見つけられるように接続する必要がある

  ```shell
  go mod edit -replace github.com/Yashikab/go_practice/tutorials/create_a_go_module/greetings=../greetings
  go mod tidy  # 同期する
  ```
