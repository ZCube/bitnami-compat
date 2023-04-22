#!/bin/bash
set -e

export VERSION_MAJOR=${VERSION_MAJOR:-10}
export VERSION_DIR=${VERSION_DIR:-./}
export REPOSITORY="deb http://apt.postgresql.org/pub/repos/apt/ bullseye-pgdg main $VERSION_MAJOR"

set -ex
key='B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8'
export GNUPGHOME="$(mktemp -d)"
gpg --batch --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys "$key"
gpg --batch --export "$key" > /etc/apt/trusted.gpg.d/postgres.gpg
command -v gpgconf > /dev/null && gpgconf --kill all
rm -rf "$GNUPGHOME"
apt-key list

set -e
echo "${REPOSITORY}" > /etc/apt/sources.list.d/pgdg.list

apt -qq update

export VERSION_ORIGINAL=$(apt-cache policy postgresql-${VERSION_MAJOR} | grep Candidate | awk '{print $2}' | sed -E "s/(.*)[.].*$/\\1/")
cat ${VERSION_DIR}.bitnami_components.json | jq -M --arg VERSION_ORIGINAL "$VERSION_ORIGINAL" '.postgresql.version = $VERSION_ORIGINAL' > /tmp/components.json
cp -rf /tmp/components.json ${VERSION_DIR}.bitnami_components.json
find ${VERSION_DIR}.bitnami_components.json -exec touch -a -m -t 202301010000.00 {} \;
