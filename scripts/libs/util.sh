#!/usr/bin/env bash

util::kill_by_port() {
  local PORT="$1"
  local PIDS=$(lsof -ti:$PORT | tr '\n' ' ')

  echo "=> detect port :$1 and find process: ${PIDS}"
  if [ -z "$PIDS" ]; then
    return
  fi
  kill -9 $PIDS

  if [ $? -eq 0 ]; then
    echo "=> succeed to terminal process"
  else
    echo "=> failt to terminal process"
  fi
}