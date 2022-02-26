#!/bin/bash
export ROOT_DIR=$(realpath "$(dirname $0)/..")
echo $ROOT_DIR


/bin/bash -c "cd bitnami-docker-$1 && git checkout ."
/bin/bash -c "$ROOT_DIR/scripts/update.sh bitnami-docker-$1"

