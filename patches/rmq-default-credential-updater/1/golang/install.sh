#!/bin/bash
set -e
export PACKAGE=rmq-default-credential-updater
export TARGET_DIR=rmq-default-credential-updater
export VERSION={{{VERSION}}}
export REF=v{{{VERSION}}}
export CGO_ENABLED=0 
rm -rf ${PACKAGE} || true
mkdir -p ${PACKAGE}
git clone -b "${REF}" https://github.com/rabbitmq/default-user-credential-updater ${PACKAGE}
pushd ${PACKAGE}
go mod download
go build -v -ldflags '-d -s -w' -tags timetzdata -o default-user-credential-updater .
mkdir -p /opt/bitnami/${TARGET_DIR}/licenses
mkdir -p /opt/bitnami/${TARGET_DIR}/bin
cp -f LICENSE /opt/bitnami/${TARGET_DIR}/licenses/${PACKAGE}-${VERSION}.txt
cp -f default-user-credential-updater /opt/bitnami/${TARGET_DIR}/bin/
popd
rm -rf ${PACKAGE}
