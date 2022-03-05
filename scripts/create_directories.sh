#!/bin/bash
echo $1
# Prepare directories
export package_dir=$(dirname $1)
export versions=( $(cat ${package_dir}/prebuildfs/opt/bitnami/.bitnami_components.json | jq  -r 'to_entries[] | "\(.value.version)"' | sed -E -e 's/-([0-9]+)//g') )
export versions_major_minor=( $(cat ${package_dir}/prebuildfs/opt/bitnami/.bitnami_components.json | jq  -r 'to_entries[] | "\(.value.version)"' | sed -E -e 's/-([0-9]+)//g' | sed -E -e 's/\.([^.]+)$//g') )
export versions_major=( $(cat ${package_dir}/prebuildfs/opt/bitnami/.bitnami_components.json | jq  -r 'to_entries[] | "\(.value.version)"' | sed -E -e 's/-([0-9]+)//g' | sed -E -e 's/\.([^.]+)$//g' | sed -E -e 's/\.([^.]+)$//g') )
export versions_with_revision=( $(cat ${package_dir}/prebuildfs/opt/bitnami/.bitnami_components.json | jq  -r 'to_entries[] | "\(.value.version)"') )
export packages=( $(cat ${package_dir}/prebuildfs/opt/bitnami/.bitnami_components.json | jq  -r 'to_entries[] | "\(.key)"') )

for (( i=0; i<${#packages[@]}; i++ )); do
echo "${packages[i]}/${versions[i]}" 
mkdir -p patches/${packages[i]}/${versions[i]}/golang
mkdir -p patches/${packages[i]}/${versions[i]}/docker
mkdir -p patches/${packages[i]}/${versions[i]}/bash
mkdir -p patches/${packages[i]}/${versions_major_minor[i]}/golang
mkdir -p patches/${packages[i]}/${versions_major_minor[i]}/docker
mkdir -p patches/${packages[i]}/${versions_major_minor[i]}/bash
mkdir -p patches/${packages[i]}/${versions_major[i]}/golang
mkdir -p patches/${packages[i]}/${versions_major[i]}/docker
mkdir -p patches/${packages[i]}/${versions_major[i]}/bash
done

# Download archives
mkdir -p archive

export OS_NAME=linux
export OS_ARCH=amd64
export OS_FLAVOUR=debian-10
export DOWNLOAD_URL="https://downloads.bitnami.com/files/stacksmith"
pushd archive
for (( i=0; i<${#packages[@]}; i++ )); do
echo "${packages[i]}-${versions_with_revision[i]}"
if [[ ! -f "${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz" ]]; then
    echo "${DOWNLOAD_URL}/${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz"
    curl --remote-name --silent --show-error --fail -L "${DOWNLOAD_URL}/${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz" -o "${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz"
fi
if [[ -f "${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz" && ! -f "${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz.unzipped" ]]; then
    tar -zxkf "${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz"
    touch "${packages[i]}-${versions_with_revision[i]}-${OS_NAME}-${OS_ARCH}-${OS_FLAVOUR}.tar.gz.unzipped"
fi
done
popd
