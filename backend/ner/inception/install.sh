#!/usr/bin/env bash
set -e
echo -e "[INFO] Downloading inception annotation tool if missing"
[ -f inception-app-standalone-0.17.0.jar ] || wget https://github.com/inception-project/inception/releases/download/inception-app-0.17.0/inception-app-standalone-0.17.0.jar
echo -e "[INFO] Building docker image"
docker build -t repo.treescale.com/dave-lopeur/kubebeber/inception
echo -e "[INFO] Pushing it to container registry"
docker push repo.treescale.com/dave-lopeur/kubebeber/inception
echo -e "[INFO] Drop existing kube deployment if exist"
k3s kubectl delete deployment inception || true
echo -e "[INFO] Create app kube deployment"
k3s kubectl apply -f deployment.yaml