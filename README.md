# Bitnami-Compat

[![Inactively Maintained](https://img.shields.io/badge/Maintenance%20Level-Inactively%20Maintained-yellowgreen.svg)](https://gist.github.com/cheerfulstoic/d107229326a01ff0f333a1d3476e068d)
[![Update bitnami-dockers](https://github.com/ZCube/bitnami-compat/actions/workflows/update.yml/badge.svg)](https://github.com/ZCube/bitnami-compat/actions/workflows/update.yml)
[![Build](https://github.com/ZCube/bitnami-compat/actions/workflows/build-on-push.yml/badge.svg)](https://github.com/ZCube/bitnami-compat/actions/workflows/build-on-push.yml)
## What is Bitnami-Compat?

* This repository builds multi-arch (`linux/amd64` + `linux/arm64`) docker images for a selection of the Bitnami repositories.

* This is a workaround for https://github.com/bitnami/charts/issues/7305.

* These images are NOT intended for production use.

  * I don't have time to test every images and I haven't built the infrastructure to test them.
  * In other words, this repository has no release management process and no support for it.
  * This repository is mainly maintained for use on my Raspberry Pi cluster.

* These images use Bitnami's scripts. Binaries refer to the official Docker or official distribution binaries of the software or Docker recipes for Docker.

# Usage #1

* I know I can't make the images in this repository trustworthy, and I know I can't fix it. This repository is just a personal repository.
* If you want to use the images, It is recommended not to use the image directly, but to create Dockerfile.arm64 by following the instructions below and then verify and use it.

* Create Dockerfile.arm64 for bitnami-dockers
  1. install task https://taskfile.dev/installation/
  2. install golang https://go.dev/doc/install
  3. clone
  ```bash
  git clone https://github.com/ZCube/bitnami-compat
  cd bitnami-compat
  task generate
  ```
  4. verify Dockerfile.arm64. Official images or build scripts referenced for build are listed in the table below.
  5. build & use.

* PS. Rabbitmq and fluentd builds are too long. Depending on the machine, it is best to prepare for 1-2 hours.

# Usage #2

* TL;DR

from
```yaml
image:
  registry: docker.io
  repository: bitnami/postgresql
  tag: 11.16.0-debian-11-r14
```
to
```yaml
image:
  registry: ghcr.io
  repository: zcube/bitnami-compat/postgresql
  tag: 11.16.0-debian-11-r15
```

## Supported bitnami charts

```
✔️  cert-manager
✔️  common
✔️  consul
✔️  elasticsearch
✔️  etcd
✔️  fluentd
✔️  kafka
✔️  keycloak
✔️  mariadb
✔️  mariadb-galera
✔️  memcached
✔️  minio
✔️  mongodb
✔️  mysql
✔️  postgresql
✔️  postgresql-ha
✔️  rabbitmq
✔️  redis
✔️  redis-cluster
✔️  sealed-secrets
✔️  tomcat
✔️  zookeeper
❓  mongodb-sharded
  ❌  mongodb-sharded
  ✔️  bitnami-shell
  ✔️  mongodb-exporter
❓  nginx
  ✔️  nginx
  ❌  git
  ✔️  nginx-exporter
❓  nginx-ingress-controller
  ❌  nginx-ingress-controller
  ✔️  nginx
❓  nginx-intel
  ❌  nginx-intel
  ❌  git
  ✔️  nginx-exporter
❓  rabbitmq-cluster-operator
  ✔️  rabbitmq
  ❌  rmq-default-credential-updater
  ❌  rabbitmq-cluster-operator
  ❌  rmq-messaging-topology-operator
❓  wordpress
  ✔️  wordpress
  ✔️  bitnami-shell
  ❌  apache-exporter
❓  wordpress-intel
  ❌  wordpress-intel
  ✔️  bitnami-shell
  ✔️  nginx-exporter
```
## Binary References

|Package|Type|Reference|Base|
|--|--|--|--|
|acmesolver|DockerImage|https://quay.io/jetstack/cert-manager-acmesolver|bitnami/minideb|
|cainjector|DockerImage|https://quay.io/jetstack/cert-manager-cainjector|bitnami/minideb|
|cert-manager|DockerImage|https://quay.io/jetstack/cert-manager-controller|bitnami/minideb|
|cert-manager-webhook|DockerImage|https://quay.io/jetstack/cert-manager-webhook|bitnami/minideb|
|consul|DockerImage|https://hub.docker.com/_/consul|bitnami/minideb|
|consul-exporter|Dockerfile|https://hub.docker.com/r/prom/consul-exporter|bitnami/minideb|
|elasticsearch|Install|https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html|bitnami/minideb|
|etcd|Dockerfile|https://gcr.io/etcd-development/etcd|bitnami/minideb|
|fluentd|Dockerfile|https://github.com/fluent/fluentd-docker-image|bitnami/minideb|
|jmx-exporter|Install|https://repo1.maven.org/maven2/io/prometheus/jmx|bitnami/minideb|
|kafka|Install|https://downloads.apache.org/kafka|bitnami/minideb|
|kubectl|Install|https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/|bitnami/minideb|
|keycloak|Dockerfile|https://github.com/keycloak/keycloak-containers|bitnami/minideb|
|mariadb|Dockerfile|https://github.com/MariaDB/mariadb-docker|bitnami/minideb|
|mariadb-galera|Dockerfile|https://github.com/MariaDB/mariadb-docker|bitnami/minideb|
|memcached|Dockerfile|https://github.com/docker-library/memcached|bitnami/minideb|
|minio|DockerImage|https://hub.docker.com/r/minio/minio|bitnami/minideb|
|minio-client|DockerImage|https://hub.docker.com/r/minio/mc|bitnami/minideb|
|mongodb|Dockerfile|https://github.com/docker-library/mongo|bitnami/minideb|
|mysql|DockerImage|https://hub.docker.com/r/mysql|mysql/mysql-server|
|mysqld-exporter|Dockerfile|https://hub.docker.com/r/prom/mysqld-exporter|bitnami/minideb|
|nginx|Dockerfile|https://github.com/nginxinc/docker-nginx|bitnami/minideb|
|pgpool|Dockerfile|https://github.com/docker-library/postgres|bitnami/minideb|
|postgresql|Dockerfile|https://github.com/docker-library/postgres|bitnami/minideb|
|postgresql-repmgr|Dockerfile|https://github.com/docker-library/postgres|bitnami/minideb|
|rabbitmq|Dockerfile|https://github.com/docker-library/rabbitmq|bitnami/minideb|
|redis|Dockerfile|https://github.com/docker-library/redis|bitnami/minideb|
|redis-cluster|Dockerfile|https://github.com/docker-library/redis|bitnami/minideb|
|redis-exporter|DockerImage|https://hub.docker.com/r/oliver006/redis_exporter|bitnami/minideb|
|redis-sentinel|Dockerfile|https://github.com/docker-library/redis|bitnami/minideb|
|tomcat|Install|https://downloads.apache.org/tomcat|bitnami/minideb|
|wordpress|||bitnami/minideb|
|zookeeper|Install|https://downloads.apache.org/zookeeper|bitnami/minideb|

## Supported images and tags

* acmesolver : [`1`, `1-debian-11`, `1.9.1`, `1.9.1-debian-11`, `1.9.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Facmesolver)
* apache : [`2.4`, `2.4-debian-11`, `2.4.54`, `2.4.54-debian-11`, `2.4.54-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fapache)
* bitnami-shell : [`11`, `11-debian-11`, `11`, `11-debian-11`, `11-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fbitnami-shell)
* cainjector : [`1`, `1-debian-11`, `1.9.1`, `1.9.1-debian-11`, `1.9.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcainjector)
* cert-manager : [`1`, `1-debian-11`, `1.9.1`, `1.9.1-debian-11`, `1.9.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager)
* cert-manager-webhook : [`1`, `1-debian-11`, `1.9.1`, `1.9.1-debian-11`, `1.9.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fcert-manager-webhook)
* consul : [`1`, `1-debian-11`, `1.12.3`, `1.12.3-debian-11`, `1.12.3-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul)
* consul-exporter : [`0`, `0-debian-11`, `0.8.0`, `0.8.0-debian-11`, `0.8.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fconsul-exporter)
* elasticsearch : [`7`, `7-debian-11`, `7.17.5`, `7.17.5-debian-11`, `7.17.5-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch)
* elasticsearch : [`8`, `8-debian-11`, `8.3.3`, `8.3.3-debian-11`, `8.3.3-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch)
* elasticsearch-exporter : [`1`, `1-debian-11`, `1.5.0`, `1.5.0-debian-11`, `1.5.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Felasticsearch-exporter)
* etcd : [`3.3`, `3.3-debian-11`, `3.3.27`, `3.3.27-debian-11`, `3.3.27-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* etcd : [`3.4`, `3.4-debian-11`, `3.4.19`, `3.4.19-debian-11`, `3.4.19-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* etcd : [`3.5`, `3.5-debian-11`, `3.5.4`, `3.5.4-debian-11`, `3.5.4-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fetcd)
* fluentd : [`1`, `1-debian-11`, `1.15.1`, `1.15.1-debian-11`, `1.15.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ffluentd)
* java : [`1.8`, `1.8-debian-11`, `1.8.342`, `1.8.342-debian-11`, `1.8.342-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjava)
* java : [`11`, `11-debian-11`, `11.0.15-1`, `11.0.15-1-debian-11`, `11.0.15-1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjava)
* java : [`17`, `17-debian-11`, `17.0.3-1`, `17.0.3-1-debian-11`, `17.0.3-1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjava)
* java : [`18`, `18-debian-11`, `18.0.1-1`, `18.0.1-1-debian-11`, `18.0.1-1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjava)
* jmx-exporter : [`0`, `0-debian-11`, `0.17.0`, `0.17.0-debian-11`, `0.17.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fjmx-exporter)
* kafka : [`3.0`, `3.0-debian-11`, `3.0.1`, `3.0.1-debian-11`, `3.0.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka : [`3.1`, `3.1-debian-11`, `3.1.1`, `3.1.1-debian-11`, `3.1.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka : [`3.2`, `3.2-debian-11`, `3.2.0`, `3.2.0-debian-11`, `3.2.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka)
* kafka-exporter : [`1`, `1-debian-11`, `1.4.2`, `1.4.2-debian-11`, `1.4.2-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkafka-exporter)
* keycloak : [`18`, `18-debian-11`, `18.0.2`, `18.0.2-debian-11`, `18.0.2-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkeycloak)
* keycloak-config-cli : [`5`, `5-debian-11`, `5.2.2`, `5.2.2-debian-11`, `5.2.2-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkeycloak-config-cli)
* kubectl : [`1.22`, `1.22-debian-11`, `1.22.12`, `1.22.12-debian-11`, `1.22.12-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkubectl)
* kubectl : [`1.23`, `1.23-debian-11`, `1.23.9`, `1.23.9-debian-11`, `1.23.9-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkubectl)
* kubectl : [`1.24`, `1.24-debian-11`, `1.24.3`, `1.24.3-debian-11`, `1.24.3-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fkubectl)
* mariadb : [`10.5`, `10.5-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.6`, `10.6-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.7`, `10.7-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb : [`10.8`, `10.8-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb)
* mariadb-galera : [`10.5`, `10.5-debian-11`, `10.5.16`, `10.5.16-debian-11`, `10.5.16-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.6`, `10.6-debian-11`, `10.6.8`, `10.6.8-debian-11`, `10.6.8-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.7`, `10.7-debian-11`, `10.7.4`, `10.7.4-debian-11`, `10.7.4-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* mariadb-galera : [`10.8`, `10.8-debian-11`, `10.8.3`, `10.8.3-debian-11`, `10.8.3-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmariadb-galera)
* memcached : [`1`, `1-debian-11`, `1.6.15`, `1.6.15-debian-11`, `1.6.15-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmemcached)
* memcached-exporter : [`0`, `0-debian-11`, `0.10.0`, `0.10.0-debian-11`, `0.10.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmemcached-exporter)
* minio : [`2022`, `2022-debian-11`, `2022.7.26`, `2022.7.26-debian-11`, `2022.7.26-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio)
* minio-client : [`2022`, `2022-debian-11`, `2022.7.24`, `2022.7.24-debian-11`, `2022.7.24-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fminio-client)
* mongodb : [`5.0`, `5.0-debian-11`, `5.0.10`, `5.0.10-debian-11`, `5.0.10-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb)
* mongodb-exporter : [`0`, `0-debian-11`, `0.33.0`, `0.33.0-debian-11`, `0.33.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb-exporter)
* mongodb-sharded : [`5.0`, `5.0-debian-11`, `5.0.10`, `5.0.10-debian-11`, `5.0.10-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmongodb-sharded)
* mysql : [`8.0`, `8.0-debian-11`, `8.0.30`, `8.0.30-debian-11`, `8.0.30-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysql)
* mysqld-exporter : [`0`, `0-debian-11`, `0.14.0`, `0.14.0-debian-11`, `0.14.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fmysqld-exporter)
* nginx : [`1.23`, `1.23-debian-11`, `1.23.1`, `1.23.1-debian-11`, `1.23.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fnginx)
* nginx-exporter : [`0`, `0-debian-11`, `0.10.0`, `0.10.0-debian-11`, `0.10.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fnginx-exporter)
* pgpool : [`4`, `4-debian-11`, `4.3.2`, `4.3.2-debian-11`, `4.3.2-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpgpool)
* php-fpm : [`7.4`, `7.4-debian-11`, `7.4.30`, `7.4.30-debian-11`, `7.4.30-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fphp-fpm)
* php-fpm : [`8.0`, `8.0-debian-11`, `8.0.21`, `8.0.21-debian-11`, `8.0.21-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fphp-fpm)
* php-fpm : [`8.1`, `8.1-debian-11`, `8.1.8`, `8.1.8-debian-11`, `8.1.8-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fphp-fpm)
* postgres-exporter : [`0`, `0-debian-11`, `0.10.1`, `0.10.1-debian-11`, `0.10.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgres-exporter)
* postgresql : [`10`, `10-debian-11`, `10.21.0`, `10.21.0-debian-11`, `10.21.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`11`, `11-debian-11`, `11.16.0`, `11.16.0-debian-11`, `11.16.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`12`, `12-debian-11`, `12.11.0`, `12.11.0-debian-11`, `12.11.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`13`, `13-debian-11`, `13.7.0`, `13.7.0-debian-11`, `13.7.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql : [`14`, `14-debian-11`, `14.4.0`, `14.4.0-debian-11`, `14.4.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql)
* postgresql-repmgr : [`10`, `10-debian-11`, `10.21.0`, `10.21.0-debian-11`, `10.21.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`11`, `11-debian-11`, `11.16.0`, `11.16.0-debian-11`, `11.16.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`12`, `12-debian-11`, `12.11.0`, `12.11.0-debian-11`, `12.11.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`13`, `13-debian-11`, `13.7.0`, `13.7.0-debian-11`, `13.7.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* postgresql-repmgr : [`14`, `14-debian-11`, `14.4.0`, `14.4.0-debian-11`, `14.4.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fpostgresql-repmgr)
* rabbitmq : [`3.10`, `3.10-debian-11`, `3.10.6`, `3.10.6-debian-11`, `3.10.6-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* rabbitmq : [`3.9`, `3.9-debian-11`, `3.9.21`, `3.9.21-debian-11`, `3.9.21-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Frabbitmq)
* redis : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis : [`7.0`, `7.0-debian-11`, `7.0.4`, `7.0.4-debian-11`, `7.0.4-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis)
* redis-cluster : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-cluster : [`7.0`, `7.0-debian-11`, `7.0.4`, `7.0.4-debian-11`, `7.0.4-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-cluster)
* redis-exporter : [`1`, `1-debian-11`, `1.43.0`, `1.43.0-debian-11`, `1.43.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-exporter)
* redis-sentinel : [`6.2`, `6.2-debian-11`, `6.2.7`, `6.2.7-debian-11`, `6.2.7-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* redis-sentinel : [`7.0`, `7.0-debian-11`, `7.0.4`, `7.0.4-debian-11`, `7.0.4-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fredis-sentinel)
* ruby : [`2.7`, `2.7-debian-11`, `2.7.6`, `2.7.6-debian-11`, `2.7.6-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fruby)
* ruby : [`3.1`, `3.1-debian-11`, `3.1.2`, `3.1.2-debian-11`, `3.1.2-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fruby)
* sealed-secrets : [`0`, `0-scratch`, `0.18.1-1`, `0.18.1-1-scratch`, `0.18.1-1-scratch-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fsealed-secrets)
* tomcat : [`10`, `10-debian-11`, `10.0.23`, `10.0.23-debian-11`, `10.0.23-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`8.5`, `8.5-debian-11`, `8.5.81`, `8.5.81-debian-11`, `8.5.81-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* tomcat : [`9.0`, `9.0-debian-11`, `9.0.65`, `9.0.65-debian-11`, `9.0.65-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Ftomcat)
* wordpress : [`6`, `6-debian-11`, `6.0.1`, `6.0.1-debian-11`, `6.0.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fwordpress)
* wordpress-nginx : [`6`, `6-debian-11`, `6.0.1`, `6.0.1-debian-11`, `6.0.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fwordpress-nginx)
* zookeeper : [`3.6`, `3.6-debian-11`, `3.6.3`, `3.6.3-debian-11`, `3.6.3-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.7`, `3.7-debian-11`, `3.7.1`, `3.7.1-debian-11`, `3.7.1-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)
* zookeeper : [`3.8`, `3.8-debian-11`, `3.8.0`, `3.8.0-debian-11`, `3.8.0-debian-11-r19`](https://github.com/zcube/bitnami-compat/pkgs/container/bitnami-compat%2Fzookeeper)

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
✔️  mongodb:5.0.9
✔️  mongodb-exporter:0.33.0
❌  mysql:5.7.38
  ❌  mysql.5.7.38-150
✔️  mysql:8.0.29
✔️  mysqld-exporter:0.14.0
❌  nginx:1.21.6
  ❌  nginx.1.21.6-153
❌  nginx:1.22.0
  ❌  nginx.1.22.0-153
✔️  nginx:1.23.0
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

Example : mariadb/10.8

1. Add submodule
  ```bash
  cd bitnami-dockers
  git submodule add https://github.com/bitnami/bitnami-docker-mariadb
  cd ..
  go run main.go makeDirs
  ```

2. Write build scripts
  * patches/golang/mariadb/10.8/install.sh
  * patches/docker/mariadb/10.8/Dockerfile.from
  * patches/docker/mariadb/10.8/Dockerfile.install
  * patches/bash/mariadb/10.8/install.sh

3. Apply and test
  ```bash
  task generate
  cd bitnami-dockers/bitnami-docker-mariadb/10.8/debian-11
  docker buildx create --name multiarchbuilder
  docker buildx inspect multiarchbuilder --bootstrap
  docker buildx use multiarchbuilder
  docker buildx build --platform linux/amd64,linux/arm64 -t testimage . -f Dockerfile.arm64
  docker run --rm -ti testimage
  ```

## Recommneded Images for Bitnami Charts

|Charts|Image|Reason|
|------|-----|------|
|thanos|docker.io/thanosio/thanos:v0.27.0|Official images are compatible.|

## License

Apache 2.0

## License for generator (Golang scripts)

MIT License

## License for patches

check References

## License for Bitnami docker scripts

Apache 2.0

## Warning

This repository contains sources, binaries and recipes from Bitnami and other software. Each software has its own license.

These images are not well tested.

This repository is considered for personal use only. 

This repository lacks support for production use.

This repository is provided as is.

It might be dangerous for production use. Use after testing
