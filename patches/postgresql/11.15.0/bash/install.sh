#!/bin/bash
sed -i "s/psql --version/psql --version | awk '{print \$3}' /g" rootfs/opt/bitnami/scripts/libpostgresql.sh