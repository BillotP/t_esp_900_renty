#!/usr/bin/env bash
# exit when any command fails
set -e
source .credential || source ../../.credential
docker build --build-arg CI_USER=$GOGET_USER --build-arg CI_TOKEN=$GOGET_TOKEN \
    -t repo.treescale.com/dave-lopeur/kubebeber/api \
    .
docker push repo.treescale.com/dave-lopeur/kubebeber/api
k3s kubectl delete deployment api || true
k3s kubectl apply -f deployment.yaml
