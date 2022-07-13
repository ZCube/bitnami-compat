# bitnami-compat
## What is Bitnami-Compat?
This repository builds multi-arch (`linux/amd64` + `linux/arm64`) docker images for a selection of the Bitnami repositories.
This is a workaround for https://github.com/bitnami/charts/issues/7305.

These images are NOT intended for production use.

These images use Bitnami's scripts. Binaries refer to the official Docker or official distribution binaries of the software or Docker recipes for Docker.

## Supported images and tags

* docker-acmesolver : [`bitnami-docker-acmesolver`, `bitnami-docker-acmesolver-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-acmesolver)
* docker-cainjector : [`bitnami-docker-cainjector`, `bitnami-docker-cainjector-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-cainjector)
* docker-cert-manager : [`bitnami-docker-cert-manager`, `bitnami-docker-cert-manager-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-cert-manager)
* docker-cert-manager-webhook : [`bitnami-docker-cert-manager-webhook`, `bitnami-docker-cert-manager-webhook-debian-11`, `1.8.2`, `1.8.2-debian-11`, `1.8.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-cert-manager-webhook)
* docker-consul : [`bitnami-docker-consul`, `bitnami-docker-consul-debian-11`, `1.12.2`, `1.12.2-debian-11`, `1.12.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-consul)
* docker-consul-exporter : [`bitnami-docker-consul-exporter`, `bitnami-docker-consul-exporter-debian-11`, `0.8.0`, `0.8.0-debian-11`, `0.8.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-consul-exporter)
* docker-elasticsearch : [`bitnami-docker-elasticsearch`, `bitnami-docker-elasticsearch-debian-11`, `7.17.5`, `7.17.5-debian-11`, `7.17.5-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-elasticsearch)
* docker-elasticsearch : [`bitnami-docker-elasticsearch`, `bitnami-docker-elasticsearch-debian-11`, `8.3.2`, `8.3.2-debian-11`, `8.3.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-elasticsearch)
* docker-etcd : [`bitnami-docker-etcd`, `bitnami-docker-etcd-debian-11`, `3.3.27`, `3.3.27-debian-11`, `3.3.27-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-etcd)
* docker-etcd : [`bitnami-docker-etcd`, `bitnami-docker-etcd-debian-11`, `3.4.18`, `3.4.18-debian-11`, `3.4.18-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-etcd)
* docker-etcd : [`bitnami-docker-etcd`, `bitnami-docker-etcd-debian-11`, `3.5.4`, `3.5.4-debian-11`, `3.5.4-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-etcd)
* docker-fluentd : [`bitnami-docker-fluentd`, `bitnami-docker-fluentd-debian-11`, `1.15.0`, `1.15.0-debian-11`, `1.15.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-fluentd)
* docker-jmx-exporter : [`bitnami-docker-jmx-exporter`, `bitnami-docker-jmx-exporter-debian-11`, `0.17.0`, `0.17.0-debian-11`, `0.17.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-jmx-exporter)
* docker-kafka : [`bitnami-docker-kafka`, `bitnami-docker-kafka-debian-11`, `3.0.1`, `3.0.1-debian-11`, `3.0.1-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-kafka)
* docker-kafka : [`bitnami-docker-kafka`, `bitnami-docker-kafka-debian-11`, `3.1.1`, `3.1.1-debian-11`, `3.1.1-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-kafka)
* docker-kafka : [`bitnami-docker-kafka`, `bitnami-docker-kafka-debian-11`, `3.2.0`, `3.2.0-debian-11`, `3.2.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-kafka)
* docker-keycloak : [`bitnami-docker-keycloak`, `bitnami-docker-keycloak-debian-11`, `18.0.2`, `18.0.2-debian-11`, `18.0.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-keycloak)
* docker-mariadb : [`bitnami-docker-mariadb`, `bitnami-docker-mariadb-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb)
* docker-mariadb : [`bitnami-docker-mariadb`, `bitnami-docker-mariadb-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb)
* docker-mariadb : [`bitnami-docker-mariadb`, `bitnami-docker-mariadb-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb)
* docker-mariadb : [`bitnami-docker-mariadb`, `bitnami-docker-mariadb-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb)
* docker-mariadb-galera : [`bitnami-docker-mariadb-galera`, `bitnami-docker-mariadb-galera-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb-galera)
* docker-mariadb-galera : [`bitnami-docker-mariadb-galera`, `bitnami-docker-mariadb-galera-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb-galera)
* docker-mariadb-galera : [`bitnami-docker-mariadb-galera`, `bitnami-docker-mariadb-galera-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb-galera)
* docker-mariadb-galera : [`bitnami-docker-mariadb-galera`, `bitnami-docker-mariadb-galera-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mariadb-galera)
* docker-memcached : [`bitnami-docker-memcached`, `bitnami-docker-memcached-debian-11`, `1.6.15`, `1.6.15-debian-11`, `1.6.15-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-memcached)
* docker-minio : [`bitnami-docker-minio`, `bitnami-docker-minio-debian-11`, `2022.7.8`, `2022.7.8-debian-11`, `2022.7.8-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-minio)
* docker-minio-client : [`bitnami-docker-minio-client`, `bitnami-docker-minio-client-debian-11`, `2022.7.6`, `2022.7.6-debian-11`, `2022.7.6-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-minio-client)
* docker-mongodb : [`bitnami-docker-mongodb`, `bitnami-docker-mongodb-debian-10`, `4.2.21`, `4.2.21-debian-10`, `4.2.21-debian-10-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mongodb)
* docker-mongodb : [`bitnami-docker-mongodb`, `bitnami-docker-mongodb-debian-10`, `4.4.15`, `4.4.15-debian-10`, `4.4.15-debian-10-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mongodb)
* docker-mongodb : [`bitnami-docker-mongodb`, `bitnami-docker-mongodb-debian-11`, `5.0.9`, `5.0.9-debian-11`, `5.0.9-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mongodb)
* docker-mysql : [`bitnami-docker-mysql`, `bitnami-docker-mysql-debian-11`, `8.0.29`, `8.0.29-debian-11`, `8.0.29-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mysql)
* docker-mysqld-exporter : [`bitnami-docker-mysqld-exporter`, `bitnami-docker-mysqld-exporter-debian-11`, `0.14.0`, `0.14.0-debian-11`, `0.14.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-mysqld-exporter)
* docker-nginx : [`bitnami-docker-nginx`, `bitnami-docker-nginx-debian-11`, `1.23.0`, `1.23.0-debian-11`, `1.23.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-nginx)
* docker-pgpool : [`bitnami-docker-pgpool`, `bitnami-docker-pgpool-debian-11`, `4.3.2`, `4.3.2-debian-11`, `4.3.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-pgpool)
* docker-postgresql : [`bitnami-docker-postgresql`, `bitnami-docker-postgresql-debian-11`, `10.21.0`, `10.21.0-debian-11`, `10.21.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql)
* docker-postgresql : [`bitnami-docker-postgresql`, `bitnami-docker-postgresql-debian-11`, `11.16.0`, `11.16.0-debian-11`, `11.16.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql)
* docker-postgresql : [`bitnami-docker-postgresql`, `bitnami-docker-postgresql-debian-11`, `12.11.0`, `12.11.0-debian-11`, `12.11.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql)
* docker-postgresql : [`bitnami-docker-postgresql`, `bitnami-docker-postgresql-debian-11`, `13.7.0`, `13.7.0-debian-11`, `13.7.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql)
* docker-postgresql : [`bitnami-docker-postgresql`, `bitnami-docker-postgresql-debian-11`, `14.4.0`, `14.4.0-debian-11`, `14.4.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql)
* docker-postgresql-repmgr : [`bitnami-docker-postgresql-repmgr`, `bitnami-docker-postgresql-repmgr-debian-11`, `10.21.0`, `10.21.0-debian-11`, `10.21.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql-repmgr)
* docker-postgresql-repmgr : [`bitnami-docker-postgresql-repmgr`, `bitnami-docker-postgresql-repmgr-debian-11`, `11.16.0`, `11.16.0-debian-11`, `11.16.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql-repmgr)
* docker-postgresql-repmgr : [`bitnami-docker-postgresql-repmgr`, `bitnami-docker-postgresql-repmgr-debian-11`, `12.11.0`, `12.11.0-debian-11`, `12.11.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql-repmgr)
* docker-postgresql-repmgr : [`bitnami-docker-postgresql-repmgr`, `bitnami-docker-postgresql-repmgr-debian-11`, `13.7.0`, `13.7.0-debian-11`, `13.7.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql-repmgr)
* docker-postgresql-repmgr : [`bitnami-docker-postgresql-repmgr`, `bitnami-docker-postgresql-repmgr-debian-11`, `14.4.0`, `14.4.0-debian-11`, `14.4.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-postgresql-repmgr)
* docker-rabbitmq : [`bitnami-docker-rabbitmq`, `bitnami-docker-rabbitmq-debian-11`, `3.10.5`, `3.10.5-debian-11`, `3.10.5-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-rabbitmq)
* docker-rabbitmq : [`bitnami-docker-rabbitmq`, `bitnami-docker-rabbitmq-debian-11`, `3.8.34`, `3.8.34-debian-11`, `3.8.34-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-rabbitmq)
* docker-rabbitmq : [`bitnami-docker-rabbitmq`, `bitnami-docker-rabbitmq-debian-11`, `3.9.20`, `3.9.20-debian-11`, `3.9.20-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-rabbitmq)
* docker-redis : [`bitnami-docker-redis`, `bitnami-docker-redis-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis)
* docker-redis : [`bitnami-docker-redis`, `bitnami-docker-redis-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis)
* docker-redis : [`bitnami-docker-redis`, `bitnami-docker-redis-debian-11`, `7.0.2`, `7.0.2-debian-11`, `7.0.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis)
* docker-redis-cluster : [`bitnami-docker-redis-cluster`, `bitnami-docker-redis-cluster-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis-cluster)
* docker-redis-cluster : [`bitnami-docker-redis-cluster`, `bitnami-docker-redis-cluster-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis-cluster)
* docker-redis-cluster : [`bitnami-docker-redis-cluster`, `bitnami-docker-redis-cluster-debian-11`, `7.0.3`, `7.0.3-debian-11`, `7.0.3-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis-cluster)
* docker-redis-exporter : [`bitnami-docker-redis-exporter`, `bitnami-docker-redis-exporter-debian-11`, `1.43.0`, `1.43.0-debian-11`, `1.43.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis-exporter)
* docker-redis-sentinel : [`bitnami-docker-redis-sentinel`, `bitnami-docker-redis-sentinel-debian-11`, `6.0.16`, `6.0.16-debian-11`, `6.0.16-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis-sentinel)
* docker-redis-sentinel : [`bitnami-docker-redis-sentinel`, `bitnami-docker-redis-sentinel-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis-sentinel)
* docker-redis-sentinel : [`bitnami-docker-redis-sentinel`, `bitnami-docker-redis-sentinel-debian-11`, `7.0.2`, `7.0.2-debian-11`, `7.0.2-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-redis-sentinel)
* docker-tomcat : [`bitnami-docker-tomcat`, `bitnami-docker-tomcat-debian-11`, `10.0.22`, `10.0.22-debian-11`, `10.0.22-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-tomcat)
* docker-tomcat : [`bitnami-docker-tomcat`, `bitnami-docker-tomcat-debian-11`, `8.5.81`, `8.5.81-debian-11`, `8.5.81-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-tomcat)
* docker-tomcat : [`bitnami-docker-tomcat`, `bitnami-docker-tomcat-debian-11`, `9.0.64`, `9.0.64-debian-11`, `9.0.64-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-tomcat)
* docker-wordpress : [`bitnami-docker-wordpress`, `bitnami-docker-wordpress-debian-11`, `5.9.3`, `5.9.3-debian-11`, `5.9.3-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-wordpress)
* docker-wordpress : [`bitnami-docker-wordpress`, `bitnami-docker-wordpress-debian-11`, `6.0.0`, `6.0.0-debian-11`, `6.0.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-wordpress)
* docker-zookeeper : [`bitnami-docker-zookeeper`, `bitnami-docker-zookeeper-debian-11`, `3.6.3`, `3.6.3-debian-11`, `3.6.3-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-zookeeper)
* docker-zookeeper : [`bitnami-docker-zookeeper`, `bitnami-docker-zookeeper-debian-11`, `3.7.1`, `3.7.1-debian-11`, `3.7.1-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-zookeeper)
* docker-zookeeper : [`bitnami-docker-zookeeper`, `bitnami-docker-zookeeper-debian-11`, `3.8.0`, `3.8.0-debian-11`, `3.8.0-debian-11-r9`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fdocker-zookeeper)

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
