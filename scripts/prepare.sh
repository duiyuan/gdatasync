#!/usr/bin/env bash

set +o errexit
set -o nounset
set +o pipefail

GIT_PS1_SHOWCONFLICTSTATE=${GIT_PS1_SHOWCONFLICTSTATE:-"no"}


gen_conf_file() {
    local appname="$1"
   
    # shellcheck source=/dev/null
    source "$ROOT/scripts/genconfig.sh" "${ROOT}/scripts/environment.sh" "$ROOT/conf/$appname.yaml" > "$HOME/$PRJ/etc/$PRJ-$appname.yaml"
    echo "=> success to make YAML for $appname, $appname.yaml in ${HOME}/$PRJ/etc/$PRJ-$appname.yaml"
}