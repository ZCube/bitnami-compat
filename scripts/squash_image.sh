#!/bin/bash
set -e

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
declare -a imageRevisions=$( yq e 'with_entries(.) | .[].revision' ./versioninfo.yaml )
imageRevisions=( ${imageRevisions[*]} )

len=${#imageNames[@]}

for (( i=0; i<$len; i++ ));
do
  if [ "${imageNames[$i]}" != "${APP}" ]; then
    continue
  fi

  export IMAGE_TAG_FROM=ghcr.io/zcube/bitnami-compat/cache/${imageNames[$i]}
  export IMAGE_TAG=ghcr.io/zcube/bitnami-compat/${imageNames[$i]}
  echo ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}
  echo ${IMAGE_TAG}:${imageVersionFulls[$i]}

  docker pull --platform=linux/amd64 ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}
  docker tag ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]} ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}-amd64
  docker pull --platform=linux/arm64 ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}
  docker tag ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]} ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}-arm64

  docker push ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}-amd64
  docker push ${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}-arm64

  crane flatten --platform=linux/amd64 "${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}-amd64" -t ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash
  crane flatten --platform=linux/arm64 "${IMAGE_TAG_FROM}:${imageVersionFulls[$i]}-arm64" -t ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash

  docker manifest create ${IMAGE_TAG}:${imageVersions[$i]}-${imageOsFlavours[$i]}-r${imageRevisions[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push   ${IMAGE_TAG}:${imageVersions[$i]}-${imageOsFlavours[$i]}-r${imageRevisions[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersionMajors[$i]}-${imageOsFlavours[$i]}-r${imageRevisions[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push   ${IMAGE_TAG}:${imageVersionMajors[$i]}-${imageOsFlavours[$i]}-r${imageRevisions[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersions[$i]}-${imageOsFlavours[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push   ${IMAGE_TAG}:${imageVersions[$i]}-${imageOsFlavours[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersions[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push   ${IMAGE_TAG}:${imageVersions[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersionMajors[$i]}-${imageOsFlavours[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push   ${IMAGE_TAG}:${imageVersionMajors[$i]}-${imageOsFlavours[$i]}-squash

  docker manifest create ${IMAGE_TAG}:${imageVersionMajors[$i]}-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-amd64-squash ${IMAGE_TAG}:${imageVersionFulls[$i]}-arm64-squash
  docker manifest push   ${IMAGE_TAG}:${imageVersionMajors[$i]}-squash

  docker image prune -a -f

done
