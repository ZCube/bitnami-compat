
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

This section has moved to the [Wiki](https://github.com/ZCube/bitnami-compat/wiki/Creating-patch-guide).

## Compatible Images for Bitnami Charts

* Recommended for production use.

|Charts|Image|Status|
|------|-----|------|
|thanos|docker.io/thanosio/thanos:v0.27.0|Official images are compatible.|
|external-dns|registry.k8s.io/external-dns/external-dns:v0.12.2|Official images are compatible.|

## License

Apache 2.0

Copyright &copy; 2022 ZCube(zcubekr@gmail.com) 

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

### License for generator (Golang scripts)

MIT License

### License for patches

check References

### License for Bitnami's code.

https://github.com/bitnami/charts

https://github.com/bitnami/containers

Copyright &copy; 2022 Bitnami

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.


