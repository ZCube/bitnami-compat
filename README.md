# bitnami-compat
## What is Bitnami-Compat?
This repository builds multi-arch (`linux/amd64` + `linux/arm64`) docker images for a selection of the Bitnami repositories.
This is a workaround for https://github.com/bitnami/charts/issues/7305.

These images are NOT intended for production use.

These images use Bitnami's scripts. Binaries refer to the official Docker or official distribution binaries of the software or Docker recipes for Docker.


## Supported images and tags

* etcd : [`3`, `3-debian-10`, `3.5.1`, `3.5.1-debian-10`, `3.5.1-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* bitnami-shell : [`10`, `10-debian-10`, `10`, `10-debian-10`, `10-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fbitnami-shell)
* consul : [`1`, `1-debian-10`, `1.11.2`, `1.11.2-debian-10`, `1.11.2-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul)
* postgresql : [`10`, `10-debian-10`, `10.20.0`, `10.20.0-debian-10`, `10.20.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`11`, `11-debian-10`, `11.15.0`, `11.15.0-debian-10`, `11.15.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`12`, `12-debian-10`, `12.10.0`, `12.10.0-debian-10`, `12.10.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`13`, `13-debian-10`, `13.6.0`, `13.6.0-debian-10`, `13.6.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`14`, `14-debian-10`, `14.2.0`, `14.2.0-debian-10`, `14.2.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* mariadb : [`10.3`, `10.3-debian-10`, `10.3.32`, `10.3.32-debian-10`, `10.3.32-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.4`, `10.4-debian-10`, `10.4.22`, `10.4.22-debian-10`, `10.4.22-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.5`, `10.5-debian-10`, `10.5.13`, `10.5.13-debian-10`, `10.5.13-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.6`, `10.6-debian-10`, `10.6.5`, `10.6.5-debian-10`, `10.6.5-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb-galera : [`10.3`, `10.3-debian-10`, `10.3.32`, `10.3.32-debian-10`, `10.3.32-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.4`, `10.4-debian-10`, `10.4.22`, `10.4.22-debian-10`, `10.4.22-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.5`, `10.5-debian-10`, `10.5.13`, `10.5.13-debian-10`, `10.5.13-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.6`, `10.6-debian-10`, `10.6.5`, `10.6.5-debian-10`, `10.6.5-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mysql : [`8.0`, `8.0-debian-10`, `8.0.28`, `8.0.28-debian-10`, `8.0.28-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysql) (not debian. oraclelinux)
* rabbitmq : [`3.9`, `3.9-debian-10`, `3.9.13`, `3.9.13-debian-10`, `3.9.13-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* minio : [`2022`, `2022-debian-10`, `2022.1.8`, `2022.1.8-debian-10`, `2022.1.8-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio)
* minio-client : [`2022`, `2022-debian-10`, `2022.1.25`, `2022.1.25-debian-10`, `2022.1.25-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio-client)
* mysqld-exporter : [`0`, `0-debian-10`, `0.13.0`, `0.13.0-debian-10`, `0.13.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysqld-exporter)
* consul-exporter : [`0`, `0-debian-10`, `0.7.1`, `0.7.1-debian-10`, `0.7.1-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul-exporter)
* cert-manager : [`1`, `1-debian-10`, `1.7.0`, `1.7.0-debian-10`, `1.7.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager)
* cert-manager-webhook : [`1`, `1-debian-10`, `1.7.0`, `1.7.0-debian-10`, `1.7.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager-webhook)
* acmesolver : [`1`, `1-debian-10`, `1.7.0`, `1.7.0-debian-10`, `1.7.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Facmesolver)
* cainjector : [`1`, `1-debian-10`, `1.7.0`, `1.7.0-debian-10`, `1.7.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcainjector)
* redis : [`5.0`, `5.0-debian-10`, `5.0.14`, `5.0.14-debian-10`, `5.0.14-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis : [`6.0`, `6.0-debian-10`, `6.0.16`, `6.0.16-debian-10`, `6.0.16-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis : [`6.2`, `6.2-debian-10`, `6.2.6`, `6.2.6-debian-10`, `6.2.6-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis-sentinel : [`5.0`, `5.0-debian-10`, `5.0.14`, `5.0.14-debian-10`, `5.0.14-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-sentinel : [`6.0`, `6.0-debian-10`, `6.0.16`, `6.0.16-debian-10`, `6.0.16-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-sentinel : [`6.2`, `6.2-debian-10`, `6.2.6`, `6.2.6-debian-10`, `6.2.6-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-cluster : [`5.0`, `5.0-debian-10`, `5.0.14`, `5.0.14-debian-10`, `5.0.14-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-cluster : [`6.0`, `6.0-debian-10`, `6.0.16`, `6.0.16-debian-10`, `6.0.16-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-cluster : [`6.2`, `6.2-debian-10`, `6.2.6`, `6.2.6-debian-10`, `6.2.6-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-exporter : [`1`, `1-debian-10`, `1.34.1`, `1.34.1-debian-10`, `1.34.1-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-exporter)
* tomcat : [`8.5`, `8.5-debian-10`, `8.5.75`, `8.5.75-debian-10`, `8.5.75-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`9.0`, `9.0-debian-10`, `9.0.58`, `9.0.58-debian-10`, `9.0.58-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`10`, `10-debian-10`, `10.0.16`, `10.0.16-debian-10`, `10.0.16-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* elasticsearch : [`7`, `7-debian-10`, `7.16.3`, `7.16.3-debian-10`, `7.16.3-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch)
* kafka : [`2`, `2-debian-10`, `2.8.1`, `2.8.1-debian-10`, `2.8.1-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka : [`3`, `3-debian-10`, `3.1.0`, `3.1.0-debian-10`, `3.1.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* keycloak : [`16`, `16-debian-10`, `16.1.1`, `16.1.1-debian-10`, `16.1.1-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkeycloak)
* memcached : [`1`, `1-debian-10`, `1.6.13`, `1.6.13-debian-10`, `1.6.13-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmemcached)
* fluentd : [`1`, `1-debian-10`, `1.14.4`, `1.14.4-debian-10`, `1.14.4-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ffluentd)
* jmx-exporter : [`0`, `0-debian-10`, `0.16.1`, `0.16.1-debian-10`, `0.16.1-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjmx-exporter)
* zookeeper : [`3.5`, `3.5-debian-10`, `3.5.9`, `3.5.9-debian-10`, `3.5.9-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.6`, `3.6-debian-10`, `3.6.3`, `3.6.3-debian-10`, `3.6.3-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.7`, `3.7-debian-10`, `3.7.0`, `3.7.0-debian-10`, `3.7.0-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* mongodb : [`4.2`, `4.2-debian-10`, `4.2.18`, `4.2.18-debian-10`, `4.2.18-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mongodb : [`4.4`, `4.4-debian-10`, `4.4.12`, `4.4.12-debian-10`, `4.4.12-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mongodb : [`5.0`, `5.0-debian-10`, `5.0.6`, `5.0.6-debian-10`, `5.0.6-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb) (MongoDB 5.0+ requires ARMv8.2-A or higher or a CPU with AVX support)
* nginx : [`1.20`, `1.20-debian-10`, `1.20.2`, `1.20.2-debian-10`, `1.20.2-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fnginx)
* nginx : [`1.21`, `1.21-debian-10`, `1.21.6`, `1.21.6-debian-10`, `1.21.6-debian-10-r0`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fnginx)

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
