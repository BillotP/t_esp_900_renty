#!/usr/bin/env bash
set -e
###
echo -e "[INFO] Setup private registry secret in OpenFAAS namespace"
source ../../.credential || source .credential
k3s kubectl create secret docker-registry gitlab-registry \
   -n openfaas-fn \
   --docker-username=$GITLAB_USER \
   --docker-password=$GITLAB_TOKEN \
   --docker-email=lopeurd@gmail.com \
   --docker-server=registry.gitlab.com || true
k3s kubectl create secret docker-registry treescale-registry \
   -n openfaas-fn \
   --docker-username=$TREESCALE_USER \
   --docker-password=$TREESCALE_TOKEN \
   --docker-email=registry@therentyapp.com \
   --docker-server=repo.treescale.com || true
