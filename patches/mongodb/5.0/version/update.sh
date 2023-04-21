#/bin/bash
set -e

export VERSION=${VERSION:-5.0}
export VERSION_DIR=${VERSION_DIR:-./}
export GPG_KEYS=F5679A222C647C87527C2F8CB00A0BD1E2C63C11

export MONGO_PACKAGE=mongodb-org
export MONGO_REPO=repo.mongodb.org
export MONGO_MAJOR=${VERSION_MAJOR_MINOR}

export REPOSITORY="http://$MONGO_REPO/apt/ubuntu focal/${MONGO_PACKAGE%-unstable}/$MONGO_MAJOR multiverse"

set -ex; \
export GNUPGHOME="$(mktemp -d)"; \
for key in $GPG_KEYS; do \
    gpg --batch --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys "$key"; \
done; \
gpg --batch --export $GPG_KEYS > /etc/apt/trusted.gpg.d/mongodb.gpg; \
command -v gpgconf > /dev/null && gpgconf --kill all || :; \
rm -fr "$GNUPGHOME"; \
apt-key list

set -e;\
echo "deb [ signed-by=/etc/apt/trusted.gpg.d/mongodb.gpg ] http://$MONGO_REPO/apt/ubuntu focal/${MONGO_PACKAGE%-unstable}/$MONGO_MAJOR multiverse" | tee "/etc/apt/sources.list.d/${MONGO_PACKAGE%-unstable}.list"

apt update
# apt -qq update

export VERSION_ORIGINAL=$(apt-cache policy mongodb-org-server | grep Candidate | awk '{print $2}' | cut  -d ':' -f2 | cut  -d '+' -f1)
cat ${VERSION_DIR}.bitnami_components.json | jq -M --arg VERSION_ORIGINAL "$VERSION_ORIGINAL" '.mongodb.version = $VERSION_ORIGINAL' > /tmp/components.json
cp -rf /tmp/components.json ${VERSION_DIR}.bitnami_components.json
