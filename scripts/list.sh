#!/bin/bash
cat .github/workflows/build-on-push.yml | yq -r '.jobs."build-bitnami-compat".strategy.matrix.include[]  | "* \(.APP) : [`\(.VERSION_MAJOR)`, `\(.VERSION_MAJOR)-\(.OS_FLAVOUR)`, `\(.VERSION)`, `\(.VERSION)-\(.OS_FLAVOUR)`, `\(.VERSION)-\(.OS_FLAVOUR)-r\(.REVISION)`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2F\(.APP))"'
