# Bitnami-Compat

[![Inactively Maintained](https://img.shields.io/badge/Maintenance%20Level-Inactively%20Maintained-yellowgreen.svg)](https://gist.github.com/cheerfulstoic/d107229326a01ff0f333a1d3476e068d)
[![Project-Stage](https://img.shields.io/badge/Project%20Stage-Experimental-yellow.svg)](https://img.shields.io/badge/Project%20Stage-Experimental-yellow.svg)
[![Production Ready](https://img.shields.io/badge/Production%20Ready-No-red.svg)](https://img.shields.io/badge/Production%20Ready-No-red.svg)
[![Update bitnami-dockers](https://github.com/ZCube/bitnami-compat/actions/workflows/update.yml/badge.svg)](https://github.com/ZCube/bitnami-compat/actions/workflows/update.yml)
[![Build](https://github.com/ZCube/bitnami-compat/actions/workflows/build-on-push.yml/badge.svg)](https://github.com/ZCube/bitnami-compat/actions/workflows/build-on-push.yml)
[![Testing](https://img.shields.io/badge/Testing-No-red.svg)](https://img.shields.io/badge/Testing-No-red.svg)


## What is Bitnami-Compat?

* This repository builds multi-arch (`linux/amd64` + `linux/arm64`) docker images for some bitnami charts.

* This is a workaround for https://github.com/bitnami/charts/issues/7305.

* These images are NOT intended for production use.
  * This repository cannot be stable. Changes in multiple repositories (apt, bitnami, etc...) are reflected so images can break at any time.
  * This repository is mainly maintained for my homelab and development.

* This repository does not support all bitnami images.
  * I believe that someday bitnami will support arm64.
  * I don't use all images from bitnami.
  * I prefer database or other development tools, not for general apps. (e.g. wordpress, odoo, discourse, ...)
  * For general apps, if arm64 is already supported (not for bitnami), I recommended to use it. This repository's may not reach production ready.
  * If anyone wants to make an image patch themselves, I can help with that. Please make an issue.

* These images use Bitnami's scripts. Binaries refer to the official Docker or official distribution binaries of the software or Docker recipes for Docker.

## Update bitnami-dockers schedule

* Every friday.

* :warning: Recently deleted for old images older than 30 days.
  * I think I'm using too much storage (I'm GitHub Free plan user...), so I'm cleaning it up. In the case of an old image, it may have been deleted, and in that case, a manifest unknown error occurs when docker pull is performed. Since this is a multi-arch image, the layer connected to the tag is deleted and only the manifest remains.

## Usage #1

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
  or
  go run main.go generate --app=mongodb
  ```
  4. verify Dockerfile.arm64. Official images or build scripts referenced for build are listed in the table below.
  5. build & use. (docker buildx needed)
  ```
  task build
  or (current platform only)
  go run main.go build --app=mongodb --tag ghcr.io/zcube/bitnami-compat/
  or (current platform only && push)
  go run main.go build --app=mongodb --tag ghcr.io/zcube/bitnami-compat/ --push
  or (docker buildx needed && push)
  go run main.go build --app=mongodb --tag ghcr.io/zcube/bitnami-compat/ --push --platforms=linux/amd64,linux/arm64 

  docker run --rm -ti ghcr.io/zcube/bitnami-compat/mongodb:6.0
  ```

* PS. Rabbitmq and fluentd builds are too long. Depending on the machine, it is best to prepare for 1-2 hours.

## Usage #2

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
✔️  contour
✔️  contour-operator
✔️  discourse
✔️  elasticsearch
✔️  etcd
✔️  external-dns
✔️  fluentd
✔️  kafka
✔️  keycloak
✔️  mariadb
✔️  mariadb-galera
✔️  memcached
✔️  minio
✔️  mongodb
✔️  mongodb-sharded
✔️  mysql
✔️  nginx
✔️  node
✔️  postgresql
✔️  postgresql-ha
✔️  rabbitmq
✔️  rabbitmq-cluster-operator
✔️  redis
✔️  redis-cluster
✔️  sealed-secrets
✔️  solr
✔️  thanos
✔️  tomcat
✔️  wordpress
✔️  zookeeper
❓  apache
  ❌  apache
  ✔️  git
  ✔️  apache-exporter
❓  nginx-ingress-controller
  ❌  nginx-ingress-controller
  ✔️  nginx
❓  nginx-intel
  ❌  nginx-intel
  ✔️  git
  ✔️  nginx-exporter
❓  wordpress-intel
  ❌  wordpress-intel
  ✔️  bitnami-shell
  ✔️  nginx-exporter
```

* For wordpress-intel charts, try using wordpress-nginx instead of wordpress-intel images. (wordpress-intel = wordpress-nginx)

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

* :warning: For RabbitMQ for ARM64, JIT is disabled for building on AMD64.

* :warning: For some images such as node and python, apache, java, the version can only match the major version. This is mostly the case for non-main apps, and I don't care because they're mostly dependencies for other apps.

* :warning: The mongodb provided by this repository does not currently work on Raspberry Pi 4. Use the image I distribute at https://github.com/ZCube/mongodb-for-armv8.0-a.

