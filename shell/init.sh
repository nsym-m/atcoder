#!/bin/bash

default_url="https://atcoder.jp/contests"
current=$(pwd)

read -p "コンテストID: " contest
read -p "問題ID: " task
echo ""

url="${default_url}/${contest}/tasks/${task}"

status=$(curl -I "${url}" 2>/dev/null | head -n 1 | cut -d$' ' -f2)

# if分を正しいbashの書き方に修正
if [ $status -ne 200 ]; then
    echo "問題を取得できません。"
    exit 1
fi

cd contests

# $contestディレクトリの存在確認。あればスキップ、なければディレクトリ作成、作成失敗したらechoでエラー吐いて終了
if [ ! -d "$contest" ]; then
    mkdir "$contest" || {
        echo "ディレクトリ作成に失敗しました: ${contest}"
        exit 1
    }
fi

# $contestディレクトリ配下に移動
cd "$contest"

# $taskディレクトリの存在確認。あればスキップ、なければディレクトリ作成、作成失敗したらechoでエラー吐いて終了
if [ ! -d "$task" ]; then
    mkdir "$task" || {
        echo "ディレクトリ作成に失敗しました: ${task}"
        exit 1
    }
fi

# $taskディレクトリ配下に移動
cd "$task"

contestDir="contests/$contest/$task"

# TODO: 下記のif文を正しい形に修正
if [ -d "test" ] && [ -n "$(find test -name '*.in')" ] && [ -n "$(find test -name '*.out')" ]; then
    if [ ! -f "main.go" ]; then
        cp $current/template/main.go .
    fi
    cd $current
    code "$current/$contestDir/main.go"
    make r
    exit
fi

oj d $url

# templateディレクトリにあるmain.goをカレントディレクトリにコピー
cp $current/template/main.go .

echo "$contest" >$current/state
echo "$task" >>$current/state

# 元のディレクトリに移動
cd $current

code "$current/$contestDir/main.go"

make r
