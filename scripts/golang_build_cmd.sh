#!/bin/bash
if test -f "${ROOT_DIR}/golang/$1/install.sh"; then
    cat "${ROOT_DIR}/golang/$1/install.sh"
fi