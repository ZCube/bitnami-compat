#!/bin/bash
export ROOT_DIR=$(realpath "$(dirname $0)/..")
echo $ROOT_DIR
find bitnami-docker-* -type f ! -name "archive" ! -name ".git" -name "Dockerfile" | xargs -I "{}" -n 1 $(dirname "$0")/create_directories.sh {}
