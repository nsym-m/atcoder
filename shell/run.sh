#!/bin/bash

# stateから現在のコンテストの状態を取得
contest=$(head -n 1 state)
task=$(sed -n '2p' state)
current=$(pwd)
contestDir="contests/$contest/$task"

cd $contestDir

open "$current/$contestDir/main.go"

# 対象になっているmain.goの変更を監視 && 変更があったらビルドしてテスト実行
echo main.go | entr -s "goimports -w main.go && go build main.go && cd $current && make t" 2>&1
