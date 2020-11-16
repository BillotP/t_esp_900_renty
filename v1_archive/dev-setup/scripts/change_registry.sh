#!/usr/bin/env bash
set -e
source .credential || source ../../.credential
A=$GITLAB_REGISTRY_URL
B=$TREESCALE_REGISTRY_URL
PS3='Please enter your choice: '
options=("$A" "$B" "Quit")
secretnameA=gitlab-registry 
secretnameB=treescale-registry
files=("backend/api/deployment.yaml" "backend/api/up.sh" "frontend/renty-web/deployment.yaml" "frontend/renty-web/up.sh" "backend/openfaas/stack.yml")
update_repo() {
    old=$1
    new=$2
    for f in ${files[@]}
    do
        echo -e "[INFO] Will update file $f"
	    sed -i "s;$old;$new;g" $f
        if [ $new = $A ]
        then
            sed -i "s;$secretnameB;$secretnameA;g" $f
        else
            sed -i "s;$secretnameA;$secretnameB;g" $f
        fi
    done
    echo "[INFO] All done"
}
select opt in "${options[@]}"
do
    case $opt in
        "$A")
            update_repo $B $A
            break
            ;;
        "$B")
            update_repo $A $B
            break
            ;;
        "Quit")
            break
            ;;
        *) echo "invalid option $REPLY";;
    esac
done