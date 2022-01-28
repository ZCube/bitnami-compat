#!/bin/bash
export ROOT_DIR=$(realpath "$(dirname $0)/..")
echo $ROOT_DIR
find $1 -type f | grep "[/]Dockerfile$" | xargs -I "{}" -n 1 $(dirname "$0")/create_directories.sh {}
