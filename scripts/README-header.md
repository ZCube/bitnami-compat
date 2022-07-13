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

