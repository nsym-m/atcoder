#!/bin/bash

contest=$(head -n 1 state)
task=$(sed -n '2p' state)
default_url="https://atcoder.jp/contests"
url="${default_url}/${contest}/tasks/${task}"

current=$(pwd)

cd "contests/$contest/$task"

# URL付きで問題提出
oj s $url main.go

cd $current
