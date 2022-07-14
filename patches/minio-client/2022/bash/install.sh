#!/bin/bash
apt update
apt install -y wget
echo ${VERSION_MAJOR}-0?${VERSION_MINOR}-0?${VERSION_PATCH}
export TAG=$(wget -q https://registry.hub.docker.com/v1/repositories/minio/mc/tags -O -  | sed -e 's/[][]//g' -e 's/"//g' -e 's/ //g' | tr '}' '\n'  | awk -F: '{print $3}' | grep -E ${VERSION_MAJOR}-0?${VERSION_MINOR}-0?${VERSION_PATCH} | grep -v .fips | sort | tail -n 1)
echo ${TAG}
sed -i "s/MC_VERSIONSTRING/$TAG/g" Dockerfile.arm64
