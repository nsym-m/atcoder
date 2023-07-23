# atcoder

## 必須環境

### OS

- Linux / Unix

### Software

- [go](https://go.dev/dl/)
- [goimports](https://pkg.go.dev/golang.org/x/tools/cmd/goimports)
- [oj](https://github.com/online-judge-tools/oj/blob/master/docs/INSTALL.ja.md)
- [entr](https://manpages.ubuntu.com/manpages/impish/man1/entr.1.html)

## セットアップ

URL: https://atcoder.jp/contests/{contestId}/tasks/{taskId}

下記{contestId}と{taskId}は上記URLのものを利用

```
$ make i
./shell/init.sh
コンテストID: {contestId}
問題ID: {taskId}
```

`make i`を実行すると下記が実行されます。

- テンプレートコードの生成してファイルを開く
- テストデータのダウンロード
- 生成した`main.go`の変更監視
- `main.go`変更時のビルド＆テスト実行

## その他

- 既にダウンロード済みの問題を実行する時は`make r`を実行
- 問題を提出する際は`make t`を実行
