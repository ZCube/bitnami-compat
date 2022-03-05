#!/bin/bash
export ROOT_DIR=$(realpath "$(dirname $0)/..")
echo $ROOT_DIR

export packages=( $(find bitnami-docker-* -type f ! -name "archive" ! -name  ".git" -name "Dockerfile" | grep "debian-10" | sed -E 's/.*bitnami-docker-([^/]+)\/.*/\1/') )
export versions_major=( $(find bitnami-docker-* -type f ! -name "archive" ! -name  ".git" -name "Dockerfile" | grep "debian-10" | sed -E 's/.*bitnami-docker-[^/]+\/([^/]+)\/.*/\1/') )
export versions=( $(find bitnami-docker-* -type f ! -name "archive" ! -name  ".git" -name "Dockerfile" | xargs -I "{}" -n 1 cat {} | grep BITNAMI_IMAGE_VERSION= | sed -E 's/.*="([0-9.]+)-debian.*/\1/') )

# TODO: use template engine.
echo ${#packages[@]}
echo ${#versions_major[@]}
echo ${#versions[@]}

if [[  ${#packages[@]} == ${#versions_major[@]} &&  ${#packages[@]} == ${#versions[@]}  ]]; then
    cat ${ROOT_DIR}/scripts/build-on-push-header > ${ROOT_DIR}/.github/workflows/build-on-push.yml
    for (( i=0; i<${#packages[@]}; i++ )); do
        echo "${packages[i]}-${versions_major[i]}-${versions[i]}"

        VERSION="${versions[i]}"
        VERSION="${VERSION#[vV]}"
        MAJOR="${VERSION%%\.*}"
        MINOR="${VERSION#*.}"
        MINOR="${MINOR%.*}"
        PATCH="${VERSION##*.}"

        echo ${MAJOR} ${MINOR} ${PATCH}
        PACKAGE=${packages[i]}

        # ALIAS
        if [ "$PACKAGE" = "redis-cluster" ]; then
            PACKAGE="redis"
        fi

        passed=0
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}.${MINOR}.${PATCH}/golang/install.sh"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}.${MINOR}/golang/install.sh"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}/golang/install.sh"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}.${MINOR}.${PATCH}/docker/Dockerfile.install"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}.${MINOR}/docker/Dockerfile.install"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}/docker/Dockerfile.install"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}.${MINOR}.${PATCH}/bash/install.sh"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}.${MINOR}/bash/install.sh"; then
            passed=1
        fi
        if test -f "${ROOT_DIR}/patches/${PACKAGE}/${MAJOR}/bash/install.sh"; then
            passed=1
        fi

        if [ $passed -eq 1 ]; then
            cat ${ROOT_DIR}/scripts/build-on-push-body | \
            sed -E -e "s/\{\{\{APP\}\}\}/${packages[i]}/g" | sed -E -e "s/\{\{\{VERSION\}\}\}/${versions[i]}/g" | sed -E -e "s/\{\{\{VERSION_MAJOR_MINOR\}\}\}/${versions_major_minor[i]}/g" | sed -E -e "s/\{\{\{VERSION_MAJOR\}\}\}/${versions_major[i]}/g" \
            >> ${ROOT_DIR}/.github/workflows/build-on-push.yml
        fi

    done
    cat ${ROOT_DIR}/scripts/build-on-push-tail >> ${ROOT_DIR}/.github/workflows/build-on-push.yml
fi
