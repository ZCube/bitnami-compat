#!/bin/bash
if test -f "${ROOT_DIR}/docker/$1/Dockerfile.install"; then
    cat "${ROOT_DIR}/docker/$1/Dockerfile.install"
fi