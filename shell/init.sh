#!/bin/bash

default_url="https://atcoder.jp/contests"
current=$(pwd)
contest=$1
task=$2
diff=$3

# 通常のコンテストはタスク名が"コンテスト名_a~g"になっているが、たまに違うものがあるため違う時だけdiffをyesにして実行する
if [ "$diff" != "yes" ]; then
    task="${contest}_${task}"
fi

url="${default_url}/${contest}/tasks/${task}"

cd contests

if [ ! -d "$contest" ]; then
    mkdir "$contest" || {
        echo "ディレクトリ作成に失敗しました: ${contest}"
        exit 1
    }
fi

cd "$contest"

if [ ! -d "$task" ]; then
    mkdir "$task" || {
        echo "ディレクトリ作成に失敗しました: ${task}"
        exit 1
    }
fi

cd "$task"

contestDir="contests/$contest/$task"

# テストとmain.goが既に存在するかチェック
if [ -d "test" ] && [ -n "$(find test -name '*.in')" ] && [ -n "$(find test -name '*.out')" ]; then
    if [ ! -f "main.go" ]; then
        cp $current/template/main.go .
    fi
    echo "$contest" >$current/state
    echo "$task" >>$current/state
    cd $current
    open "$current/$contestDir/main.go"
    make r
    exit
fi

# テストをダウンロード
oj d $url

cp $current/template/main.go .

# 対象のコンテストをstateに持たせる
echo "$contest" >$current/state
echo "$task" >>$current/state

cd $current

open "$current/$contestDir/main.go"

make r
