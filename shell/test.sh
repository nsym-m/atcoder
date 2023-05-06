#!/bin/bash

contest=$(head -n 1 state)
task=$(sed -n '2p' state)

current=$(pwd)

cd "contests/$contest/$task"

# テスト実行
oj t -c "./main"

cd $current
