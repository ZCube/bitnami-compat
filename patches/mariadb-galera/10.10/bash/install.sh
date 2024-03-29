#!/bin/bash
# https://mariadb.com/kb/en/galera-cluster-system-variables/#wsrep_replicate_myisam
# over 10.7
cat << EOF > rootfs/opt/bitnami/scripts/mariadb-galera/patch.sh
#!/bin/bash
sed -i -e "s/wsrep_replicate_myisam=ON/wsrep_mode=REPLICATE_MYISAM/g" "\${DB_CONF_DIR}/my.cnf"
EOF
chmod +x rootfs/opt/bitnami/scripts/mariadb-galera/patch.sh
sed -i -e "s/  mysql_install_db/  \/opt\/bitnami\/scripts\/mariadb-galera\/patch.sh;mysql_install_db/g" rootfs/opt/bitnami/scripts/libmariadbgalera.sh
sed -i -e "s/export LD_PRELOAD/export LD_PRELOAD;\/opt\/bitnami\/scripts\/mariadb-galera\/patch.sh/g" rootfs/opt/bitnami/scripts/mariadb-galera/run.sh

find rootfs/opt/bitnami/scripts/mariadb-galera/patch.sh -exec touch -a -m -t 202301010000.00 {} \;
find rootfs/opt/bitnami/scripts/libmariadbgalera.sh -exec touch -a -m -t 202301010000.00 {} \;
find rootfs/opt/bitnami/scripts/mariadb-galera/run.sh -exec touch -a -m -t 202301010000.00 {} \;
