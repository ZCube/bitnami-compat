# Bitnami-Compat

[![Inactively Maintained](https://img.shields.io/badge/Maintenance%20Level-Inactively%20Maintained-yellowgreen.svg)](https://gist.github.com/cheerfulstoic/d107229326a01ff0f333a1d3476e068d)
[![Project-Stage](https://img.shields.io/badge/Project%20Stage-Development-yellowgreen.svg)](https://img.shields.io/badge/Project%20Stage-Development-yellowgreen.svg)
[![Production Ready](https://img.shields.io/badge/Production%20Ready-No-red.svg)](https://img.shields.io/badge/Production%20Ready-No-red.svg)
[![Update bitnami-dockers](https://github.com/ZCube/bitnami-compat/actions/workflows/update.yml/badge.svg)](https://github.com/ZCube/bitnami-compat/actions/workflows/update.yml)
[![Build](https://github.com/ZCube/bitnami-compat/actions/workflows/build-on-push.yml/badge.svg)](https://github.com/ZCube/bitnami-compat/actions/workflows/build-on-push.yml)
[![Chart Testing](https://github.com/ZCube/bitnami-compat/actions/workflows/test-charts.yaml/badge.svg)](https://github.com/ZCube/bitnami-compat/actions/workflows/test-charts.yaml)

## What is Bitnami-Compat?

* Build https://github.com/bitnami/containers locally ( without Bitnami's StackSmith binaries )

### Goal

* [x] Build https://github.com/bitnami/containers 
  * [x] Support Arm64.
  * [x] Comment StackSmith download scripts.
  * [x] Patch with component-based build script.
  * [x] Synchronize with Bitnami containers version.
  * [x] Build and deploy using Github Actions.
  * [x] Listing supported images.

* [x] Support https://github.com/bitnami/charts on Arm64
  * [x] Listing supported charts.
  * [x] Helm chart testing using Github Actions.
    * [x] Goss
    * [x] Cypress
    * [ ] Ginko
  * [x] Helm chart test with multiple images.
    * [x] Bitnami Amd64
    * [x] Bitnami-compat Amd64
    * [x] Bitnami-compat Arm64 (Mac Mini M1)
    * [ ] Official image Amd64
    * [ ] Official image Arm64

### Not Goal

* Support every images and charts on Bitnami's
  * Can't afford the time and cost of development and maintenance.

* Production ready or commercial support.
  * This repository is just open source and maintained by very few people.

### More detail

* These images use Bitnami's scripts. Binaries refer to the official Docker or official distribution binaries of the software or Docker recipes for Docker.

  * ![Bitnami-Compat](https://www.plantuml.com/plantuml/svg/XPH1pjem48NtSmgBJLUE4XTQLK88jOkgPwYw6E81B18xiXsXgkhTQySafFv_4bREctdFOhvCZiwZzNtR0BYRqXrQR9asaXnrR0SVdg95VOkYA07EGI2s-Quy8VlHueNOwJVxghp6LcqOEdQ99HKh5Vi3jMaxeqbxjWc5iQ5Y-xqeFdufIbv6l_DOtroh_FLmo2o4pzHeguhgSx5TiHdN9rSuF4oVrkvIMl6pQcWmXeTuJjvjkAzQibhIaYptr7OD-fIGj55AGULQq1sLvsTZUMTi6bPKWX2VnupTSeQhhUgy6zoFStG6nobH2kr_uKXQ1lOhGFcta8i-uofUSXKPL-py8b2H0PNchXk3abbg21sn2661hSZPcjxMuG_qgf7hT4L6zwLvdQz8RDQZ5WcBZ72UiiHOl6KCmwRDVjogozNp0njOGTgf5-K01uRrUTyVXm_JbkJrQSAGhy7KaWyGvc_w_s7p2S7qIkTtJeC0c7s2KW__ULLIaluL-vvL0UCU_0C0)

  * Some example : https://gist.github.com/ZCube/7e3045b1f15b1c1223f58267fc738e57

## Update bitnami-dockers schedule

* Every friday.

## Usage #1

* I know I can't make the images in this repository trustworthy, and I know I can't fix it. This repository is just a personal repository.
* If you want to use the images, It is recommended not to use the image directly, but to create Dockerfile.arm64 by following the instructions below and then verify and use it.

* Create Dockerfile.arm64 for bitnami-dockers
  1. install task https://taskfile.dev/installation/
  2. install golang https://go.dev/doc/install
  3. clone
  ```bash
  git clone --recursive https://github.com/ZCube/bitnami-compat
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
or
```bash
curl -LO https://raw.githubusercontent.com/ZCube/bitnami-compat/main/values/values_bitnami_compat_postgresql.yaml
helm install db bitnami/postgresql -f values_bitnami_compat_postgresql.yaml
```

## Test

* https://github.com/ZCube/bitnami-compat-test

* legend

  * O : test pass
  * X : test pass fail
  * Empty : not tested

* note
  * cypress v9 does not support arm64 so amd64 machine with ssh, docker needed to test.
  * arm64 test was done on mac mini m1

* result

    | App    | Tests   | bitnami (amd64) | bitnami-compat (amd64) | bitnami-compat (arm64) | official (amd64) | official (arm64) | official link |
    | :----- | :------ | :-------------: | :--------------------: | :--------------------: | :--------------: | :--------------: | :-----------: |
    | cert-manager  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | consul  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | consul  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | contour  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | discourse  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | discourse  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :x:     |                  |                  |               |
    | elasticsearch  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | elasticsearch  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | etcd  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | external-dns  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | fluentd  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | gitea  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | gitea  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | kafka  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | keycloak  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | keycloak  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | mariadb  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | mariadb-galera  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | memcached  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | minio  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | minio  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | mongodb  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | mongodb-sharded  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | mysql  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | nginx  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | nginx  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | postgresql  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | postgresql-ha  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | rabbitmq  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | rabbitmq  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | rabbitmq-cluster-operator  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | rabbitmq-cluster-operator  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | redis  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | redis-cluster  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | sealed-secrets  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | solr  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | solr  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | tomcat  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | tomcat  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | wordpress  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | wordpress  | Cypress |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |
    | zookeeper  | Goss    |  :heavy_check_mark: | :heavy_check_mark: | :heavy_check_mark:     |                  |                  |               |

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
✔️  gitea
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

