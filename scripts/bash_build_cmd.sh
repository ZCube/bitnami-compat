#!/bin/bash
if test -f "${ROOT_DIR}/bash/$1/install.sh"; then
    if test -d "$2"; then
        pushd $2
            bash "${ROOT_DIR}/bash/$1/install.sh"
        popd
    fi
fi