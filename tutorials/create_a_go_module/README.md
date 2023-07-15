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

## Return and handle an error

[参考](https://go.dev/doc/tutorial/handle-errors)

greeting moduleからエラーを返し、callerでそれをハンドルする

1. greetings/greetings.goでnameが空だったときにエラーを返却するようにする

    - nilはエラーがなかったことを意味する。

2. hello/hello.goを編集しerrorをハンドルするようにする

## Return a random greeting

- 何個か用意された挨拶パターンのどれかを返すように改修する。
- これを行いためにGoのsliceを使う。
- sliceは配列のようなもので、動的なサイズ変更は行えない。

### コードの特徴

- randomFormat関数はパッケージ内でのみ使用する関数なので、小文字から始める
- 初期値を記述する場合は、配列のサイズをブラケットから省略できる
  - これはスライスの根底にある配列サイズが動的に変更することができることがわかる
- randのシードはデフォルトで現在時刻となる

## Return greetings for multiple people

[参考](https://go.dev/doc/tutorial/greetings-multiple-people)

- 1個のリクエストで複数人への挨拶ができるようにする。
- 言い換えれば、複数の値をインプットして複数の値で出力する
- これをしようとすると、関数の型を変更することになる
- もしモジュールがすでに世に公開されていてHelloでの呼び出しを記述していた場合、変更により彼らのプログラムを壊してしまう
- こういう場合は、違う名前を持つ新しい関数で書くことである
- 新しい関数は複数のパラメータを持ち、後方互換性を持つ古い関数を保持する

## Add a Test

[参考](https://go.dev/doc/tutorial/add-a-test)

- Go ではbuildin moduleとして単体テストができるようになっている
- suffixが `_test.go` となるファイルを、testコマンドでは捕捉する

### 概要

- テスト関数の実装はテストしたいコードと同じパッケージに実装する
- test関数はtestingパッケージからポインタを受け取る。Tタイプをパラメータとする。
- Tのパラメータメソッドをテストのレポートやログに使用できる。
