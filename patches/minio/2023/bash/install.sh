#!/bin/bash
echo ${VERSION_MAJOR}-0?${VERSION_MINOR}-0?${VERSION_PATCH}
export TAG=$(wget -q 'https://registry.hub.docker.com/v2/repositories/minio/minio/tags?page_size=30' -O -  | jq  -r '.results[].name' | grep -E ${VERSION_MAJOR}-0?${VERSION_MINOR}-0?${VERSION_PATCH} | grep -v .fips | grep -v hotfix | sort | tail -n 1)
echo ${TAG}
sed -i "s/MINIO_VERSIONSTRING/$TAG/g" Dockerfile.arm64
