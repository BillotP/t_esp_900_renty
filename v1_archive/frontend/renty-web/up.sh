#!/usr/bin/env bash
# exit when any command fails
set -e
docker build -t repo.treescale.com/dave-lopeur/kubebeber/renty-web .
docker push repo.treescale.com/dave-lopeur/kubebeber/renty-web
k3s kubectl delete deployment renty-web || true
k3s kubectl apply -f deployment.yaml
