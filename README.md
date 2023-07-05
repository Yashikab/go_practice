# go_practice

## Install

- 自分の場合は chromebook or ubuntuなのでアーキテクチャを調べる

```shell
uname -m
```

- x86-64 だったので、これにあったファイルをDLする。<https://go.dev/dl/>
- <https://go.dev/doc/install> を参考にインストールする。(現行 1.20.5)

```shell
 rm -rf /usr/local/go && sudo tar -C /usr/local -xzf path_to/go1.20.5.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```
