#!/bin/bash
sed -i "s/psql --version/psql --version | awk '{print \$3}' /g" rootfs/opt/bitnami/scripts/libpostgresql.sh
find rootfs/opt/bitnami/scripts/libpostgresql.sh -exec touch -a -m -t 202301010000.00 {} \;
