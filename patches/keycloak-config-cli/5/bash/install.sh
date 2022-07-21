#!/bin/bash
sed -i "s/RUN wget -nc /#RUN wget -nc /g" Dockerfile.arm64
sed -i "s/echo /#echo /g" Dockerfile.arm64
sed -i "s/tar -zxf \/tmp\/bitnami/#tar -zxf \/tmp\/bitnami/g" Dockerfile.arm64
sed -i "s/rm -rf \/tmp\/bitnami/#rm -rf \/tmp\/bitnami/g" Dockerfile.arm64
sed -i -E "s/#(.*)keycloak-config-cli/\1keycloak-config-cli/g" Dockerfile.arm64
