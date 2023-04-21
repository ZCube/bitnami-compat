#/bin/bash
set -e

export VERSION=${VERSION:-10.5}
export VERSION_DIR=${VERSION_DIR:-./}
export GPG_KEYS=177F4010FE56CA3336300305F1656F24C74CD1D8
export REPOSITORY="http://archive.mariadb.org/mariadb-${VERSION}/repo/debian/ bullseye main"

set -ex; \
export GNUPGHOME="$(mktemp -d)"; \
for key in $GPG_KEYS; do \
    gpg --batch --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys "$key"; \
done; \
gpg --batch --export $GPG_KEYS > /etc/apt/trusted.gpg.d/mariadb.gpg; \
command -v gpgconf > /dev/null && gpgconf --kill all || :; \
rm -fr "$GNUPGHOME"; \
apt-key list

set -e;\
echo "deb ${REPOSITORY}" > /etc/apt/sources.list.d/mariadb.list; \
{ \
    echo 'Package: *'; \
    echo 'Pin: release o=MariaDB'; \
    echo 'Pin-Priority: 999'; \
} > /etc/apt/preferences.d/mariadb

apt -qq update

export VERSION_ORIGINAL=$(apt-cache policy mariadb-server | grep Candidate | awk '{print $2}' | cut  -d ':' -f2 | cut  -d '+' -f1)
cat ${VERSION_DIR}.bitnami_components.json | jq -M --arg VERSION_ORIGINAL "$VERSION_ORIGINAL" '."mariadb-galera".version = $VERSION_ORIGINAL' > /tmp/components.json
cp -rf /tmp/components.json ${VERSION_DIR}.bitnami_components.json
