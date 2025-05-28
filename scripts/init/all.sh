#!/usr/bin/env bash

set -u 
set -e

ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

app=datasync

for a in $app; do 
    # shellcheck source=/dev/null
    source "$ROOT/scripts/init/$a.sh"    
done;