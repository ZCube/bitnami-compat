#!/bin/bash
set -e
export PACKAGE=sealed-secrets
export TARGET_DIR=sealed-secrets
export VERSION={{{VERSION}}}
export REF=v${VERSION}
export CGO_ENABLED=0 
rm -rf ${PACKAGE} || true
mkdir -p ${PACKAGE}
git clone -b "${REF}" https://github.com/bitnami-labs/sealed-secrets ${PACKAGE}
pushd ${PACKAGE}
go mod download
make all
#go build -v -ldflags '-d -s -w' .
mkdir -p /opt/bitnami/${TARGET_DIR}/licenses
mkdir -p /opt/bitnami/${TARGET_DIR}/bin
cp -f LICENSE /opt/bitnami/${TARGET_DIR}/licenses/${PACKAGE}-${VERSION}.txt
cp -f controller /opt/bitnami/${TARGET_DIR}/bin/controller
cp -f kubeseal /opt/bitnami/${TARGET_DIR}/bin/kubeseal
popd
