#!/bin/bash
sed -i -E "s/#RUN (\. \/opt\/bitnami\/scripts\/libcomponent.sh) && (component_unpack \"${PACKAGE}\")/RUN \1 \\&\\& OS_ARCH=amd64 \2/g" Dockerfile.arm64 
