#!/usr/bin/env bash
# source .dev.env
k3s kubectl create secret generic jwt-public-key \
  --from-file=jwt-public-key=./rsa-2048bit-key-pair.pub \
  --namespace openfaas-fn
