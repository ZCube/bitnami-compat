#!/bin/bash
sed -i "s/docker\.io\/bitnami\/minideb:bullseye/docker.io\/mysql:${VERSION}-oracle/g" Dockerfile.arm64
sed -i "s/apt-get update -qq/microdnf update -y/g" prebuildfs/usr/sbin/install_packages
sed -i "s/apt-get install -y --no-install-recommends/microdnf install -y/g" prebuildfs/usr/sbin/install_packages
sed -i "s/rm -r \/var\/lib\/apt\/lists \/var\/cache\/apt\/archives/microdnf clean all/g" prebuildfs/usr/sbin/install_packages
sed -i "s/RUN install_packages .*/RUN install_packages acl ca-certificates curl gzip procps psmisc tar/g" Dockerfile.arm64
sed -i "s/RUN apt-get update/#RUN apt-get update/g" Dockerfile.arm64
sed -i "s/rm -r \/var\/lib\/apt/#rm -r \/var\/lib\/apt/g" Dockerfile.arm64
sed -i "s/RUN mkdir \/docker-entrypoint-initdb.d/RUN mkdir -p \/docker-entrypoint-initdb.d/g" Dockerfile.arm64
