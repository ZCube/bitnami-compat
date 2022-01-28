#!/bin/bash
echo $1
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 mkdir -p golang/{}
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 mkdir -p docker/{}
