#!/bin/bash

#!/bin/bash

declare -a imageNames=$( yq e 'with_entries(.) | .[].name' ./versioninfo.yaml )
imageNames=( ${imageNames[*]} )
declare -a imageVersionFulls=$( yq e 'with_entries(.) | .[].version_full' ./versioninfo.yaml )
imageVersionFulls=( ${imageVersionFulls[*]} )
declare -a imageVersionMajors=$( yq e 'with_entries(.) | .[].version_major' ./versioninfo.yaml )
imageVersionMajors=( ${imageVersionMajors[*]} )
declare -a imageVersions=$( yq e 'with_entries(.) | .[].version' ./versioninfo.yaml )
imageVersions=( ${imageVersions[*]} )
declare -a imageOsFlavours=$( yq e 'with_entries(.) | .[].os_flavour' ./versioninfo.yaml )
imageOsFlavours=( ${imageOsFlavours[*]} )

len=${#imageNames[@]}

notFound=false
for (( i=0; i<$len; i++ ));
do
  export IMAGE_TAG=ghcr.io/zcube/bitnami-compat/${imageNames[$i]}
  echo ${IMAGE_TAG}:${imageVersionFulls[$i]}

  docker pull --platform=linux/amd64 ${IMAGE_TAG}:${imageVersionFulls[$i]}
  docker tag ${IMAGE_TAG}:${imageVersionFulls[$i]} ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64
  docker pull --platform=linux/arm64 ${IMAGE_TAG}:${imageVersionFulls[$i]}
  docker tag ${IMAGE_TAG}:${imageVersionFulls[$i]} ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64

  docker-squash --tmp-dir "${PWD}/tmp" "${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64" -t ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash &
  docker-squash --tmp-dir "${PWD}/tmp" "${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64" -t ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash &
  wait

  docker push ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash &
  docker push ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash &
  wait

  docker manifest create ${IMAGE_TAG}:${imageVersionFulls[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push ${IMAGE_TAG}:${imageVersionFulls[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersions[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push ${IMAGE_TAG}:${imageVersions[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersions[$i]}-${imageOsFlavours[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push ${IMAGE_TAG}:${imageVersions[$i]}-${imageOsFlavours[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersionMajors[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push ${IMAGE_TAG}:${imageVersionMajors[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersionMajors[$i]}-${imageOsFlavours[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push ${IMAGE_TAG}:${imageVersionMajors[$i]}-${imageOsFlavours[$i]}-squash
done
