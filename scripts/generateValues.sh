#!/bin/bash

rm -rf values/*.yaml
for d in ./charts/bitnami/*/ ; do
    [ -L "${d%/}" ] && continue
    [[ ${d} == *"common"* ]] && continue
    [[ ${d} == *"appsmith"* ]] && continue
    [[ ${d} == *"bitnami-shell"* ]] && continue
    HELM_APP_NAME=$(basename $d)
    yq e '(.. | select(has("tag")) | pick(["registry","repository","tag","digest"])) as $i ireduce({}; setpath($i | path; $i))' ./charts/bitnami/${HELM_APP_NAME}/values.yaml > tmp.yaml

    declare -a imageKeys=( $(yq e '(...|select(has("tag"))) |.repository | (. | parent | path | join("."))' tmp.yaml) )
    imageKeys=(${imageKeys[*]})

    declare -a imageRepositories=( $(yq e '(...|select(has("tag"))) |.repository | (. | sub("bitnami/(.*)", "${1}"))' tmp.yaml) )
    imageRepositories=(${imageRepositories[*]})

    len=${#imageKeys[@]}

    notFound=false
    for (( i=0; i<$len; i++ ));
    do 
        if [[ ! -z "${imageRepositories[$i]}" ]]
        then
            APP_TAG=$(APP_NAME=${imageRepositories[$i]} yq e 'with_entries(select(.[].name == env(APP_NAME))) | .[].version_full' ./versioninfo.yaml | tail -1);
            # echo "${imageKeys[$i]} : ${imageRepositories[$i]} : ${APP_TAG}";
            if [[ ! -z "${imageRepositories[$i]}" ]]
            then
                if [[ ${APP_TAG} == "null" ]]
                then
                    # echo "null found"
                    notFound=true
                fi
                APP_NAME=${imageRepositories[$i]} APP_TAG=${APP_TAG} yq e '(...|select(.repository == "bitnami/" + env(APP_NAME))) |= {"registry": "ghcr.io", "repository": "zcube/bitnami-compat/" + env(APP_NAME), "tag": env(APP_TAG), "digest":""} ' tmp.yaml > tmp.yaml_
                mv tmp.yaml_ tmp.yaml
            fi
        fi
    done

    if [ $notFound = true ]; then
        continue
    fi
    echo "================================================================"
    echo "$(basename $d)"
    echo "================================================================"
    mv tmp.yaml values/values_bitnami_compat_${HELM_APP_NAME}.yaml
    cat values/values_bitnami_compat_${HELM_APP_NAME}.yaml
done

