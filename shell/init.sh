#!/bin/bash

default_url="https://atcoder.jp/contests"
current=$(pwd)

read -p "コンテストID: " contest
read -p "問題ID: " task
echo ""

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
    code "$current/$contestDir/main.go"
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

code "$current/$contestDir/main.go"

make r
