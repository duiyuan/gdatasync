#!/usr/bin/env bash

set +o errexit
set -o nounset
set +o pipefail

ROOT=$(dirname "${BASH_SOURCE[0]}")/../..

GIT_PS1_SHOWCONFLICTSTATE=${GIT_PS1_SHOWCONFLICTSTATE:-"no"}


wwt::gen() {
    server="$1"
    source ${ROOT}/scripts/genconfig.sh ${ROOT}/scripts/environment.sh $ROOT/conf/${server}.yaml > $HOME/${PRJ}/etc/${server}.yaml

    bee::log::info "=> mv ${server}.yaml to ${HOME}/${PRJ}/etc/${server}.yaml"
}