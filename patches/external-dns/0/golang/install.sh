#!/bin/bash
set -e
export PACKAGE=external-dns
export TARGET_DIR=external-dns
export VERSION={{{VERSION}}}
export REF={{{VERSION}}}
export CGO_ENABLED=0 
rm -rf ${PACKAGE} || true
mkdir -p ${PACKAGE}
git clone -b "v${REF}" https://github.com/kubernetes-sigs/external-dns ${PACKAGE}
pushd ${PACKAGE}
make
cp build/external-dns .
mkdir -p /opt/bitnami/${TARGET_DIR}/licenses
mkdir -p /opt/bitnami/${TARGET_DIR}/bin
cp -f LICENSE /opt/bitnami/${TARGET_DIR}/licenses/${PACKAGE}-${VERSION}.txt
cp -f ${PACKAGE} /opt/bitnami/${TARGET_DIR}/bin/external-dns
popd
rm -rf ${PACKAGE}
