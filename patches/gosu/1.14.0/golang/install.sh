#!/bin/bash
export PACKAGE=gosu
export TARGET_DIR=common
export REF=${VERSION_MAJOR_MINOR}
export CGO_ENABLED=0 
rm -rf ${PACKAGE} || true
mkdir -p ${PACKAGE}
git clone -b "${REF}" https://github.com/tianon/gosu ${PACKAGE}
pushd ${PACKAGE}
go mod download
go build -v -ldflags '-d -s -w' .
mkdir -p /opt/bitnami/${TARGET_DIR}/licenses
mkdir -p /opt/bitnami/${TARGET_DIR}/bin
cp -f LICENSE /opt/bitnami/${TARGET_DIR}/licenses/${PACKAGE}-${VERSION}.txt
cp -f ${PACKAGE} /opt/bitnami/${TARGET_DIR}/bin/gosu
popd
rm -rf ${PACKAGE}
