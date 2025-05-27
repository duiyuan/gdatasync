#!/usr/bin/env bash

set -u
set -e
set -o pipefail


ROOT=$(dirname "${BASH_SOURCE[0]}")/../..


source "$ROOT/scripts/init/datasync.sh"

echo "=> binary builder start..."

output=$INSTALL_PATH/datasync
go build -o $output $ROOT/cmd/datasync/main.go 

echo "=> run program with nohup"
nohup $output > /dev/null 2>&1 &