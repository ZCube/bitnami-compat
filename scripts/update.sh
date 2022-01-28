#!/bin/bash
find $1 -type f | grep Dockerfile | xargs -I "{}" -n 1 sed -i 's/amd64/${TARGETARCH}/g' {}
