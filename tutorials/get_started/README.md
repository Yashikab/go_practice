## Write Some code

- ほかのモジュールに含まれるパッケージをインポートするときは、これらの依存関係を自分のモジュールを通して管理する。
- モジュールは go.modファイルに定義される。
- 依存関係のトラッキングを有効にするために go mod init コマンドを実行する。
- 実際の開発では、モジュールパスはコードをを管理するリポジトリの場所にある。
- チュートリアルの目的では、 example/helloを使用する

```shell
go mod init example/hello
```

- hello.go ファイルを作成し、次のコードを保存する
  - main packageの記述
  - テキスト出力のフォーマットなどを行う fmt のインポート
  - main関数の実装。main関数はmain packageを実行するときにデフォルトで実行される

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

## Call Code in an external package

- 外部にあるquote packageを使って出力する。

```go
package main

import "fmt"

import "rsc.io/quote"

func main() {
    fmt.Println(quote.Go())
}
```

- quoteモジュールをrequirementに加える
  - モジュールの認証のために、go.sumファイルが同時に作成される
- `go run .`で実行する
