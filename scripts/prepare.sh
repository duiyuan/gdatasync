#!/usr/bin/env bash

set +o errexit
set -o nounset
set +o pipefail

ROOT=$(dirname "${BASH_SOURCE[0]}")/..

GIT_PS1_SHOWCONFLICTSTATE=${GIT_PS1_SHOWCONFLICTSTATE:-"no"}


gen() {
    local f="$1"
    source ${ROOT}/scripts/genconfig.sh ${ROOT}/scripts/environment.sh $ROOT/conf/$f.yaml > $HOME/${PRJ}/etc/$f.yaml

    echo "=> mv $f.yaml to ${HOME}/${PRJ}/etc/$f.yaml"
}