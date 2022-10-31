#!/bin/bash
sed -i -E "s/ (\"version_module\")/#\1/g" rootfs/opt/bitnami/scripts/apache/postunpack.sh || true