#!/bin/bash
# sed -i -E "s/^php_major_version=.*/php_major_version=\"\$\(php -v \| grep \^PHP \| cut -d\' \' -f2 \| cut -d\'.\' -f1\).\$\(php -v \| grep \^PHP \| cut -d\' \' -f2 \| cut -d\'.\' -f2\)\"/g" rootfs/opt/bitnami/scripts/apache-modphp/postunpack.sh || true
