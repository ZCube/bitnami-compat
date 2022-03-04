#!/bin/bash
export ROOT_DIR=$(realpath "$(dirname $0)/..")
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

# Prepare directories
export versions=( $(cat ${filename} | awk -F'/' '{print $2}' | sed -E -e 's/-([0-9]+)//g') )
export versions_with_revision=( $(cat ${filename} | awk -F'/' '{print $2}') )
export packages=( $(cat ${filename} | awk -F'/' '{print $1}') )


mkdir -p patches/${packages[i]}/${versions[i]}/golang
mkdir -p patches/${packages[i]}/${versions[i]}/docker
mkdir -p patches/${packages[i]}/${versions[i]}/bash

# Dockerfile
cat ${ROOT_DIR}/scripts/Dockerfile.head > _Dockerfile
for (( i=0; i<${#packages[@]}; i++ )); do
echo "${packages[i]}/${versions[i]}" 
if test -f "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/docker/Dockerfile.from"; then
    cat "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/docker/Dockerfile.from" >> _Dockerfile
fi
done

cat $(dirname $1)/Dockerfile >> _Dockerfile

cat /dev/null > $(dirname $1)/packages.sh
for (( i=0; i<${#packages[@]}; i++ )); do
echo "${packages[i]}/${versions[i]}" 
if test -f "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/golang/install.sh"; then
    cat "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/golang/install.sh" >> $(dirname $1)/packages.sh
fi
done

# 1. golang
echo "COPY --from=golang-builder /opt/bitnami/ /opt/bitnami/" > _install
# 2. docker
for (( i=0; i<${#packages[@]}; i++ )); do
echo "${packages[i]}/${versions[i]}" 
if test -f "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/docker/Dockerfile.install"; then
    cat "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/docker/Dockerfile.install" >> _install
fi
done

# Comments
sed -i 's/RUN \. \/opt\/bitnami\/scripts\/libcomponent.sh/#RUN . \/opt\/bitnami\/scripts\/libcomponent.sh/g' _Dockerfile 

sed '/RUN apt-get update && apt-get upgrade -y && \\/e cat _install' -i _Dockerfile

cp -f _Dockerfile $(dirname $1)/Dockerfile

# 3. bash
for (( i=0; i<${#packages[@]}; i++ )); do
echo "${packages[i]}/${versions[i]}" 
if test -f "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/bash/install.sh"; then
    if test -d "$2"; then
        pushd $2
            bash "${ROOT_DIR}/patches/${packages[i]}/${versions[i]}/bash/install.sh"
        popd
    fi
fi
done
