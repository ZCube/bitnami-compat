#!/bin/bash
export ROOT_DIR=$(realpath "$(dirname $0)/..")
echo $ROOT_DIR

export packages=( $(find bitnami-docker-* -type f ! -name "archive" ! -name  ".git" -name "Dockerfile" | grep "debian-10" | sed -E 's/.*bitnami-docker-([^/]+)\/.*/\1/') )
export major_versions=( $(find bitnami-docker-* -type f ! -name "archive" ! -name  ".git" -name "Dockerfile" | grep "debian-10" | sed -E 's/.*bitnami-docker-[^/]+\/([^/]+)\/.*/\1/') )
export versions=( $(find bitnami-docker-* -type f ! -name "archive" ! -name  ".git" -name "Dockerfile" | xargs -I "{}" -n 1 cat {} | grep BITNAMI_IMAGE_VERSION= | sed -E 's/.*="([0-9.]+)-debian.*/\1/') )

# TODO: use template engine.
echo ${#packages[@]}
echo ${#major_versions[@]}
echo ${#versions[@]}

if [[  ${#packages[@]} == ${#major_versions[@]} &&  ${#packages[@]} == ${#versions[@]}  ]]; then
    cat ${ROOT_DIR}/scripts/build-on-push-header > ${ROOT_DIR}/.github/workflows/build-on-push.yml
    for (( i=0; i<${#packages[@]}; i++ )); do
        echo "${packages[i]}-${major_versions[i]}-${versions[i]}"
        cat ${ROOT_DIR}/scripts/build-on-push-body | \
        sed -E -e "s/\{\{\{APP\}\}\}/${packages[i]}/g" | sed -E -e "s/\{\{\{VERSION\}\}\}/${versions[i]}/g" | sed -E -e "s/\{\{\{VERSION_MAJOR_MINOR\}\}\}/${versions_major_minor[i]}/g" | sed -E -e "s/\{\{\{VERSION_MAJOR\}\}\}/${versions_major[i]}/g" \
        >> ${ROOT_DIR}/.github/workflows/build-on-push.yml
    done
    cat ${ROOT_DIR}/scripts/build-on-push-tail >> ${ROOT_DIR}/.github/workflows/build-on-push.yml
fi
