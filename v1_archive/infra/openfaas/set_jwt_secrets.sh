#!/usr/bin/env bash
set -e
source ../../backend/openfaas/.dev.env || source backend/openfaas/.dev.env

k3s kubectl create secret generic jwt-public-key \
  --from-file=jwt-public-key=./backend/openfaas/rsa-2048bit-key-pair.pub \
  --namespace openfaas-fn || true

k3s kubectl create secret generic jwt-private-key \
  --from-file=jwt-private-key=./backend/openfaas/rsa-2048bit-key-pair.pem \
  --namespace openfaas-fn || true

encoded=$(echo -n $JWT_ISSUER | base64)
k3s kubectl create secret generic jwt-issuer \
  --from-literal jwt-issuer=$JWT_ISSUER \
  -n openfaas-fn || kubectl patch secret jwt-issuer \
  -n openfaas-fn -p="{\"data\":{\"jwt-issuer\": \"$encoded\"}}" -v=1

encoded=$(echo -n $JWT_AUDIENCE | base64)
k3s kubectl create secret generic jwt-audience \
  --from-literal jwt-audience=$JWT_AUDIENCE \
  -n openfaas-fn || kubectl patch secret jwt-audience \
  -n openfaas-fn -p="{\"data\":{\"jwt-audience\": \"$encoded\"}}" -v=1

encoded=$(echo -n $JWT_EXPIRACY | base64)
k3s kubectl create secret generic jwt-expiracy \
  --from-literal jwt-expiracy=$JWT_EXPIRACY \
  -n openfaas-fn || kubectl patch secret jwt-expiracy \
  -n openfaas-fn -p="{\"data\":{\"jwt-expiracy\": \"$encoded\"}}" -v=1

encoded=$(echo -n $JWT_KEYID | base64)
k3s kubectl create secret generic jwt-key-id \
  --from-literal jwt-key-id=$JWT_KEYID \
  -n openfaas-fn || kubectl patch secret jwt-key-id \
  -n openfaas-fn -p="{\"data\":{\"jwt-key-id\": \"$encoded\"}}" -v=1

encoded=$(echo -n $XSRF_TOKEN_SEED | base64)
k3s kubectl create secret generic xsrf-token-seed \
  --from-literal xsrf-token-seed=$XSRF_TOKEN_SEED \
  -n openfaas-fn || k3s kubectl patch secret xsrf-token-seed \
  -n openfaas-fn -p="{\"data\":{\"xsrf-token-seed\": \"$encoded\"}}" -v=1
