#!/bin/bash

contest=$(head -n 1 state)
task=$(sed -n '2p' state)

current=$(pwd)

cd "contests/$contest/$task"

# 問題提出
oj s main.go

cd $current
