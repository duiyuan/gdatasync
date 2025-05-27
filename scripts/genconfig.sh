#!/usr/bin/env bash
# shellcheck disable="SC2013"

ENV_FILE="$1"
TEMP_FILE="$2"

if [ $# -ne 2 ];then
    echo "Usage: genconfig.sh scripts/environment.sh conf/xxx.yaml"
    exit 1
fi

# shellcheck source=/dev/null
source "$ENV_FILE"

set +u

for env in $(sed -n 's/^[^#].*${\(.*\)}.*/\1/p' "$TEMP_FILE"); do
    if [[ -z "$(eval echo \$${env})" && "${env}" != *"REDIS"* ]];then
        echo "environment variable '${env}' not set"
        missing=true
    fi
done

if [ "${missing}" ];then
    echo "You may run 'source scripts/environment.sh' to set these environment"
    exit 1
fi

eval "cat << EOF
$(cat "$TEMP_FILE")
EOF"
