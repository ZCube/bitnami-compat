#!/bin/bash
export ROOT_DIR=$(realpath "$(dirname $0)/..")
echo $ROOT_DIR
find  -maxdepth 1 -type d | grep bitnami-docker | xargs -I "{}" /bin/bash -c 'cd {} && git checkout .'
find  -maxdepth 1 -type d | grep bitnami-docker | xargs -I "{}" /bin/bash -c "$ROOT_DIR/scripts/update.sh {}"

