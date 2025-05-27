#!/usr/bin/env bash

ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

source "$ROOT/scripts/environment.sh"
source "$ROOT/scripts/prepare.sh"

gen_conf_file "datasync"