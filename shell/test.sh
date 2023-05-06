#!/bin/bash

contest=$(head -n 1 state)
task=$(sed -n '2p' state)

current=$(pwd)

cd "contests/$contest/$task"

oj t -c "./main"

cd $current
