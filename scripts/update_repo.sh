#!/bin/bash
sed -i 's/amd64/${TARGETARCH}/g' $1
echo $1
echo $(dirname $1)
cat ${ROOT_DIR}/scripts/Dockerfile.head > _Dockerfile

rm _list*
cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' | tr -d "\"" > _list1
cat $(dirname $1)/Dockerfile | grep component_unpack | awk '{print $6"/"$7}' | tr -d "\""  > _list2
filename=$(ls _list* -l  --time-style=+""  | awk '{print $5"\t"$6}' | sort -h | tail -n 1 | awk '{print $2}')

list_size1=`stat -c%s "_list1"`
list_size2=`stat -c%s "_list2"`

if [[ "$list_size1" -lt "$list_size2" ]]; then
    filename="_list2"
elif [[ "$list_size1" -gt "$list_size2" ]]; then
    filename="_list1"
elif [[ "$list_size1" -eq "$list_size2" ]]; then
    filename="_list2"
fi

#cat $(dirname $1)/prebuildfs/opt/bitnami/.bitnami_components.json | jq 'to_entries[] | "\(.key)/\(.value.version)"' |
#cat $(dirname $1)/Dockerfile | grep component_unpack | awk '{print $6"/"$7}' | tr -d "\""
cat ${filename} | tr -d "\"" | xargs -I {} -n 1 /bin/bash ${ROOT_DIR}/scripts/docker_froms.sh {} >> _Dockerfile
cat $(dirname $1)/Dockerfile >> _Dockerfile
cat ${filename} | tr -d "\"" | xargs -I {} -n 1 /bin/bash ${ROOT_DIR}/scripts/golang_build_cmd.sh {} > $(dirname $1)/packages.sh

# 1. golang
echo "COPY --from=golang-builder /opt/bitnami/ /opt/bitnami/" > _install
# 2. docker
cat ${filename} | tr -d "\"" | xargs -I {} -n 1 /bin/bash ${ROOT_DIR}/scripts/docker_installs.sh {} >> _install
#echo "RUN apt-get update && apt-get upgrade -y && \\" >> _install

sed -i 's/RUN \. \/opt\/bitnami\/scripts\/libcomponent.sh/#RUN . \/opt\/bitnami\/scripts\/libcomponent.sh/g' _Dockerfile 

sed '/RUN apt-get update && apt-get upgrade -y && \\/e cat _install' -i _Dockerfile

cp -f _Dockerfile $(dirname $1)/Dockerfile

# 3. bash
cat ${filename} | tr -d "\"" | xargs -I {} -n 1 /bin/bash ${ROOT_DIR}/scripts/bash_build_cmd.sh {}  $(dirname $1)
