
# TL;DR
from
```
image:
  registry: docker.io
  repository: bitnami/postgresql
  tag: 11.14.0-debian-10-r28
```
to
```
image:
  registry: ghcr.io
  repository: zcube/bitnami-compat/postgresql
  tag: 11.14.0-debian-10
```

## How to add a new image

Example : mariadb/10.6.5-3

1. Add submodule
  ```
  git submodule add https://github.com/bitnami/bitnami-docker-mariadb
  ```

2. Create directories and check filelist
  ```
  chmod +x scripts/create.sh
  ./scripts/create.sh
  ```
  * Check file list on archive/mariadb-10.6.5-3-linux-amd64-debian-10

3. Write build scripts
  * bash/mariadb/10.6.5-3/install.sh
  * docker/mariadb/10.6.5-3/Dockerfile.from
  * docker/mariadb/10.6.5-3/Dockerfile.install
  * golang/mariadb/10.6.5-3/install.sh

4. Apply and test
  ```bash
  ./scripts/updates_with_checkout.sh 
  cd bitnami-docker-mariadb/10.6/debian-10
  docker buildx create --name multiarchbuilder
  docker buildx inspect multiarchbuilder --bootstrap
  docker buildx use multiarchbuilder
  docker buildx build --platform linux/amd64,linux/arm64 -t bitnami-test .
  docker run --rm -ti bitnami-test
  ```

## License

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

## Warning

This repository contains sources, binaries and recipes from Bitnami and other software. Each software has its own license.

These images are not well tested.

The images do not automatically receive security patches.

This repository is considered a proof of concept only.

Use at your own risk.
