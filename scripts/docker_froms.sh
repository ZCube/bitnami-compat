#!/bin/bash
if test -f "${ROOT_DIR}/docker/$1/Dockerfile.from"; then
    cat "${ROOT_DIR}/docker/$1/Dockerfile.from"
fi