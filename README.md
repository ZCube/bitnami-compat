# Bitnami-Compat

[![Inactively Maintained](https://img.shields.io/badge/Maintenance%20Level-Inactively%20Maintained-yellowgreen.svg)](https://gist.github.com/cheerfulstoic/d107229326a01ff0f333a1d3476e068d)
## What is Bitnami-Compat?

* This repository builds multi-arch (`linux/amd64` + `linux/arm64`) docker images for a selection of the Bitnami repositories.

* This is a workaround for https://github.com/bitnami/charts/issues/7305.

* These images are NOT intended for production use.

  * I don't have time to test every images and I haven't built the infrastructure to test them.
  * In other words, this repository has no release management process and no support for it.
  * This repository is mainly maintained for use on my Raspberry Pi cluster.

* These images use Bitnami's scripts. Binaries refer to the official Docker or official distribution binaries of the software or Docker recipes for Docker.

## Binary References

|Package|Type|Reference|
|--|--|--|
|acmesolver|DockerImage|https://quay.io/jetstack/cert-manager-acmesolver|
|cainjector|DockerImage|https://quay.io/jetstack/cert-manager-cainjector|
|cert-manager|DockerImage|https://quay.io/jetstack/cert-manager-controller|
|cert-manager-webhook|DockerImage|https://quay.io/jetstack/cert-manager-webhook|
|consul|DockerImage|https://hub.docker.com/_/consul|
|consul-exporter|Dockerfile|https://hub.docker.com/r/prom/consul-exporter|
|elasticsearch|Install|https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html|
|etcd|Dockerfile|https://gcr.io/etcd-development/etcd|
|fluentd|Dockerfile|https://github.com/fluent/fluentd-docker-image|
|jmx-exporter|Install|https://repo1.maven.org/maven2/io/prometheus/jmx|
|kafka|Install|https://downloads.apache.org/kafka|
|kubectl|Install|https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/|
|keycloak|Dockerfile|https://github.com/keycloak/keycloak-containers|
|mariadb|Dockerfile|https://github.com/MariaDB/mariadb-docker|
|mariadb-galera|Dockerfile|https://github.com/MariaDB/mariadb-docker|
|memcached|Dockerfile|https://github.com/docker-library/memcached|
|minio|DockerImage|https://hub.docker.com/r/minio/minio|
|minio-client|DockerImage|https://hub.docker.com/r/minio/mc|
|mongodb|Dockerfile|https://github.com/docker-library/mongo|
|mysql|DockerImage|https://hub.docker.com/r/mysql|
|mysqld-exporter|Dockerfile|https://hub.docker.com/r/prom/mysqld-exporter|
|nginx|Dockerfile|https://github.com/nginxinc/docker-nginx|
|pgpool|Dockerfile|https://github.com/docker-library/postgres|
|postgresql|Dockerfile|https://github.com/docker-library/postgres|
|postgresql-repmgr|Dockerfile|https://github.com/docker-library/postgres|
|rabbitmq|Dockerfile|https://github.com/docker-library/rabbitmq|
|redis|Dockerfile|https://github.com/docker-library/redis|
|redis-cluster|Dockerfile|https://github.com/docker-library/redis|
|redis-exporter|DockerImage|https://hub.docker.com/r/oliver006/redis_exporter|
|redis-sentinel|Dockerfile|https://github.com/docker-library/redis|
|tomcat|Install|https://downloads.apache.org/tomcat|
|wordpress|||
|zookeeper|Install|https://downloads.apache.org/zookeeper|

## Supported images and tags

* acmesolver : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Facmesolver)
* bitnami-shell : [`11`, `11-debian-11`, `11`, `11-debian-11`, `11-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fbitnami-shell)
* cainjector : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcainjector)
* cert-manager : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager)
* cert-manager-webhook : [`1`, `1-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager-webhook)
* consul : [`1`, `1-debian-11`, `1.12.2`, `1.12.2-debian-11`, `1.12.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul)
* consul-exporter : [`0`, `0-debian-11`, `0.8.0`, `0.8.0-debian-11`, `0.8.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul-exporter)
* elasticsearch : [`7`, `7-debian-11`, `7.17.5`, `7.17.5-debian-11`, `7.17.5-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch)
* elasticsearch : [`8`, `8-debian-11`, `8.3.2`, `8.3.2-debian-11`, `8.3.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch)
* etcd : [`3.3`, `3.3-debian-11`, `3.3.27`, `3.3.27-debian-11`, `3.3.27-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* etcd : [`3.4`, `3.4-debian-11`, `3.4.19`, `3.4.19-debian-11`, `3.4.19-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* etcd : [`3.5`, `3.5-debian-11`, `3.5.4`, `3.5.4-debian-11`, `3.5.4-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* fluentd : [`1`, `1-debian-11`, `1.15.0`, `1.15.0-debian-11`, `1.15.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ffluentd)
* jmx-exporter : [`0`, `0-debian-11`, `0.17.0`, `0.17.0-debian-11`, `0.17.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjmx-exporter)
* kafka : [`3.0`, `3.0-debian-11`, `3.0.1`, `3.0.1-debian-11`, `3.0.1-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka : [`3.1`, `3.1-debian-11`, `3.1.1`, `3.1.1-debian-11`, `3.1.1-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka : [`3.2`, `3.2-debian-11`, `3.2.0`, `3.2.0-debian-11`, `3.2.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* keycloak : [`18`, `18-debian-11`, `18.0.2`, `18.0.2-debian-11`, `18.0.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkeycloak)
* kubectl : [`1.22`, `1.22-debian-11`, `1.22.11`, `1.22.11-debian-11`, `1.22.11-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkubectl)
* kubectl : [`1.23`, `1.23-debian-11`, `1.23.8`, `1.23.8-debian-11`, `1.23.8-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkubectl)
* kubectl : [`1.24`, `1.24-debian-11`, `1.24.2`, `1.24.2-debian-11`, `1.24.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkubectl)
* mariadb : [`10.5`, `10.5-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.6`, `10.6-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.7`, `10.7-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.8`, `10.8-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb-galera : [`10.5`, `10.5-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.6`, `10.6-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.7`, `10.7-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.8`, `10.8-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* memcached : [`1`, `1-debian-11`, `1.6.15`, `1.6.15-debian-11`, `1.6.15-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmemcached)
* minio : [`2022`, `2022-debian-11`, `2022.7.8`, `2022.7.8-debian-11`, `2022.7.8-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio)
* minio-client : [`2022`, `2022-debian-11`, `2022.7.6`, `2022.7.6-debian-11`, `2022.7.6-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio-client)
* mongodb : [`4.2`, `4.2-debian-10`, `4.2.21`, `4.2.21-debian-10`, `4.2.21-debian-10-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mongodb : [`4.4`, `4.4-debian-10`, `4.4.15`, `4.4.15-debian-10`, `4.4.15-debian-10-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mongodb : [`5.0`, `5.0-debian-11`, `5.0.9`, `5.0.9-debian-11`, `5.0.9-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mysql : [`8.0`, `8.0-debian-11`, `8.0.29`, `8.0.29-debian-11`, `8.0.29-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysql)
* mysqld-exporter : [`0`, `0-debian-11`, `0.14.0`, `0.14.0-debian-11`, `0.14.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysqld-exporter)
* nginx : [`1.23`, `1.23-debian-11`, `1.23.0`, `1.23.0-debian-11`, `1.23.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fnginx)
* pgpool : [`4`, `4-debian-11`, `4.3.2`, `4.3.2-debian-11`, `4.3.2-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpgpool)
* postgresql : [`10`, `10-debian-11`, `10.21.0`, `10.21.0-debian-11`, `10.21.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`11`, `11-debian-11`, `11.16.0`, `11.16.0-debian-11`, `11.16.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`12`, `12-debian-11`, `12.11.0`, `12.11.0-debian-11`, `12.11.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`13`, `13-debian-11`, `13.7.0`, `13.7.0-debian-11`, `13.7.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`14`, `14-debian-11`, `14.4.0`, `14.4.0-debian-11`, `14.4.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql-repmgr : [`10`, `10-debian-11`, `10.21.0`, `10.21.0-debian-11`, `10.21.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`11`, `11-debian-11`, `11.16.0`, `11.16.0-debian-11`, `11.16.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`12`, `12-debian-11`, `12.11.0`, `12.11.0-debian-11`, `12.11.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`13`, `13-debian-11`, `13.7.0`, `13.7.0-debian-11`, `13.7.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`14`, `14-debian-11`, `14.4.0`, `14.4.0-debian-11`, `14.4.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* rabbitmq : [`3.10`, `3.10-debian-11`, `3.10.5`, `3.10.5-debian-11`, `3.10.5-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* rabbitmq : [`3.8`, `3.8-debian-11`, `3.8.34`, `3.8.34-debian-11`, `3.8.34-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* rabbitmq : [`3.9`, `3.9-debian-11`, `3.9.20`, `3.9.20-debian-11`, `3.9.20-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* redis : [`6.0`, `6.0-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis : [`7.0`, `7.0-debian-11`, `7.0.3`, `7.0.3-debian-11`, `7.0.3-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis-cluster : [`6.0`, `6.0-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-cluster : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-cluster : [`7.0`, `7.0-debian-11`, `7.0.3`, `7.0.3-debian-11`, `7.0.3-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-exporter : [`1`, `1-debian-11`, `1.43.0`, `1.43.0-debian-11`, `1.43.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-exporter)
* redis-sentinel : [`6.0`, `6.0-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-sentinel : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-sentinel : [`7.0`, `7.0-debian-11`, `7.0.3`, `7.0.3-debian-11`, `7.0.3-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* tomcat : [`10`, `10-debian-11`, `10.0.22`, `10.0.22-debian-11`, `10.0.22-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`8.5`, `8.5-debian-11`, `8.5.81`, `8.5.81-debian-11`, `8.5.81-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`9.0`, `9.0-debian-11`, `9.0.64`, `9.0.64-debian-11`, `9.0.64-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* wordpress : [`5`, `5-debian-11`, `5.9.3`, `5.9.3-debian-11`, `5.9.3-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fwordpress)
* wordpress : [`6`, `6-debian-11`, `6.0.0`, `6.0.0-debian-11`, `6.0.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fwordpress)
* zookeeper : [`3.6`, `3.6-debian-11`, `3.6.3`, `3.6.3-debian-11`, `3.6.3-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.7`, `3.7-debian-11`, `3.7.1`, `3.7.1-debian-11`, `3.7.1-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.8`, `3.8-debian-11`, `3.8.0`, `3.8.0-debian-11`, `3.8.0-debian-11-r11`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)

# TL;DR
from
```yaml
image:
  registry: docker.io
  repository: bitnami/postgresql
  tag: 11.14.0-debian-10-r28
```
to
```yaml
image:
  registry: ghcr.io
  repository: zcube/bitnami-compat/postgresql
  tag: 11.14.0-debian-10
```

## Install Task

* using task https://taskfile.dev/installation/
```bash
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
```

## List docker image patch status

```bash
go run main.go list
```

```
...
(o) mongodb:5.0.9
(x) mysql:5.7.38
 * mysql.5.7.38-150 patch needed
(o) mysql:8.0.29
(o) mysqld-exporter:0.14.0
...
```

## Generate Dockerfile.arm64
```bash
task generate
```

## Upgrade bitnami docker repositories
```bash
task upgrade
```

## How to add a new image

Example : mariadb/10.6.5-3

1. Add submodule
  ```bash
  cd bitnami-dockers
  git submodule add https://github.com/bitnami/bitnami-docker-mariadb
  ```

2. Write build scripts
  * patches/bash/mariadb/10.6.5-3/install.sh
  * patches/docker/mariadb/10.6.5-3/Dockerfile.from
  * patches/docker/mariadb/10.6.5-3/Dockerfile.install
  * patches/golang/mariadb/10.6.5-3/install.sh

3. Apply and test
  ```bash
  task generate
  cd bitnami-dockers/bitnami-docker-mariadb/10.6/debian-11
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

## License for generator

MIT License

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

## Warning

This repository contains sources, binaries and recipes from Bitnami and other software. Each software has its own license.

These images are not well tested.

The images do not automatically receive security patches.

This repository is considered a proof of concept only.

Use at your own risk.
