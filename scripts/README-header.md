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

## Supported bitnami charts

```
✔️  cert-manager
✔️  common
✔️  consul
✔️  etcd
✔️  fluentd
✔️  mariadb
✔️  mariadb-galera
✔️  minio
✔️  mysql
✔️  postgresql
✔️  postgresql-ha
✔️  rabbitmq
✔️  redis
✔️  redis-cluster
✔️  tomcat
✔️  zookeeper
❓  elasticsearch
  ✔️  elasticsearch
  ❌  elasticsearch-exporter
  ✔️  bitnami-shell
  ✔️  bitnami-shell
❓  kafka
  ✔️  kafka
  ✔️  kubectl
  ✔️  bitnami-shell
  ❌  kafka-exporter
  ✔️  jmx-exporter
❓  keycloak
  ✔️  keycloak
  ❌  keycloak-config-cli
❓  memcached
  ✔️  memcached
  ✔️  bitnami-shell
  ❌  memcached-exporter
❓  mongodb
  ✔️  mongodb
  ✔️  nginx
  ✔️  kubectl
  ✔️  bitnami-shell
  ❌  mongodb-exporter
❓  mongodb-sharded
  ❌  mongodb-sharded
  ✔️  bitnami-shell
  ❌  mongodb-exporter
❓  nginx
  ✔️  nginx
  ❌  git
  ❌  nginx-exporter
❓  nginx-ingress-controller
  ❌  nginx-ingress-controller
  ✔️  nginx
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
  ❌  nginx-exporter
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

