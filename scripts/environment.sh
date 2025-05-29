#!/usr/bin/env bash

export PRJ="godemo"

export ROOT=$(dirname "${BASH_SOURCE[0]}")/..
export LOG_PATH=$HOME/$PRJ/logs
export ETC_PATH=$HOME/$PRJ/etc
export INSTALL_PATH=$HOME/$PRJ/install

mkdir -p "$ETC_PATH"
mkdir -p "$LOG_PATH/datasync"
mkdir -p "$INSTALL_PATH"

DATASYNC_WS_PROVIDER=${DATASYNC_WS_PROVIDER:-"ws://192.168.218.203:62222/api"}
DATASYNC_HTTP_PROVIDER=${DATASYNC_HTTP_PROVIDER:-"http://192.168.218.203:62222/api"}
