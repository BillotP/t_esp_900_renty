#!/usr/bin/env bash
# exit when any command fails
set -e
source .credential || source ../../.credential
docker build --build-arg CI_USER=$GITHUB_USER --build-arg CI_TOKEN=$GITHUB_GOGET_TOKEN \
    -t repo.treescale.com/dave-lopeur/kubebeber/api \
    .
docker push repo.treescale.com/dave-lopeur/kubebeber/api
k3s kubectl delete deployment api || true
k3s kubectl apply -f deployment.yaml
