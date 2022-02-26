#!/bin/bash
echo $1
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 mkdir -p golang/{}
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 mkdir -p docker/{}
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 mkdir -p bash/{}
mkdir -p archive

export OS_NAME=linux
export OS_ARCH=amd64
export OS_FLAVOUR=debian-10
export DOWNLOAD_URL="https://downloads.bitnami.com/files/stacksmith"
pushd archive
cat ../$(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)-\(.value.version)"' | xargs -I {} -n 1 /bin/bash -c '[ -f "{}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz" ] || curl --remote-name --silent --show-error --fail -L "${DOWNLOAD_URL}/{}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz" -o "{}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz"'
cat ../$(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)-\(.value.version)"' | xargs -I {} -n 1 /bin/bash -c '[ -d "{}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}" ] || tar -zxkvf "{}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz"'
popd
