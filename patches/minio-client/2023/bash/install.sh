#!/bin/bash
apt update
apt install -y wget jq
echo ${VERSION_MAJOR}-0?${VERSION_MINOR}-0?${VERSION_PATCH}
export TAG=$(wget -q https://registry.hub.docker.com/v2/repositories/minio/mc/tags -O -  | jq  -r '.results[].name' | grep -E ${VERSION_MAJOR}-0?${VERSION_MINOR}-0?${VERSION_PATCH} | grep -v .fips | sort | tail -n 1)
echo ${TAG}
sed -i "s/MC_VERSIONSTRING/$TAG/g" Dockerfile.arm64
