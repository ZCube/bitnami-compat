# bitnami-compat
## What is Bitnami-Compat?
This repository builds multi-arch (`linux/amd64` + `linux/arm64`) docker images for a selection of the Bitnami repositories.
This is a workaround for https://github.com/bitnami/charts/issues/7305.

These images are NOT intended for production use.

These images use Bitnami's scripts. Binaries refer to the official Docker or official distribution binaries of the software or Docker recipes for Docker.


## Supported images and tags

* acmesolver : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Facmesolver)
* cainjector : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcainjector)
* cert-manager : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager)
* cert-manager-webhook : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager-webhook)
* consul : [`1`, `1-debian-11`, `1.12.2`, `1.12.2-debian-11`, `1.12.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul)
* consul-exporter : [`0`, `0-debian-11`, `0.8.0`, `0.8.0-debian-11`, `0.8.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul-exporter)
* elasticsearch : [`7`, `7-debian-11`, `7.17.5`, `7.17.5-debian-11`, `7.17.5-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch)
* elasticsearch : [`8`, `8-debian-11`, `8.3.2`, `8.3.2-debian-11`, `8.3.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch)
* etcd : [`3.3`, `3.3-debian-11`, `3.3.27`, `3.3.27-debian-11`, `3.3.27-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* etcd : [`3.4`, `3.4-debian-11`, `3.4.18`, `3.4.18-debian-11`, `3.4.18-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* etcd : [`3.5`, `3.5-debian-11`, `3.5.4`, `3.5.4-debian-11`, `3.5.4-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* fluentd : [`1`, `1-debian-11`, `1.15.0`, `1.15.0-debian-11`, `1.15.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ffluentd)
* jmx-exporter : [`0`, `0-debian-11`, `0.17.0`, `0.17.0-debian-11`, `0.17.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjmx-exporter)
* kafka : [`3.0`, `3.0-debian-11`, `3.0.1`, `3.0.1-debian-11`, `3.0.1-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka : [`3.1`, `3.1-debian-11`, `3.1.1`, `3.1.1-debian-11`, `3.1.1-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka : [`3.2`, `3.2-debian-11`, `3.2.0`, `3.2.0-debian-11`, `3.2.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* keycloak : [`18`, `18-debian-11`, `18.0.2`, `18.0.2-debian-11`, `18.0.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkeycloak)
* mariadb : [`10.5`, `10.5-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.6`, `10.6-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.7`, `10.7-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.8`, `10.8-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb-galera : [`10.5`, `10.5-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.6`, `10.6-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.7`, `10.7-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.8`, `10.8-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* memcached : [`1`, `1-debian-11`, `1.6.15`, `1.6.15-debian-11`, `1.6.15-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmemcached)
* minio : [`2022`, `2022-debian-11`, `2022.7.8`, `2022.7.8-debian-11`, `2022.7.8-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio)
* minio-client : [`2022`, `2022-debian-11`, `2022.7.6`, `2022.7.6-debian-11`, `2022.7.6-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio-client)
* mongodb : [`4.2`, `4.2-debian-10`, `4.2.21`, `4.2.21-debian-10`, `4.2.21-debian-10-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mongodb : [`4.4`, `4.4-debian-10`, `4.4.15`, `4.4.15-debian-10`, `4.4.15-debian-10-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mongodb : [`5.0`, `5.0-debian-11`, `5.0.9`, `5.0.9-debian-11`, `5.0.9-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mysql : [`8.0`, `8.0-debian-11`, `8.0.29`, `8.0.29-debian-11`, `8.0.29-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysql)
* mysqld-exporter : [`0`, `0-debian-11`, `0.14.0`, `0.14.0-debian-11`, `0.14.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysqld-exporter)
* nginx : [`1.23`, `1.23-debian-11`, `1.23.0`, `1.23.0-debian-11`, `1.23.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fnginx)
* postgresql : [`10`, `10-debian-11`, `10.21.0`, `10.21.0-debian-11`, `10.21.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`11`, `11-debian-11`, `11.16.0`, `11.16.0-debian-11`, `11.16.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`12`, `12-debian-11`, `12.11.0`, `12.11.0-debian-11`, `12.11.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`13`, `13-debian-11`, `13.7.0`, `13.7.0-debian-11`, `13.7.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`14`, `14-debian-11`, `14.4.0`, `14.4.0-debian-11`, `14.4.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* rabbitmq : [`3.10`, `3.10-debian-11`, `3.10.5`, `3.10.5-debian-11`, `3.10.5-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* rabbitmq : [`3.8`, `3.8-debian-11`, `3.8.34`, `3.8.34-debian-11`, `3.8.34-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* rabbitmq : [`3.9`, `3.9-debian-11`, `3.9.20`, `3.9.20-debian-11`, `3.9.20-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* redis : [`6.0`, `6.0-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis : [`7.0`, `7.0-debian-11`, `7.0.2`, `7.0.2-debian-11`, `7.0.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis-cluster : [`6.0`, `6.0-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-cluster : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-cluster : [`7.0`, `7.0-debian-11`, `7.0.3`, `7.0.3-debian-11`, `7.0.3-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-exporter : [`1`, `1-debian-11`, `1.43.0`, `1.43.0-debian-11`, `1.43.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-exporter)
* redis-sentinel : [`6.0`, `6.0-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-sentinel : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-sentinel : [`7.0`, `7.0-debian-11`, `7.0.2`, `7.0.2-debian-11`, `7.0.2-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* tomcat : [`10`, `10-debian-11`, `10.0.22`, `10.0.22-debian-11`, `10.0.22-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`8.5`, `8.5-debian-11`, `8.5.81`, `8.5.81-debian-11`, `8.5.81-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`9.0`, `9.0-debian-11`, `9.0.64`, `9.0.64-debian-11`, `9.0.64-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* wordpress : [`5`, `5-debian-11`, `5.9.3`, `5.9.3-debian-11`, `5.9.3-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fwordpress)
* wordpress : [`6`, `6-debian-11`, `6.0.0`, `6.0.0-debian-11`, `6.0.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fwordpress)
* zookeeper : [`3.6`, `3.6-debian-11`, `3.6.3`, `3.6.3-debian-11`, `3.6.3-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.7`, `3.7-debian-11`, `3.7.1`, `3.7.1-debian-11`, `3.7.1-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.8`, `3.8-debian-11`, `3.8.0`, `3.8.0-debian-11`, `3.8.0-debian-11-r5`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)

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
