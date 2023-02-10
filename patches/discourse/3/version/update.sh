#/bin/bash
set -e

export VERSION_MAJOR=${VERSION_MAJOR:-10}
export VERSION_DIR=${VERSION_DIR:-./}

export VERSION_ORIGINAL=16.0.0-0
cat ${VERSION_DIR}.bitnami_components.json | jq -M --arg VERSION_ORIGINAL "$VERSION_ORIGINAL" '.node.version = $VERSION_ORIGINAL' > /tmp/components.json
cp -rf /tmp/components.json ${VERSION_DIR}.bitnami_components.json
