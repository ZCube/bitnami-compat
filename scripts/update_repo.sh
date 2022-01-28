#!/bin/bash
sed -i 's/amd64/${TARGETARCH}/g' $1
echo $1
echo $(dirname $1)
cat ${ROOT_DIR}/scripts/Dockerfile.head > _Dockerfile
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 /bin/bash ${ROOT_DIR}/scripts/docker_froms.sh {} >> _Dockerfile
cat $(dirname $1)/Dockerfile >> _Dockerfile
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 /bin/bash ${ROOT_DIR}/scripts/golang_build_cmd.sh {} > $(dirname $1)/packages.sh

cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | xargs -I {} -n 1 /bin/bash ${ROOT_DIR}/scripts/docker_installs.sh {} > _install
echo "COPY --from=golang-builder /opt/bitnami/ /opt/bitnami/" >> _install
#echo "RUN apt-get update && apt-get upgrade -y && \\" >> _install

sed -i 's/RUN \. \/opt\/bitnami\/scripts\/libcomponent.sh/#RUN . \/opt\/bitnami\/scripts\/libcomponent.sh/g' _Dockerfile 

sed '/RUN apt-get update && apt-get upgrade -y && \\/e cat _install' -i _Dockerfile

cp -f _Dockerfile $(dirname $1)/Dockerfile
