#!/bin/bash

contest=$(head -n 1 state)
task=$(sed -n '2p' state)
current=$(pwd)
contestDir="contests/$contest/$task"

cd $contestDir

code "$current/$contestDir/main.go"

echo main.go | entr -s "go build main.go && cd $current && make t" 2>&1
