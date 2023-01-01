#!/bin/bash
sed -i -E "s/ (\"version_module\")/#\1/g" rootfs/opt/bitnami/scripts/apache/postunpack.sh || true
find rootfs/opt/bitnami/scripts/apache/postunpack.sh -exec touch -a -m -t 202301010000.00 {} \; || true

